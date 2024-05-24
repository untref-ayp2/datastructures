package avltree

import (
	"errors"

	"github.com/untref-ayp2/data-structures/types"
)

type AVLTree[T types.Ordered] struct {
	root *AVLNode[T]
}

// Constructor. Devuelve un árbol AVL con la raíz inicializada en nil.
func NewAVLTree[T types.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{root: nil}
}

func (avl *AVLTree[T]) String() string {
	return avl.root.string()
}

func (avl *AVLTree[T]) GetRoot() *AVLNode[T] {
	return avl.root
}

func (avl *AVLTree[T]) Insert(k T) {
	avl.root = avl.root.insert(k)
}

func (avl *AVLTree[T]) Remove(k T) {
	avl.root = avl.root.remove(k)
}

func (avl *AVLTree[T]) Search(k T) bool {
	return avl.root.search(k)
}

func (avl *AVLTree[T]) FindMin() (T, error) {
	if avl.root == nil {
		var emptyValue T

		return emptyValue, errors.New("árbol vacío")
	}

	minNode := avl.root.findMin()

	return minNode.GetData(), nil
}

func (avl *AVLTree[T]) FindMax() (T, error) {
	if avl.root == nil {
		var emptyValue T

		return emptyValue, errors.New("árbol vacío")
	}

	maxNode := avl.root.findMax()

	return maxNode.GetData(), nil
}

func (avl *AVLTree[T]) GetHeight() int {
	return avl.root.getHeight()
}

func (avl *AVLTree[T]) GetBalance() int {
	return avl.root.getBalance()
}

func (avl *AVLTree[T]) IsEmpty() bool {
	return avl.root == nil
}

func (avl *AVLTree[T]) Clear() {
	avl.root = nil
}

func (avl *AVLTree[T]) InOrder() string {
	return avl.root.inOrder()
}

func (avl *AVLTree[T]) Iterator() types.Iterator[T] {
	return NewAVLInOrderIterator[T](avl.root)
}
