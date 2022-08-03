# Goerli Testnet setup

## Deployment

Relayers list must be the set of accounts previously set up for the relayers
```bash
./bridge evm-cli --url 'http://localhost:8545' --private-key ${ADMIN_KEY} deploy --all --relayer-threshold 1 --domain 1 --relayers '0x2fb940fa2de638c694672d34b2f01f03fd9401c7,0xe6129866a913045c2fb99f6fa5150229690a5816,0xf2736b5fa1eb42d08bf98f350c8f15303ea1619c' --erc20-name ENGI --erc20-symbol ENGI --fee 0 
```

Deployed contracts:
```console
	Deployed contracts
=========================================================
Bridge: 0xCEE4a38919e4BB66a7B4aBd834fd2D79563aA013
---------------------------------------------------------
ERC20 Token: 0xb7D43b3c22389889A964f89F141b12D5fb1CA804
---------------------------------------------------------
ERC20 Handler: 0x52eEEFfD54d6e2dD7DF1Ce354623Ec67E4E4DEC5
---------------------------------------------------------
ERC721 Token: 0x767482AEF544F84E244c1eFe6BcB68875DbcA3a0
---------------------------------------------------------
ERC721 Handler: 0x4618522A2a6A826aEB2287e032E474B569007f19
---------------------------------------------------------
Generic Handler: 0x315978a6b7228832b16AdaaE418Ac6fa0c2D6Ce0
=========================================================
```

Deploy the vendor as well:
```bash
cd contracts
yarn; MNEMONIC_GOERLI='secret mnemonic ...' yarn hardhat --network goerli deploy
```

Vendor Contract:
```console
0x893eE238147a93DD957626c6eC3cEd97FB44e52b
```

Register a resource ID (can be arbirtary, except the lower byte identifies the 'home' chain):
Note: The resource MUST be a 32 byte value, or substrate will not see it the same as ethereum
```bash
./bridge --url 'http://localhost:8545' --private-key ${ADMIN_KEY} evm-cli bridge register-resource --bridge 0xCEE4a38919e4BB66a7B4aBd834fd2D79563aA013 --handler 0x52eEEFfD54d6e2dD7DF1Ce354623Ec67E4E4DEC5 --target 0xb7D43b3c22389889A964f89F141b12D5fb1CA804 --resource 0x00000000000000000000000000000063822bbd62abfb4ab9c92210c193e71b01
```

```conole
Registering resource
Handler address: 0x52eEEFfD54d6e2dD7DF1Ce354623Ec67E4E4DEC5
Resource ID: 0x00000000000000000000000000000063822bbd62abfb4ab9c92210c193e71b01
Target address: 0xb7D43b3c22389889A964f89F141b12D5fb1CA804
Bridge address: 0xCEE4a38919e4BB66a7B4aBd834fd2D79563aA013
```

Set ENGI as mintable/burnable:
```bash
./bridge evm-cli bridge set-burn --url 'http://localhost:8545' --private-key ${ADMIN_KEY} --handler 0x52eEEFfD54d6e2dD7DF1Ce354623Ec67E4E4DEC5 --bridge 0xCEE4a38919e4BB66a7B4aBd834fd2D79563aA013 --token-contract 0xb7D43b3c22389889A964f89F141b12D5fb1CA804
```

Allow ERC20 contract to mint:
```bash
./bridge evm-cli erc20 add-minter --url 'http://localhost:8545' --private-key ${ADMIN_KEY} --contract  0xb7D43b3c22389889A964f89F141b12D5fb1CA804 --minter 0x52eEEFfD54d6e2dD7DF1Ce354623Ec67E4E4DEC5
```

Mint some ENGI into Vendor contract!:
```bash
./bridge evm-cli erc20 mint --url 'http://localhost:8545' --private-key ${ADMIN_KEY} --amount 1000000 --decimals 18 --contract 0xb7D43b3c22389889A964f89F141b12D5fb1CA804 --recipient 0x893eE238147a93DD957626c6eC3cEd97FB44e52b
```

These steps performed by ENGI Vendor contract, but for testing purposes, here they are:

Authorize spending (recipient is actually the ERC20 handler contract):
```bash
./bridge --private-key ${ADMIN_KEY} --url http://localhost:8545 evm-cli erc20 approve --amount 100 --recipient 0x0 --decimals 18 --contract 0xb7D43b3c22389889A964f89F141b12D5fb1CA804  --recipient 0x52eEEFfD54d6e2dD7DF1Ce354623Ec67E4E4DEC5
```

Send some ENGI:
```bash
./bridge --private-key ${ADMIN_KEY} --url http://localhost:8545 evm-cli erc20 deposit --bridge 0xCEE4a38919e4BB66a7B4aBd834fd2D79563aA013 --recipient 0xe9623f3ca3CcC1c1D415F3196D4B75007B316aC3 --amount 2 --domain 2 --resource 0x00000000000000000000000000000063822bbd62abfb4ab9c92210c193e71b01 --decimals 12
```
