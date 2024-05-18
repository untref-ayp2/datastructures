package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecorridosDiapo(t *testing.T) {
	mas := NewBinaryTree("+")
	menos := NewBinaryTree("-")
	mult := NewBinaryTree("*")
	a := NewBinaryTree("a")
	b := NewBinaryTree("b")
	c := NewBinaryTree("c")
	d := NewBinaryTree("d")
	menos.InsertLeft(b)
	menos.InsertRight(c)
	mult.InsertLeft(menos)
	mult.InsertRight(d)
	mas.InsertLeft(a)
	mas.InsertRight(mult)
	// Recorrido en Pre-Order
	assert.Equal(t, "+a*-bcd", mas.StringPreOrder())
	// Recorrido en In-Order
	assert.Equal(t, "a+b-c*d", mas.StringInOrder())
	// Recorrido en Post-Order
	assert.Equal(t, "abc-d*+", mas.StringPostOrder())
}

func TestSizeHeightEmptyDiapo(t *testing.T) {
	mas := NewBinaryTree("+")
	menos := NewBinaryTree("-")
	mult := NewBinaryTree("*")
	a := NewBinaryTree("a")
	b := NewBinaryTree("b")
	c := NewBinaryTree("c")
	d := NewBinaryTree("d")
	menos.InsertLeft(b)
	menos.InsertRight(c)
	mult.InsertLeft(menos)
	mult.InsertRight(d)
	mas.InsertLeft(a)
	mas.InsertRight(mult)

	// Size del arbol(cantidad de nodos)
	assert.Equal(t, 7, mas.Size())

	// Height del arbol
	assert.Equal(t, 3, mas.Height())

	// isEmpty Test
	assert.False(t, mas.IsEmpty())

	// Vaciar el arbol
	mas.Empty()
	assert.True(t, mas.IsEmpty())
}
