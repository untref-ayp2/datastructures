// Package avltree implementa un árbol AVL.
package avltree

import (
	"errors"

	"github.com/untref-ayp2/data-structures/types"
)

type AVLTree[T types.Ordered] struct {
	root *AVLNode[T]
}

// NewAVLTree crea un árbol AVL vacío.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//
// Retorna:
//   - un puntero a un árbol AVL vacío.
func NewAVLTree[T types.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{root: nil}
}

// String devuelve una representación en cadena del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.String()
//
// Retorna:
//   - una representación en cadena del árbol AVL.
func (avl *AVLTree[T]) String() string {
	return avl.root.string()
}

// GetRoot devuelve la raíz del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.GetRoot()
//
// Retorna:
//   - la raíz del árbol AVL.
func (avl *AVLTree[T]) GetRoot() *AVLNode[T] {
	return avl.root
}

// Insert agrega un elemento al árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//
// Parámetros:
//   - `k` el elemento a agregar.
func (avl *AVLTree[T]) Insert(k T) {
	avl.root = avl.root.insert(k)
}

// Remove elimina un elemento del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.Remove(3)
//
// Parámetros:
//   - `k` el elemento a eliminar.
func (avl *AVLTree[T]) Remove(k T) {
	avl.root = avl.root.remove(k)
}

// Search comprueba si un elemento está en el árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.Search(3)
//
// Parámetros:
//   - `k` el elemento a buscar.
//
// Retorna:
//   - `true` si el elemento está en el árbol, `false` en caso contrario.
func (avl *AVLTree[T]) Search(k T) bool {
	return avl.root.search(k)
}

// FindMin devuelve el menor elemento del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.FindMin() // 3
//
// Retorna:
//   - el menor elemento del árbol AVL.
func (avl *AVLTree[T]) FindMin() (T, error) {
	if avl.root == nil {
		var emptyValue T

		return emptyValue, errors.New("árbol vacío")
	}

	minNode := avl.root.findMin()

	return minNode.GetData(), nil
}

// FindMax devuelve el mayor elemento del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.FindMax() // 7
//
// Retorna:
//   - el mayor elemento del árbol AVL.
func (avl *AVLTree[T]) FindMax() (T, error) {
	if avl.root == nil {
		var emptyValue T

		return emptyValue, errors.New("árbol vacío")
	}

	maxNode := avl.root.findMax()

	return maxNode.GetData(), nil
}

// GetHeight devuelve la altura del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.GetHeight() // 1
//
// Retorna:
//   - la altura del árbol AVL.
func (avl *AVLTree[T]) GetHeight() int {
	return avl.root.getHeight()
}

// GetBalance devuelve el balance del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.GetBalance() // 0
//
// Retorna:
//   - el balance del árbol AVL.
func (avl *AVLTree[T]) GetBalance() int {
	return avl.root.getBalance()
}

// IsEmpty comprueba si el árbol AVL está vacío.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.IsEmpty() // true
//
// Retorna:
//   - `true` si el árbol AVL está vacío, `false` en caso contrario.
func (avl *AVLTree[T]) IsEmpty() bool {
	return avl.root == nil
}

// Clear elimina todos los elementos del árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.Clear()
func (avl *AVLTree[T]) Clear() {
	avl.root = nil
}

// InOrder devuelve una representación en cadena del árbol AVL en InOrder.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	avl.Insert(5)
//	avl.Insert(3)
//	avl.Insert(7)
//	avl.InOrder() // "3 5 7"
//
// Retorna:
//   - una representación en cadena del árbol AVL en InOrder.
func (avl *AVLTree[T]) InOrder() string {
	return avl.root.inOrder()
}

// Iterator devuelve un iterador para recorrer el árbol AVL en orden.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	// ...
//	iterator := avl.Iterator()
//	for iterator.HasNext() {
//		fmt.Println(iterator.Next())
//	}
//
// Retorna:
//   - un iterador para recorrer el árbol AVL en orden.
func (avl *AVLTree[T]) Iterator() types.Iterator[T] {
	return NewAVLInOrderIterator[T](avl.root)
}
