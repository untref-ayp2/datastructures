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

func TestLinkedListPrependOnEmptyList(t *testing.T) {
	list := NewLinkedList[string]()

	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestLinkedListPrependOnNonEmptyList(t *testing.T) {
	list := NewLinkedList[string]()
	list.Append("1")

	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestLinkedListAppendOnEmptyList(t *testing.T) {
	list := NewLinkedList[string]()

	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestLinkedListAppendOnNonEmptyList(t *testing.T) {
	list := NewLinkedList[string]()
	list.Append("1")

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

func TestLinkedListRemoveFirst(t *testing.T) {
	list := NewLinkedList[int]()

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

func TestLinkedListRemoveLast(t *testing.T) {
	list := NewLinkedList[int]()

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

func TestLinkedListRemoveOnLastElement(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Remove(4)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
}

func TestLinkedListRemoveNotExists(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Remove(4)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
}

func TestLinkedListRemoveFirstOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
}

func TestLinkedListRemoveLastOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
}

func TestLinkedListStringOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Equal(t, "LinkedList: []", list.String())
}

func TestLinkedListString(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.Equal(t, "LinkedList: [1] → [2] → [3]", list.String())
}
