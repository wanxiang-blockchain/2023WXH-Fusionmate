import { expect } from 'chai';

import { ethers, deployments, getNamedAccounts } from 'hardhat';

describe('MockAssistantFactory contract', function () {
    it('create assistant should work', async function () {
        await deployments.fixture(['MockAssistantFactory']);
        const { deployer, trader } = await getNamedAccounts();
        const factory = await deployments.get('MockAssistantFactory');

        const botFactory = await ethers.getContractAt('MockAssistantFactory', factory.address);
        console.log('MockAssistantFactory contract address is', botFactory.address);

        // const registry = await botFactory.registry();
        // const accountImpl = await botFactory.accountImpl();
        // const fmToken = await botFactory.fmToken();

        const mintPrice = ethers.utils.parseUnits('100');
        const tx1 = await botFactory.createAssistantMock('bot1', 'AI1', 'https://fusionMate.io', 10, mintPrice, 11);
        tx1.wait();
        const traderSign = await ethers.getSigner(trader);
        const tx2 = await botFactory.connect(traderSign).createAssistantMock('bot2', 'AI2', 'https://fusionMate.io', 5, mintPrice, 222);
        tx2.wait();

        const bot1Address = await botFactory.assistantsMap(11);
        const bot2Address = await botFactory.assistantsMap(222);
        const botNotExist = await botFactory.assistantsMap(333); // address(0)
        // console.log('bot1 address: ', bot1Address);
        // console.log('bot2 address: ', bot2Address);

        const maker1 = await botFactory.assistantMaker(bot1Address);
        const maker2 = await botFactory.assistantMaker(bot2Address);
        expect(maker1).to.equal(deployer);
        expect(maker2).to.equal(trader);
    });
    it('create account should work', async function () {
        await deployments.fixture(['MockAssistantFactory']);
        const { deployer, trader } = await getNamedAccounts();
        const factory = await deployments.get('MockAssistantFactory');

        const botFactory = await ethers.getContractAt('MockAssistantFactory', factory.address);

        const tokenId = 888;
        const botAddress = '0xd8058efe0198ae9dD7D563e1b4938Dcbc86A1F81';
        const tx = await botFactory.createAccount(tokenId, botAddress);
        tx.wait();

        const tba = await botFactory.getAccount(tokenId, botAddress);
        console.log(tba);
    });
});
