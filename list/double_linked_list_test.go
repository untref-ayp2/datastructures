package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDoubleLinkedList(t *testing.T) {
	list := NewDoubleLinkedList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestDoubleLinkedListPrepend(t *testing.T) {
	list := NewDoubleLinkedList[string]()

	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())

	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestDoubleLinkedListAppend(t *testing.T) {
	list := NewDoubleLinkedList[string]()

	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())

	list.Append("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "2", list.Tail().Data())
}

func TestDoubleLinkedListClear(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Clear()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestDoubleLinkedListIsEmpty(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	assert.True(t, list.IsEmpty())

	list.Append(1)
	assert.False(t, list.IsEmpty())
}

func TestDoubleLinkedListSize(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	assert.Equal(t, 0, list.Size())

	list.Append(1)
	assert.Equal(t, 1, list.Size())

	list.Append(2)
	assert.Equal(t, 2, list.Size())
}

func TestDoubleLinkedListHead(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	assert.Nil(t, list.Head())

	list.Append(1)
	assert.Equal(t, 1, list.Head().Data())

	list.Append(2)
	assert.Equal(t, 1, list.Head().Data())
}

func TestDoubleLinkedListTail(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	assert.Nil(t, list.Tail())

	list.Append(1)
	assert.Equal(t, 1, list.Tail().Data())

	list.Append(2)
	assert.Equal(t, 2, list.Tail().Data())
}

func TestDoubleLinkedListRemove(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.Remove(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.Remove(3)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestDoubleLinkedListRemoveLastElement(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Append(1)
	list.Append(2)

	list.Remove(2)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())
}

func TestDoubleLinkedListRemoveFirst(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveFirst()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.RemoveFirst()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestDoubleLinkedListRemoveLast(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveLast()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 2, list.Tail().Data())

	list.RemoveLast()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())

	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestDoubleLinkedListAllRemovesOnEmptyList(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())

	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())

	list.Remove(1)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestDoubleLinkedListFind(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	nodo := list.Find(2)
	assert.NotNil(t, nodo)
	assert.Equal(t, 2, nodo.Data())

	nodo = list.Find(4)
	assert.Nil(t, nodo)
}

func TestDoubleLinkedListStringOnEmpty(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	assert.Equal(t, "DoubleLinkedList: []", list.String())
}

func TestDoubleLinkedListString(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.Equal(t, "DoubleLinkedList: [1] ↔ [2] ↔ [3]", list.String())
}
