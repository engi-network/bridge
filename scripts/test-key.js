const ethers = require('ethers');

const data = {
    "address": "449f9748f5a19154e6fd012b3ae5ec5ad1a042f7",
    "crypto": {
      "cipher": "aes-128-ctr",
      "ciphertext": "f0b514e51be9ce06b79a83b0ab84092bea9b8ee76fb4b9768a8900b002324617",
      "cipherparams": {
        "iv": "944a59a3b119a80967d328e488503a45"
      },
      "kdf": "scrypt",
      "kdfparams": {
        "dklen": 32,
        "n": 262144,
        "p": 1,
        "r": 8,
        "salt": "d681a2352886f0c4f31841e27df0f345a374010982d300a011554e9357929868"
      },
      "mac": "6a39d7a1ce228a6541b97f7e9518fc92b2684f350ee3cd52853a4454f40c9b19"
    },
    "id": "5ed8e592-b88e-4013-ab28-58d189fd1d57",
    "version": 3
};

let walletJson = JSON.stringify(data);

const wallet = ethers.Wallet.fromEncryptedJson(walletJson, 'super_secret').then(function (wallet) {
	console.log(wallet.privateKey);
});

