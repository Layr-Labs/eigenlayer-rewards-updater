# rewards-updater

Generates proofs for posting rewards to chain that stakers and operators can claim.

The rewards updater consumes a json line file containing the cumulative rewards for each staker and operator, and generates a proof for each staker and operator (called "earners") to claim their rewards.

The proof is generated by building a merkle tree of ordered earner/token pairs. Each leaf node represents an earner, from which a nested tree representing all of the earner's rewards tokens and their associated amounts. Since the earner and token addresses are ordered on each merkleization process, the indices for earners and tokens are not fixed across roots.

## Building

```bash
git clone git@github.com:Layr-Labs/eigenlayer-rewards-updater.git

cd eigenlayer-rewards-updater

# Download and install dependencies
make deps

# Build eigenlayer-rewards-updater in your local OS/arch
make

# Run the test suite
make test
```

### Dockerfile

```bash
make docker

docker run eigenlayer-rewards-updater:latest [updater]
```

## Running

### Public docker image

A publicly hosted Docker image can be found at:

```bash
public.ecr.aws/z6g0f8n7/eigenlayer-rewards-updater:latest
```

This image is a multi-arch capable image, currently built for `linux/amd64` and `linux/arm64`. `latest` is the only tag available at the moment with plans to ship versioned images soon.

### With config

To run the `rewards-updater` with a config file, copy the `config.yml.tpl` template to your desired location and fill in the specified fields.

Then, pass the path of the config file using the `--config=<path>` flag.

### Command line args

```bash
./bin/eigenlayer-rewards-updater updater \
    --debug true \
    --environment "dev" \
    --network "devnet" \ 
    --rpc-url "http://...." \
    --private-key "<ethereum private key>" \
    --rewards-coordinator-address "<contract address>" \
    --proof-store-base-url "http://...."
```

### docker-compose

To run with docker compose, simply execute:

```bash
docker compose up <service>

docker compose up updater
```

By default, docker compose will be looking for a config file in the root of this project.

If you wish to provide the parameters through flags, update the `updater-cli-args` service and run:

```bash
docker compose up updater-cli-args
```

### Kubernetes (Helm chart)

The rewards updater comes with a helm chart that will deploy a cron job to run the updater at a specified interval.

To run the helm chart, copy the `values.yaml` file to your desired location and fill in the specified fields (the config file value needs to be updated with values).

Then, run the following command:

```bash
helm upgrade --install \                                                                                                                                                                                                                                                       (sm-readmeUpdate✱) 
      --atomic \    
      --cleanup-on-fail \
      --timeout 2m \
      --force \
      --debug \
      --wait  \
      --version=$(date +%s) \
      -f ./eigenlayer-rewards-updater/<values file> \
      rewards-updater ./eigenlayer-rewards-updater
```

# Flags and arguments

## Global flags

### `--config`

EnvVar: `EIGENLAYER_CONFIG`

Optional, path to a file based config to use

### `--debug`

EnvVar: `EIGENLAYER_DEBUG`

*Values:* `true, false`

Enables debug logging

### --enable-statsd

EnvVar: `EIGENLAYER_ENABLE_STATSD`

Enable/disable statsd metrics collection. Defaults to `true` and will attempt to auto-detect the DataDog statsd collection agent from the environment.

### --enable-tracing

EnvVar: `EIGENLAYER_ENABLE_TRACING`

Enable/disable tracing. Defaults to `true` and will attempt to auto-detect the DataDog tracing agent from the environment.

### `--pushgateway-enabled`

EnvVar: `EIGENLAYER_PUSHGATEWAY_ENABLED`

Enable/disable pushgateway metrics collection. Defaults to `false`.

### `--pushgateway-url`

EnvVar: `EIGENLAYER_PUSHGATEWAY_URL`

The URL of the pushgateway to send metrics to.

## Updater

### `--environment`

EnvVar: `EIGENLAYER_ENVIRONMENT`

The target environment

*Values:* `preprod, testnet, mainnet`

### `--network`

EnvVar: `EIGENLAYER_NETWORK`

The Ethereum network to target

*Values:* `holesky, ethereum`

### `--rpc-url`

EnvVar: `EIGENLAYER_RPC_URL`

Fully qualified URL to an Ethereum RPC node.

_Example_

```bash
https://ethereum-holesky-rpc.publicnode.com
```

### `--private-key`

EnvVar: `EIGENLAYER_PRIVATE_KEY`

An Ethereum account private key, in hexidecimal form.

### `--rewards-coordinator-address`

EnvVar: `EIGENLAYER_REWARDS_COORDINATOR_ADDRESS`

The contract address of the target rewards coordinator contract used to post reward proofs

_Example_

```bash
0x56c119bD92Af45eb74443ab14D4e93B7f5C67896
```

### `--proof-store-base-url`

EnvVar: `EIGENLAYER_PROOF_STORE_BASE_URL`

The base URL of where rewards data is stored.

e.g.

```bash
https://eigenlabs-rewards-testnet-holesky.s3.amazonaws.com
```

### `--root-index`

EnvVar: `EIGENLAYER_ROOT_INDEX`

The index of the root to disable

e.g.

```bash
10
```


The proof store will fetch two files from this URL with the following paths:

```bash
# Recent snapshots
<base url>/<environment>/<network>/recent-snapshots.json

# Claim amounts
<base url>/<environment>/<network>/<snapshot date>/claim-amounts.json

```

## Deploying to Kubernetes with Helm

This repo comes with a helm chart that enables you to deploy the updater in an automated fashion, running it as a cronjob.

This chart follows the standard patterns of all other helm charts; values can be overridden in a separate yaml file or passed as flags when invoking helm.

Example:
```bash
helm upgrade --install \
  --atomic \
  --cleanup-on-fail \
  --timeout 2m \
  --force \
  --debug \
  --wait  \
  --version=$(date +%s) \
  -f ./eigenlayer-rewards-updater/values.yaml \
  eigenlayer-rewards-updater ./eigenlayer-rewards-updater
```

# Environment-specific values

### `testnet-holesky`

```yaml
environment: testnet
network: holesky
rpc_url: https://ethereum-holesky-rpc.publicnode.com
proof_store_base_url: https://eigenlabs-rewards-testnet-holesky.s3.amazonaws.com
rewards_coordinator_address: 0xAcc1fb458a1317E886dB376Fc8141540537E68fE
```

### `mainnet-ethereum`

```yaml
environment: mainnet
network: ethereum
rpc_url: https://ethereum-rpc.publicnode.com
proof_store_base_url: https://eigenlabs-rewards-mainnet-ethereum.s3.amazonaws.com
rewards_coordinator_address: 0x7750d328b314effa365a0402ccfd489b80b0adda
```

## CLI Examples

### Disable root

```bash
eigenlayer-rewards-updater disable-root \
    --environment testnet \
    --network  holesky \
    --rpc-url https://ethereum-holesky-rpc.publicnode.com \
    --rewards-coordinator-address "0xAcc1fb458a1317E886dB376Fc8141540537E68fE" \
    --private-key "<ethereum private key>" \
    --root-index 10
```
