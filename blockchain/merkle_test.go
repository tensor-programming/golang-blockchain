package blockchain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMerkleNode(t *testing.T) {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
		[]byte("node4"),
		[]byte("node5"),
		[]byte("node6"),
		[]byte("node7"),
	}

	// level 1
	mn1 := NewMerkleNode(nil, nil, data[0])
	mn2 := NewMerkleNode(nil, nil, data[1])
	mn3 := NewMerkleNode(nil, nil, data[2])
	mn4 := NewMerkleNode(nil, nil, data[3])
	mn5 := NewMerkleNode(nil, nil, data[4])
	mn6 := NewMerkleNode(nil, nil, data[5])
	mn7 := NewMerkleNode(nil, nil, data[6])
	mn8 := NewMerkleNode(nil, nil, data[6])

	// level 2
	mn9 := NewMerkleNode(mn1, mn2, nil)
	mn10 := NewMerkleNode(mn3, mn4, nil)
	mn11 := NewMerkleNode(mn5, mn6, nil)
	mn12 := NewMerkleNode(mn7, mn8, nil)

	//level 3

	mn13 := NewMerkleNode(mn9, mn10, nil)
	mn14 := NewMerkleNode(mn11, mn12, nil)

	//level 4

	mn15 := NewMerkleNode(mn13, mn14, nil)

	root := fmt.Sprintf("%x", mn15.Data)
	tree := NewMerkleTree(data)

	assert.Equal(t, root, fmt.Sprintf("%x", tree.RootNode.Data), "Merkle node root has is equal")

}
