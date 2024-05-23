package avltree

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/types"
)

// Nodo del árbol AVL, además del dato y los hijos, registra la altura.
type AVLNode[T types.Ordered] struct {
	data   T           // dato
	height int         // altura
	left   *AVLNode[T] // hijo izquierdo
	right  *AVLNode[T] // hijo derecho
}

func newAVLNode[T types.Ordered](data T, left *AVLNode[T], right *AVLNode[T]) *AVLNode[T] {
	return &AVLNode[T]{left: left, right: right, data: data, height: 0}
}

func (n *AVLNode[T]) GetData() T {
	return n.data
}

func (n *AVLNode[T]) string() string {
	return fmt.Sprintf("%v", n.data)
}

func (n *AVLNode[T]) getLeft() *AVLNode[T] {
	return n.left
}

func (n *AVLNode[T]) getRight() *AVLNode[T] {
	return n.right
}

func (n *AVLNode[T]) getHeight() int {
	if n == nil {
		return -1
	}

	return n.height
}

func (n *AVLNode[T]) getBalance() int {
	if n == nil {
		return 0
	}

	return n.left.getHeight() - n.right.getHeight()
}

func (n *AVLNode[T]) updateHeight() {
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
}

func (n *AVLNode[T]) insert(value T) *AVLNode[T] {
	// Si el nodo es nil, lo crea
	if n == nil {
		return newAVLNode[T](value, nil, nil)
	}
	// Primero inserta el nodo como si fuera un BST común
	switch {
	case value < n.data:
		n.left = n.left.insert(value)
	case value > n.data:
		n.right = n.right.insert(value)
	default: // el elemento ya se encuentra en el árbol
		return n
	}
	// Actualiza la altura del nodo, y si es necesario, aplica rotaciones
	n.updateHeight()

	return n.applyRotation()
}

func (n *AVLNode[T]) rotateRight() *AVLNode[T] {
	y := n.left   // y es el hijo izquierdo de n
	t2 := y.right // t2 es el hijo derecho de y

	// reasignamos los punteros
	y.right = n
	n.left = t2

	// Actualizamos las alturas
	n.updateHeight()
	y.updateHeight()

	return y
}

func (n *AVLNode[T]) rotateLeft() *AVLNode[T] {
	x := n.right // x es el hijo derecho de n
	t2 := x.left // t2 es el hijo izquierdo de x

	// reasignamos los punteros
	x.left = n
	n.right = t2

	// Actualizamos las alturas
	n.updateHeight()
	x.updateHeight()

	return x
}

func (n *AVLNode[T]) remove(value T) *AVLNode[T] {
	if n == nil {
		return n
	}

	// Primero elimina el nodo como si fuera un BST común
	switch {
	case value < n.data:
		n.left = n.left.remove(value)
	case value > n.data:
		n.right = n.right.remove(value)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		temp := n.right.findMin()
		n.data = temp.data
		n.right = n.right.remove(temp.data)
	}

	// Actualiza la altura del nodo, y si es necesario, aplica rotaciones
	n.updateHeight()

	return n.applyRotation()
}

func (n *AVLNode[T]) applyRotation() *AVLNode[T] {
	balance := n.getBalance()

	// Si |balance| > 1, el árbol está desbalanceado
	// Debemos aplicar rotaciones para balancearlo

	// Desbalanceado a la izquierda -> rotación simple a derecha
	if balance > 1 {
		// Si además el hijo izquierdo está desbalanceado a la derecha,
		// aplicamos una rotación a la izquierda resultando en un
		// desbalanceo izquierda-derecha
		if n.left.getBalance() < 0 {
			n.left = n.left.rotateLeft()
		}

		return n.rotateRight()
	}

	// Desbalanceado a la derecha -> rotación simple a izquierda
	if balance < -1 {
		// Si además el hijo derecho está desbalanceado a la izquierda,
		// aplicamos una rotación a la derecha resultando en un
		// desbalanceo derecha-izquierda
		if n.right.getBalance() > 0 {
			n.right = n.right.rotateRight()
		}

		return n.rotateLeft()
	}

	return n
}

func (n *AVLNode[T]) findMin() *AVLNode[T] {
	if n.left == nil {
		return n
	}

	return n.left.findMin()
}

func (n *AVLNode[T]) findMax() *AVLNode[T] {
	if n.right == nil {
		return n
	}

	return n.right.findMax()
}

func (n *AVLNode[T]) search(k T) bool {
	if n == nil {
		return false
	}

	if n.data > k {
		return n.left.search(k)
	}

	if n.data < k {
		return n.right.search(k)
	}

	return true
}

func (n *AVLNode[T]) inOrder() string {
	if n == nil {
		return ""
	}

	return n.left.inOrder() + " " + n.string() + " " + n.right.inOrder()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
