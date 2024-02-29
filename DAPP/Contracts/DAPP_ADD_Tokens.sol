// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract DAppToken is ERC20, Ownable {
    constructor() ERC20("DAPPContentToken", "DCT") {
        // Mint the fixed supply of 1,000,000 tokens to the deployer's address
        _mint(msg.sender, 1000000 * 10 ** uint(decimals()));
    }

    // Optional: Implement any additional functionality you need for token distribution, etc.
}