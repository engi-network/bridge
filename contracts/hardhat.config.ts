import {HardhatUserConfig} from 'hardhat/types';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';
import "@nomiclabs/hardhat-ethers";
import {accounts} from './utils/network';

const config: HardhatUserConfig = {
  solidity: {
    version: '0.8.4',
  },
  namedAccounts: {
    deployer: 0,
  },
  networks: {
    local: {
        url: 'http://127.0.0.1:9989',
    },
    goerli: {
        url: 'https://goerli.infura.io/v3/506381d31443434ba331ada55b1eb07e',
        accounts: accounts('goerli'),
    }
  },
  paths: {
    sources: 'src',
  },
};
export default config;

