package main

import (
	"Video_MerkleTree/pkg/merkletree" // Adjust with your actual module path
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func writeTreeToFile(node *merkletree.MerkleNode, file *os.File, level int) {
	if node == nil {
		return
	}
	indent := strings.Repeat("  ", level)
	_, err := file.WriteString(fmt.Sprintf("%s%x\n", indent, node.Data))
	if err != nil {
		fmt.Printf("Failed to write to file: %s\n", err)
		return
	}
	writeTreeToFile(node.Left, file, level+1)
	writeTreeToFile(node.Right, file, level+1)
}

func main() {
	// Hardcoded path to the video file
	videoPath := "data/TestVideo.mp4" // Adjust this path as needed

	videoData, err := ioutil.ReadFile(videoPath)
	if err != nil {
		fmt.Printf("Failed to read video file: %s\n", err)
		return
	}

	chunkSize := 1024 // Adjust based on needs
	var chunks [][]byte
	for i := 0; i < len(videoData); i += chunkSize {
		end := i + chunkSize
		if end > len(videoData) {
			end = len(videoData)
		}
		chunks = append(chunks, videoData[i:end])
	}

	tree := merkletree.NewMerkleTree(chunks)

	err = ioutil.WriteFile("data/merkleroot.txt", tree.Root.Data, 0644) // Adjust with your actual module path
	if err != nil {
		fmt.Printf("Failed to write Merkle root to file: %s\n", err)
		return
	}
	fmt.Println("Merkle root written to merkleroot.txt")

	file, err := os.Create("data/merkletree.txt") // Adjust with your actual module path
	if err != nil {
		fmt.Printf("Failed to create merkletree.txt: %s\n", err)
		return
	}
	defer file.Close()

	writeTreeToFile(tree.Root, file, 0)
	fmt.Println("Merkle tree structure written to merkletree.txt")
}
