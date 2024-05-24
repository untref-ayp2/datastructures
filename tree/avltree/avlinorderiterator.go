package avltree

import (
	"errors"

	"github.com/untref-ayp2/data-structures/stack"
	"github.com/untref-ayp2/data-structures/types"
)

type AVLInOrderIterator[T types.Ordered] struct {
	stack *stack.Stack[AVLNode[T]] // pila de nodos
}

// NewAVLInOrderIterator crea un iterador en orden para un árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLInOrderIterator(avl.GetRoot())
//
// Parámetros:
//   - `root` la raíz del árbol AVL.
//
// Retorna:
//   - un iterador en orden para el árbol AVL.
func NewAVLInOrderIterator[T types.Ordered](root *AVLNode[T]) *AVLInOrderIterator[T] {
	iterator := &AVLInOrderIterator[T]{
		stack: stack.NewStack[AVLNode[T]](),
	}
	iterator.stackLeftChildren(root)

	return iterator
}

// stackLeftChildren apila los nodos hijos izquierdos de un nodo.
//
// Parámetros:
//   - `node` el nodo a partir del cual se apilarán los nodos hijos izquierdos.
func (it *AVLInOrderIterator[T]) stackLeftChildren(node *AVLNode[T]) {
	for node != nil {
		it.stack.Push(*node)
		node = node.getLeft()
	}
}

// HasNext indica si hay más elementos para recorrer.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLInOrderIterator(avl.GetRoot())
//	iterator.HasNext()
//
// Retorna:
//   - `true` si hay más elementos para recorrer, `false` en caso contrario.
func (it *AVLInOrderIterator[T]) HasNext() bool {
	return !it.stack.IsEmpty()
}

// Next devuelve el siguiente elemento del recorrido.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLInOrderIterator(avl.GetRoot())
//	iterator.Next()
//
// Retorna:
//   - el siguiente elemento del recorrido.
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
