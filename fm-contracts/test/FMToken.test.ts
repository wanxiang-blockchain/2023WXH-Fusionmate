import { expect } from 'chai';

import { ethers, deployments, getNamedAccounts } from 'hardhat';

describe('FMToken contract', function () {
    it('Deployment should assign the total supply of tokens to the owner', async function () {
        await deployments.fixture(['FMToken']);
        const { deployer, trader } = await getNamedAccounts();
        const Token = await deployments.get('FMToken');
        const fm_token = await ethers.getContractAt('FMToken', Token.address);
        console.log('FMToken contract address is', Token.address);
        const beforeBalance = await fm_token.balanceOf(deployer);
        expect(beforeBalance).to.equal(0);
        // console.log('before mint balance of deployer is', beforeBalance.toString());
        const tx1 = await fm_token.mint();
        tx1.wait();
        const afterBalance = await fm_token.balanceOf(deployer);
        expect(afterBalance).to.equal(ethers.utils.parseUnits('1000'));
        // console.log('after mint balance of deployer is', afterBalance.toString());

        const traderSign = await ethers.getSigner(trader);
        const tx2 = await fm_token.connect(traderSign).mintForTBA(trader, ethers.utils.parseUnits('10000'));
        tx2.wait();
        const traderBalance = await fm_token.balanceOf(trader);
        expect(traderBalance).to.equal(ethers.utils.parseUnits('10000'));
        // console.log('trader balance is ', traderBalance.toString());
        const supply = await fm_token.totalSupply();
        // console.log('totalSupply is', supply.toString());
        expect(supply).to.equal(afterBalance.add(traderBalance));
    });
});
