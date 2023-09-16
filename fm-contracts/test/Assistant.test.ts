import { expect } from 'chai';

import { ethers, deployments, getNamedAccounts } from 'hardhat';

describe('Assistant contract', function () {
    it('assistant should work', async function () {
        await deployments.fixture(['MockAssistantFactory']);
        const { deployer, trader } = await getNamedAccounts();
        const factory = await deployments.get('MockAssistantFactory');

        const botFactory = await ethers.getContractAt('MockAssistantFactory', factory.address);
        console.log('MockAssistantFactory contract address is', botFactory.address);

        // create two assistant
        const mintPrice = ethers.utils.parseUnits('100');
        const tx1 = await botFactory.createAssistantMock('bot1', 'AI1', 'https://fusionMate.io', 10, mintPrice, 11);
        tx1.wait();
        const traderSign = await ethers.getSigner(trader);
        const tx2 = await botFactory.connect(traderSign).createAssistantMock('bot2', 'AI2', 'https://fusionMate.io', 5, mintPrice, 222);
        tx2.wait();

        // obtain assistant contract address according to collection id
        const bot1Address = await botFactory.assistantsMap(11);
        const bot2Address = await botFactory.assistantsMap(222);
        // console.log('bot1', bot1Address);
        // console.log('bot2', bot2Address);
        const maker1 = await botFactory.assistantMaker(bot1Address); // deployer address
        const maker2 = await botFactory.assistantMaker(bot2Address); // trader address

        // get Assistant contract and invoke
        const assistant1 = await ethers.getContractAt('MockAssistant', bot1Address);
        console.log('MockAssistant1 contract address is', assistant1.address);
        const assistant2 = await ethers.getContractAt('MockAssistant', bot2Address);
        console.log('MockAssistant2 contract address is', assistant2.address);

        // for assistant1 function
        expect(await assistant1.maxSupply()).to.equal(10);
        expect(await assistant1.mintPrice()).to.equal(mintPrice);
        expect(await assistant1.collectionId()).to.equal(11);
        expect(await assistant1.factory()).to.equal(factory.address);
        expect(await assistant1.baseURI()).to.equal('https://fusionMate.io');

        // for assistant2 function
        expect(await assistant2.maxSupply()).to.equal(5);
        expect(await assistant2.mintPrice()).to.equal(mintPrice);
        expect(await assistant2.collectionId()).to.equal(222);
        expect(await assistant2.factory()).to.equal(factory.address);
        expect(await assistant2.baseURI()).to.equal('https://fusionMate.io');

        // get FMToken instance by factory
        const fmTokenAddr = await botFactory.fmToken();
        const fmToken = await ethers.getContractAt('FMToken', fmTokenAddr);

        // make sure deployer and trader has enough fmToken to mint nft
        let tx3 = await fmToken.mint();
        tx3.wait();
        let tx4 = await fmToken.connect(traderSign).mint();
        tx4.wait();
        expect(await fmToken.balanceOf(deployer)).to.equal(ethers.utils.parseUnits('1000'));
        expect(await fmToken.balanceOf(trader)).to.equal(ethers.utils.parseUnits('1000'));

        // before deployer and trader invoke mint()
        // maker1 = deployer, maker2 = trader;
        expect(await fmToken.balanceOf(maker1)).to.equal(await fmToken.balanceOf(deployer));
        expect(await fmToken.balanceOf(maker2)).to.equal(await fmToken.balanceOf(trader));

        // trader approve assistant1 to transfer token
        // trader mint: transfer fmToken to assistant1 contract -> maker1 address
        let tx5 = await fmToken.connect(traderSign).approve(bot1Address, mintPrice.mul(5));
        tx5.wait();
        expect(await fmToken.allowance(trader, bot1Address)).to.equal(mintPrice.mul(5));

        // trader invoke assistant1 rewardMaker(), namely trader transfer mintPrice to maker1/deployer
        let tx6 = await assistant1.connect(traderSign).mockRewardMaker();
        tx6.wait();
        // trader: 1000 - minPrice(100), maker1 1000 + mintPrice
        expect(await fmToken.balanceOf(trader)).to.equal(ethers.utils.parseUnits('900'));
        expect(await fmToken.balanceOf(maker1)).to.equal(ethers.utils.parseUnits('1100'));

        // trader generate 2 NFT
        let tx7 = await assistant1.connect(traderSign).mockGenerateNFT();
        tx7.wait();

        let tx8 = await assistant1.connect(traderSign).mockGenerateNFT();
        tx8.wait();

        // balanceOf(trader) = 2
        expect(await assistant1.balanceOf(trader)).to.equal(2);
        // ownerOf(0) = ownerOf(1) = trader
        expect(await assistant1.ownerOf(0)).to.equal(trader);
        expect(await assistant1.ownerOf(1)).to.equal(trader);
        // name = 'bot1'
        expect(await assistant1.name()).to.equal('bot1');
        // symbol = 'AI1'
        expect(await assistant1.symbol()).to.equal('AI1');
        // tokenURI(0) =https://fusionMate.io0
        // tokenURI(1) =https://fusionMate.io1
        expect(await assistant1.tokenURI(1)).to.equal('https://fusionMate.io1');
        // tokenOfOwnerByIndex(owner, index) = tokenId
        expect(await assistant1.tokenOfOwnerByIndex(trader, 1)).to.equal(1);
        // totalSupply()
        expect(await assistant1.totalSupply()).to.equal(2);
        // tokenByIndex(index) = tokenId
        expect(await assistant1.tokenByIndex(0)).to.equal(0);

        // deployer invoke mint function for assistant1
        // deployer must approve assistant first
        let tx9 = await fmToken.approve(bot1Address, mintPrice.mul(5));
        tx9.wait();
        expect(await fmToken.allowance(deployer, bot1Address)).to.equal(mintPrice.mul(5));
        // tokenId = 2
        let tx10 = await assistant1.mint();
        tx10.wait();

        // deployer = maker1, balance is not change
        expect(await fmToken.balanceOf(deployer)).to.equal(ethers.utils.parseUnits('1100'));
        expect(await assistant1.ownerOf(2)).to.equal(deployer);

        const tba = await botFactory.getAccount(2, bot1Address);
        expect(await assistant1.tokenBoundAccountsList(2)).to.equal(tba);

        // test mockHavestForTBA
        // tokenId = 2 is deployer's nft
        const amount = ethers.utils.parseUnits('99');
        let tx11 = await assistant1.mockHavestForTBA(2, amount);
        tx11.wait();
        expect(await fmToken.balanceOf(tba)).to.equal(amount);
    });
});
