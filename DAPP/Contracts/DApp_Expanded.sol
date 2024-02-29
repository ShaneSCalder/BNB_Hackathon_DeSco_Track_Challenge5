// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract MonthlyContentToken is ERC20, Ownable, Pausable, AccessControl {
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    mapping(address => uint256) public subscriptions;

    constructor() ERC20("MonthlyContentToken", "MCT") {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(PAUSER_ROLE, msg.sender);
        _setupRole(MINTER_ROLE, msg.sender);
    }

    function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function setSubscription(address subscriber, uint256 endTimestamp) public onlyOwner {
        subscriptions[subscriber] = endTimestamp;
    }

    function hasActiveSubscription(address subscriber) public view returns (bool) {
        return subscriptions[subscriber] >= block.timestamp;
    }

    // Add event logging for subscription updates and token minting
    event SubscriptionUpdated(address subscriber, uint256 endTimestamp);
    event TokenMinted(address to, uint256 amount);

    // Override necessary functions to integrate Pausable
    function _beforeTokenTransfer(address from, address to, uint256 amount)
        internal
        override
        whenNotPaused
    {
        super._beforeTokenTransfer(from, to, amount);
    }
}
