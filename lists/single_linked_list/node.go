package single_linked_list

import "cmp"

type Node[T cmp.Ordered] struct {
	data T
	next *Node[T]
}

func NewNode[T cmp.Ordered](data T) *Node[T] {
	return &Node[T]{data: data}
}

func (n *Node[T]) Data() T {
	return n.data
}

func (n *Node[T]) SetData(newData T) {
	n.data = newData
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

func (n *Node[T]) SetNext(newNext *Node[T]) {
	if newNext != nil {
		n.next = newNext
	} else {
		n.next = nil
	}
}
