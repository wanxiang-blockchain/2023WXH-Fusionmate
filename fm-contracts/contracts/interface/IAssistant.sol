// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import '@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol';
import '@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol';

interface IAssistant is IERC721Enumerable, IERC721Metadata {
    // Assistant function
    function baseURI() external view returns (string memory);

    function maxSupply() external returns (uint256);

    function mintPrice() external returns (uint256);

    // get token bound account address of the tokenId
    function tokenBoundAccountsList(uint256 tokenId) external returns (address);

    function mint() external returns (uint256);

    function harvestForTBA(bytes memory signature, uint256 tokenId, uint256 amount, uint256 collectionId) external;
}
