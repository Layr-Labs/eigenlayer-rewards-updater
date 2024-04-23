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
