package list

// LinkedNode representa un nodo de una lista enlazada simple.
type LinkedNode[T comparable] struct {
	data T
	next *LinkedNode[T]
}

// NewDoubleLinkedListNode crea un nuevo nodo de lista enlazada doble con el dato especificado.
func NewLinkedListNode[T comparable](data T) *LinkedNode[T] {
	return &LinkedNode[T]{data: data}
}

// SetData establece el dato almacenado en el nodo.
func (n *LinkedNode[T]) SetData(data T) {
	n.data = data
}

// Data devuelve el dato almacenado en el nodo.
func (n *LinkedNode[T]) Data() T {
	return n.data
}

// SetNext establece el nodo siguiente al nodo actual.
func (n *LinkedNode[T]) SetNext(newNext *LinkedNode[T]) {
	n.next = newNext
}

// Next devuelve el nodo siguiente al nodo actual.
func (n *LinkedNode[T]) Next() *LinkedNode[T] {
	return n.next
}

// HasNext devuelve true si el nodo actual tiene asignado un nodo siguiente.
func (n *LinkedNode[T]) HasNext() bool {
	return n.next != nil
}
