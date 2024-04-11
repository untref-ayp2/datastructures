package main

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/list"
)

func main() {
	// Crear una lista enlazada simple de enteros
	listaEntero := list.NewLinkedList[int]()

	// Insertar elementos en la lista
	listaEntero.Append(1)
	listaEntero.Append(2)
	listaEntero.Append(3)
	listaEntero.Append(4)

	// Iterar la lista e imprimir los elementos
	fmt.Println("Iterando lista simple de enteros")
	i := 0
	for actual := listaEntero.Head(); actual != nil; actual = actual.Next() {
		fmt.Printf("Elemento %d: %d\n", i, actual.Data())
		i++
	}

	listaEntero.Clear()

	// Crear una lista doblemente enlazada de enteros
	listaDoble := list.NewDoubleLinkedList[int]()

	// Insertar elementos en la lista
	listaDoble.Prepend(1)
	listaDoble.Prepend(2)
	listaDoble.Prepend(3)
	listaDoble.Prepend(4)

	// Iterar la lista e imprimir los elementos
	fmt.Println("Iterando lista doble de enteros")
	i = listaDoble.Size() - 1
	for actual := listaDoble.Tail(); actual != nil; actual = actual.Prev() {
		fmt.Printf("Elemento %d: %d\n", i, actual.Data())
		i--
	}

	// Crear una lista circular de enteros
	listaCircular := list.NewCircularList[int]()

	// Insertar elementos en la lista
	listaCircular.Prepend(1)
	listaCircular.Prepend(2)
	listaCircular.Prepend(3)
	listaCircular.Prepend(4)

	// Iterar la lista e imprimir los elementos
	fmt.Println("Iterando lista circular de enteros")
	current := listaCircular.Head()
	for i := 0; i < listaCircular.Size(); i++ {
		fmt.Printf("Elemento %d: %d\n", i, current.Data())
		current = current.Next()
	}
}
