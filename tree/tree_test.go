package tree

import (
	"testing"
)

// Verifica si BinarySearchTree implementa la interfaz Set.
func TestBSTCompliesTreeInterface(_ *testing.T) {
	var _ Tree[int] = NewBinarySearchTree[int]()
}

// Verifica si AVLTree implementa la interfaz Tree.
func TestAVLTreeCompliesTreeInterface(_ *testing.T) {
	var _ Tree[int] = NewAVLTree[int]()
}
