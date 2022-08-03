// Example showing how to buy ENGI from ethereum chain.
// User must submit their substrate address and an amount of ENGI to buy via the purchase contract.
// ENGI will then show up on the substrate side of the chain.

const uc = require('@polkadot/util-crypto');

function substrateToHex(address) {
    var ar = uc.decodeAddress(address);

    return '0x' + Array.from(ar, function (byte) {
        return ('0' + (byte & 0xFF).toString(16)).slice(-2);
    }).join('');
}

const PURCHASE_CONTRACT = '0x893eE238147a93DD957626c6eC3cEd97FB44e52b';
const abi = require('./Vendor.json');
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545');
const nemo = 'monster various october car donkey plug float kind perfect nation fog extend';
const wallet = await new ethers.Wallet(ethers.Wallet.fromMnemonic(nemo).privateKey, provider);
const contract = new ethers.Contract(PURCHASE_CONTRACT, abi.abi, wallet);
const substrate_recipient = substrateToHex('5HWXo5wsUhmg9i67szY69YaGN8PAWTgj1WhCguyk6FzRqxog');
const overrides = {
    value: ethers.utils.parseEther("1.0"),
};
contract.connect(wallet);

await contract.deposit(substrate_recipient, overrides);
