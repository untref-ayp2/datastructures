package double_linked_list

import (
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList[int]()
	if list == nil {
		t.Error("NewList() devolvió nil")
	}
	if list.Size() != 0 {
		t.Error("El tamaño de la lista debería ser 0")
	}
	if list.Head() != nil {
		t.Error("Head() debería devolver nil")
	}
	if list.Tail() != nil {
		t.Error("Tail() debería devolver nil")
	}
}

func TestPrepend(t *testing.T) {
	list := NewList[int]()
	list.Prepend(1)
	if list.Size() != 1 {
		t.Error("El tamaño de la lista debería ser 1")
	}
	if list.Head().Data() != 1 {
		t.Error("El dato del primer nodo debería ser 1")
	}
	if list.Tail().Data() != 1 {
		t.Error("El dato del último nodo debería ser 1")
	}
	list.Prepend(2)
	if list.Size() != 2 {
		t.Error("El tamaño de la lista debería ser 2")
	}
	if list.Head().Data() != 2 {
		t.Error("El dato del primer nodo debería ser 2")
	}
	if list.Tail().Data() != 1 {
		t.Error("El dato del último nodo debería ser 1")
	}
}

func TestAppend(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	if list.Size() != 1 {
		t.Error("El tamaño de la lista debería ser 1")
	}
	if list.Head().Data() != 1 {
		t.Error("El dato del primer nodo debería ser 1")
	}
	if list.Tail().Data() != 1 {
		t.Error("El dato del último nodo debería ser 1")
	}
	list.Append(2)
	if list.Size() != 2 {
		t.Error("El tamaño de la lista debería ser 2")
	}
	if list.Head().Data() != 1 {
		t.Error("El dato del primer nodo debería ser 1")
	}
	if list.Tail().Data() != 2 {
		t.Error("El dato del último nodo debería ser 2")
	}
}

func TestClear(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Clear()
	if list.Size() != 0 {
		t.Error("El tamaño de la lista debería ser 0")
	}
	if list.Head() != nil {
		t.Error("Head() debería devolver nil")
	}
	if list.Tail() != nil {
		t.Error("Tail() debería devolver nil")
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewList[int]()
	if !list.IsEmpty() {
		t.Error("IsEmpty() debería devolver true")
	}
	list.Append(1)
	if list.IsEmpty() {
		t.Error("IsEmpty() debería devolver false")
	}
}

func TestSize(t *testing.T) {
	list := NewList[int]()
	if list.Size() != 0 {
		t.Error("El tamaño de la lista debería ser 0")
	}
	list.Append(1)
	if list.Size() != 1 {
		t.Error("El tamaño de la lista debería ser 1")
	}
	list.Append(2)
	if list.Size() != 2 {
		t.Error("El tamaño de la lista debería ser 2")
	}
}

func TestHead(t *testing.T) {
	list := NewList[int]()
	if list.Head() != nil {
		t.Error("Head() debería devolver nil")
	}
	list.Append(1)
	if list.Head().Data() != 1 {
		t.Error("El dato del primer nodo debería ser 1")
	}
	list.Append(2)
	if list.Head().Data() != 1 {
		t.Error("El dato del primer nodo debería ser 1")
	}
}

func TestTail(t *testing.T) {
	list := NewList[int]()
	if list.Tail() != nil {
		t.Error("Tail() debería devolver nil")
	}
	list.Append(1)
	if list.Tail().Data() != 1 {
		t.Error("El dato del último nodo debería ser 1")
	}
	list.Append(2)
	if list.Tail().Data() != 2 {
		t.Error("El dato del último nodo debería ser 2")
	}
}

func TestRemove(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	if list.Size() != 2 {
		t.Error("El tamaño de la lista debería ser 2")
	}
	if list.Head().Data() != 1 {
		t.Error("El dato del primer nodo debería ser 1")
	}
	if list.Tail().Data() != 3 {
		t.Error("El dato del último nodo debería ser 3")
	}
	list.Remove(1)
	if list.Size() != 1 {
		t.Error("El tamaño de la lista debería ser 1")
	}
	if list.Head().Data() != 3 {
		t.Error("El dato del primer nodo debería ser 3")
	}
	if list.Tail().Data() != 3 {
		t.Error("El dato del último nodo debería ser 3")
	}
	list.Remove(3)
	if list.Size() != 0 {
		t.Error("El tamaño de la lista debería ser 0")
	}
	if list.Head() != nil {
		t.Error("Head() debería devolver nil")
	}
	if list.Tail() != nil {
		t.Error("Tail() debería devolver nil")
	}
}

func TestFind(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	nodo := list.Find(2)
	if nodo == nil {
		t.Error("Find(2) debería devolver un nodo")
	}
	if nodo.Data() != 2 {
		t.Error("El dato del nodo debería ser 2")
	}
	if list.Find(4) != nil {
		t.Error("Find(4) debería devolver nil")
	}
}