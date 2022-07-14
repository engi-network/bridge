# Goerli Testnet setup

## Deployment

Relayers list must be the set of accounts previously set up for the relayers
```bash
./chainbridge-core-example evm-cli --url 'http://localhost:8545' --private-key ${ADMIN_KEY} deploy --all --relayer-threshold 2 --domain 1 --relayers '0x2fb940fa2de638c694672d34b2f01f03fd9401c7,0xe6129866a913045c2fb99f6fa5150229690a5816,0xf2736b5fa1eb42d08bf98f350c8f15303ea1619c' --erc20-name ENGI --erc20-symbol ENGI --fee 0 
```

Deployed contracts:
```console
	Deployed contracts
=========================================================
Bridge: 0x90e630d8272074A07c453d4CB78205d2dE0521de
---------------------------------------------------------
ERC20 Token: 0xE5Ea4a579DA97CaD189E8aE9c560F4FB249E389D
---------------------------------------------------------
ERC20 Handler: 0xE6227a62c90E0b1B37163C0CB3972Cd1Cb7D3806
---------------------------------------------------------
ERC721 Token: 0x81abAb83D39Fb053e110cDFC668b72EAbDd99C99
---------------------------------------------------------
ERC721 Handler: 0xAF078A94A5B89b6407e413700332A86C3B6155c3
---------------------------------------------------------
Generic Handler: 0x420C4654e6317432c412cE1b0c401495b8e4820f
=========================================================
```

Register a resource ID (can be arbirtary, except the lower byte identifies the 'home' chain):
```bash
./chainbridge-core-example --url 'http://localhost:8545' --private-key ${ADMIN_KEY} evm-cli bridge register-resource --bridge 0x90e630d8272074A07c453d4CB78205d2dE0521de --handler 0xE6227a62c90E0b1B37163C0CB3972Cd1Cb7D3806 --target 0xE5Ea4a579DA97CaD189E8aE9c560F4FB249E389D --resource 0x454e474901
```

```conole
Registering resource
Handler address: 0xE6227a62c90E0b1B37163C0CB3972Cd1Cb7D3806
Resource ID: 0x454e474901
Target address: 0xE5Ea4a579DA97CaD189E8aE9c560F4FB249E389D
Bridge address: 0x90e630d8272074A07c453d4CB78205d2dE0521de
```

Set ENGI as mintable/burnable:
```bash
./chainbridge-core-example evm-cli bridge set-burn --url 'http://localhost:8545' --private-key ${ADMIN_KEY} --handler 0xE6227a62c90E0b1B37163C0CB3972Cd1Cb7D3806 --bridge 0x90e630d8272074A07c453d4CB78205d2dE0521de --token-contract 0xE5Ea4a579DA97CaD189E8aE9c560F4FB249E389D
```

Allow ERC20 contract to mint:
```bash
./chainbridge-core-example evm-cli erc20 add-minter --url 'http://localhost:8545' --private-key ${ADMIN_KEY} --contract 0xE5Ea4a579DA97CaD189E8aE9c560F4FB249E389D --minter 0xE6227a62c90E0b1B37163C0CB3972Cd1Cb7D380
```

Mint some ENGI!:
```bash
./chainbridge-core-example evm-cli erc20 mint --url 'http://localhost:8545' --private-key 7db559763bd2d2b25f6eccbeccf66d7c4248e72948f3acb2cefc28862bea58a6 --amount 1000 --decimals 18 --contract 0xE5Ea4a579DA97CaD189E8aE9c560F4FB249E389D --recipient 0x449f9748f5a19154e6fd012B3AE5Ec5Ad1a042f7
```
