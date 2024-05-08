package list

import "fmt"

// DoubleLinkedList implementa una lista enlazada doble genérica.
type DoubleLinkedList[T comparable] struct {
	head *DoubleLinkedNode[T]
	tail *DoubleLinkedNode[T]
	size int
}

// NewDoubleLinkedList crea una nueva lista vacía.
//
// Uso:
//
//	list := NewDoubleLinkedList[int]() // Crea una nueva lista enlazada doble.
func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

// Head devuelve el primer nodo de la lista.
//
// Uso:
//
//	head := list.Head() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (l *DoubleLinkedList[T]) Head() *DoubleLinkedNode[T] {
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
func (l *DoubleLinkedList[T]) Tail() *DoubleLinkedNode[T] {
	return l.tail
}

// Size devuelve el tamaño de la lista.
//
// Uso:
//
//	size := list.Size() // Obtiene el tamaño de la lista.
//
// Retorna:
//   - el tamaño de la lista.
func (l *DoubleLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty evalúa si la lista está vacía.
//
// Uso:
//
//	empty := list.IsEmpty() // Verifica si la lista está vacía.
//
// Retorna:
//   - true si la lista está vacía; false en caso contrario.
func (l *DoubleLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear elimina todos los nodos de la lista.
//
// Uso:
//
//	list.Clear() // Elimina todos los nodos de la lista.
func (l *DoubleLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Prepend inserta un dato al inicio de la lista.
//
// Uso:
//
//	list.Prepend(10) // Inserta el dato 10 al inicio de la lista.
//
// Parámetros:
//   - `data`: el dato a insertar al frente de la lista.
func (l *DoubleLinkedList[T]) Prepend(data T) {
	newNode := NewDoubleLinkedNode[T](data)
	if l.size == 0 {
		l.tail = newNode
	} else {
		newNode.SetNext(l.head)
		l.head.SetPrev(newNode)
	}
	l.head = newNode
	l.size++
}

// Append inserta un dato al final de la lista.
//
// Uso:
//
//	list.Append(10) // Inserta el dato 10 al final de la lista.
//
// Parámetros:
//   - `data`: el dato a insertar al final de la lista.
func (l *DoubleLinkedList[T]) Append(data T) {
	newNode := NewDoubleLinkedNode[T](data)
	if l.size == 0 {
		l.head = newNode
	} else {
		l.tail.SetNext(newNode)
		newNode.SetPrev(l.tail)
	}
	l.tail = newNode
	l.size++
}

// Find busca un dato en la lista
//
// Uso:
//
//	node := list.Find(10) // Busca el dato 10 en la lista.
//
// Parámetros:
//   - `data`: el dato a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el dato buscado; `nil` si no se encuentra.
func (l *DoubleLinkedList[T]) Find(data T) *DoubleLinkedNode[T] {
	for current := l.head; current != nil; current = current.Next() {
		if current.Data() == data {
			return current
		}
	}

	return nil
}

// RemoveFirst elimina el primer nodo de la lista.
//
// Uso:
//
//	list.RemoveFirst() // Elimina el primer nodo de la lista.
func (l *DoubleLinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.head = l.head.Next()
	l.size--

	if l.IsEmpty() {
		l.tail = nil
	} else {
		l.head.SetPrev(nil)
	}
}

// RemoveLast elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoveLast() // Elimina el último nodo de la lista.
func (l *DoubleLinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0

		return
	}

	l.tail = l.tail.Prev()
	l.tail.SetNext(nil)
	l.size--
}

// Remove elimina un la primera aparición de un dato en la lista.
//
// Uso:
//
//	list.Remove(10) // Elimina la primera aparición del dato 10 en la lista.
func (l *DoubleLinkedList[T]) Remove(data T) {
	node := l.Find(data)

	if node == nil {
		return
	}

	if node == l.head {
		l.RemoveFirst()

		return
	}

	if node == l.tail {
		l.RemoveLast()

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
func (l *DoubleLinkedList[T]) String() string {
	if l.IsEmpty() {
		return "DoubleLinkedList: []"
	}

	result := "DoubleLinkedList: "

	current := l.Head()
	for {
		result += fmt.Sprintf("[%v]", current.Data())
		if !current.HasNext() {
			break
		}
		result += " ↔ "
		current = current.Next()
	}

	return result
}
