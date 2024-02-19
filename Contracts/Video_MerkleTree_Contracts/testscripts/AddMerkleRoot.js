const fs = require('fs');

// Paths
const merkleRootPath = './merkleroot.txt'; // Path to the .txt file containing the Merkle root
const templatePath = './ContractTemplate.sol'; // Path to your Solidity contract template
const outputPath = './GeneratedContract.sol'; // Desired path for the new Solidity file

// Read the Merkle root from the .txt file
const merkleRoot = fs.readFileSync(merkleRootPath, 'utf-8').trim();

// Read the template Solidity contract
const contractTemplate = fs.readFileSync(templatePath, 'utf-8');

// Replace a placeholder in the contract template with the actual Merkle root
const newContract = contractTemplate.replace('/*MERKLE_ROOT_PLACEHOLDER*/', `'0x${merkleRoot}'`);

// Write the new Solidity contract to a file
fs.writeFileSync(outputPath, newContract);

console.log(`Generated contract with Merkle root embedded at ${outputPath}`);
