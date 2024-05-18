package binarytree

import (
	"math"

	"github.com/untref-ayp2/data-structures/types"
)

// BinaryNode implementa el tipo BinaryNode con dos campos de tipo
// *BinaryNode que son punteros a los hijos izquierdo y derecho, tambien de tipo BinaryNode,
// y un tercer campo de tipo T generico pero Ordered, por compatibilidad con BinarySerchTree
type BinaryNode[T types.Ordered] struct {
	left  *BinaryNode[T]
	right *BinaryNode[T]
	data  T
}

// NewBinaryNode crea un nuevo BinaryNode.
//
// Uso:
//
//	d := binarytree.NewBinarNode[int](data, h-izq, h-der)
//
// Parámetros:
// 'data' : el dato que guarda el nodo de tipo T
// 'left' : el nodo que será asignado como hijo izquierdo
// 'right' : el nodo que será asignado como hijo derecho
//
// Retorna un nuevo BinaryNode
func NewBinaryNode[T types.Ordered](
	data T,
	left *BinaryNode[T],
	right *BinaryNode[T],
) *BinaryNode[T] {
	return &BinaryNode[T]{left: left, right: right, data: data}
}

// Retorna el dato guardado en el nodo de tipo T
func (n *BinaryNode[T]) GetData() T {
	return n.data
}

// Retorna un slice de tipo T con los elementos del arbol en In-Order
func (n *BinaryNode[T]) GetInOrder() []T {
	s := make([]T, 0, n.Size())
	n.getInOrder(&s)
	return s
}

// funcion recursiva que agrega en el slice que recibe como parámetro
// los elementos del arbol en In-Order
// Parámetros: un puntero a un slice de tipo T
func (n *BinaryNode[T]) getInOrder(s *[]T) {
	if n.left != nil {
		n.left.getInOrder(s)
	}
	*s = append(*s, n.data)
	if n.right != nil {
		n.right.getInOrder(s)
	}
}

// Retorna un slice de tipo T con los elementos del arbol en Pre-Order
func (n *BinaryNode[T]) GetPreOrder() []T {
	s := make([]T, 0, n.Size())
	n.getPreOrder(&s)
	return s
}

// funcion recursiva que agrega en el slice que recibe como parámetro
// los elementos del arbol en Pre-Order
// Parámetros: un puntero a un slice de tipo T
func (n *BinaryNode[T]) getPreOrder(s *[]T) {
	*s = append(*s, n.data)
	if n.left != nil {
		n.left.getPreOrder(s)
	}
	if n.right != nil {
		n.right.getPreOrder(s)
	}
}

// Retorna un slice de tipo T con los elementos del arbol en Post-Order
func (n *BinaryNode[T]) GetPostOrder() []T {
	s := make([]T, 0, n.Size())
	n.getPostOrder(&s)
	return s
}

// funcion recursiva que agrega en el slice que recibe como parámetro
// los elementos del arbol en Post-Order
// Parámetros: un puntero a un slice de tipo T
func (n *BinaryNode[T]) getPostOrder(s *[]T) {
	if n.left != nil {
		n.left.getPostOrder(s)
	}
	if n.right != nil {
		n.right.getPostOrder(s)
	}
	*s = append(*s, n.data)
}

// Retorna la cantidad de nodos hacia abajo.
func (n *BinaryNode[T]) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return size
}

// Retorna la distancia al nodo mas profundo.
func (n *BinaryNode[T]) Height() int {
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.Height()
	}
	if n.right != nil {
		rightHeight = n.right.Height()
	}
	return int(1 + math.Max(float64(leftHeight), float64(rightHeight)))
}
