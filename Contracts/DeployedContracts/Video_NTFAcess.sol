// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract MerkleMetaNFT is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    mapping(uint256 => bytes32) private _merkleRoots; // Mapping from token ID to Merkle root
    mapping(uint256 => string) private _accessCodes;  // Mapping from token ID to unique access code

    constructor() ERC721("MerkleMeta", "MMT") Ownable(msg.sender) {}  // Pass msg.sender as the initialOwner{}

    // Function to mint a test NFT with hardcoded data
    function mintTestNFT() public onlyOwner {
        address recipient = 0xF023E5ddddD79c247240372fd2A8D0600dE7433d; // Specified wallet address
        string memory tokenURI = "ipfs://example_uri_for_test_nft_metadata";
        bytes32 merkleRoot = keccak256(abi.encodePacked("0x123abc456def7890001112233445566778899aabbccddeeff0011223344556677")); // Fake Merkle root // Fake Merkle root
        string memory accessCode = "TEST-ACCESS-CODE-1234"; // Fake access code

        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();

        _mint(recipient, newItemId); // Mint the NFT to the specified recipient
        _setTokenURI(newItemId, tokenURI); // Set the token URI (metadata)
        _merkleRoots[newItemId] = merkleRoot; // Associate the Merkle root with the new token ID
        _accessCodes[newItemId] = accessCode; // Associate the access code with the new token ID
    }

    // Function to get the Merkle root and access code for a token
    function getContentData(uint256 tokenId) public view returns (bytes32, string memory) {
        require(ownerOf(tokenId) != address(0), "NFT does not exist.");
        return (_merkleRoots[tokenId], _accessCodes[tokenId]);
    }
}
