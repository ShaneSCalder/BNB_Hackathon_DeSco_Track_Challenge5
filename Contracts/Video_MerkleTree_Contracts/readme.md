## Video_MerkleRoot.sol Contract Overview:

The "Video_MerkleRoot.sol" smart contract is crafted for the Binance Smart Chain to securely record Merkle roots of video content, providing a decentralized mechanism for verifying video data integrity. It allows content creators to leverage blockchain technology for immutable proof of authenticity, facilitating trustless content distribution and management. This contract serves as a foundational component in a broader ecosystem aimed at enhancing digital rights management and content monetization through blockchain.

# Video Merkle Root Toolkit

This repository contains a suite of tools designed for managing and verifying video content integrity using blockchain technology. It includes a Solidity contract (`Video_MerkleRoot.sol`), a Go program (`Video_MerkleTree`) for generating Merkle roots and proofs, and a Node.js script (`generateContractWithMerkleRoot.js`) for embedding Merkle roots into a Solidity contract.

## Components

- **`Video_MerkleRoot.sol`**: A Solidity contract for recording Merkle roots on the Binance Smart Chain, facilitating immutable proof of video content integrity.
- **`Video_MerkleTree` Go Program**: A program that generates a Merkle root and proof from video content, ensuring the content's integrity can be verified.
- **`generateContractWithMerkleRoot.js` Node.js Script**: A script that embeds a generated Merkle root into a new Solidity contract, automating the process of contract creation with specific Merkle roots.

## Prerequisites

- **Solidity Contract**: Ensure `Video_MerkleRoot.sol` is deployed on the Binance Smart Chain.
- **Go Environment**: Installed for running `Video_MerkleTree` to generate Merkle roots and proofs.
- **Node.js**: Installed for executing `generateContractWithMerkleRoot.js` to generate contracts with embedded Merkle roots.

## Usage

### Generating Merkle Root and Proof

1. **Run `Video_MerkleTree`**:
    - Navigate to the Go program directory.
    - Execute the program with your video file as input.
    - The program outputs `merkleroot.txt` containing the Merkle root and optionally a proof file.

### Embedding Merkle Root into Solidity Contract

1. **Prepare `merkleroot.txt`**:
    - Ensure the `merkleroot.txt` file is in the same directory as the `generateContractWithMerkleRoot.js` script.

2. **Run `generateContractWithMerkleRoot.js`**:
    - Execute the script to read the Merkle root from `merkleroot.txt`.
    - The script generates a new Solidity contract file with the Merkle root embedded.

### Deployment

- Deploy the generated Solidity contract to the Binance Smart Chain using your preferred development environment (e.g., Remix Truffle or Hardhat).

## Notes

- The toolkit streamlines the process of securing video content integrity on the blockchain.
- Ensure secure handling of private keys and sensitive information throughout the process.

## Troubleshooting

### General Issues

- **Solidity Deployment Issues**: Verify the contract code, ensure network connectivity, and check gas settings.
- **Go Program Errors**: Confirm the Go environment setup and validate the video file path.
- **Node.js Script Failures**: Check Node.js installation, ensure correct file paths, and verify the presence of `merkleroot.txt`.

### Specific Issues

- **Merkle Root Encoding in Contract**: If you encounter issues related to the Merkle root not being recognized or validated correctly by your contract, consider upgrading the contract to hash the Merkle root using `keccak256` for encoding. This ensures that the Merkle root is consistently handled as a `bytes32` type within the blockchain environment.

#### Upgrading Merkle Root Handling

If the direct assignment of Merkle roots in your contract (`Video_MerkleRoot.sol`) does not meet your needs or you require enhanced security, update the contract to hash the Merkle root using Solidity's `keccak256` function. This approach is beneficial for ensuring the integrity and uniform format of Merkle roots.

Replace the direct assignment of `predefinedMerkleRoot` with a hashed version:

