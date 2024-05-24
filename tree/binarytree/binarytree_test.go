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
	assert.Equal(t, "+a*-bcd", raiz.StringPreOrder())
	// Recorrido en In-Order
	assert.Equal(t, []string{"a", "+", "b", "-", "c", "*", "d"}, raiz.GetInOrder())
	assert.Equal(t, "a+b-c*d", raiz.StringInOrder())
	// Recorrido en Post-Order
	assert.Equal(t, []string{"a", "b", "c", "-", "d", "*", "+"}, raiz.GetPostOrder())
	assert.Equal(t, "abc-d*+", raiz.StringPostOrder())
}

func TestGetRoot(t *testing.T) {
	raiz := NewBinaryTree("algo")
	assert.Equal(t, "algo", raiz.GetRoot().GetData())
}

func TestRecorridosArbolVacio(t *testing.T) {
	raiz := NewBinaryTree("algo")
	raiz.Clear()
	// Recorrido en Pre-Order
	assert.Empty(t, raiz.GetPreOrder())
	assert.Equal(t, "", raiz.StringPreOrder())
	// Recorrido en In-Order
	assert.Empty(t, raiz.GetInOrder())
	assert.Equal(t, "", raiz.StringInOrder())
	// Recorrido en Post-Order
	assert.Empty(t, raiz.GetPostOrder())
	assert.Equal(t, "", raiz.StringPostOrder())
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
	mas.Clear()
	assert.True(t, mas.IsEmpty())
}

func TestSizeHeightEmptyOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree("+")
	tree.Clear()

	// Size del arbol(cantidad de nodos)
	assert.Equal(t, 0, tree.Size())

	// Height del arbol
	assert.Equal(t, -1, tree.Height())

	// isEmpty Test
	assert.True(t, tree.IsEmpty())
}

func TestInertsOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree("+")

	uno := NewBinaryTree("1")
	dos := NewBinaryTree("2")

	tree.Clear()
	tree.InsertLeft(uno)
	assert.False(t, tree.IsEmpty())

	tree.Clear()
	tree.InsertRight(dos)
	assert.False(t, tree.IsEmpty())
}

func TestIteratorUnoPorUno(t *testing.T) {
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
	tree := mas

	it := tree.Iterator()

	expectedValues := []string{"a", "+", "b", "-", "c", "*", "d"}

	for _, expected := range expectedValues {
		assert.True(t, it.HasNext())
		val, err := it.Next()
		assert.Equal(t, expected, val)
		assert.NoError(t, err)
	}
	assert.False(t, it.HasNext())
}

func TestIteratorWhenEmpty(t *testing.T) {
	btree := NewBinaryTree[int](0)
	btree.Clear()

	it := btree.Iterator()

	assert.False(t, it.HasNext())

	val, err := it.Next()

	assert.Zero(t, val)
	assert.Error(t, err)
}
