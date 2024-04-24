# payment-updater

Generates proofs for posting payments to chain that stakers and operators can claim.

## Building

```bash
git clone git@github.com:Layr-Labs/eigenlayer-payment-updater.git

cd eigenlayer-payment-updater

# Download and install dependencies
make deps

# Build eigenlayer-payment-updater in your local OS/arch
make

# Run the test suite
make test
```

### Dockerfile

```bash
make docker

docker run payments-updater:latest [updater]
```

## Running

### With config

To run the `payment-updater` with a config file, copy the `config.yml.tpl` template to your desired location and fill in the specified fields.

Then, pass the path of the config file using the `--config=<path>` flag.

### Command line args

```bash
./bin/payment-updater updater \
    --debug true \
    --environment "dev" \
    --network "devnet" \ 
    --rpc-url "http://...." \
    --private-key "<ethereum private key>" \
    --aws-access-key-id "<aws access key id>" \
    --aws-secret-access-key "<aws secret key" \
    --aws-region "us-east-1" \ 
    --s3-output-bucket "s3://<url>" \
    --payment-coordinator-address "<contract address>"
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

# Flags and arguments

## Global flags

### `--config`

Optional, path to a file based config to use

### `--debug`

*Values:* `true, false`

Enables debug logging

## Updater

### `--environment`

The target environment

*Values:* `local, dev, preprod, prod`

### `--network`

The Ethereum network to target

*Values:* `local, devnet, holesky, mainnet`

### `--rpc-url`

Fully qualified URL to an Ethereum RPC node.

_Example_

```bash
https://ethereum-holesky-rpc.publicnode.com
```

### `--private-key`

An Ethereum account private key, in hexidecimal form.

### `--aws-access-key-id` & `--aws-access-secret-key`

An AWS AccessKeyId/AccessSecretKey pair that has access to the Athena cluster containing the data to build the payment proof as well as the specified S3 output bucket

### `--aws-region`

*Default value:* `us-east-1`

The region to use for Athena queries and S3 output.

### `--s3-output-bucket`

The bucket name and path where the output of the Athena query should be put.

_Example_

```bash
s3://<bucket name>/<bucket>/<path>
```

### `--payment-coordinator-address`

The contract address of the target payment coordinator contract used to post payment proofs

_Example_

```bash
0x56c119bD92Af45eb74443ab14D4e93B7f5C67896
```