```solidity
// Original assignment
bytes32 public predefinedMerkleRoot = /*MERKLE_ROOT_PLACEHOLDER*/;

// Upgraded approach using keccak256
bytes32 public predefinedMerkleRoot = keccak256(abi.encodePacked(/*MERKLE_ROOT_PLACEHOLDER*/));
```

When generating the Solidity contract with the `generateContractWithMerkleRoot.js` script, ensure the placeholder replacement accommodates this change:

```javascript
// Ensure the replacement line in the script reflects the hashing approach
const newContract = contractTemplate.replace('/*MERKLE_ROOT_PLACEHOLDER*/', `keccak256(abi.encodePacked("0x${merkleRoot}"))`);
```

**Note**: This modification in the JS script is illustrative. You might need to adjust the Solidity contract template and script logic to properly incorporate dynamic Merkle root hashing directly within the Solidity code.

### Why This Matters

Hashing the Merkle root with `keccak256` ensures that all inputs are uniformly treated and securely encoded on the blockchain. This method provides an additional layer of integrity verification and standardization, especially important when dealing with diverse or dynamic content sources.


# NTF Video Access Contract on Binance Smart Chain

The `NTFVideoAccess` smart contract utilizes the Binance Smart Chain (BSC) to mint NFTs that grant access to exclusive video content. Each NFT is uniquely associated with a Merkle root for content verification and an access code for access control. This document outlines the deployment process, NFT minting procedure, and metadata setup.

## Contract Features

- **Merkle Root Association**: Securely link each NFT to video content via a unique Merkle root.
- **Access Code Integration**: Embed a unique access code within each NFT for content access.
- **ERC721 Compatibility**: Ensures interoperability with BSC wallets and marketplaces.

## Setup and Deployment

### Requirements

- Node.js and npm installed.
- A BSC wallet with BNB for transaction fees.
- [Truffle](https://www.trufflesuite.com/) or [Hardhat](https://hardhat.org/) for smart contract deployment (optional).

### Deploying the Contract

1. Compile the contract using your preferred development environment.
2. Deploy the contract to the Binance Smart Chain. Ensure you have enough BNB in your wallet to cover the deployment gas fees.

## Minting NFTs with `createNFT.js`

The `createNFT.js` script facilitates minting NFTs by interacting with the deployed `NTFVideoAccess` contract. Before running the script, ensure you have the contract's ABI and address.

### Preparing `createNFT.js`

- The script requires Node.js, Web3.js, and your contract's ABI.
- Update the script with your contract's address, ABI, and the wallet private key for transactions.

### Running `createNFT.js`

```bash
node createNFT.js <recipient_address> <tokenURI> <merkleRoot> <accessCode>
```

- **`<recipient_address>`**: Address to receive the NFT.
- **`<tokenURI>`**: URI pointing to the NFT's metadata.
- **`<merkleRoot>`**: Unique Merkle root associated with the video content.
- **`<accessCode>`**: Access code for video content.

## Setting Up Metadata

Each NFT requires metadata hosted on IPFS, including details like name, description, and an image. Here's an example of the metadata format:

```json
{
  "name": "Exclusive Video Content",
  "description": "This NFT grants access to unique video content.",
  "image": "ipfs://<Your_IPFS_Image_Hash_Here>",
  "attributes": [
    {
      "trait_type": "Access Code",
      "value": "ACCESSCODE1234"
    },
    {
      "trait_type": "Merkle Root",
      "value": "0x123...abc"
    }
  ]
}
```

### Uploading Metadata to IPFS

1. Create a JSON file with the above structure.
2. Upload the JSON file to IPFS using a service like [Pinata](https://pinata.cloud/) or the IPFS desktop application.
3. Use the resulting IPFS URI as the `tokenURI` when minting your NFT.

## Conclusion

The `NTFVideoAccess` contract on BSC, combined with the `createNFT.js` script and properly set up metadata, provides a robust system for issuing NFTs that grant access to exclusive video content. By leveraging BSC's efficient and cost-effective network, this approach ensures a scalable and user-friendly experience for content creators and consumers alike.
