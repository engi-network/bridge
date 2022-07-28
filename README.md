# Engi cross-chain bridge

## Requirements:
  - Docker
  - golang
  - nodejs
  - npm

## For Mac M1 compatibility
Needed for the chainbridge cli tools, since they build to amd64
```bash
softwareupdate --install-rosetta
```

## Build cli tool:
```bash
go build
```

Key setup for relayers, [see here](docs/account_setup.md)

ETH contract deployment [here](docs/goerli.md)

Substrate [resource registration](https://chainbridge.chainsafe.io/local/#register-relayers).
Note on the Method name: it's not immediately clear, but this turns out to be <pallet_name>.<method_name>


### Spin up a local test network

```bash
docker network create eth-goerli-test
```

This will bring up an eth light node connected to the Goerli testnet and a parity chain with the bridge pallet.
```bash
docker-compose -f ./configs/docker/goerli-chain.yaml up
```

Bring up a chainbridge:
```bash
docker-compose -f ./configs/docker/chainbridge-goerli.yaml up
```

There will be a metrics endpoint at `http://localhost:8001/metrics`.

Teardown:
```bash
docker-compose -f configs/docker/simple-local-chain.yaml down
```

If needed, an eth console can be brought up in a separate shell window with:
```bash
docker exec -it eth1 geth attach http://127.0.0.1:8545
```

```console
Welcome to the Geth JavaScript console!

instance: Geth/v1.10.20-stable-8f2416a8/linux-arm64/go1.18.3
coinbase: 0x0305af826e28b51e171bf9f41202ba428938bea6
at block: 0 (Thu Jan 01 1970 00:00:00 GMT+0000 (UTC))
 modules: eth:1.0 net:1.0 personal:1.0 rpc:1.0 web3:1.0

To exit, press ctrl-d or type exit
> eth.chainId()
"0x5"
```
