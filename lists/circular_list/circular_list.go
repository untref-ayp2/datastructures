package circular_list

import "cmp"

// CircularList implementa una lista enlazada circular genérica.
type CircularList[T cmp.Ordered] struct {
	head *Node[T]
	size int
}

// NewList crea una nueva lista enlazada circular.
func NewList[T cmp.Ordered]() *CircularList[T] {
	return &CircularList[T]{}
}

func (cl *CircularList[T]) Head() *Node[T] {
	return cl.head
}

func (cl *CircularList[T]) Tail() *Node[T] {
	if cl.size == 0 {
		return nil
	} else {
		return cl.head.Prev()
	}
}

func (cl *CircularList[T]) Size() int {
	return cl.size
}

func (cl *CircularList[T]) IsEmpty() bool {
	return cl.size == 0
}

func (cl *CircularList[T]) Clear() {
	cl.head = nil
	cl.size = 0
}

func (cl *CircularList[T]) Prepend(data T) {
	node := NewNode(data)
	if cl.size == 0 {
		cl.head = node
		cl.head.SetNext(cl.head)
		cl.head.SetPrev(cl.head)
	} else {
		node.SetNext(cl.head)
		node.SetPrev(cl.head.Prev())
		cl.head.Prev().SetNext(node)
		cl.head.SetPrev(node)
		cl.head = node
	}
	cl.size++
}

func (cl *CircularList[T]) Append(data T) {
	node := NewNode(data)
	if cl.size == 0 {
		cl.head = node
		cl.head.SetNext(cl.head)
		cl.head.SetPrev(cl.head)
	} else {
		node.SetNext(cl.head)
		node.SetPrev(cl.head.Prev())
		cl.head.Prev().SetNext(node)
		cl.head.SetPrev(node)
	}
	cl.size++
}

func (cl *CircularList[T]) Find(data T) *Node[T] {
	if cl.size == 0 {
		return nil
	}

	node := cl.head
	for i := 0; i < cl.size; i++ {
		if node.Data() == data {
			return node
		}
		node = node.Next()
	}

	return nil
}

func (cl *CircularList[T]) RemoveFirst() {
	if cl.size == 0 {
		return
	}

	if cl.size == 1 {
		cl.head = nil
		cl.size--
		return
	}

	cl.head.Prev().SetNext(cl.head.Next())
	cl.head.Next().SetPrev(cl.head.Prev())
	cl.head = cl.head.Next()
	cl.size--
}

func (cl *CircularList[T]) RemoveLast() {
	if cl.size == 0 {
		return
	}

	if cl.size == 1 {
		cl.head = nil
		cl.size--
		return
	}

	cl.head.Prev().Prev().SetNext(cl.head)
	cl.head.SetPrev(cl.head.Prev().Prev())
	cl.size--
}

func (cl *CircularList[T]) Remove(data T) {
	node := cl.Find(data)
	if node == nil {
		return
	}

	if node == cl.head {
		cl.RemoveFirst()
		return
	}

	node.Prev().SetNext(node.Next())
	node.Next().SetPrev(node.Prev())
	cl.size--
}
