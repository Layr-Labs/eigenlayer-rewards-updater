package txsigner

// AWS KMS Ethereum transaction signer.
//
// Adapted from github.com/Layr-Labs/multichain-go/pkg/txSigner/awsKmsSigner.go
// (MIT). AWS KMS has no Ethereum awareness; it only exposes raw secp256k1 ECDSA
// signing over a digest. This wrapper bridges the gaps so a KMS key can sign
// Ethereum transactions:
//
//   - KMS key must use spec ECC_SECG_P256K1 (same curve as Ethereum) and is
//     signed with algorithm ECDSA_SHA_256 + MessageType=DIGEST, so KMS signs the
//     already-keccak256'd tx hash as-is (no re-hashing).
//   - KMS returns the signature ASN.1/DER-encoded; we parse out (r, s).
//   - We normalize to low-S (EIP-2) to avoid malleability.
//   - KMS does not return the recovery id v; we brute-force v in {0,1} and keep
//     whichever recovers the known KMS public key.
//   - The address is derived from the KMS public key (the private key never
//     leaves AWS).

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	secp256k1N     = crypto.S256().Params().N
	secp256k1HalfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

// AWSKMSSigner implements ITransactionSigner using AWS KMS.
type AWSKMSSigner struct {
	kmsClient *kms.KMS
	keyID     string
	publicKey *ecdsa.PublicKey
	address   common.Address
}

// NewAWSKMSSigner connects to AWS KMS, fetches the public key for keyID, and
// derives the Ethereum address. region may be empty to defer to the standard
// AWS config/credential chain (env vars, ~/.aws/config, instance role, etc.).
func NewAWSKMSSigner(keyID, region string) (*AWSKMSSigner, error) {
	awsConfig := aws.NewConfig()
	if region != "" {
		awsConfig = awsConfig.WithRegion(region)
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // honor ~/.aws/config
		Config:            *awsConfig,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	signer := &AWSKMSSigner{
		kmsClient: kms.New(sess),
		keyID:     keyID,
	}

	pubKey, err := signer.getPubKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get public key from KMS: %w", err)
	}
	signer.publicKey = pubKey
	signer.address = crypto.PubkeyToAddress(*pubKey)

	return signer, nil
}

func (k *AWSKMSSigner) GetAddress() common.Address {
	return k.address
}

// SignerFn returns a bind.SignerFn that signs transaction hashes via KMS.
func (k *AWSKMSSigner) SignerFn(chainID *big.Int) bind.SignerFn {
	expectedPubKey := crypto.FromECDSAPub(k.publicKey) // 65-byte 0x04||X||Y
	signer := types.LatestSignerForChainID(chainID)

	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != k.address {
			return nil, fmt.Errorf("address mismatch: expected %s, got %s", k.address.Hex(), address.Hex())
		}

		txHashBytes := signer.Hash(tx).Bytes()

		rBytes, sBytes, err := k.getSignatureFromKms(txHashBytes)
		if err != nil {
			return nil, err
		}

		// Normalize S to the lower half of the curve order (EIP-2).
		sBigInt := new(big.Int).SetBytes(sBytes)
		if sBigInt.Cmp(secp256k1HalfN) > 0 {
			sBytes = new(big.Int).Sub(secp256k1N, sBigInt).Bytes()
		}

		signature, err := recoverEthereumSignature(expectedPubKey, txHashBytes, rBytes, sBytes)
		if err != nil {
			return nil, err
		}

		return tx.WithSignature(signer, signature)
	}
}

func (k *AWSKMSSigner) getPublicKeyDerBytesFromKMS() ([]byte, error) {
	out, err := k.kmsClient.GetPublicKey(&kms.GetPublicKeyInput{
		KeyId: aws.String(k.keyID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get public key from KMS: %w", err)
	}

	var asn1pubk asn1EcPublicKey
	if _, err := asn1.Unmarshal(out.PublicKey, &asn1pubk); err != nil {
		return nil, fmt.Errorf("failed to parse ASN.1 public key: %w", err)
	}
	return asn1pubk.PublicKey.Bytes, nil
}

func (k *AWSKMSSigner) getPubKey() (*ecdsa.PublicKey, error) {
	pkBytes, err := k.getPublicKeyDerBytesFromKMS()
	if err != nil {
		return nil, err
	}
	pubkey, err := crypto.UnmarshalPubkey(pkBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %w", err)
	}
	return pubkey, nil
}

func (k *AWSKMSSigner) getSignatureFromKms(txHashBytes []byte) ([]byte, []byte, error) {
	out, err := k.kmsClient.Sign(&kms.SignInput{
		KeyId:            aws.String(k.keyID),
		SigningAlgorithm: aws.String("ECDSA_SHA_256"),
		MessageType:      aws.String("DIGEST"),
		Message:          txHashBytes,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("KMS sign failed: %w", err)
	}

	var sigAsn1 asn1EcSig
	if _, err := asn1.Unmarshal(out.Signature, &sigAsn1); err != nil {
		return nil, nil, fmt.Errorf("failed to parse ASN.1 signature: %w", err)
	}
	return sigAsn1.R.Bytes, sigAsn1.S.Bytes, nil
}

// recoverEthereumSignature assembles r||s||v, brute-forcing v in {0,1} until the
// recovered public key matches the KMS key.
func recoverEthereumSignature(expectedPublicKeyBytes, txHash, r, s []byte) ([]byte, error) {
	rsSignature := append(adjustSignatureLength(r), adjustSignatureLength(s)...)

	for _, v := range []byte{0, 1} {
		signature := append(append([]byte{}, rsSignature...), v)
		recovered, err := crypto.Ecrecover(txHash, signature)
		if err != nil {
			continue
		}
		if bytes.Equal(recovered, expectedPublicKeyBytes) {
			return signature, nil
		}
	}
	return nil, fmt.Errorf("no recovery id matched expected KMS public key %s", hex.EncodeToString(expectedPublicKeyBytes))
}

func adjustSignatureLength(buffer []byte) []byte {
	buffer = bytes.TrimLeft(buffer, "\x00")
	for len(buffer) < 32 {
		buffer = append([]byte{0}, buffer...)
	}
	return buffer
}

// ASN.1 structures returned by KMS.
type asn1EcSig struct {
	R asn1.RawValue
	S asn1.RawValue
}

type asn1EcPublicKey struct {
	EcPublicKeyInfo asn1EcPublicKeyInfo
	PublicKey       asn1.BitString
}

type asn1EcPublicKeyInfo struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.ObjectIdentifier
}
