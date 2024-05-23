package avltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLInOrderIteratorVacio(t *testing.T) {
	avl := NewAVLTree[int]()
	iterator := NewAVLInOrderIterator[int](avl.GetRoot())

	assert.False(t, iterator.HasNext())
}

func TestAVLInOrderIterator(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(4)
	avl.Insert(2)
	avl.Insert(6)
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(5)
	avl.Insert(7)

	iterator := NewAVLInOrderIterator[int](avl.GetRoot())

	expected := []int{1, 2, 3, 4, 5, 6, 7}

	for _, expVal := range expected {
		assert.True(t, iterator.HasNext())
		val, _ := iterator.Next()
		assert.Equal(t, expVal, val)
	}
	assert.False(t, iterator.HasNext())
}

func TestAVLInOrderIteratorNextOverflow(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)

	iterator := NewAVLInOrderIterator[int](avl.GetRoot())

	_, _ = iterator.Next()
	_, _ = iterator.Next()
	_, _ = iterator.Next()

	assert.False(t, iterator.HasNext())
	_, err := iterator.Next()
	assert.Error(t, err)
}
