// Example showing how to buy ENGI from ethereum chain.
// User must submit their substrate address and an amount of ENGI to buy via the purchase contract.
// ENGI will then show up on the substrate side of the chain.


const PURCHASE_CONTRACT = '0xf054Be6AC90377Fa13d80e4a7528425945d13F8E';
const abi = require('./Vendor.json');
const provider = new ethers.providers.JsonRpcProvider('https://goerli.infura.io/v3/506381d31443434ba331ada55b1eb07e');
const nemo = 'monster various october car donkey plug float kind perfect nation fog extend';
const wallet = await new ethers.Wallet(ethers.Wallet.fromMnemonic(nemo).privateKey, provider);

const contract = new ethers.Contract(PURCHASE_CONTRACT, abi.abi, wallet);
contract.connect(wallet);

await contract.setHandler('0xA56E0DE18617C8dee8ffc0137f2c5B568150B5A7');
