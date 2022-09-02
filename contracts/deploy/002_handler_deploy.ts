import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {DeployFunction} from 'hardhat-deploy/types';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const {deployments, getNamedAccounts} = hre;
  const {deploy} = deployments;

  const {deployer, tokenOwner} = await getNamedAccounts();

  const Vendor = await deployments.get("Vendor");

  const handler = await deploy('ERC20Handler', {
    from: deployer,
    args: [
        // 1st arg, address of bridge contract
        '0xCEE4a38919e4BB66a7B4aBd834fd2D79563aA013',
        // 4th resourceID
        Vendor.address,
    ],
    log: true,
  });
};
export default func;
func.tags = ['ERC20Handler'];
