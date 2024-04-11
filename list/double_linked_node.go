package list

// DoubleLinkedNode representa un nodo de una lista enlazada doble.
type DoubleLinkedNode[T comparable] struct {
	data T
	next *DoubleLinkedNode[T]
	prev *DoubleLinkedNode[T]
}

// NewDoubleLinkedListNode crea un nuevo nodo de lista enlazada doble con el dato especificado.
func NewDoubleLinkedListNode[T comparable](data T) *DoubleLinkedNode[T] {
	return &DoubleLinkedNode[T]{data: data}
}

// SetData establece el dato almacenado en el nodo.
func (n *DoubleLinkedNode[T]) SetData(data T) {
	n.data = data
}

// Data devuelve el dato almacenado en el nodo.
func (n *DoubleLinkedNode[T]) Data() T {
	return n.data
}

// SetNext establece el nodo siguiente al nodo actual.
func (n *DoubleLinkedNode[T]) SetNext(next *DoubleLinkedNode[T]) {
	n.next = next
}

// Next devuelve el nodo siguiente al nodo actual.
func (n *DoubleLinkedNode[T]) Next() *DoubleLinkedNode[T] {
	return n.next
}

// HasNext devuelve true si el nodo actual tiene asignado un nodo siguiente.
func (n *DoubleLinkedNode[T]) HasNext() bool {
	return n.next != nil
}

// setPrev establece el nodo anterior al nodo actual.
func (n *DoubleLinkedNode[T]) SetPrev(newPrev *DoubleLinkedNode[T]) {
	n.prev = newPrev
}

// Prev devuelve el nodo anterior al nodo actual.
func (n *DoubleLinkedNode[T]) Prev() *DoubleLinkedNode[T] {
	return n.prev
}

// HasPrev devuelve true si el nodo actual tiene asignado un nodo previo.
func (n *DoubleLinkedNode[T]) HasPrev() bool {
	return n.prev != nil
}
