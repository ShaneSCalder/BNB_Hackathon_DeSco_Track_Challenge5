package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TargetHashData struct {
	TargetHash string `json:"targetHash"`
}

type MerkleNode struct {
	Hash  string      `json:"hash"`
	Left  *MerkleNode `json:"left,omitempty"`
	Right *MerkleNode `json:"right,omitempty"`
}

// ReadMerkleTree reads the Merkle tree structure from a JSON file and returns the root node.
func ReadMerkleTree(filename string) (*MerkleNode, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading Merkle tree file: %w", err)
	}
	var root MerkleNode
	if err := json.Unmarshal(bytes, &root); err != nil {
		return nil, fmt.Errorf("error unmarshaling Merkle tree: %w", err)
	}
	return &root, nil
}

// ReadTargetHash reads the target hash from a JSON file.
func ReadTargetHash(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading target hash file: %w", err)
	}
	var data TargetHashData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return "", fmt.Errorf("error unmarshaling target hash: %w", err)
	}
	return data.TargetHash, nil
}

// Simplified GenerateProof function.
func GenerateProof(node *MerkleNode, targetHash string, proof *[]string) bool {
	if node == nil {
		return false
	}
	if node.Hash == targetHash {
		// Target hash found, start building proof.
		return true
	}
	// Traverse left subtree.
	if GenerateProof(node.Left, targetHash, proof) {
		if node.Right != nil { // If there's a right sibling, add its hash to the proof.
			*proof = append(*proof, node.Right.Hash)
		}
		return true
	}
	// Traverse right subtree.
	if GenerateProof(node.Right, targetHash, proof) {
		if node.Left != nil { // If there's a left sibling, add its hash to the proof.
			*proof = append(*proof, node.Left.Hash)
		}
		return true
	}
	return false
}

func main() {
	merkleTreePath := "./data/merkletree.json"
	targetHashPath := "./data/targethash.json"

	root, err := ReadMerkleTree(merkleTreePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetHash, err := ReadTargetHash(targetHashPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var proof []string
	found := GenerateProof(root, targetHash, &proof)
	if !found {
		fmt.Println("Target hash not found in the tree")
		return
	}

	proofJSON, err := json.MarshalIndent(map[string]interface{}{
		"targetHash": targetHash,
		"proof":      proof,
	}, "", "    ")
	if err != nil {
		fmt.Println("Failed to marshal proof to JSON:", err)
		return
	}

	err = ioutil.WriteFile("data/proof.json", proofJSON, 0644)
	if err != nil {
		fmt.Println("Failed to write proof to file:", err)
		return
	}

	fmt.Println("Merkle proof successfully written to data/proof.json")
}
