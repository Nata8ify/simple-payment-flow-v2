// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MTHBToken is ERC20 {

    constructor(uint256 initialSupply) public ERC20("Mock Thai Baht", "MTHB") {
        _mint(msg.sender, initialSupply);
    }

}
