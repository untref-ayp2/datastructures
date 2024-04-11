package list

// LinkedList se implementa con un nodo que contiene un dato y un puntero al siguiente nodo.
type LinkedList[T comparable] struct {
	head *LinkedNode[T]
	tail *LinkedNode[T]
	size int
}

// NewLinkedList crea una nueva lista vacía.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Head devuelve el primer nodo de la lista.
func (l *LinkedList[T]) Head() *LinkedNode[T] {
	return l.head
}

// Tail devuelve el último nodo de la lista.
func (l *LinkedList[T]) Tail() *LinkedNode[T] {
	return l.tail
}

// Size devuelve el tamaño de la lista.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// IsEmpty devuelve true si la lista está vacía.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear elimina todos los nodos de la lista.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Prepend inserta un dato al inicio de la lista.
func (l *LinkedList[T]) Prepend(data T) {
	newNode := NewLinkedListNode(data)
	if l.IsEmpty() {
		l.tail = newNode
	} else {
		newNode.SetNext(l.head)
	}
	l.head = newNode
	l.size++
}

// Append inserta un dato al final de la lista.
func (l *LinkedList[T]) Append(data T) {
	newNode := NewLinkedListNode(data)
	if l.IsEmpty() {
		l.head = newNode
	} else {
		l.tail.SetNext(newNode)
	}
	l.tail = newNode
	l.size++
}

// Find busca un dato en la lista, si lo encuentra devuelve el nodo
// correspondiente, si no lo encuentra devuelve nil
func (l *LinkedList[T]) Find(data T) *LinkedNode[T] {
	for current := l.head; current != nil; current = current.Next() {
		if current.Data() == data {
			return current
		}
	}
	return nil
}

// RemoveFirst elimina el primer nodo de la lista.
func (l *LinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.head = l.head.Next()

	if l.head == nil {
		l.tail = nil
	}

	l.size--
}

// RemoveLast elimina el último nodo de la lista.
func (l *LinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		current := l.head
		for current.Next() != l.tail {
			current = current.Next()
		}
		current.SetNext(nil)
		l.tail = current
	}
	l.size--
}

// Remove elimina un la primera aparición de un dato en la lista.
func (l *LinkedList[T]) Remove(data T) {
	node := l.Find(data)

	if node == nil {
		return
	}

	if node == l.head {
		l.RemoveFirst()
		return
	}

	current := l.Head()

	for current.Next() != node {
		current = current.Next()
	}

	current.SetNext(node.Next())

	if node == l.tail {
		l.tail = current
	}
	l.size--
}
