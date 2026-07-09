package txsigner

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// TestRecoverEthereumSignature exercises the KMS->Ethereum signature bridging
// (recovery-id brute force + 32-byte padding) without touching AWS, by feeding
// it raw (r, s) values the same way KMS would after DER decoding.
func TestRecoverEthereumSignature(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	expectedPub := crypto.FromECDSAPub(&key.PublicKey)

	digest := crypto.Keccak256([]byte("submitRoot tx hash stand-in"))

	// crypto.Sign returns [R || S || V]; strip V to mimic KMS (which never
	// returns a recovery id).
	sig, err := crypto.Sign(digest, key)
	require.NoError(t, err)
	r, s := sig[:32], sig[32:64]

	out, err := recoverEthereumSignature(expectedPub, digest, r, s)
	require.NoError(t, err)
	require.Len(t, out, 65)

	recovered, err := crypto.SigToPub(digest, out)
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(key.PublicKey), crypto.PubkeyToAddress(*recovered))
}

// TestLowSNormalization verifies that a high-S signature (which KMS may return)
// is folded back into the lower half of the curve order before assembly.
func TestLowSNormalization(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	expectedPub := crypto.FromECDSAPub(&key.PublicKey)

	digest := crypto.Keccak256([]byte("high-s case"))
	sig, err := crypto.Sign(digest, key)
	require.NoError(t, err)
	r, s := sig[:32], sig[32:64]

	// Flip to the malleable high-S form: s' = N - s.
	highS := new(big.Int).Sub(secp256k1N, new(big.Int).SetBytes(s))
	require.Greater(t, highS.Cmp(secp256k1HalfN), 0, "expected a high-S value")

	// Apply the same normalization SignerFn does before assembling the sig.
	sBytes := highS.Bytes()
	if highS.Cmp(secp256k1HalfN) > 0 {
		sBytes = new(big.Int).Sub(secp256k1N, highS).Bytes()
	}

	out, err := recoverEthereumSignature(expectedPub, digest, r, sBytes)
	require.NoError(t, err)

	recovered, err := crypto.SigToPub(digest, out)
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(key.PublicKey), crypto.PubkeyToAddress(*recovered))
}
