package single_linked_list

import (
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// El estado inicial de la lista debe ser lista vacía
	if !list.IsEmpty() {
		t.Error("La lista debería estar vacía")
	}

	// La lista debería tener tamaño 0
	if list.Size() != 0 {
		t.Error("El tamaño de la lista debería ser 0")
	}

	// Agregar elementos a la lista
	list.InsertAt(0, 1)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)

	if list.IsEmpty() {
		t.Error("La lista no debería estar vacía")
	}

	if list.Size() != 3 {
		t.Error("El tamaño de la lista debería ser 3")
	}

	// el primer elemento debería ser 1
	if value, _ := list.Get(0); value != 1 {
		t.Error("El primer elemento de la lista debería ser 1")
	}

	// el segundo elemento debería ser 2
	if value, _ := list.Get(1); value != 2 {
		t.Error("El segundo elemento de la lista debería ser 2")
	}

	// el tercer elemento debería ser 3
	if value, _ := list.Get(2); value != 3 {
		t.Error("El tercer elemento de la lista debería ser 3")
	}

	// Test removing elements from the list
	_, err := list.RemoveAt(1)
	if err != nil {
		t.Error("No se puede remover el elemento en la posición 1")
	}

	if list.Size() != 2 {
		t.Error("El tamaño de la lista debería ser 2 después de remover un elemento")
	}

	// el segundo elemento ahora debería ser 3
	if value, _ := list.Get(1); value != 3 {
		t.Error("El segundo elemento de la lista debería ser 3 después de remover un elemento")
	}

	// Test clearing the list
	list.Clear()

	if !list.IsEmpty() {
		t.Error("La lista debería estar vacía después de limpiarla")
	}
}

func TestSingleLinkedListErrors(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Tratar de obtener un elemento de una lista vacía debería devolver un error
	_, err := list.Get(0)
	if err == nil {
		t.Error("Debería devolver un error al intentar obtener un elemento de una lista vacía")
	}

	// Tratar de remover un elemento de una lista vacía debería devolver un error
	_, err = list.RemoveAt(0)
	if err == nil {
		t.Error("Debería devolver un error al intentar remover un elemento de una lista vacía")
	}

	// Tratar de remover un elemento en una posición fuera de rango debería devolver un error
	list.InsertAt(0, 1)
	_, err = list.RemoveAt(1)
	if err == nil {
		t.Error("Debería devolver un error al intentar remover un elemento en una posición fuera de rango")
	}
}

func TestSingleLinkedListInsertAt(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Insertar un elemento en la posición 0 debería agregar el elemento al inicio de la lista
	list.InsertAt(0, 1)
	if value, _ := list.Get(0); value != 1 {
		t.Error("El primer elemento de la lista debería ser 1")
	}

	// Insertar un elemento en la posición 0 debería agregar el elemento al inicio de la lista
	list.InsertAt(0, 2)
	if value, _ := list.Get(0); value != 2 {
		t.Error("El primer elemento de la lista debería ser 2")
	}

	// Insertar un elemento en la posición 1 debería agregar el elemento al final de la lista
	list.InsertAt(1, 3)
	if value, _ := list.Get(1); value != 3 {
		t.Error("El segundo elemento de la lista debería ser 3")
	}

	// Insertar un elemento en la posición 1 debería agregar el elemento en la posición 1
	list.InsertAt(1, 4)
	if value, _ := list.Get(1); value != 4 {
		t.Error("El segundo elemento de la lista debería ser 4")
	}
}

func TestSingleLinkedListSet(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Insertar elementos a la lista
	list.InsertAt(0, 1)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)

	// Reemplazar el primer elemento de la lista
	list.Set(0, 4)
	if value, _ := list.Get(0); value != 4 {
		t.Error("El primer elemento de la lista debería ser 4")
	}

	// Reemplazar el último elemento de la lista
	list.Set(2, 5)
	if value, _ := list.Get(2); value != 5 {
		t.Error("El último elemento de la lista debería ser 5")
	}
}

func TestSingleLinkedListRemoveAt(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Insertar elementos a la lista
	list.InsertAt(0, 1)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)

	// Remover el primer elemento de la lista
	list.RemoveAt(0)
	value, _ := list.Get(0)
	if value != 2 {
		t.Error("El primer elemento de la lista debería ser 2")
	}

	// Remover el último elemento de la lista
	list.RemoveAt(1)
	value, _ = list.Get(list.Size() - 1)
	if value != 2 {
		t.Error("El último elemento de la lista debería ser 2")
	}
}

func TestSingleLinkedListInsertAtErrors(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Tratar de insertar un elemento en una posición fuera de rango debería devolver un error
	err := list.InsertAt(1, 1)
	if err == nil {
		t.Error("Debería devolver un error al intentar insertar un elemento en una posición fuera de rango")
	}
}

func TestSingleLinkedListRemoveAtErrors(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Tratar de remover un elemento en una posición fuera de rango debería devolver un error
	list.InsertAt(0, 1)
	_, err := list.RemoveAt(1)
	if err == nil {
		t.Error("Debería devolver un error al intentar remover un elemento en una posición fuera de rango")
	}
}

func TestSingleLinkedListIterate(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Insertar elementos a la lista
	list.InsertAt(0, 1)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)

	// Iterar sobre la lista
	var values []int
	for value := range list.Iterate() {
		values = append(values, value)
	}

	// Los valores de la lista deberían ser 1, 2, 3
	if values[0] != 1 || values[1] != 2 || values[2] != 3 {
		t.Error("Los valores de la lista deberían ser 1, 2, 3")
	}
}

func TestSingleLinkedListIterateEmpty(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Iterar sobre una lista vacía debería devolver un canal cerrado
	for _ = range list.Iterate() {
		t.Error("El canal debería estar cerrado")
	}
}

func TestSingleLinkedListIndexOf(t *testing.T) {
	// Crea una lista de enteros
	list := NewSingleLinkedList[int]()

	// Insertar elementos a la lista
	list.InsertAt(0, 1)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)

	// Buscar un elemento que está en la lista
	index := list.IndexOf(2)
	if index != 1 {
		t.Error("El índice del elemento 2 debería ser 1")
	}

	// Buscar un elemento que no está en la lista
	index = list.IndexOf(4)
	if index != -1 {
		t.Error("El índice del elemento 4 debería ser -1")
	}
}
