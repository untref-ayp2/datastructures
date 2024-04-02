package double_linked_list

import "cmp"

// node representa un nodo de una lista enlazada doble.
type node[T cmp.Ordered] struct {
	data T
	next *node[T]
	prev *node[T]
}

// NewNode crea un nuevo nodo de lista enlazada doble con el dato especificado.
func newNode[T cmp.Ordered](data T) *node[T] {
	return &node[T]{data: data}
}

// getData devuelve el dato almacenado en el nodo.
func (n *node[T]) getData() T {
	return n.data
}

// setData establece el dato almacenado en el nodo.
func (n *node[T]) setData(newData T) {
	n.data = newData
}

// getNext devuelve el nodo siguiente al nodo actual.
func (n *node[T]) getNext() *node[T] {
	return n.next
}

// setNext establece el nodo siguiente al nodo actual.
func (n *node[T]) setNext(newNext *node[T]) {
	if newNext != nil {
		n.next = newNext
	} else {
		n.next = nil
	}
}

// getPrev devuelve el nodo anterior al nodo actual.
func (n *node[T]) getPrev() *node[T] {
	return n.prev
}

// setPrev establece el nodo anterior al nodo actual.
func (n *node[T]) setPrev(newPrev *node[T]) {
	if newPrev != nil {
		n.prev = newPrev
	} else {
		n.prev = nil
	}
}