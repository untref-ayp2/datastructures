package avltree

import (
	"errors"

	"github.com/untref-ayp2/data-structures/queue"
	"github.com/untref-ayp2/data-structures/types"
)

type AVLLevelIterator[T types.Ordered] struct {
	queue *queue.Queue[AVLNode[T]] // cola de nodos
}

// NewAVLLevelIterator crea un iterador de nivel para un árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLLevelIterator(avl.GetRoot())
//
// Parámetros:
//   - `root` la raíz del árbol AVL.
//
// Retorna:
//   - un iterador de nivel para el árbol AVL.
func NewAVLLevelIterator[T types.Ordered](root *AVLNode[T]) *AVLLevelIterator[T] {
	iterator := &AVLLevelIterator[T]{
		queue: queue.NewQueue[AVLNode[T]](),
	}
	if root != nil {
		iterator.queue.Enqueue(*root)
	}

	return iterator
}

// HasNext indica si hay más elementos para recorrer.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLLevelIterator(avl.GetRoot())
//	iterator.HasNext()
//
// Retorna:
//   - `true` si hay más elementos para recorrer, `false` en caso contrario.
func (it *AVLLevelIterator[T]) HasNext() bool {
	return !it.queue.IsEmpty()
}

// Next devuelve el siguiente elemento del recorrido.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLLevelIterator(avl.GetRoot())
//	iterator.Next()
//
// Retorna:
//   - el siguiente elemento del recorrido.
func (it *AVLLevelIterator[T]) Next() (T, error) {
	var data T
	if it.queue.IsEmpty() {
		return data, errors.New("no hay más elementos")
	}
	next, _ := it.queue.Dequeue()
	if next.getLeft() != nil {
		it.queue.Enqueue(*next.getLeft())
	}
	if next.getRight() != nil {
		it.queue.Enqueue(*next.getRight())
	}

	return next.data, nil
}
