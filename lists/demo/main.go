package main

import (
	"fmt"

	single_linked_list "github.com/untref-ayp2/data-structures/lists/single-linked-list"
	double_linked_list "github.com/untref-ayp2/data-structures/lists/double-linked-list"
	lists "github.com/untref-ayp2/data-structures/lists"
)

func main() {
	// Crear una lista enlazada simple de enteros
	var listaEntero lists.List[int]
	listaEntero = single_linked_list.NewSingleLinkedList[int]()
	// Insertar elementos en la lista
	listaEntero.InsertAt(0, 1)
	listaEntero.InsertAt(0, 2)
	listaEntero.InsertAt(0, 3)
	listaEntero.InsertAt(0, 4)
	// Iterar la lista e imprimir los elementos
	fmt.Println("Iterando lista simple de enteros")
	i := 0
	for dato := range listaEntero.Iterate() {
		fmt.Printf("Elemento %d: %d\n", i, dato)
		i++
	}

	listaEntero.Clear()

	listaEntero = double_linked_list.NewDoubleLinkedList[int]()
	// Insertar elementos en la lista
	listaEntero.InsertAt(0, 1)
	listaEntero.InsertAt(0, 2)
	listaEntero.InsertAt(0, 3)
	listaEntero.InsertAt(0, 4)
	// Iterar la lista e imprimir los elementos
	fmt.Println("Iterando lista doble de enteros")
	i = 0
	for dato := range listaEntero.Iterate() {
		fmt.Printf("Elemento %d: %d\n", i, dato)
		i++
	}
}
