// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import '@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol';
import '@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol';
import '@openzeppelin/contracts/utils/Counters.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';
import './interface/IAssistantFactory.sol';
import './interface/IERC6551Registry.sol';
import './interface/IFMToken.sol';

// a single AI assistant NFT instance
contract Assistant is ERC721Enumerable, ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private tokenId;

    // A NFT class can mint maximuim nft
    uint256 public maxSupply;
    uint256 public mintPrice;
    uint256 public collectionId;
    address public factory;
    string public baseURI;

    event TokenBoundAccountCreated(uint256 indexed tokenId, address indexed tba);

    // tokenId -> tba
    mapping(uint256 => address) public tokenBoundAccountsList;

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _baseURI,
        uint256 _maxSupply,
        uint256 _mintPrice,
        address _factory,
        uint256 _collectionId
    ) ERC721(_name, _symbol) {
        _setBaseURI(_baseURI);
        maxSupply = _maxSupply;
        mintPrice = _mintPrice;
        factory = _factory;
        collectionId = _collectionId;
    }

    function _setBaseURI(string memory _baseURI) internal {
        baseURI = _baseURI;
    }

    function getFMToken() public returns (IERC20) {
        address token = IAssistantFactory(factory).fmToken();
        return IERC20(token);
    }

    /// @dev any EOA user can invoke mint if this user transfer it's fmToken to factory contract
    /// fmToken locked in factory will reward maker
    function mint() public returns (uint256) {
        rewardMaker();
        uint256 newestTokenId = generateNFT();

        address tba = IAssistantFactory(factory).createAccount(newestTokenId, address(this));
        tokenBoundAccountsList[newestTokenId] = tba;
        emit TokenBoundAccountCreated(newestTokenId, tba);
        return newestTokenId;
    }

    function generateNFT() internal returns (uint256) {
        // generate nft for msg.sender
        uint256 newestTokenId = tokenId.current();
        require(newestTokenId < maxSupply, 'nft tokenId is execeed maxSupply');
        _safeMint(msg.sender, newestTokenId);
        // because we use ERC721URIStorage we must set tokenURI = tokenURI(newestTokenId)
        _setTokenURI(newestTokenId, tokenURI(newestTokenId));
        tokenId.increment();
        return newestTokenId;
    }

    function rewardMaker() internal {
        IERC20 fmToken = getFMToken();
        uint256 fmBalance = fmToken.balanceOf(msg.sender);
        require(fmBalance >= mintPrice, 'fm token balance is insufficient');
        // approve maker to get msg.sender fm token
        address maker = IAssistantFactory(factory).assistantMaker(address(this));
        require(fmToken.allowance(msg.sender, address(this)) >= mintPrice, 'allowance for maker is insufficient');
        fmToken.transferFrom(msg.sender, maker, mintPrice);
    }

    // tokenId -> TBA map
    function harvestForTBA(bytes memory _signature, uint256 _tokenId, uint256 _amount, uint256 _collectionId) public {
        // verify backend signature, signature msg should include FMToken amount, owner
        bool sigRes = IAssistantFactory(factory).harvestVerify(_signature, _amount, _collectionId, _tokenId);
        require(sigRes, 'backend signature is invalid');
        address owner = ownerOf(_tokenId);
        require(owner == msg.sender, 'not the owner');
        address tba = tokenBoundAccountsList[_tokenId];
        IFMToken fmToken = IFMToken(IAssistantFactory(factory).fmToken());
        fmToken.mintForTBA(tba, _amount);
    }

    // ============ OVERRIDE FUNCTIONS ============
    function tokenURI(uint256 _tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        _requireMinted(_tokenId);
        return string(abi.encodePacked(baseURI, Strings.toString(_tokenId)));
    }

    function _burn(uint256 _tokenId) internal override(ERC721, ERC721URIStorage) {
        super._burn(_tokenId);
    }

    // TODO: set _batchSize = 1
    function _beforeTokenTransfer(address _from, address _to, uint256 _tokenId, uint256 _batchSize) internal override(ERC721Enumerable, ERC721) {
        super._beforeTokenTransfer(_from, _to, _tokenId, _batchSize);
    }

    function supportsInterface(bytes4 interfaceId) public view override(ERC721Enumerable, ERC721URIStorage) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
