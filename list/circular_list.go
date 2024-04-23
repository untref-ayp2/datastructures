package list

import "fmt"

// CircularList implementa una lista enlazada circular genérica.
type CircularList[T comparable] struct {
	head *DoubleLinkedNode[T]
	size int
}

// NewCircularList crea una nueva lista enlazada circular.
//
// Uso:
//
//	list := NewCircularList[int]() // Crea una nueva lista enlazada circular.
func NewCircularList[T comparable]() *CircularList[T] {
	return &CircularList[T]{}
}

// Head devuelve el primer nodo de la lista.
//
// Uso:
//
//	head := list.Head() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (l *CircularList[T]) Head() *DoubleLinkedNode[T] {
	return l.head
}

// Tail devuelve el último nodo de la lista.
//
// Uso:
//
//	tail := list.Tail() // Obtiene el último nodo de la lista.
//
// Retorna:
//   - el último nodo de la lista.
func (l *CircularList[T]) Tail() *DoubleLinkedNode[T] {
	if l.size == 0 {
		return nil
	} else {
		return l.head.Prev()
	}
}

// Size devuelve el tamaño de la lista.
//
// Uso:
//
//	size := list.Size() // Obtiene el tamaño de la lista.
//
// Retorna:
//   - el tamaño de la lista.
func (l *CircularList[T]) Size() int {
	return l.size
}

// IsEmpty evalúa si la lista está vacía.
//
// Uso:
//
//	empty := list.IsEmpty() // Verifica si la lista está vacía.
//
// Retorna:
//   - `true` si la lista está vacía; `false` en caso contrario.
func (l *CircularList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear elimina todos los elementos de la lista.
//
// Uso:
//
//	list.Clear() // Elimina todos los elementos de la lista.
func (l *CircularList[T]) Clear() {
	l.head = nil
	l.size = 0
}

// Prepend agrega un nuevo nodo al principio de la lista.
//
// Uso:
//
//	list.Prepend(10) // Agrega el valor 10 al principio de la lista.
//
// Parámetros:
//   - `data`: el valor a agregar al principio de la lista.
func (l *CircularList[T]) Prepend(data T) {
	node := NewDoubleLinkedListNode(data)
	if l.size == 0 {
		l.head = node
		l.head.SetNext(l.head)
		l.head.SetPrev(l.head)
	} else {
		node.SetNext(l.head)
		node.SetPrev(l.head.Prev())
		l.head.Prev().SetNext(node)
		l.head.SetPrev(node)
		l.head = node
	}
	l.size++
}

// Append agrega un nuevo nodo al final de la lista.
//
// Uso:
//
//	list.Append(10) // Agrega el valor 10 al final de la lista.
//
// Parámetros:
//   - `data`: el valor a agregar al final de la lista.
func (l *CircularList[T]) Append(data T) {
	node := NewDoubleLinkedListNode(data)
	if l.size == 0 {
		l.head = node
		l.head.SetNext(l.head)
		l.head.SetPrev(l.head)
	} else {
		node.SetNext(l.head)
		node.SetPrev(l.head.Prev())
		l.head.Prev().SetNext(node)
		l.head.SetPrev(node)
	}
	l.size++
}

// Find busca un nodo en la lista.
//
// Uso:
//
//	node := list.Find(10) // Busca el valor 10 en la lista.
//
// Parámetros:
//   - `data`: el valor a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el valor buscado; `nil` si el valor no se encuentra en la lista.
func (l *CircularList[T]) Find(data T) *DoubleLinkedNode[T] {
	if l.size == 0 {
		return nil
	}

	node := l.head
	for i := 0; i < l.size; i++ {
		if node.Data() == data {
			return node
		}
		node = node.Next()
	}

	return nil
}

// RemoveFirst elimina el primer nodo de la lista.
//
// Uso:
//
//	list.RemoveFirst() // Elimina el primer nodo de la lista.
func (l *CircularList[T]) RemoveFirst() {
	if l.size == 0 {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.size--
		return
	}

	l.head.Prev().SetNext(l.head.Next())
	l.head.Next().SetPrev(l.head.Prev())
	l.head = l.head.Next()
	l.size--
}

// RemoveLast elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoveLast() // Elimina el último nodo de la lista.
func (l *CircularList[T]) RemoveLast() {
	if l.size == 0 {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.size--
		return
	}

	l.head.Prev().Prev().SetNext(l.head)
	l.head.SetPrev(l.head.Prev().Prev())
	l.size--
}

// Remove elimina un la primera aparición de un dato en la lista.
//
// Uso:
//
//	list.Remove(10) // Elimina la primera aparición del dato 10 en la lista.
//
// Parámetros:
//   - `data`: el dato a eliminar de la lista.
func (l *CircularList[T]) Remove(data T) {
	node := l.Find(data)
	if node == nil {
		return
	}

	if node == l.head {
		l.RemoveFirst()
		return
	}

	node.Prev().SetNext(node.Next())
	node.Next().SetPrev(node.Prev())
	l.size--
}

// String devuelve una representación en cadena de la lista.
//
// Uso:
//
//	fmt.Println(list) // Imprime la representación en cadena de la lista.
//
// Retorna:
//   - una representación en cadena de la lista.
func (l *CircularList[T]) String() string {
	if l.IsEmpty() {
		return "CircularList: {}"
	}

	result := "CircularList: {\n"

	current := l.Head()
	for {
		result += fmt.Sprintf("  %v\n", current.Data())
		if current == l.Tail() {
			break
		}
		current = current.Next()
	}

	return result + "}"
}
