// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract MerkleRootManager is Ownable {
    mapping(uint256 => bytes32) private _merkleRoots;

    event MerkleRootAdded(uint256 indexed itemId, bytes32 indexed merkleRoot);

    // Constructor is simplified to rely on Ownable's constructor for owner setup.
    constructor() Ownable() {}

    // Function to add a Merkle root for a specific item.
    function addMerkleRoot(uint256 itemId, bytes32 merkleRoot) public onlyOwner {
        require(_merkleRoots[itemId] == bytes32(0), "Item already has a Merkle root.");
        _merkleRoots[itemId] = merkleRoot;
        emit MerkleRootAdded(itemId, merkleRoot);
    }

    // Function to retrieve the Merkle root of an item.
    function getMerkleRoot(uint256 itemId) public view returns (bytes32) {
        require(_merkleRoots[itemId] != bytes32(0), "Merkle root does not exist.");
        return _merkleRoots[itemId];
    }
}
