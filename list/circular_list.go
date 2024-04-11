package list

// CircularList implementa una lista enlazada circular genérica.
type CircularList[T comparable] struct {
	head *DoubleLinkedNode[T]
	size int
}

// NewCircularList crea una nueva lista enlazada circular.
func NewCircularList[T comparable]() *CircularList[T] {
	return &CircularList[T]{}
}

func (l *CircularList[T]) Head() *DoubleLinkedNode[T] {
	return l.head
}

func (l *CircularList[T]) Tail() *DoubleLinkedNode[T] {
	if l.size == 0 {
		return nil
	} else {
		return l.head.Prev()
	}
}

func (l *CircularList[T]) Size() int {
	return l.size
}

func (l *CircularList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *CircularList[T]) Clear() {
	l.head = nil
	l.size = 0
}

func (l *CircularList[T]) Prepend(data T) {
	node := NewDoubleLinkedListNode(data)
	if l.size == 0 {
		l.head = node
		l.head.SetNext(l.head)
		l.head.SetPrev(l.head)
	} else {
		node.SetNext(l.head)
		node.SetPrev(l.head.Prev())
		l.head.Prev().SetNext(node)
		l.head.SetPrev(node)
		l.head = node
	}
	l.size++
}

func (l *CircularList[T]) Append(data T) {
	node := NewDoubleLinkedListNode(data)
	if l.size == 0 {
		l.head = node
		l.head.SetNext(l.head)
		l.head.SetPrev(l.head)
	} else {
		node.SetNext(l.head)
		node.SetPrev(l.head.Prev())
		l.head.Prev().SetNext(node)
		l.head.SetPrev(node)
	}
	l.size++
}

func (l *CircularList[T]) Find(data T) *DoubleLinkedNode[T] {
	if l.size == 0 {
		return nil
	}

	node := l.head
	for i := 0; i < l.size; i++ {
		if node.Data() == data {
			return node
		}
		node = node.Next()
	}

	return nil
}

func (l *CircularList[T]) RemoveFirst() {
	if l.size == 0 {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.size--
		return
	}

	l.head.Prev().SetNext(l.head.Next())
	l.head.Next().SetPrev(l.head.Prev())
	l.head = l.head.Next()
	l.size--
}

func (l *CircularList[T]) RemoveLast() {
	if l.size == 0 {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.size--
		return
	}

	l.head.Prev().Prev().SetNext(l.head)
	l.head.SetPrev(l.head.Prev().Prev())
	l.size--
}

func (l *CircularList[T]) Remove(data T) {
	node := l.Find(data)
	if node == nil {
		return
	}

	if node == l.head {
		l.RemoveFirst()
		return
	}

	node.Prev().SetNext(node.Next())
	node.Next().SetPrev(node.Prev())
	l.size--
}
