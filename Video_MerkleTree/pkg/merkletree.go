package merkletree

import (
	"crypto/sha256"
	"io/ioutil"
	"math"
)

// MerkleNode represents a node in the Merkle tree
type MerkleNode struct {
	Left  *MerkleNode // Left child node
	Right *MerkleNode // Right child node
	Data  []byte      // Data stored in the node
}

// MerkleTree represents the Merkle tree
type MerkleTree struct {
	Root *MerkleNode // Root node of the Merkle tree
}

// NewMerkleNode creates a new Merkle node with the given data
func NewMerkleNode(data []byte) *MerkleNode {
	return &MerkleNode{nil, nil, data}
}

// calculateChunkSize calculates the chunk size based on the video size
func calculateChunkSize(videoSize int64) int {
	// Define maximum chunk size
	const maxChunkSize = 1024 * 1024 // 1 MB

	// Calculate chunk size based on video size and other factors
	chunkSize := int(math.Sqrt(float64(videoSize))) // Example dynamic chunking algorithm

	// Ensure chunk size does not exceed maximum chunk size
	if chunkSize > maxChunkSize {
		chunkSize = maxChunkSize
	}

	return chunkSize
}

// NewMerkleTree constructs a Merkle tree from the given data
func NewMerkleTree(data []byte) *MerkleTree {
	var nodes []MerkleNode

	// Calculate dynamic chunk size based on video size
	chunkSize := calculateChunkSize(int64(len(data)))

	// Chunk the video data dynamically
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		nodes = append(nodes, *NewMerkleNode(data[i:end]))
	}

	// Construct the Merkle tree using dynamically sized chunks
	for len(nodes) > 1 {
		var level []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			node := MerkleNode{
				Left:  &nodes[i],
				Right: nil,
				Data:  nil,
			}
			if i+1 < len(nodes) {
				node.Right = &nodes[i+1]
			}
			data := append(node.Left.Data, node.Right.Data...)
			hash := sha256.Sum256(data)
			node.Data = hash[:]
			level = append(level, node)
		}
		nodes = level
	}

	return &MerkleTree{Root: &nodes[0]}
}

// GetRoot returns the Merkle root hash of the tree
func (mt *MerkleTree) GetRoot() []byte {
	return mt.Root.Data
}

// GenerateProof generates a Merkle proof for a specific chunk index
func (mt *MerkleTree) GenerateProof(chunkIndex int) [][]byte {
	var proof [][]byte
	currentNode := mt.Root

	// Traverse the tree to generate the proof
	for depth := 0; depth < chunkIndex; depth++ {
		if chunkIndex&(1<<depth) == 0 {
			proof = append(proof, currentNode.Right.Data)
			currentNode = currentNode.Left
		} else {
			proof = append(proof, currentNode.Left.Data)
			currentNode = currentNode.Right
		}
	}
	return proof
}

// NewMerkleRoot returns the Merkle root hash of the tree
func NewMerkleRoot(data []byte) []byte {
	mt := NewMerkleTree(data)
	return mt.GetRoot()
}

// WriteRootToFile writes the Merkle root to a file
func WriteRootToFile(root []byte, filename string) error {
	return ioutil.WriteFile(filename, root, 0644)
}

// WriteProofToFile writes the Merkle proof to a file
func WriteProofToFile(proof [][]byte, filename string) error {
	// Combine proof slices into a single byte slice
	var combinedProof []byte
	for _, data := range proof {
		combinedProof = append(combinedProof, data...)
	}
	return ioutil.WriteFile(filename, combinedProof, 0644)
}
