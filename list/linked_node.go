package list

// LinkedNode representa un nodo de una lista enlazada simple.
type LinkedNode[T comparable] struct {
	data T
	next *LinkedNode[T]
}

// NewLinkedListNode crea un nuevo nodo de lista enlazada con el dato especificado.
//
// Uso:
//
//	node := list.NewLinkedListNode(10) // Crea un nuevo nodo con el dato 10.
//
// Parámetros:
//   - `data`: el dato a almacenar en el nodo.
func NewLinkedListNode[T comparable](data T) *LinkedNode[T] {
	return &LinkedNode[T]{data: data}
}

// SetData establece el dato almacenado en el nodo.
//
// Uso:
//
//	node.SetData(20) // Establece el dato del nodo a 20.
//
// Parámetros:
//   - `data`: el dato a almacenar en el nodo.
func (n *LinkedNode[T]) SetData(data T) {
	n.data = data
}

// Data devuelve el dato almacenado en el nodo.
//
// Uso:
//
//	data := node.Data() // Obtiene el dato almacenado en el nodo.
//
// Retorna:
//   - el dato almacenado en el nodo.
func (n *LinkedNode[T]) Data() T {
	return n.data
}

// SetNext establece el nodo siguiente al nodo actual.
//
// Uso:
//
//	node.SetNext(newNode) // Establece el nodo siguiente al nodo actual.
//
// Parámetros:
//   - `newNext`: el nodo siguiente al nodo actual.
func (n *LinkedNode[T]) SetNext(newNext *LinkedNode[T]) {
	n.next = newNext
}

// Next devuelve el nodo siguiente al nodo actual.
//
// Uso:
//
//	nextNode := node.Next() // Obtiene el nodo siguiente al nodo actual.
func (n *LinkedNode[T]) Next() *LinkedNode[T] {
	return n.next
}

// HasNext evalúa si el nodo actual tiene asignado un nodo siguiente.
//
// Uso:
//
//	if node.HasNext() {
//		fmt.Println("El nodo tiene un nodo siguiente.")
//	}
//
// Retorna:
//   - `true` si el nodo tiene un nodo siguiente; `false` en caso contrario.
func (n *LinkedNode[T]) HasNext() bool {
	return n.next != nil
}
