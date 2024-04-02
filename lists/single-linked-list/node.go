package single_linked_list

import "cmp"

type node[T cmp.Ordered] struct {
	data T
	next *node[T]
}

func newNode[T cmp.Ordered](data T) *node[T] {
	return &node[T]{data: data, next: nil}
}

func (n *node[T]) getData() T {
	return n.data
}

func (n *node[T]) setData(newData T) {
	n.data = newData
}

func (n *node[T]) getNext() *node[T] {
	return n.next
}

func (n *node[T]) setNext(newNext *node[T]) {
	if newNext != nil {
		n.next = newNext
	} else {
		n.next = nil
	}
}
