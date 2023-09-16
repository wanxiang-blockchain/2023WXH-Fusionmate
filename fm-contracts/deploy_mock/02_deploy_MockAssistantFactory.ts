import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
    const { deployments, getNamedAccounts } = hre;
    const { deploy } = deployments;

    const { deployer } = await getNamedAccounts();
    const fmToken = await deployments.get('FMToken');
    const accountImpl = await deployments.get('ERC6551Account');
    const registry = await deployments.get('ERC6551Registry');

    await deploy('MockAssistantFactory', {
        from: deployer,
        args: [fmToken.address, accountImpl.address, registry.address],
        log: true,
    });
};
export default func;
func.tags = ['MockAssistantFactory'];
func.dependencies = ['FMToken', 'ERC6551Account', 'ERC6551Registry'];
