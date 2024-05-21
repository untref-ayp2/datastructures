package tree

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
	iterator.apilarHijosIzquierdos(root)

	return iterator
}

func (it *AVLInOrderIterator[T]) apilarHijosIzquierdos(nodo *AVLNode[T]) {
	for nodo != nil {
		it.stack.Push(*nodo)
		nodo = nodo.getLeft()
	}
}

func (it *AVLInOrderIterator[T]) HasNext() bool {
	return !it.stack.IsEmpty()
}

func (it *AVLInOrderIterator[T]) Next() (AVLNode[T], error) {
	var dato T
	nodo := newAVLNode[T](dato, nil, nil)
	if it.stack.IsEmpty() {
		return *nodo, errors.New("no hay más elementos")
	}
	next, _ := it.stack.Pop()
	if next.getRight() != nil {
		it.apilarHijosIzquierdos(next.getRight())
	}

	return next, nil
}
