package binarytree

import (
	"fmt"
	"math"
	"strings"

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

func (n *BinaryNode[T]) GetData() T {
	return n.data
}

// Muestra por consola su recorrido en Pre-Order.
func (n *BinaryNode[T]) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

// Retorna un string con el recorrido en Pre-Order
func (n *BinaryNode[T]) StringPreOrder() string {
	sb := strings.Builder{}
	n.stringPreOrder(&sb)
	return sb.String()
}

// funcion recursiva que concatena en un String Builder el recorrido en Pre-Order
// desde el nodo.
// Parámetros: un puntero a un objeto StringsBuilder
// Retorna: un string
func (n *BinaryNode[T]) stringPreOrder(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%v", n.data))
	if n.left != nil {
		n.left.stringPreOrder(sb)
	}
	if n.right != nil {
		n.right.stringPreOrder(sb)
	}
}

// Muestra por consola su recorrido en In-Order.
func (n *BinaryNode[T]) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

// Retorna un string con el recorrido en In-Order
func (n *BinaryNode[T]) StringInOrder() string {
	sb := strings.Builder{}
	n.stringInOrder(&sb)
	return sb.String()
}

// funcion recursiva que concatena en un String Builder el recorrido en In-Order
// desde el nodo.
// Parámetros: un puntero a un objeto StringsBuilder
// Retorna: un string
func (n *BinaryNode[T]) stringInOrder(sb *strings.Builder) {
	if n.left != nil {
		n.left.stringInOrder(sb)
	}
	sb.WriteString(fmt.Sprintf("%v", n.data))
	if n.right != nil {
		n.right.stringInOrder(sb)
	}
}

// Muestra por consola su recorrido en Post-Order.
func (n *BinaryNode[T]) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

// Retorna un string con el recorrido en Post-Order
func (n *BinaryNode[T]) StringPostOrder() string {
	sb := strings.Builder{}
	n.stringPostOrder(&sb)
	return sb.String()
}

// funcion recursiva que concatena en un String Builder el recorrido en Post-Order
// desde el nodo.
// Parámetros: un puntero a un objeto StringsBuilder
// Retorna: un string
func (n *BinaryNode[T]) stringPostOrder(sb *strings.Builder) {
	if n.left != nil {
		n.left.stringPostOrder(sb)
	}
	if n.right != nil {
		n.right.stringPostOrder(sb)
	}
	sb.WriteString(fmt.Sprintf("%v", n.data))
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
