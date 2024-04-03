package main

import (
	"fmt"

	cl "github.com/untref-ayp2/data-structures/lists/circular_list"
	dl "github.com/untref-ayp2/data-structures/lists/double_linked_list"
	sl "github.com/untref-ayp2/data-structures/lists/single_linked_list"
)

func main() {
	// Crear una lista enlazada simple de enteros

	listaEntero := sl.NewList[int]()
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

	listaDoble := dl.NewList[int]()
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

	listaCircular := cl.NewList[int]()
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
