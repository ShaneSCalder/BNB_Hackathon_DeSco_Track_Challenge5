package main

import (
	"Video_MerkleTree/pkg/merkletree" // Adjust this import path as necessary
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Node represents a node in the Merkle tree, including its hash and the hashes of its children.
type Node struct {
	Hash  string `json:"hash"`
	Left  *Node  `json:"left,omitempty"`
	Right *Node  `json:"right,omitempty"`
}

// buildNode recursively constructs the structured representation of the Merkle tree.
func buildNode(node *merkletree.MerkleNode) *Node {
	if node == nil {
		return nil
	}
	newNode := &Node{
		Hash: hex.EncodeToString(node.Data),
	}
	if node.Left != nil {
		newNode.Left = buildNode(node.Left)
	}
	if node.Right != nil {
		newNode.Right = buildNode(node.Right)
	}
	return newNode
}

func main() {
	videoPath := "data/TestVideo.mp4"
	videoData, err := ioutil.ReadFile(videoPath)
	if err != nil {
		fmt.Printf("Failed to read video file: %s\n", err)
		return
	}

	chunkSize := 1024 // Example chunk size
	var chunks [][]byte
	for i := 0; i < len(videoData); i += chunkSize {
		end := i + chunkSize
		if end > len(videoData) {
			end = len(videoData)
		}
		chunks = append(chunks, videoData[i:end])
	}

	tree := merkletree.NewMerkleTree(chunks) // Construct the Merkle tree
	structuredRoot := buildNode(tree.Root)   // Build the structured representation

	// Serialize the structured tree to JSON
	treeJSON, err := json.MarshalIndent(structuredRoot, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal Merkle tree to JSON: %s\n", err)
		return
	}

	// Write the structured tree to a JSON file
	err = ioutil.WriteFile("data/merkletree.json", treeJSON, 0644)
	if err != nil {
		fmt.Printf("Failed to write Merkle tree to file: %s\n", err)
		return
	}
	fmt.Println("Merkle tree structure written to merkletree.json")

	// Serialize the Merkle root to JSON and write it to a file
	merkleRootJSON, err := json.MarshalIndent(map[string]string{"merkleRoot": hex.EncodeToString(tree.Root.Data)}, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal Merkle root to JSON: %s\n", err)
		return
	}

	err = ioutil.WriteFile("data/merkleroot.json", merkleRootJSON, 0644)
	if err != nil {
		fmt.Printf("Failed to write Merkle root to file: %s\n", err)
		return
	}
	fmt.Println("Merkle root written to merkleroot.json")
}
