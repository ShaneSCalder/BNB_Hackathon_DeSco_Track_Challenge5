package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Utility function to hash data using SHA-256
func hash(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

// Function to recursively build the Merkle tree and return the root hash
// Additionally, it captures the hashes at each level of the tree
func buildMerkleTree(leaves [][]byte, tree *[][]string) []byte {
	if len(leaves) == 1 {
		return leaves[0]
	}

	var parentLeaves [][]byte
	var currentLevel []string
	for i := 0; i < len(leaves); i += 2 {
		var combinedHash []byte
		if i+1 < len(leaves) {
			combinedHash = hash(append(leaves[i], leaves[i+1]...))
		} else {
			// For an odd number of leaves, duplicate the last leaf
			combinedHash = hash(append(leaves[i], leaves[i]...))
		}
		parentLeaves = append(parentLeaves, combinedHash)
		currentLevel = append(currentLevel, hex.EncodeToString(combinedHash))
	}
	*tree = append(*tree, currentLevel) // Append the current level of hashes to the tree
	return buildMerkleTree(parentLeaves, tree)
}

func main() {
	fileContent, err := ioutil.ReadFile("data/kycData.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var kycData map[string]string
	if err := json.Unmarshal(fileContent, &kycData); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	var leaves [][]byte
	for _, data := range kycData {
		leaves = append(leaves, hash([]byte(data)))
	}

	var merkleTree [][]string // To capture the entire Merkle tree
	merkleRoot := buildMerkleTree(leaves, &merkleTree)

	// Output the Merkle root
	err = ioutil.WriteFile("data/merkleroot.txt", []byte(hex.EncodeToString(merkleRoot)), 0644)
	if err != nil {
		fmt.Println("Error writing Merkle root to file:", err)
		return
	}

	// Output the entire Merkle tree
	treeJSON, err := json.MarshalIndent(merkleTree, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling Merkle tree to JSON:", err)
		return
	}
	err = ioutil.WriteFile("data/merkletree.json", treeJSON, 0644)
	if err != nil {
		fmt.Println("Error writing Merkle tree to file:", err)
		return
	}

	fmt.Println("Merkle root and tree written to merkleroot.txt and merkletree.json respectively")
}
