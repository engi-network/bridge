# Goerli Testnet setup

## Deployment

Get chainbridge-deploy cli tool, these steps do not need to be repeated for Goerli:
```bash
git clone https://github.com/chainsafe/chainbridge-deploy
cd chainbridge-deploy/cb-sol-cli
make install
```

Back in the root of this repo:
``bash
cb-sol-cli --url http://rpc.goerli.mudit.blog --networkId 5 --privateKey $(node scripts/test-key.js ) deploy --all --chainId 5 --config
```

Contracts are deployed here:
```console
================================================================
Url:        http://rpc.goerli.mudit.blog
Deployer:   0x449f9748f5a19154e6fd012B3AE5Ec5Ad1a042f7
Gas Limit:   8000000
Gas Price:   20000000
Deploy Cost: 0.00029567156

Options
=======
Chain Id:    5
Threshold:   2
Relayers:    0xff93B45308FD417dF303D6515aB04D9e89a750Ca,0x8e0a907331554AF72563Bd8D43051C2E64Be5d35,0x24962717f8fA5BA3b931bACaF9ac03924EB475a0,0x148FfB2074A9e59eD58142822b3eB3fcBffb0cd7,0x4CEEf6139f00F9F4535Ad19640Ff7A0137708485
Bridge Fee:  0
Expiry:      100

Contract Addresses
================================================================
Bridge:             0xC63A1A2c5C5C3913d84707BFD1Ee83B2a6d9bF70
----------------------------------------------------------------
Erc20 Handler:      0x0cd4370317aD014170708F5c7378f09be54F93A2
----------------------------------------------------------------
Erc721 Handler:     0x12048bC6158DC104F1B9827364b0c19ab171eb0D
----------------------------------------------------------------
Generic Handler:    0x222f0900A8157d4c189F7B7dD1781D9A982Fc3F2
----------------------------------------------------------------
Erc20:              0x097A32b3560FA230BC9a5c64AaB4Ea28E5AfC3ef
----------------------------------------------------------------
Erc721:             0x60C6A313930C059FA0F103141ab9254E3496aAb0
----------------------------------------------------------------
Centrifuge Asset:   Not Deployed
----------------------------------------------------------------
WETC:               Not Deployed
================================================================
```
