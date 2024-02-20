const Web3 = require('web3');
const fs = require('fs');
require('dotenv').config();

// Configuration: Set these variables
const contractABI = JSON.parse(fs.readFileSync('./NTF_VideoAccess_ABI.json', 'utf-8'));
const contractAddress = 'YOUR_CONTRACT_ADDRESS_HERE';
const providerURL = process.env.PROVIDER_URL; // e.g., Infura URL
const privateKey = process.env.PRIVATE_KEY; // Owner's private key for minting
const accountAddress = process.env.ACCOUNT_ADDRESS; // Owner's account address

// Initialize web3 instance
const web3 = new Web3(providerURL);

// Contract instance
const nftContract = new web3.eth.Contract(contractABI, contractAddress);

// Read Merkle root and access code from files
const merkleRoot = fs.readFileSync('./merkleroot.txt', 'utf-8').trim();
const accessCode = fs.readFileSync('./accesscode.txt', 'utf-8').trim();

// Recipient address and token URI (set these as needed)
const recipientAddress = 'RECIPIENT_WALLET_ADDRESS';
const tokenURI = 'ipfs://TOKEN_METADATA_URI';

async function mintNFT() {
    const nonce = await web3.eth.getTransactionCount(accountAddress, 'latest'); // Get nonce
    const tx = {
        from: accountAddress,
        to: contractAddress,
        nonce: nonce,
        gas: 500000,
        data: nftContract.methods.mintNFT(recipientAddress, tokenURI, web3.utils.asciiToHex(merkleRoot), accessCode).encodeABI(),
    };

    // Sign the transaction
    const signedTx = await web3.eth.accounts.signTransaction(tx, privateKey);

    // Send the transaction
    const receipt = await web3.eth.sendSignedTransaction(signedTx.rawTransaction);

    console.log('NFT minted successfully:', receipt.transactionHash);
}

// Execute the minting process
mintNFT().catch(console.error);
