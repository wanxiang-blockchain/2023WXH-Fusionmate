import { ethers } from 'hardhat';

async function setBackend() {
    // AssistantFactory deployed on goerli
    const assistantFactory = '0x0D818DC92584C1dd7c585E4065800A226af1109A';
    const factory = await ethers.getContractAt('AssistantFactory', assistantFactory);

    // todo: must modify backend
    const backend = '0x0D818DC92584C1dd7c585E4065800A226af1109A';
    let tx = await factory.setBackend(backend);
    tx.wait();
    console.log(tx.hash);
}

setBackend();
