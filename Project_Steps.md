# Blockchain-Based Subscription with KYC Verification

This guide outlines the steps for integrating Know Your Customer (KYC) verification with a token-based subscription service on a blockchain using Solidy, Goland and Javascript. It ensures that subscriptions can only be purchased with DApp tokens by users who have completed KYC verification.

## Overview

The process involves collecting KYC data and generating a Merkle tree off-chain, submitting the Merkle root to the blockchain, purchasing subscriptions with DApp tokens contingent on KYC verification, and verifying KYC using Merkle proofs to grant subscription access.

## Requirements

- Wallet (e.g., MetaMask)
- DApp tokens for subscription purchase - you can utalize contracts provided
- Web3 provider and development tools (e.g., Truffle, Hardhat, remix)
- Access to a backend server for off-chain KYC processing - merkle root go programs provided - these can be complied for implentation

## Implementation Steps

### Step 1: KYC Data Collection and Merkle Tree Generation

1. **Collect KYC Data:** Users submit their KYC information through a secure platform interface.
2. **Verify KYC Data:** The service provider verifies the submitted KYC data off-chain. 
3. **Generate Merkle Tree:** After verification, a Merkle tree is generated from the KYC data. The Merkle root represents the verified information, with individual data points hashed to form the leaves.
   - go programs can be adapted to provide different data outputs (txt, JSON etc.)

### Step 2: Merkle Root Submission - initiated prior to purchase

1. **Inform User:** Once KYC is verified, inform the user of their unique Merkle root.
2. **Connect Wallet:** Users connect their wallet to the platform.
3. **Submit Merkle Root:** Users submit their Merkle root to the KYCVerification smart contract via a transaction, linking their wallet to their KYC status.

### Step 3: Subscription Purchase

1. **Token Acquisition:** Ensure users have enough DApp tokens in their wallet for the subscription. i.e. 1 token = 1 month subscription. 
2. **Initiate Purchase:** Users initiate a subscription purchase through the platform. tokens would be used for a month at a time - also allowing user to purchase when convient and avaliable to watch. 
3. **Provide Merkle Proof:** As part of the purchase process, users must provide a Merkle proof that validates their KYC data against the stored Merkle root. Security verification check. 
4. **Verify and Process:** The subscription contract verifies the Merkle proof with the KYCVerification contract. Upon successful verification, the contract processes the token payment and activates the subscription.

### Step 4: Accessing Subscribed Services

1. **Service Access:** Users access the subscribed services through the platform.
2. **Verification Check:** The platform checks the subscription status and KYC verification before granting access.

## Security Considerations

- Ensure secure handling and storage of KYC data off-chain.
- Conduct thorough smart contract audits to prevent vulnerabilities.
- Implement secure web3 interactions for transactions and data submissions.

## User UI / UX improvements
- As this is a POC and not a full product we will need to provide easier integration points to reduce steps for the users
- As we are not developing data wallets and platforms at this time there is a lot of work to be completed to ensure a quick and easy solution for future users.

## Developers / Content creators
- the intent is to build out a full solution over time with testing, feedback and interaction
- comments and suggestions are appreciated
- back end merkle trees can be secured in blocks and then encrypted - we will be working on solutions for encryption i.e. symmetric, asymmetric, lattice etc.
- Kyc encryption options future modules. 

## Future builds
- data wallet - KYC & content management for users
- Merkle Tree - for large data sets, videos and other content
- AML and audit tools

