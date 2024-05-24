package avltree

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/types"
	"github.com/untref-ayp2/data-structures/utils"
)

// AVLNode es un nodo del árbol AVL, además del dato y los hijos, registra la altura.
type AVLNode[T types.Ordered] struct {
	data   T           // dato
	height int         // altura
	left   *AVLNode[T] // hijo izquierdo
	right  *AVLNode[T] // hijo derecho
}

// newAVLNode crea un nuevo nodo AVL con el dato, y los hijos izquierdo y derecho pasados como parámetros.
//
// Parámetros:
//   - data: dato del nodo.
//   - left: hijo izquierdo.
//   - right: hijo derecho.
//
// Retorna:
//   - un puntero al nodo creado.
func newAVLNode[T types.Ordered](data T, left *AVLNode[T], right *AVLNode[T]) *AVLNode[T] {
	return &AVLNode[T]{left: left, right: right, data: data, height: 0}
}

// GetData retorna el dato del nodo.
//
// Uso:
//
//	data := node.GetData()
//
// Retorna:
//   - el dato del nodo.
func (n *AVLNode[T]) GetData() T {
	return n.data
}

// string retorna el dato del nodo en formato string.
//
// Retorna:
//   - el dato del nodo en formato string.
func (n *AVLNode[T]) string() string {
	return fmt.Sprintf("%v", n.data)
}

// getLeft retorna el hijo izquierdo del nodo.
//
// Retorna:
//   - el hijo izquierdo del nodo.
func (n *AVLNode[T]) getLeft() *AVLNode[T] {
	return n.left
}

// getRight retorna el hijo derecho del nodo.
//
// Retorna:
//   - el hijo derecho del nodo.
func (n *AVLNode[T]) getRight() *AVLNode[T] {
	return n.right
}

// getHeight retorna la altura del nodo.
//
// Retorna:
//   - la altura del nodo.
func (n *AVLNode[T]) getHeight() int {
	if n == nil {
		return -1
	}

	return n.height
}

// getBalance retorna el balance del nodo.
//
// Retorna:
//   - el balance del nodo.
func (n *AVLNode[T]) getBalance() int {
	if n == nil {
		return 0
	}

	return n.left.getHeight() - n.right.getHeight()
}

// updateHeight actualiza la altura del nodo.
func (n *AVLNode[T]) updateHeight() {
	n.height = 1 + utils.Max(n.left.getHeight(), n.right.getHeight())
}

// insert inserta un nuevo nodo en el árbol AVL.
//
// Parámetros:
//   - value: valor a insertar.
//
// Retorna:
//   - un puntero al nodo insertado.
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

// rotateRight realiza una rotación simple a la derecha.
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

// rotateLeft realiza una rotación simple a la izquierda.
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

// remove elimina un nodo del árbol AVL.
//
// Parámetros:
//   - value: valor a eliminar.
//
// Retorna:
//   - un puntero al nodo eliminado.
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

// applyRotation aplica rotaciones al árbol AVL para balancearlo.
//
// Retorna:
//   - un puntero al nodo raiz del sub-arbol, resultante de aplicar las rotaciones.
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

// findMin retorna el nodo con el valor mínimo del árbol AVL.
//
// Retorna:
//   - un puntero al nodo con el valor mínimo.
func (n *AVLNode[T]) findMin() *AVLNode[T] {
	if n.left == nil {
		return n
	}

	return n.left.findMin()
}

// findMax retorna el nodo con el valor máximo del árbol AVL.
//
// Retorna:
//   - un puntero al nodo con el valor máximo.
func (n *AVLNode[T]) findMax() *AVLNode[T] {
	if n.right == nil {
		return n
	}

	return n.right.findMax()
}

// search busca un valor en el árbol AVL.
//
// Parámetros:
//   - k: valor a buscar.
//
// Retorna:
//   - true si el valor se encuentra en el árbol, false en caso contrario.
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

// inOrder devuelve una representación en orden del árbol AVL.
//
// Retorna:
//   - una representación en orden del árbol AVL.
func (n *AVLNode[T]) inOrder() string {
	if n == nil {
		return ""
	}

	return n.left.inOrder() + " " + n.string() + " " + n.right.inOrder()
}
