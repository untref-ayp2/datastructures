package tree

import (
	"testing"
)

// FIX: Verificar si AVLTree implementa la interfaz Set.

// Verifica si AVLTree implementa la interfaz Tree.
func TestAVLTreeCompliesTreeInterface(_ *testing.T) {
	var _ Tree[int] = NewAVLTree[int]()
}
