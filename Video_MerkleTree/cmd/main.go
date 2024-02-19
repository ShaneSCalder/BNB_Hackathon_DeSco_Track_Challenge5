package main

import (
	"fmt"
	"os"
	"your_module_path/merkletree" // Replace "your_module_path" with your module's actual path
)

func main() {
	// Read MP4 video file into memory
	videoBytes, err := os.ReadFile("your_video.mp4")
	if err != nil {
		fmt.Printf("Error reading video file: %v\n", err)
		return
	}

	// Create a Merkle tree from the video data
	mt := merkletree.NewMerkleTree(videoBytes)

	// Get the Merkle root
	merkleRoot := mt.GetRoot()

	// Write the Merkle root to a file
	err = merkletree.WriteRootToFile(merkleRoot, "merkle_root.txt")
	if err != nil {
		fmt.Printf("Error writing Merkle root to file: %v\n", err)
		return
	}
	fmt.Println("Merkle root written to merkle_root.txt")

	// Generate a Merkle proof for a specific chunk index (e.g., chunkIndex)
	chunkIndex := 3 // Example chunk index; consider making this configurable
	proof := mt.GenerateProof(chunkIndex)

	// Write the Merkle proof to a file
	err = merkletree.WriteProofToFile(proof, "merkle_proof.txt")
	if err != nil {
		fmt.Printf("Error writing Merkle proof to file: %v\n", err)
		return
	}
	fmt.Println("Merkle proof written to merkle_proof.txt")
}
