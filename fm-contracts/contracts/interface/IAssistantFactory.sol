// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface IAssistantFactory {
    // get backend address which will be used to verify signature
    function backend() external returns (address);

    // get test scores contract address
    function fmToken() external returns (address);

    // get a nft contract address by index
    function assistants(uint256 index) external returns (address);

    // obtain the maker address of assistant instance
    function assistantMaker(address assistant) external returns (address);

    function harvestVerify(bytes memory signature, uint256 amount, uint256 collectionId, uint tokenId) external returns (bool);

    //======================the following functions are ERC6551 ========================
    function registry() external returns (address);

    function accountImpl() external returns (address);

    function createAccount(uint256 tokenId, address tokenContract) external returns (address);

    function getAccount(uint256 tokenId, address tokenContract) external returns (address);
}
