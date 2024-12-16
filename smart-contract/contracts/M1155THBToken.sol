// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract M1155THBToken is ERC1155 {

    uint256 public constant PHASE_1_MTHB = 0;
    uint256 public constant PHASE_2_MTHB = 1;
    uint256 public constant PHASE_3_MTHB = 2;
    uint256 public constant PHASE_4_MTHB = 3;
    uint256 public nPhase = 4;

    address public owner;

    constructor() ERC1155("https://mock.example/api/phase/{id}.json") {
        _mint(msg.sender, PHASE_1_MTHB, (10 * (10**18)), "");
        _mint(msg.sender, PHASE_2_MTHB, (20 * (10**18)), "");
        _mint(msg.sender, PHASE_3_MTHB, (30 * (10**18)), "");
        _mint(msg.sender, PHASE_4_MTHB, (40 * (10**18)), "");
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    function phaseMint(address _to, uint256 _amount) public onlyOwner {
        _mint(_to, nPhase, (_amount ** 18), "");
    }

}
