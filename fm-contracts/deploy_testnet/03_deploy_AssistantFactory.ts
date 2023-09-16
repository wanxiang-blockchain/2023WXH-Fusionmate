import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
    const { deployments, getNamedAccounts } = hre;
    const { deploy } = deployments;

    const { deployer } = await getNamedAccounts();
    const fmToken = await deployments.get('FMToken');
    const accountImpl = await deployments.get('ERC6551Account');

    await deploy('AssistantFactory', {
        from: deployer,
        args: [fmToken.address, accountImpl.address],
        log: true,
    });
};
export default func;
func.tags = ['AssistantFactory'];
func.dependencies = ['FMToken', 'ERC6551Account'];
