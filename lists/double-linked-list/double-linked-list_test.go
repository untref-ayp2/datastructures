package double_linked_list

import (
	"testing"
)

func TestDoubleLinkedList(t *testing.T){
	// Crear una lista de enteros
	lista := NewDoubleLinkedList[int]()

	// El estado inicial de la lista debe ser lista vacía
	if !lista.IsEmpty() {
		t.Error("La lista debería estar vacía")
	}

	// La lista debería tener tamaño 0
	if lista.Size() != 0 {
		t.Error("El tamaño de la lista debería ser 0")
	}

	// Agregar elementos a la lista
	lista.InsertAt(0, 1)
	lista.InsertAt(1, 2)
	lista.InsertAt(2, 3)

	if lista.IsEmpty() {
		t.Error("La lista no debería estar vacía")
	}

	if lista.Size() != 3 {
		t.Error("El tamaño de la lista debería ser 3")
	}

	// el primer elemento debería ser 1
	if value, _ := lista.Get(0); value != 1 {
		t.Error("El primer elemento de la lista debería ser 1")
	}

	// el segundo elemento debería ser 2
	if value, _ := lista.Get(1); value != 2 {
		t.Error("El segundo elemento de la lista debería ser 2")
	}

	// el tercer elemento debería ser 3
	if value, _ := lista.Get(2); value != 3 {
		t.Error("El tercer elemento de la lista debería ser 3")
	}

	// Test removing elements from the list
	_, err := lista.RemoveAt(1)
	if err != nil {
		t.Error("No se puede remover el elemento en la posición 1")
	}

	if lista.Size() != 2 {
		t.Error("El tamaño de la lista debería ser 2 después de remover un elemento")
	}

	// el segundo elemento ahora debería ser 3
	if value, _ := lista.Get(1); value != 3 {
		t.Error("El segundo elemento de la lista debería ser 3 después de remover un elemento")
	}
}

func TestDoubleLinkedListErrors(t *testing.T){
	// Crear una lista de enteros
	lista := NewDoubleLinkedList[int]()

	// Agregar elementos a la lista
	lista.InsertAt(0, 1)
	lista.InsertAt(1, 2)
	lista.InsertAt(2, 3)

	// Test removing elements from the list
	_, err := lista.RemoveAt(3)
	if err == nil {
		t.Error("Debería haber un error al intentar remover un elemento fuera de rango")
	}

	_, err = lista.Get(3)
	if err == nil {
		t.Error("Debería haber un error al intentar obtener un elemento fuera de rango")
	}

	err = lista.Set(3, 4)
	if err == nil {
		t.Error("Debería haber un error al intentar reemplazar un elemento fuera de rango")
	}
}

func TestDoubleLinkedListIterate(t *testing.T){
	// Crear una lista de enteros
	lista := NewDoubleLinkedList[int]()

	// Agregar elementos a la lista
	lista.InsertAt(0, 1)
	lista.InsertAt(1, 2)
	lista.InsertAt(2, 3)

	// Recorrer la lista e imprimir los elementos
	i := 0
	for dato := range lista.Iterate() {
		if i == 0 && dato != 1 {
			t.Error("El primer elemento de la lista debería ser 1")
		}

		if i == 1 && dato != 2 {
			t.Error("El segundo elemento de la lista debería ser 2")
		}

		if i == 2 && dato != 3 {
			t.Error("El tercer elemento de la lista debería ser 3")
		}

		i++
	}
}

func TestDoubleLinkedListClear(t *testing.T){
	// Crear una lista de enteros
	lista := NewDoubleLinkedList[int]()

	// Agregar elementos a la lista
	lista.InsertAt(0, 1)
	lista.InsertAt(1, 2)
	lista.InsertAt(2, 3)

	// Limpiar la lista
	lista.Clear()

	if !lista.IsEmpty() {
		t.Error("La lista debería estar vacía después de limpiarla")
	}

	if lista.Size() != 0 {
		t.Error("El tamaño de la lista debería ser 0 después de limpiarla")
	}
}

func TestDoubleLinkedListInsertAt(t *testing.T){
	// Crear una lista de enteros
	lista := NewDoubleLinkedList[int]()

	// Agregar elementos a la lista
	lista.InsertAt(0, 1)
	lista.InsertAt(1, 2)
	lista.InsertAt(2, 3)

	// Insertar un elemento en la posición 1
	lista.InsertAt(1, 4)

	// el segundo elemento debería ser 4
	if value, _ := lista.Get(1); value != 4 {
		t.Error("El segundo elemento de la lista debería ser 4")
	}

	// el tercer elemento debería ser 2
	if value, _ := lista.Get(2); value != 2 {
		t.Error("El tercer elemento de la lista debería ser 2")
	}

	// el cuarto elemento debería ser 3
	if value, _ := lista.Get(3); value != 3 {
		t.Error("El cuarto elemento de la lista debería ser 3")
	}
}

func TestDoubleLinkedListInsertAtErrors(t *testing.T){
	// Crear una lista de enteros
	lista := NewDoubleLinkedList[int]()

	// Agregar elementos a la lista
	lista.InsertAt(0, 1)
	lista.InsertAt(1, 2)
	lista.InsertAt(2, 3)

	// Insertar un elemento en la posición 4
	err := lista.InsertAt(4, 4)
	if err == nil {
		t.Error("Debería haber un error al intentar insertar un elemento fuera de rango")
	}
}

func TestDoubleLinkedListSet(t *testing.T){
	// Crear una lista de enteros
	lista := NewDoubleLinkedList[int]()

	// Agregar elementos a la lista
	lista.InsertAt(0, 1)
	lista.InsertAt(1, 2)
	lista.InsertAt(2, 3)

	// Reemplazar el segundo elemento por 4
	lista.Set(1, 4)

	// el segundo elemento debería ser 4
	if value, _ := lista.Get(1); value != 4 {
		t.Error("El segundo elemento de la lista debería ser 4")
	}
}
