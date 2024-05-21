package tree

import (
	"fmt"
	"math"
	"strings"

	"github.com/untref-ayp2/data-structures/types"
)

// BinaryNode implementa el tipo BinaryNode con dos campos de tipo
// *BinaryNode que son punteros a los hijos izquierdo y derecho, tambien de tipo BinaryNode,
// y un tercer campo de tipo T generico pero Ordered, por compatibilidad con BinarySerchTree.
type BinaryNode[T types.Ordered] struct {
	left  *BinaryNode[T]
	right *BinaryNode[T]
	data  T
}

// NewBinaryNode crea un nuevo BinaryNode.
//
// Uso:
//
//	d := binarytree.NewBinarNode[int](data, hIzq, hDer)
//
// Parámetros:
//   - 'data' : el dato que guarda el nodo de tipo T
//   - 'left' : el nodo que será asignado como hijo izquierdo
//   - 'right' : el nodo que será asignado como hijo derecho
//
// Retorna:
//   - un puntero a un nuevo BinaryNode.
func NewBinaryNode[T types.Ordered](
	data T,
	left *BinaryNode[T],
	right *BinaryNode[T],
) *BinaryNode[T] {
	return &BinaryNode[T]{left: left, right: right, data: data}
}

// Retorna el dato guardado en el nodo de tipo T.
//
// Retorna:
//   - el dato guardado en el nodo de tipo T.
func (n *BinaryNode[T]) GetData() T {
	return n.data
}

// Retorna el hijo izquierdo del nodo.
//
// Retorna:
//   - un puntero al hijo izquierdo del nodo.
func (n *BinaryNode[T]) GetLeft() *BinaryNode[T] {
	return n.left
}

// Retorna el hijo derecho del nodo.
//
// Retorna:
//   - un puntero al hijo derecho del nodo.
func (n *BinaryNode[T]) GetRight() *BinaryNode[T] {
	return n.right
}

// Retorna un los elementos del árbol en In-Order.
//
// Retorna:
//   - un slice de tipo T con los elementos del árbol en In-Order.
func (n *BinaryNode[T]) GetInOrder() []T {
	s := make([]T, 0, n.Size())
	n.getInOrder(&s)

	return s
}

// Función recursiva que agrega en el slice que recibe como parámetro
// los elementos del árbol en In-Order.
//
// Parámetros:
//   - 's' un puntero a un slice de tipo T.
func (n *BinaryNode[T]) getInOrder(s *[]T) {
	if n.left != nil {
		n.left.getInOrder(s)
	}
	*s = append(*s, n.data)
	if n.right != nil {
		n.right.getInOrder(s)
	}
}

// Retorna los elementos del árbol en Pre-Order.
//
// Retorna:
//   - un slice de tipo T con los elementos del árbol en Pre-Order.
func (n *BinaryNode[T]) GetPreOrder() []T {
	s := make([]T, 0, n.Size())
	n.getPreOrder(&s)

	return s
}

// Función recursiva que agrega en el slice que recibe como parámetro
// los elementos del árbol en Pre-Order
//
// Parámetros:
//   - 's' un puntero a un slice de tipo T.
func (n *BinaryNode[T]) getPreOrder(s *[]T) {
	*s = append(*s, n.data)
	if n.left != nil {
		n.left.getPreOrder(s)
	}
	if n.right != nil {
		n.right.getPreOrder(s)
	}
}

// Retorna los elementos del árbol en Post-Order.
//
// Retorna:
//   - un slice de tipo T con los elementos del árbol en Post-Order.
func (n *BinaryNode[T]) GetPostOrder() []T {
	s := make([]T, 0, n.Size())
	n.getPostOrder(&s)

	return s
}

// Función recursiva que agrega en el slice que recibe como parámetro
// los elementos del árbol en Post-Order
//
// Parámetros:
//   - 's' un puntero a un slice de tipo T.
func (n *BinaryNode[T]) getPostOrder(s *[]T) {
	if n.left != nil {
		n.left.getPostOrder(s)
	}
	if n.right != nil {
		n.right.getPostOrder(s)
	}
	*s = append(*s, n.data)
}

// Representación en String recorriendo el árbol en Pre-Order.
//
// Retorna:
//   - un string con el recorrido en Pre-Order.
func (n *BinaryNode[T]) StringPreOrder() string {
	return n.stringPreOrder()
}

// Función recursiva que concatena en un String Builder el recorrido en Pre-Order
// desde el nodo.
//
// Retorna:
//   - el string con el recorrido en Pre-Order.
func (n *BinaryNode[T]) stringPreOrder() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%v", n.data))
	if n.left != nil {
		sb.WriteString(n.left.stringPreOrder())
	}
	if n.right != nil {
		sb.WriteString(n.right.stringPreOrder())
	}

	return sb.String()
}

// Representación en String recorriendo el árbol en In-Order.
//
// Retorna:
//   - un string con el recorrido en In-Order.
func (n *BinaryNode[T]) StringInOrder() string {
	return n.stringInOrder()
}

// Función recursiva que concatena en un String Builder el recorrido en In-Order
// desde el nodo.
//
// Retorna:
//   - el string con el recorrido en In-Order.
func (n *BinaryNode[T]) stringInOrder() string {
	sb := strings.Builder{}
	if n.left != nil {
		sb.WriteString(n.left.stringInOrder())
	}
	sb.WriteString(fmt.Sprintf("%v", n.data))
	if n.right != nil {
		sb.WriteString(n.right.stringInOrder())
	}

	return sb.String()
}

// Representación en String recorriendo el árbol en Post-Order.
//
// Retorna:
//   - un string con el recorrido en Post-Order.
func (n *BinaryNode[T]) StringPostOrder() string {
	return n.stringPostOrder()
}

// funcion recursiva que concatena en un String Builder el recorrido en Post-Order
// desde el nodo.
//
// Parámetros:
//   - 'sb' un puntero a un objeto StringsBuilder
//
// Retorna: un string.
func (n *BinaryNode[T]) stringPostOrder() string {
	sb := strings.Builder{}
	if n.left != nil {
		sb.WriteString(n.left.stringPostOrder())
	}
	if n.right != nil {
		sb.WriteString(n.right.stringPostOrder())
	}
	sb.WriteString(fmt.Sprintf("%v", n.data))

	return sb.String()
}

// Evalúa el tamaño del árbol.
//
// Retorna:
//   - la cantidad de nodos del árbol
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

// Retorna la altura del árbol.
//
// Retorna:
//   - la altura del árbol.
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
