// Check balance of a token owner

const ethers = require('ethers');
const abi = require('./erc20.json');
const token_addr = '0xb7D43b3c22389889A964f89F141b12D5fb1CA804';
const who = '0x893eE238147a93DD957626c6eC3cEd97FB44e52b';
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545');
const contract = new ethers.Contract(token_addr, abi, provider);

(async () => {
    var bal = await contract.balanceOf(who);
    console.log(bal);
})()
