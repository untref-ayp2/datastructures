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
	raiz := mas
	// Recorrido en Pre-Order
	assert.Equal(t, []string{"+", "a", "*", "-", "b", "c", "d"}, raiz.GetPreOrder())
	// Recorrido en In-Order
	assert.Equal(t, []string{"a", "+", "b", "-", "c", "*", "d"}, raiz.GetInOrder())
	// Recorrido en Post-Order
	assert.Equal(t, []string{"a", "b", "c", "-", "d", "*", "+"}, raiz.GetPostOrder())
}

func TestRecorridosArbolVacio(t *testing.T) {
	raiz := NewBinaryTree("algo")
	raiz.Empty()
	// Recorrido en Pre-Order
	vacio := []string{}
	assert.Equal(t, vacio, raiz.GetPreOrder())
	// Recorrido en In-Order
	assert.Equal(t, vacio, raiz.GetInOrder())
	// Recorrido en Post-Order
	assert.Equal(t, vacio, raiz.GetPostOrder())
}

func TestRecorridosUnSoloNodo(t *testing.T) {
	raiz := NewBinaryTree(123)
	// Recorrido en Pre-Order
	unoSolo := []int{123}
	assert.Equal(t, unoSolo, raiz.GetPreOrder())
	// Recorrido en In-Order
	assert.Equal(t, unoSolo, raiz.GetInOrder())
	// Recorrido en Post-Order
	assert.Equal(t, unoSolo, raiz.GetPostOrder())
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
