package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTNuevo(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.Equal(t, 0, bstree.Size())
	assert.Nil(t, bstree.GetRoot())
	assert.Nil(t, bstree.FindMin())
}

func TestBSTInsertarUnElemento(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())
}

func TestBSTObtenerRaiz(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 4, bstree.GetRoot().GetData())
}

func TestBSTBuscarElementoExistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	found := bstree.Search(2)

	assert.NotNil(t, found)
	assert.Equal(t, 2, found.GetData())

	found = bstree.Search(5)

	assert.NotNil(t, found)
	assert.Equal(t, 5, found.GetData())
}

func TestBSTBuscarElementoInexistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	found := bstree.Search(6)

	assert.Nil(t, found)
}

func TestBSTBuscarMinimo(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	min := bstree.FindMin()

	assert.Equal(t, 1, min.GetData())
}

func TestBSTRecorrerInOrder(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.Equal(t, "1 2 3 4 5 ", bstree.InOrder())
}

func TestBSTRepresentacionEnString(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.Equal(t, "1 2 3 4 5 ", bstree.String())
}

func TestBSTBorrarRaiz(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.NotNil(t, bstree.Search(4))
	assert.Equal(t, 5, bstree.Size())

	bstree.Remove(4)

	assert.Nil(t, bstree.Search(4))
	assert.Equal(t, 4, bstree.Size())
}

func TestBSTBorrarTodosDeAUno(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.Equal(t, 5, bstree.Size())

	bstree.Remove(1)
	assert.Nil(t, bstree.Search(1))
	assert.Equal(t, 4, bstree.Size())

	bstree.Remove(2)
	assert.Nil(t, bstree.Search(2))
	assert.Equal(t, 3, bstree.Size())

	bstree.Remove(3)
	assert.Nil(t, bstree.Search(3))
	assert.Equal(t, 2, bstree.Size())

	bstree.Remove(5)
	assert.Nil(t, bstree.Search(5))
	assert.Equal(t, 1, bstree.Size())

	bstree.Remove(4)
	assert.Nil(t, bstree.Search(4))
	assert.Equal(t, 0, bstree.Size())
}

func TestBSTBorrarInexistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.Equal(t, 0, bstree.Size())

	bstree.Remove(1)

	assert.Equal(t, 0, bstree.Size())
}

func TestBSTBorrarUnicoElemento(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())

	bstree.Remove(4)

	assert.Equal(t, 0, bstree.Size())
}

func TestBSTBorrarRaizConUnHijo(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(3)

	assert.Equal(t, 2, bstree.Size())

	bstree.Remove(4)

	assert.Nil(t, bstree.Search(4))
	assert.Equal(t, 1, bstree.Size())
}
