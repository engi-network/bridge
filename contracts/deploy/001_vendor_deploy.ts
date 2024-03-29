import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {DeployFunction} from 'hardhat-deploy/types';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const {deployments, getNamedAccounts} = hre;
  const {deploy} = deployments;

  const {deployer, tokenOwner} = await getNamedAccounts();
console.log(`MY ADDER ${deployer}`);

  await deploy('Vendor', {
    from: deployer,
    args: [
        // 1st arg, address of bridge contract
        '0xCEE4a38919e4BB66a7B4aBd834fd2D79563aA013',
        // 2nd arg, address of ERC20 token
        '0xb7D43b3c22389889A964f89F141b12D5fb1CA804',
        // 3th resourceID
        '0x00000000000000000000000000000063822bbd62abfb4ab9c92210c193e71b01',
    ],
    log: true,
  });
};
export default func;
func.tags = ['Vendor'];
