// Package list proporciona implementaciones de listas enlazadas simples, dobles y circulares.
// Expone las estructuras LinkedList, DoubleLinkedList y CircularList y sus métodos para manipular listas.
package list

import "fmt"

// LinkedList se implementa con un nodo que contiene un dato y un puntero al siguiente nodo.
// Los elementos deben ser de un tipo comparable.
type LinkedList[T comparable] struct {
	head *LinkedNode[T]
	tail *LinkedNode[T]
	size int
}

// NewLinkedList crea una nueva lista vacía.
//
// Uso:
//
//	list := list.NewLinkedList[int]() // Crea una nueva lista vacía.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Head devuelve el primer nodo de la lista.
//
// Uso:
//
//	head := list.Head() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (l *LinkedList[T]) Head() *LinkedNode[T] {
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
func (l *LinkedList[T]) Tail() *LinkedNode[T] {
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
func (l *LinkedList[T]) Size() int {
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
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear elimina todos los nodos de la lista.
//
// Uso:
//
//	list.Clear() // Elimina todos los nodos de la lista.
func (l *LinkedList[T]) Clear() {
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
//   - `data`: el dato a insertar en la lista.
func (l *LinkedList[T]) Prepend(data T) {
	newNode := NewLinkedListNode(data)
	if l.IsEmpty() {
		l.tail = newNode
	} else {
		newNode.SetNext(l.head)
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
//   - `data`: el dato a insertar en la lista.
func (l *LinkedList[T]) Append(data T) {
	newNode := NewLinkedListNode(data)
	if l.IsEmpty() {
		l.head = newNode
	} else {
		l.tail.SetNext(newNode)
	}
	l.tail = newNode
	l.size++
}

// Find busca un dato en la lista, si lo encuentra devuelve el nodo
// correspondiente, si no lo encuentra devuelve nil
//
// Uso:
//
//	node := list.Find(10) // Busca el dato 10 en la lista.
//
// Parámetros:
//   - `data`: el dato a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el dato; `nil` si el dato no se encuentra.
func (l *LinkedList[T]) Find(data T) *LinkedNode[T] {
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
func (l *LinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.head = l.head.Next()

	if l.head == nil {
		l.tail = nil
	}

	l.size--
}

// RemoveLast elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoveLast() // Elimina el último nodo de la lista.
func (l *LinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		current := l.head
		for current.Next() != l.tail {
			current = current.Next()
		}
		current.SetNext(nil)
		l.tail = current
	}
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
func (l *LinkedList[T]) Remove(data T) {
	node := l.Find(data)

	if node == nil {
		return
	}

	if node == l.head {
		l.RemoveFirst()

		return
	}

	current := l.Head()

	for current.Next() != node {
		current = current.Next()
	}

	current.SetNext(node.Next())

	if node == l.tail {
		l.tail = current
	}
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
func (l *LinkedList[T]) String() string {
	if l.IsEmpty() {
		return "LinkedList: []"
	}

	result := "LinkedList: "

	current := l.Head()
	for {
		result += fmt.Sprintf("[%v]", current.Data())
		if !current.HasNext() {
			break
		}
		result += " → "
		current = current.Next()
	}

	return result
}
