package list

// DoubleLinkedList implementa una lista enlazada doble genérica.
type DoubleLinkedList[T comparable] struct {
	head *DoubleLinkedNode[T]
	tail *DoubleLinkedNode[T]
	size int
}

// NewDoubleLinkedList crea una nueva lista vacía.
func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

// Head devuelve el primer nodo de la lista.
func (l *DoubleLinkedList[T]) Head() *DoubleLinkedNode[T] {
	return l.head
}

// Tail devuelve el último nodo de la lista.
func (l *DoubleLinkedList[T]) Tail() *DoubleLinkedNode[T] {
	return l.tail
}

// Size devuelve el tamaño de la lista.
func (l *DoubleLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty devuelve true si la lista está vacía.
func (l *DoubleLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear elimina todos los nodos de la lista.
func (l *DoubleLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Prepend inserta un dato al inicio de la lista.
func (l *DoubleLinkedList[T]) Prepend(data T) {
	newNode := NewDoubleLinkedListNode[T](data)
	if l.size == 0 {
		l.tail = newNode
	} else {
		newNode.SetNext(l.head)
		l.head.SetPrev(newNode)
	}
	l.head = newNode
	l.size++
}

// Append inserta un dato al final de la lista.
func (l *DoubleLinkedList[T]) Append(data T) {
	newNode := NewDoubleLinkedListNode[T](data)
	if l.size == 0 {
		l.head = newNode
	} else {
		l.tail.SetNext(newNode)
		newNode.SetPrev(l.tail)
	}
	l.tail = newNode
	l.size++
}

// Find busca un dato en la lista, si lo encuentra devuelve el nodo
// correspondiente, si no lo encuentra devuelve nil
func (l *DoubleLinkedList[T]) Find(data T) *DoubleLinkedNode[T] {
	for current := l.head; current != nil; current = current.Next() {
		if current.Data() == data {
			return current
		}
	}
	return nil
}

// RemoveFirst elimina el primer nodo de la lista.
func (l *DoubleLinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.head = l.head.Next()
	l.size--

	if l.IsEmpty() {
		l.tail = nil
	} else {
		l.head.SetPrev(nil)
	}
}

// RemoveLast elimina el último nodo de la lista.
func (l *DoubleLinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}

	l.tail = l.tail.Prev()
	l.tail.SetNext(nil)
	l.size--
}

// Remove elimina un la primera aparición de un dato en la lista.
func (l *DoubleLinkedList[T]) Remove(data T) {
	node := l.Find(data)

	if node == nil {
		return
	}

	if node == l.head {
		l.RemoveFirst()
		return
	}

	if node == l.tail {
		l.RemoveLast()
		return
	}

	node.Prev().SetNext(node.Next())
	node.Next().SetPrev(node.Prev())
	l.size--
}
