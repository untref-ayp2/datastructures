// Package binarytree provee una implementación de árbol binario.
package binarytree

import "github.com/untref-ayp2/data-structures/types"

type BinaryTree[T types.Ordered] struct {
	root *BinaryNode[T]
}

// NewBinaryTree crea un nuevo BinaryTree a partir de un dato de tipo T.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//
// Parámetros:
//   - 'data' : el dato de tipo T que guarda el nodo raíz
//
// Retorna:
//   - un puntero a un nuevo BinaryTree.
func NewBinaryTree[T types.Ordered](data T) *BinaryTree[T] {
	node := NewBinaryNode(data, nil, nil)

	return &BinaryTree[T]{root: node}
}

// Inserta del lado izquierdo de la raíz, el árbol que se pasa por parámetro
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//
// Parámetros:
//   - `bt` un puntero a un BinaryTree.
func (t *BinaryTree[T]) InsertLeft(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.left = bt.root
	}
}

// Inserta del lado derecho de la raíz, el árbol que se pasa por parámetro
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertRight(bt2)
//
// Parámetros:
//   - `bt` un puntero a un BinaryTree.
func (t *BinaryTree[T]) InsertRight(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.right = bt.root
	}
}

// Retorna un slice de T con el recorrido en Post-Order del árbol.
// La responsabilidad la delega en nodo raíz, que sabe obtener su
// recorrido en Post-Order.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//	bt.GetPostOrder()
//
// Retorna:
//   - un slice de T con el recorrido en Post-Order del árbol.
func (t *BinaryTree[T]) GetPostOrder() []T {
	s := []T{}
	if t.root != nil {
		s = t.root.GetPostOrder()
	}

	return s
}

// Retorna un slice de T con el recorrido en Pre-Order del árbol.
// La responsabilidad la delega en nodo raíz, que sabe obtener su
// recorrido en Pre-Order.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//	bt.GetPreOrder()
//
// Retorna:
//   - un slice de T con el recorrido en Pre-Order del árbol.
func (t *BinaryTree[T]) GetPreOrder() []T {
	s := []T{}
	if t.root != nil {
		s = t.root.GetPreOrder()
	}

	return s
}

// Retorna un slice de T con el recorrido en In-Order del árbol.
// La responsabilidad la delega en nodo raíz, que sabe obtener su
// recorrido en In-Order.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//	bt.GetInOrder()
//
// Retorna:
//   - un slice de T con el recorrido en In-Order del árbol.
func (t *BinaryTree[T]) GetInOrder() []T {
	s := []T{}
	if t.root != nil {
		s = t.root.GetInOrder()
	}

	return s
}

// Retorna un string que representa el recorrido en Pre-Order de todo el árbol.
// La responsabilidad la delega en nodo raíz,
// que sabe obtener su recorrido en Pre-Order.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//	fmt.Println(bt.StringPreOrder())
//
// Retorna:
//   - un string que representa el recorrido en Pre-Order de todo el árbol.
func (t *BinaryTree[T]) StringPreOrder() string {
	if t.root != nil {
		return t.root.StringPreOrder()
	}

	return ""
}

// Retorna un string que representa el recorrido en Post-Order de todo el árbol.
// La responsabilidad la delega en nodo raíz,
// que sabe obtener su recorrido en Post-Order.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//	fmt.Println(bt.StringPostOrder())
//
// Retorna:
//   - un string que representa el recorrido en Post-Order de todo el árbol.
func (t *BinaryTree[T]) StringPostOrder() string {
	if t.root != nil {
		return t.root.StringPostOrder()
	}

	return ""
}

// Retorna un string que representa el recorrido en In-Order de todo el árbol.
// La responsabilidad la delega en nodo raíz,
// que sabe obtener su recorrido en In-Order.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//	fmt.Println(bt.StringInOrder())
//
// Retorna:
//   - un string que representa el recorrido en In-Order de todo el árbol.
func (t *BinaryTree[T]) StringInOrder() string {
	if t.root != nil {
		return t.root.StringInOrder()
	}

	return ""
}

// Limpia el árbol poniendo la raíz en nil.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Empty()
func (t *BinaryTree[T]) Empty() {
	t.root = nil
}

// Evalúa si el árbol está vacío.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.IsEmpty()
//
// Retorna:
//   - `true` si el árbol está vacío; `false` si no lo está.
func (t *BinaryTree[T]) IsEmpty() bool {
	return t.root == nil
}

// Retorna la cantidad de nodos del árbol.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Size()
//
// Retorna:
//   - la cantidad de nodos del árbol.
func (t *BinaryTree[T]) Size() int {
	if t.root != nil {
		return t.root.Size()
	}

	return 0
}

// Retorna la altura del árbol, o sea, la distancia desde la raíz al nodo más profundo.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Height()
//
// Retorna:
//   - la altura del árbol.
func (t *BinaryTree[T]) Height() int {
	if t.root != nil {
		return t.root.Height()
	}

	return -1
}
