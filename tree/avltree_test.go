package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTreeVacio(t *testing.T) {
	avl := NewAVLTree[int]()

	assert.True(t, avl.IsEmpty())
	assert.Equal(t, 0, avl.GetHeight())
	assert.Equal(t, 0, avl.GetBalance())
}

func TestAVLTreeClear(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)

	assert.False(t, avl.IsEmpty())

	avl.Clear()

	assert.True(t, avl.IsEmpty())
	assert.Equal(t, 0, avl.GetHeight())
	assert.Equal(t, 0, avl.GetBalance())
}

func TestAVLTreeInsert(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)

	assert.False(t, avl.IsEmpty())
	assert.Equal(t, 2, avl.GetHeight())
	assert.Equal(t, 0, avl.GetBalance())
}

func TestAVLTreeRemove(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)

	avl.Remove(2)

	assert.False(t, avl.IsEmpty())
	assert.Equal(t, 2, avl.GetHeight())
	assert.Equal(t, 1, avl.GetBalance())
}

func TestAVLTreeSearchMinMax(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)
	avl.Insert(4)
	avl.Insert(5)

	assert.Equal(t, 1, avl.FindMin())
	assert.Equal(t, 5, avl.FindMax())
}

func TestAVLTreePrintInOrder(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)
	avl.Insert(4)
	avl.Insert(5)

	assert.Equal(t, " 1  2  3  4  5 ", avl.InOrder())
}

func TestAVLTreeSearch(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)
	avl.Insert(4)
	avl.Insert(5)

	assert.True(t, avl.Search(3))
	assert.False(t, avl.Search(6))
}

func TestAVLTreeString(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)
	avl.Insert(4)
	avl.Insert(5)

	assert.Equal(t, "2", avl.String())
}

func TestAVLTreeGetRoot(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)
	avl.Insert(4)
	avl.Insert(5)

	assert.Equal(t, 2, avl.GetRoot().GetData())
}

func TestAVLTreeCreateReverseOrder(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(5)
	avl.Insert(4)
	avl.Insert(3)
	avl.Insert(2)
	avl.Insert(1)

	assert.Equal(t, 3, avl.GetHeight())
}

func TestAVLTreeInsertRepeated(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(1)
	avl.Insert(1)
	avl.Insert(1)
	avl.Insert(1)

	assert.Equal(t, 1, avl.GetHeight())
}

func TestAVLTreeTodasLasRotaciones(t *testing.T) {
	// rotacion simple a izquierda
	avl := NewAVLTree[int]()
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(30)

	assert.Equal(t, 2, avl.GetHeight())

	// rotacion simple a derecha
	avl = NewAVLTree[int]()
	avl.Insert(30)
	avl.Insert(20)
	avl.Insert(10)

	assert.Equal(t, 2, avl.GetHeight())

	// rotacion izquierda-derecha
	avl = NewAVLTree[int]()
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(20)

	assert.Equal(t, 2, avl.GetHeight())

	// rotacion derecha-izquierda
	avl = NewAVLTree[int]()
	avl.Insert(30)
	avl.Insert(10)
	avl.Insert(20)

	assert.Equal(t, 2, avl.GetHeight())
}

func TestAVLTreeTodasLasRotacionesAlRemover(t *testing.T) {
	// rotacion izquierda-derecha
	avl := NewAVLTree[int]()
	avl.Insert(20)
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(15)
	avl.Remove(30)

	assert.Equal(t, 2, avl.GetHeight())

	// rotacion derecha-izquierda
	avl = NewAVLTree[int]()
	avl.Insert(20)
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(25)
	avl.Remove(10)

	assert.Equal(t, 2, avl.GetHeight())

	// rotacion izquierda-izquierda
	avl = NewAVLTree[int]()
	avl.Insert(20)
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(5)
	avl.Remove(30)

	assert.Equal(t, 2, avl.GetHeight())

	// rotacion derecha-izquierda
	avl = NewAVLTree[int]()
	avl.Insert(20)
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(35)
	avl.Remove(10)

	assert.Equal(t, 2, avl.GetHeight())
}

func TestAVLTreeRemoverDatoSinHijoDerecho(t *testing.T) {
	avl := NewAVLTree[int]()
	avl.Insert(20)
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(25)
	avl.Remove(30)

	assert.Equal(t, 2, avl.GetHeight())
}

func TestAVLTreeRemoverUltimoDato(t *testing.T) {
	avl := NewAVLTree[int]()
	avl.Insert(20)
	avl.Remove(20)

	assert.Equal(t, 0, avl.GetHeight())
}

func TestAVLTreeRemoverDatosInexistentesAIzquierdaYDerecha(t *testing.T) {
	avl := NewAVLTree[int]()
	avl.Insert(20)
	avl.Remove(10)
	avl.Remove(30)

	assert.Equal(t, 1, avl.GetHeight())
}

func TestAVLTreeIterator(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(4)
	avl.Insert(2)
	avl.Insert(6)
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(5)
	avl.Insert(7)

	iterator := avl.Iterator()

	expected := []int{1, 2, 3, 4, 5, 6, 7}

	for _, expVal := range expected {
		assert.True(t, iterator.HasNext())
		val, _ := iterator.Next()
		assert.Equal(t, expVal, val)
	}
	assert.False(t, iterator.HasNext())
}
