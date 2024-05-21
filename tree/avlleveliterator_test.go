package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLLevelIteratorVacio(t *testing.T) {
	avl := NewAVLTree[int]()
	iterator := NewAVLLevelIterator[int](avl.GetRoot())

	assert.False(t, iterator.HasNext())
}

func TestAVLLevelIterator(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(4)
	avl.Insert(2)
	avl.Insert(6)
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(5)
	avl.Insert(7)

	iterator := NewAVLLevelIterator[int](avl.GetRoot())

	expected := []int{4, 2, 6, 1, 3, 5, 7}

	for _, expVal := range expected {
		assert.True(t, iterator.HasNext())
		val, _ := iterator.Next()
		assert.Equal(t, expVal, val.GetData())
	}
	assert.False(t, iterator.HasNext())
}

func TestAVLLevelIteratorNextOverflow(t *testing.T) {
	avl := NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)

	iterator := NewAVLLevelIterator[int](avl.GetRoot())

	_, _ = iterator.Next()
	_, _ = iterator.Next()
	_, _ = iterator.Next()

	assert.False(t, iterator.HasNext())
	_, err := iterator.Next()
	assert.Error(t, err)
}
