// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import './../Assistant.sol';

contract MockAssistant is Assistant {
    constructor(
        string memory _name,
        string memory _symbol,
        string memory _baseURI,
        uint256 _maxSupply,
        uint256 _mintPrice,
        address _factory,
        uint256 _collectionId
    ) Assistant(_name, _symbol, _baseURI, _maxSupply, _mintPrice, _factory, _collectionId) {}

    function mockGenerateNFT() public returns (uint256) {
        return generateNFT();
    }

    function mockRewardMaker() public {
        rewardMaker();
    }

    function mockHavestForTBA(uint256 _tokenId, uint256 _amount) public {
        address owner = ownerOf(_tokenId);
        require(owner == msg.sender, 'not the owner');
        address tba = tokenBoundAccountsList[_tokenId];
        IFMToken fmToken = IFMToken(IAssistantFactory(factory).fmToken());
        fmToken.mintForTBA(tba, _amount);
    }
}
