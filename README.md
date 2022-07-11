# Engi cross-chain bridge

## Requirements:
  - Docker
  - nodejs
  - npm


### Spin up a local test network

This will bring up an eth light node connected to the Goerli testnet and a parity chain with the bridge pallet.
```bash
docker-compose -f ./configs/docker/goerli-chain.yaml up
```

For details of the bridge contract deployment, [see here](docs/goerli.md).

Bring up a chainbridge:
```bash
docker-compose -f ./configs/docker/chainbridge-goerli.yaml up
```

Teardown:
```bash
docker-compose -f configs/docker/simple-local-chain.yaml down -v
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
