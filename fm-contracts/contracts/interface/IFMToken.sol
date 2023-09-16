// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

// todo: IFMToken is extend IERC20
interface IFMToken {
    function mintForTBA(address tab, uint256 amount) external;
}
