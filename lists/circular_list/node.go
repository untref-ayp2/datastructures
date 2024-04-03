package circular_list

import "github.com/untref-ayp2/data-structures/types"

// Node representa un nodo de una lista enlazada doble.
// La lista circular será doblemente enlazada.
type Node[T types.Ordered] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

// NewNode crea un nuevo nodo de lista enlazada doble con el dato especificado.
func NewNode[T types.Ordered](data T) *Node[T] {
	return &Node[T]{data: data}
}

// Data devuelve el dato almacenado en el nodo.
func (n *Node[T]) Data() T {
	return n.data
}

// setData establece el dato almacenado en el nodo.
func (n *Node[T]) SetData(newData T) {
	n.data = newData
}

// Next devuelve el nodo siguiente al nodo actual.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// setNext establece el nodo siguiente al nodo actual.
func (n *Node[T]) SetNext(newNext *Node[T]) {
	n.next = newNext
}

// Prev devuelve el nodo anterior al nodo actual.
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// setPrev establece el nodo anterior al nodo actual.
func (n *Node[T]) SetPrev(newPrev *Node[T]) {
	n.prev = newPrev
}

func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

func (n *Node[T]) HasPrev() bool {
	return n.prev != nil
}
