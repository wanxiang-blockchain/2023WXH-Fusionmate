// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';

contract FMToken is ERC20 {
    uint256 perAmount = 1000 * (10 ** decimals());

    constructor() ERC20('FusionMate Token', 'FM') {}

    function mint() external {
        _mint(msg.sender, perAmount);
    }

    function mintForTBA(address tba, uint256 amount) public {
        _mint(tba, amount);
    }
}
