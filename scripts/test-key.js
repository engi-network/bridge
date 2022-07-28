const fs = require('fs');
const ethers = require('ethers');

const walletJson = fs.readFileSync(process.argv[2]).toString();

const wallet = ethers.Wallet.fromEncryptedJson(walletJson, 'super_secret').then(function (wallet) {
	console.log(wallet.privateKey);
});

