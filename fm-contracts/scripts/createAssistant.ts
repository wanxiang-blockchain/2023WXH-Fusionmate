import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';

async function mockCreateAssistant() {
    // mockAssistantFactory deployed on goerli
    const mockFactory = '0xe68f06593d0cf94193144b451cdb55d9c0011cfb';
    const mockAssistantFactory = await ethers.getContractAt('MockAssistantFactory', mockFactory);
    // name, symbol, baseURI, maxSupply, mintPrice, collectionId
    const mintPrice = ethers.utils.parseUnits('100');
    let tx = await mockAssistantFactory.createAssistantMock('bot2', 'AI2', 'https://fusionMate.io', 10, mintPrice, 888, {
        gasLimit: BigNumber.from(4000000),
    });
    tx.wait();
    console.log(tx.hash);
}

mockCreateAssistant();
