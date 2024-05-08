package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircularList(t *testing.T) {
	list := NewCircularList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestCircularListPrepend(t *testing.T) {
	list := NewCircularList[string]()

	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())

	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestCircularListAppend(t *testing.T) {
	list := NewCircularList[string]()

	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())

	list.Append("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "2", list.Tail().Data())
}

func TestCircularListClear(t *testing.T) {
	list := NewCircularList[int]()

	list.Append(1)
	list.Append(2)
	list.Clear()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestCircularListIsEmpty(t *testing.T) {
	list := NewCircularList[int]()

	assert.True(t, list.IsEmpty())

	list.Append(1)
	assert.False(t, list.IsEmpty())
}

func TestCircularListSize(t *testing.T) {
	list := NewCircularList[int]()

	assert.Equal(t, 0, list.Size())

	list.Append(1)
	assert.Equal(t, 1, list.Size())

	list.Append(2)
	assert.Equal(t, 2, list.Size())
}

func TestCircularListHead(t *testing.T) {
	list := NewCircularList[int]()

	assert.Nil(t, list.Head())

	list.Append(1)
	assert.Equal(t, 1, list.Head().Data())

	list.Append(2)
	assert.Equal(t, 1, list.Head().Data())
}

func TestCircularListTail(t *testing.T) {
	list := NewCircularList[int]()

	assert.Nil(t, list.Tail())

	list.Append(1)
	assert.Equal(t, 1, list.Tail().Data())

	list.Append(2)
	assert.Equal(t, 2, list.Tail().Data())
}

func TestCircularListRemove(t *testing.T) {
	list := NewCircularList[int]()

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

func TestCircularListRemoveLast(t *testing.T) {
	list := NewCircularList[int]()

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

func TestCircularListRemoveFirst(t *testing.T) {
	list := NewCircularList[int]()

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

func TestCircularListAllRemoveOnEmptyList(t *testing.T) {
	list := NewCircularList[int]()

	list.Remove(1)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())

	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())

	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestCircularListFind(t *testing.T) {
	list := NewCircularList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	nodo := list.Find(2)
	assert.NotNil(t, nodo)
	assert.Equal(t, 2, nodo.Data())

	nodo = list.Find(4)
	assert.Nil(t, nodo)
}

func TestCircularListFindOnEmptyList(t *testing.T) {
	list := NewCircularList[int]()

	nodo := list.Find(1)
	assert.Nil(t, nodo)
}

func TestCircularListStringOnEmpty(t *testing.T) {
	list := NewCircularList[int]()

	assert.Equal(t, "CircularList: ⇢ [] ⇠", list.String())
}

func TestCircularListString(t *testing.T) {
	list := NewCircularList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.Equal(t, "CircularList: ⇢ [1] ↔ [2] ↔ [3] ⇠", list.String())
}
