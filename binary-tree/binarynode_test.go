package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNuevoNodo(t *testing.T) {
	nodo := NewBinaryNode(5, nil, nil)
	assert.Equal(t, 5, nodo.GetData())
}

//      4
//     / \
//    2   5
//   / \
//  1   3

func TestInOrder(t *testing.T) {
	nodo1 := NewBinaryNode(1, nil, nil)
	nodo3 := NewBinaryNode(3, nil, nil)
	nodo2 := NewBinaryNode(2, nodo1, nodo3)
	nodo5 := NewBinaryNode(5, nil, nil)
	nodo4 := NewBinaryNode(4, nodo2, nodo5)
	raiz := nodo4
	assert.Equal(t, "42135", raiz.StringPreOrder())
	assert.Equal(t, "12345", raiz.StringInOrder())
	assert.Equal(t, "13254", raiz.StringPostOrder())
}

func TestSize(t *testing.T) {
	nodo1 := NewBinaryNode(1, nil, nil)
	nodo3 := NewBinaryNode(3, nil, nil)
	nodo2 := NewBinaryNode(2, nodo1, nodo3)
	nodo5 := NewBinaryNode(5, nil, nil)
	nodo4 := NewBinaryNode(4, nodo2, nodo5)
	raiz := nodo4
	assert.Equal(t, 5, raiz.Size())
}

func TestHeight(t *testing.T) {
	nodo1 := NewBinaryNode(1, nil, nil)
	nodo3 := NewBinaryNode(3, nil, nil)
	nodo2 := NewBinaryNode(2, nodo1, nodo3)
	nodo5 := NewBinaryNode(5, nil, nil)
	nodo4 := NewBinaryNode(4, nodo2, nodo5)
	raiz := nodo4
	assert.Equal(t, 2, raiz.Height())
}
