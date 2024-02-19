// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract NTFVideoAccess is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    // Mapping from token ID to Merkle root
    mapping(uint256 => bytes32) private _merkleRoots;
    // Mapping from token ID to unique access code
    mapping(uint256 => string) private _accessCodes;

    constructor() ERC721("NTFVideoAccess", "NVA") {}

    // Function to mint NFT with dynamic Merkle root and access code
    function mintNFT(
        address recipient,
        string memory tokenURI,
        bytes32 merkleRoot,
        string memory accessCode
    ) public onlyOwner {
        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();

        _mint(recipient, newItemId); // Mint the NFT to the specified recipient
        _setTokenURI(newItemId, tokenURI); // Set the token URI (metadata)
        _merkleRoots[newItemId] = merkleRoot; // Associate the Merkle root with the new token ID
        _accessCodes[newItemId] = accessCode; // Associate the access code with the new token ID

        // Additional logic or events can be added here
    }

    // Function to get the Merkle root and access code for a token
    function getContentData(uint256 tokenId) public view returns (bytes32, string memory) {
        require(_exists(tokenId), "NFT does not exist.");
        return (_merkleRoots[tokenId], _accessCodes[tokenId]);
    }
}
