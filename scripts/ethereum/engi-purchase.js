// Example showing how to buy ENGI from ethereum chain.
// User must submit their substrate address and an amount of ENGI to buy via the purchase contract.
// ENGI will then show up on the substrate side of the chain.

// run it ->
// npx hardhat run engi-purchase.js 

const uc = require('@polkadot/util-crypto');

function substrateToHex(address) { var ar = uc.decodeAddress(address); return '0x' + Array.from(ar, function (byte) { return ('0' + (byte & 0xFF).toString(16)).slice(-2); }).join(''); }

    // testnet
    // see https://github.com/engi-network/website/blob/master/src/utils/ethereum/constants.ts
const PURCHASE_CONTRACT = '0xf054Be6AC90377Fa13d80e4a7528425945d13F8E';
const TOKEN_CONTRACT = '0xb7d43b3c22389889a964f89f141b12d5fb1ca804';
const BRIDGE_CONTRACT = '0xcee4a38919e4bb66a7b4abd834fd2d79563aa013';
const abi = require('./Vendor.json');
const provider = new ethers.providers.JsonRpcProvider('https://goerli.infura.io/v3/506381d31443434ba331ada55b1eb07e');
const nemo = 'monster various october car donkey plug float kind perfect nation fog extend';
const wallet = await new ethers.Wallet(ethers.Wallet.fromMnemonic(nemo).privateKey, provider);
const contract = new ethers.Contract(PURCHASE_CONTRACT, abi.abi, wallet);
const substrate_recipient = substrateToHex('5HWXo5wsUhmg9i67szY69YaGN8PAWTgj1WhCguyk6FzRqxog');
// 300 Gwei limmit
const overrides = { value: ethers.utils.parseEther("0.01"), gasLimit: 160000, gasPrice: ethers.utils.parseEther("0.0000003") };
contract.connect(wallet);

await contract.deposit(substrate_recipient, overrides);
})()
