// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import './MockAssistant.sol';

contract MockAssistantFactory {
    address public registry;
    address public accountImpl;
    // the salt used to generate the TBA
    uint8 salt = 0;
    address public fmToken;
    // store all NFT contracts which are created by factory
    mapping(uint => address) public assistantsMap;
    // Assistant[] public assistants;
    // map assistant contract address to maker
    mapping(address => address) public assistantMaker;

    event NewAssistantCreated(address indexed maker, address indexed astBot, uint256 indexed nftId);

    constructor(address _fmToken, address _accountImpl, address _registry) {
        fmToken = _fmToken;
        accountImpl = _accountImpl;
        registry = _registry;
    }

    function createAssistantMock(
        string memory _name,
        string memory _symbol,
        string memory _baseURI,
        uint256 _maxSupply,
        uint256 _mintPrice,
        uint256 _collectionId
    ) external {
        Assistant astBot = new MockAssistant(_name, _symbol, _baseURI, _maxSupply, _mintPrice, address(this), _collectionId);
        require(assistantsMap[_collectionId] == address(0), 'nft id exists');
        assistantsMap[_collectionId] = address(astBot);
        assistantMaker[address(astBot)] = msg.sender;
        emit NewAssistantCreated(msg.sender, address(astBot), _collectionId);
    }

    function createAccount(uint _tokenId) public returns (address) {
        return IERC6551Registry(registry).createAccount(accountImpl, block.chainid, fmToken, _tokenId, salt, '');
    }

    function getAccount(uint _tokenId) public view returns (address) {
        return IERC6551Registry(registry).account(accountImpl, block.chainid, fmToken, _tokenId, salt);
    }
}
