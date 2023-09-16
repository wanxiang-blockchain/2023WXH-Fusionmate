// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import './Assistant.sol';
import './interface/IERC6551Registry.sol';
import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';
import '@openzeppelin/contracts/utils/cryptography/EIP712.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

// create AI assistant factory
contract AssistantFactory is EIP712, Ownable {
    // todo: backend address for verify signature
    address public backend = 0x6F881627057b37B12163118e09F7e7901096Ccea;
    address public immutable registry = 0x02101dfB77FDE026414827Fdc604ddAF224F0921;
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

    constructor(address _fmToken, address _accountImpl) EIP712('FusionMate', '1') {
        fmToken = _fmToken;
        accountImpl = _accountImpl;
    }

    // todo: mint price should have limits
    function createAssistant(
        bytes memory _signature,
        string memory _name,
        string memory _symbol,
        string memory _baseURI,
        uint256 _collectionId,
        uint256 _maxSupply,
        uint256 _mintPrice
    ) external {
        bool sigRes = verify(_signature, _name, _symbol, _baseURI, msg.sender, _collectionId, _maxSupply, _mintPrice);
        require(sigRes, 'invalid backend signature');
        Assistant astBot = new Assistant(_name, _symbol, _baseURI, _maxSupply, _mintPrice, address(this), _collectionId);
        require(assistantsMap[_collectionId] == address(0), 'nft id exists');
        assistantsMap[_collectionId] = address(astBot);
        assistantMaker[address(astBot)] = msg.sender;
        emit NewAssistantCreated(msg.sender, address(astBot), _collectionId);
    }

    // todo: digest may be inclued in signature?
    // verify EIP712 signature from backend
    function verify(
        bytes memory _signature,
        string memory _name,
        string memory _symbol,
        string memory _baseURI,
        address makerAddress,
        uint256 collectionId,
        uint256 maxSupply,
        uint256 mintPrice
    ) public view returns (bool) {
        bytes32 digest = _hashTypedDataV4(
            keccak256(
                abi.encode(
                    keccak256(
                        'FusionMateNFTContractCreation(string name,string symbol,string baseURI,address makerAddress,uint256 collectionId,uint256 maxSupply,uint256 mintPrice)'
                    ),
                    keccak256(bytes(_name)),
                    keccak256(bytes(_symbol)),
                    keccak256(bytes(_baseURI)),
                    makerAddress,
                    collectionId,
                    maxSupply,
                    mintPrice
                )
            )
        );
        address recoveredAddress = ECDSA.recover(digest, _signature);
        return recoveredAddress == backend;
    }

    // @title Verify signature: Compare recovered address with backend address
    function harvestVerify(bytes memory _signature, uint256 _amount, uint256 _collectionId, uint256 _tokenId) public view returns (bool) {
        bytes32 digest = _hashTypedDataV4(
            keccak256(
                abi.encode(keccak256('FusionMateHarvest(uint256 amount,uint256 collectionId,uint256 tokenId)'), _amount, _collectionId, _tokenId)
            )
        );
        address recoveredAddress = ECDSA.recover(digest, _signature);
        return recoveredAddress == backend;
    }

    function createAccount(uint _tokenId, address _tokenContract) public returns (address) {
        return IERC6551Registry(registry).createAccount(accountImpl, block.chainid, _tokenContract, _tokenId, salt, '');
    }

    function getAccount(uint _tokenId, address _tokenContract) public view returns (address) {
        return IERC6551Registry(registry).account(accountImpl, block.chainid, _tokenContract, _tokenId, salt);
    }

    function setBackend(address _backend) public onlyOwner {
        backend = _backend;
    }
}
