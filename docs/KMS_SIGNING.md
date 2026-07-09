# Signing `submitRoot` with AWS KMS

The `updater` command reads the weekly rewards root from the sidecar
(`GenerateRewardsRoot`) and submits it on-chain via `RewardsCoordinator.submitRoot`.
By default it signs that transaction with a local private key. It can instead
sign with **AWS KMS**, so the signing key never leaves the HSM — only the
derived address and IAM-gated `Sign` calls are exposed.

Nothing about the rewards flow changes: same sidecar call, same root, same
transaction. Only *how the transaction is signed* changes.

## How KMS signs an Ethereum transaction

AWS KMS has no Ethereum awareness; it only offers raw secp256k1 ECDSA over a
digest. `pkg/txsigner/awsKmsSigner.go` (adapted from
[`multichain-go`](https://github.com/Layr-Labs/multichain-go/blob/master/pkg/txSigner/awsKmsSigner.go))
bridges the gaps:

- **Same curve.** The KMS key must use spec `ECC_SECG_P256K1`, the curve
  Ethereum uses.
- **No double hashing.** We call `Sign` with `MessageType=DIGEST` and the
  already-`keccak256` tx hash, so KMS signs it as-is (the `ECDSA_SHA_256`
  algorithm name does not cause a re-hash for `DIGEST`).
- **DER → (r, s).** KMS returns an ASN.1/DER signature; we decode `r` and `s`.
- **Low-S.** We normalize `s` to the lower half of the curve order (EIP-2).
- **Recovery id.** KMS doesn't return `v`; we try `v ∈ {0,1}` and keep whichever
  recovers the known KMS public key.
- **Address.** Derived from the KMS public key via `GetPublicKey`.

## AWS prerequisites

1. **Create a KMS key** for signing:
   - Key type: **Asymmetric**
   - Key spec: **`ECC_SECG_P256K1`**
   - Key usage: **Sign and verify**
2. **Find the signer address.** It's derived from the key's public key. (The
   updater logs `signer_address` at startup; or derive it offline from
   `GetPublicKey`.)
3. **Authorize that address** as the `rewardsUpdater` on the target
   `RewardsCoordinator`. If it isn't authorized, `submitRoot` reverts.
4. **IAM**: the host/role running the updater needs `kms:Sign` and
   `kms:GetPublicKey` on the key. Credentials/region resolve through the
   standard AWS chain (env vars, `~/.aws/config`, instance/task role). `--aws-region`
   is optional and overrides the region.

## Usage

Private key (default, unchanged):

```bash
updater updater \
  --rpc-url "$ETH_RPC_URL" \
  --sidecar-rpc-url "$SIDECAR_RPC_URL" \
  --rewards-coordinator-address 0x7750d328b314EfFa365A0402Ccfd489B80B0adda \
  --private-key "$PRIVATE_KEY"
```

AWS KMS:

```bash
updater updater \
  --rpc-url "$ETH_RPC_URL" \
  --sidecar-rpc-url "$SIDECAR_RPC_URL" \
  --rewards-coordinator-address 0x7750d328b314EfFa365A0402Ccfd489B80B0adda \
  --signer-type aws_kms \
  --kms-key-id "arn:aws:kms:us-east-1:1234567890:key/abcd-...." \
  --aws-region us-east-1
```

All flags also bind to env vars with the `EIGENLAYER_` prefix and underscores,
e.g. `EIGENLAYER_SIGNER_TYPE=aws_kms`, `EIGENLAYER_KMS_KEY_ID=...`.

> The sidecar RPC URL must point at a sidecar that supports `GenerateRewardsRoot`
> (i.e. the one your node runs). The public read-only gateway
> (`sidecar-rpc.eigenlayer.xyz/...`) only exposes GET endpoints and cannot
> generate a root to submit.

## Running at cadence

The command is one-shot (run, submit, exit) and is driven by the existing
external schedule (cron). It returns a non-zero exit code on failure so the
scheduler can alert. Submitting is idempotent: the contract rejects a root whose
`rewardsCalculationEndTimestamp` is not newer than the current one, so an extra
run is a no-op. Switching to KMS does not change any of this — only the
`--signer-type`/`--kms-key-id` flags are added to the existing invocation.
