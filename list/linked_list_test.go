package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestLinkedListPrepend(t *testing.T) {
	list := NewLinkedList[string]()

	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())

	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestLinkedListAppend(t *testing.T) {
	list := NewLinkedList[string]()

	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())

	list.Append("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "2", list.Tail().Data())
}

func TestLinkedListClear(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Clear()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestLinkedListIsEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	assert.True(t, list.IsEmpty())

	list.Append(1)
	assert.False(t, list.IsEmpty())
}

func TestLinkedListSize(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Equal(t, 0, list.Size())

	list.Append(1)
	assert.Equal(t, 1, list.Size())

	list.Append(2)
	assert.Equal(t, 2, list.Size())
}

func TestLinkedListHead(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Nil(t, list.Head())

	list.Append(1)
	assert.Equal(t, 1, list.Head().Data())

	list.Append(2)
	assert.Equal(t, 1, list.Head().Data())
}

func TestLinkedListTail(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Nil(t, list.Tail())

	list.Append(1)
	assert.Equal(t, 1, list.Tail().Data())

	list.Append(2)
	assert.Equal(t, 2, list.Tail().Data())
}

func TestLinkedListRemove(t *testing.T) {
	list := NewLinkedList[int]()

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

func TestLinkedListFind(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	nodo := list.Find(2)
	assert.NotNil(t, nodo)
	assert.Equal(t, 2, nodo.Data())

	nodo = list.Find(4)
	assert.Nil(t, nodo)
}
