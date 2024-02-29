package merkletree

import (
	"crypto/sha256"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

type MerkleTree struct {
	Root *MerkleNode
}

func NewMerkleNode(left, right *MerkleNode) *MerkleNode {
	node := &MerkleNode{Left: left, Right: right}

	if left == nil && right == nil {
		panic("Left and Right both cannot be nil for a new node.")
	} else if right == nil {
		node.Data = left.Data
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	return node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []*MerkleNode

	for _, datum := range data {
		hashedData := sha256.Sum256(datum)
		node := NewMerkleNode(&MerkleNode{Data: hashedData[:]}, nil)
		nodes = append(nodes, node)
	}

	for len(nodes) > 1 {
		var level []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			var right *MerkleNode
			left := nodes[i]

			if i+1 < len(nodes) {
				right = nodes[i+1]
			}

			node := NewMerkleNode(left, right)
			level = append(level, node)
		}

		nodes = level
	}

	return &MerkleTree{Root: nodes[0]}
}
