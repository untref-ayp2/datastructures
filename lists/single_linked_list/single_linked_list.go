package single_linked_list

import "github.com/untref-ayp2/data-structures/types"

// Lista enlazada simple
// Se implementa con un nodo que contiene un dato y un puntero al siguiente nodo
type SingleLinkedList[T types.Ordered] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// Crea una nueva lista vacía
func NewList[T types.Ordered]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

// Devuelve el primer nodo de la lista
func (sll *SingleLinkedList[T]) Head() *Node[T] {
	return sll.head
}

// Devuelve el último nodo de la lista
func (sll *SingleLinkedList[T]) Tail() *Node[T] {
	return sll.tail
}

// Devuelve el tamaño de la lista
func (sll *SingleLinkedList[T]) Size() int {
	return sll.size
}

// Devuelve True si la lista está vacía
func (sll *SingleLinkedList[T]) IsEmpty() bool {
	return sll.size == 0
}

// Elimina todos los nodos de la lista
func (ssl *SingleLinkedList[T]) Clear() {
	ssl.head = nil
	ssl.tail = nil
	ssl.size = 0
}

// Inserta un dato al inicio de la lista
func (sll *SingleLinkedList[T]) Prepend(data T) {
	newNode := NewNode(data)
	if sll.IsEmpty() {
		sll.head = newNode
		sll.tail = newNode
	} else {
		newNode.SetNext(sll.head)
		sll.head = newNode
	}
	sll.size++
}

// Inserta un dato al final de la lista
func (sll *SingleLinkedList[T]) Append(data T) {
	newNode := NewNode(data)
	if sll.IsEmpty() {
		sll.head = newNode
		sll.tail = newNode
	} else {
		sll.tail.SetNext(newNode)
		sll.tail = newNode
	}
	sll.size++
}

// Busca un dato en la lista, si lo encuentra devuelve el nodo correspondiente,
// si no lo encuentra devuelve nil
func (sll *SingleLinkedList[T]) Find(data T) *Node[T] {
	current := sll.head
	for current != nil {
		if current.Data() == data {
			return current
		}
		current = current.Next()
	}
	return nil
}

// Elimina el primer nodo de la lista
func (sll *SingleLinkedList[T]) RemoveFirst() {
	if !sll.IsEmpty() {
		sll.head = sll.head.Next()
		sll.size--
		if sll.Head() == nil {
			sll.tail = nil
		}
	}
}

// Elimina el último nodo de la lista
func (sll *SingleLinkedList[T]) RemoveLast() {
	if !sll.IsEmpty() {
		if sll.Size() == 1 {
			sll.head = nil
			sll.tail = nil
		} else {
			current := sll.head
			for current.Next() != sll.tail {
				current = current.Next()
			}
			current.SetNext(nil)
			sll.tail = current
		}
		sll.size--
	}
}

// Elimina un la primera aparición de un dato en la lista
func (sll *SingleLinkedList[T]) Remove(data T) {
	nodo := sll.Find(data)
	if nodo != nil {
		if nodo != sll.Head() {
			current := sll.Head()
			for current.Next() != nodo {
				current = current.Next()
			}
			current.SetNext(nodo.Next())
			if nodo == sll.tail {
				sll.tail = current
			}
			sll.size--
		} else {
			sll.RemoveFirst()
		}
	}
}
