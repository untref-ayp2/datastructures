package common

import (
	"testing"

	"github.com/untref-ayp2/data-structures/tree/avltree"
	"github.com/untref-ayp2/data-structures/tree/binarytree"
)

// Verifica si BinarySearchTree implementa la interfaz Set.
func TestBSTCompliesTreeInterface(_ *testing.T) {
	var _ Tree[int] = binarytree.NewBinarySearchTree[int]()
}

// Verifica si AVLTree implementa la interfaz Tree.
func TestAVLTreeCompliesTreeInterface(_ *testing.T) {
	var _ Tree[int] = avltree.NewAVLTree[int]()
}
