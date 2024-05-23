package avltree

import (
	"errors"
	"github.com/untref-ayp2/data-structures/stack"
	"github.com/untref-ayp2/data-structures/types"
)

type AVLInOrderIterator[T types.Ordered] struct {
	stack *stack.Stack[AVLNode[T]] // pila de nodos
}

func NewAVLInOrderIterator[T types.Ordered](root *AVLNode[T]) *AVLInOrderIterator[T] {
	iterator := &AVLInOrderIterator[T]{
		stack: stack.NewStack[AVLNode[T]](),
	}
	iterator.stackLeftChildren(root)

	return iterator
}

func (it *AVLInOrderIterator[T]) stackLeftChildren(node *AVLNode[T]) {
	for node != nil {
		it.stack.Push(*node)
		node = node.getLeft()
	}
}

func (it *AVLInOrderIterator[T]) HasNext() bool {
	return !it.stack.IsEmpty()
}

func (it *AVLInOrderIterator[T]) Next() (T, error) {
	var data T
	if it.stack.IsEmpty() {
		return data, errors.New("no hay más elementos")
	}
	next, _ := it.stack.Pop()
	if next.getRight() != nil {
		it.stackLeftChildren(next.getRight())
	}

	return next.data, nil
}
