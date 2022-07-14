# Account generation for relayers:

Relayer accounts will need to maintain a balance for transaction fees.

### Setup relayer accounts

This will generate an ETH address for relayer1, relayer2 and relayer3
```bash
docker run -v ${PWD}/configs/docker:/datadir ethereum/client-go:stable --datadir /datadir account new --password /datadir/passtxt.txt
docker run -v ${PWD}/configs/docker:/datadir ethereum/client-go:stable --datadir /datadir account new --password /datadir/passtxt.txt
docker run -v ${PWD}/configs/docker:/datadir ethereum/client-go:stable --datadir /datadir account new --password /datadir/passtxt.txt
```

```console
> ls configs/docker/keystore 
UTC--2022-07-13T14-50-25.683253459Z--2fb940fa2de638c694672d34b2f01f03fd9401c7
UTC--2022-07-13T14-50-36.198591715Z--e6129866a913045c2fb99f6fa5150229690a5816
UTC--2022-07-13T14-50-32.448086088Z--f2736b5fa1eb42d08bf98f350c8f15303ea1619c
```

Import ethereum account keyfiles to chainbridge format:
```bash
for kf in $(ls configs/docker/keystore)
do
  docker run -v ${PWD}/configs/docker/:/data chainsafe/chainbridge --keystore /data/keys accounts import --privateKey  $(node scripts/test-key.js ${kf} ) --password super_secret
done
```

Set up substrate relayer accounts:
```bash
docker run -v ${PWD}/configs/docker/:/data chainsafe/chainbridge --keystore /data/keys accounts generate --sr25519 --password super_secret
docker run -v ${PWD}/configs/docker/:/data chainsafe/chainbridge --keystore /data/keys accounts generate --sr25519 --password super_secret
docker run -v ${PWD}/configs/docker/:/data chainsafe/chainbridge --keystore /data/keys accounts generate --sr25519 --password super_secret
```

```console
bridge % ls -l configs/docker/keys 
0x2FB940FA2De638c694672d34b2f01F03fd9401c7.key
0xE6129866a913045c2FB99f6fA5150229690a5816.key
0xF2736b5Fa1Eb42D08bf98f350C8f15303eA1619C.key
5CVNdoFdQcc1WggXkfEHdzEMUrnKR7jYToY3ieuoVFWAmS4C.key
5DUSiZ1VrR7yaSwAko9VeSqPPC3C8NHAYMM95AbHmYQYBiCW.key
5DqffveoskJchv135hAHdLUE54YAN7mkDhRzRwR6eWCASm25.key
```

Each relayer config will need a polkadot and an eth account:
```console
bridge % cat configs/docker/configs/relayer1.json 
{
    "chains": [
        {
            "name": "eth",
            "type": "ethereum",
            "id": "5",
            "endpoint": "ws://eth1:8546",
            "from": "0x2fb940fa2de638c694672d34b2f01f03fd9401c7",
            "opts": {
                "bridge": "",
                "erc20Handler": "",
                "erc721Handler": "",
                "genericHandler": "",
                "gasLimit": "1000000",
                "maxGasPrice": "20000000",
                "startBlock": "7208056",
                "blockConfirmations": "2"
            }
        },
        {
            "name": "sub",
            "type": "substrate",
            "id": "1",
            "endpoint": "ws://sub-chain:9944",
            "from": "5CVNdoFdQcc1WggXkfEHdzEMUrnKR7jYToY3ieuoVFWAmS4C",
            "opts": {
                "useExtendedCall": "true",
                "startBlock": "0"
            }
        }
    ]
}
```

