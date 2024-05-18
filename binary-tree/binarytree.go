package binarytree

import "github.com/untref-ayp2/data-structures/types"

//"golang.org/x/exp/constraints"

type BinaryTree[T types.Ordered] struct {
	root *BinaryNode[T]
}

// NewBinaryTreee crea un nuevo BinaryTree a partir de un dato de tipo T.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//
// Parámetros:
// 'data' : el dato que guarda el nodo raiz de tipo T
//
// Retorna un puntero a un nuevo BinaryTree.
func NewBinaryTree[T types.Ordered](data T) *BinaryTree[T] {
	node := NewBinaryNode(data, nil, nil)
	return &BinaryTree[T]{root: node}
}

// Inserta del lado izquierdo de la raiz, el arbol que se pasa por parámetro
// Parámetros: un puntero a un BinaryTree
func (t *BinaryTree[T]) InsertLeft(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.left = bt.root
	}
}

// Inserta del lado derecho de la raiz, el arbol que se pasa por parámetro
// Parámetros: un puntero a un BinaryTree
func (t *BinaryTree[T]) InsertRight(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.right = bt.root
	}
}

// Muestra el recorrido en Pre-Order de todo el arbol.
// La responsabilidad la delega en en nodo raiz, que sabe mostrar su recorrido en Pre-Order.
func (t *BinaryTree[T]) PrintPreOrder() {
	if t.root != nil {
		t.root.PrintPreOrder()
	}
}

// Retorna un string que representa el recorrido en Pre-Order de todo el arbol.
// La responsabilidad la delega en en nodo raiz, que sabe obtener su recorrido en Pre-Order.
func (t *BinaryTree[T]) StringPreOrder() string {
	if t.root != nil {
		return t.root.StringPreOrder()
	}
	return ""
}

// Muestra el recorrido en In-Order de todo el arbol.
// La responsabilidad la delega en en nodo raiz, que sabe mostrar su recorrido en In-Order.
func (t *BinaryTree[T]) PrintInOrder() {
	if t.root != nil {
		t.root.PrintInOrder()
	}
}

// Retorna un string que representa el recorrido en In-Order de todo el arbol.
// La responsabilidad la delega en en nodo raiz, que sabe obtener su recorrido en In-Order.
func (t *BinaryTree[T]) StringInOrder() string {
	if t.root != nil {
		return t.root.StringInOrder()
	}
	return ""
}

// Muestra el recorrido en Post-Order de todo el arbol.
// La responsabilidad la delega en en nodo raiz, que sabe mostrar su recorrido en Post-Order.
func (t *BinaryTree[T]) PrintPostOrder() {
	if t.root != nil {
		t.root.PrintPostOrder()
	}
}

// Retorna un string que representa el recorrido en Post-Order de todo el arbol.
// La responsabilidad la delega en en nodo raiz, que sabe obtener su recorrido en Post-Order.
func (t *BinaryTree[T]) StringPostOrder() string {
	if t.root != nil {
		return t.root.StringPostOrder()
	}
	return ""
}

// Limpia el árbol poniendo la raiz en nil
func (t *BinaryTree[T]) Empty() {
	t.root = nil
}

// Devuelve true o false según la raiz apunte a un BinaryNode o a nil.
func (t *BinaryTree[T]) IsEmpty() bool {
	return t.root == nil
}

// Retorna la cantidad de nodos del árbol.
func (t *BinaryTree[T]) Size() int {
	if t.root != nil {
		return t.root.Size()
	} else {
		return -1
	}
}

// Retorna la altura del árbol, o sea, la distancia desde la raíz al nodo más profundo.
func (t *BinaryTree[T]) Height() int {
	if t.root != nil {
		return t.root.Height()
	} else {
		return 0
	}
}
