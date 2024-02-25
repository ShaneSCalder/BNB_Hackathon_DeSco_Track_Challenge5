package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// TargetHashData holds the structure for reading the target hash JSON file.
type TargetHashData struct {
	TargetHash string `json:"targetHash"`
}

// ReadMerkleTree reads the Merkle tree structure from a JSON file.
func ReadMerkleTree(filename string) ([][]string, error) {
	var tree [][]string
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading Merkle tree file: %w", err)
	}
	if err := json.Unmarshal(bytes, &tree); err != nil {
		return nil, fmt.Errorf("error unmarshaling Merkle tree: %w", err)
	}
	return tree, nil
}

// ReadTargetHash reads the target hash from a JSON file.
func ReadTargetHash(filename string) (string, error) {
	var data TargetHashData
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading target hash file: %w", err)
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return "", fmt.Errorf("error unmarshaling target hash: %w", err)
	}
	return data.TargetHash, nil
}

// GenerateProof constructs the Merkle proof for the given target hash.
func GenerateProof(tree [][]string, targetHash string) ([]string, error) {
	proof := []string{}
	var index int
	for i, hash := range tree[0] {
		if hash == targetHash {
			index = i
			break
		}
	}

	for layer := 1; layer < len(tree); layer++ {
		levelLength := len(tree[layer])
		if index%2 == 0 {
			if index+1 < levelLength {
				proof = append(proof, tree[layer][index+1])
			}
		} else {
			proof = append(proof, tree[layer][index-1])
		}
		index /= 2
	}

	return proof, nil
}

// WriteProofToJSON outputs the proof to a JSON file.
func WriteProofToJSON(proof []string, targetHash, filename string) error {
	proofData := map[string]interface{}{
		"targetLeafHash": targetHash,
		"proof":          proof,
	}
	proofJSON, err := json.MarshalIndent(proofData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling proof to JSON: %w", err)
	}
	if err := ioutil.WriteFile(filename, proofJSON, 0644); err != nil {
		return fmt.Errorf("error writing proof to file: %w", err)
	}
	return nil
}

func main() {
	// Fixed paths to the Merkle tree and target hash JSON files
	merkleTreePath := "./data/merkletree.json"
	targetHashPath := "./data/targethash.json"

	tree, err := ReadMerkleTree(merkleTreePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetHash, err := ReadTargetHash(targetHashPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	proof, err := GenerateProof(tree, targetHash)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := WriteProofToJSON(proof, targetHash, "./data/proof.json"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Merkle proof successfully written to ./data/proof.json")
}
