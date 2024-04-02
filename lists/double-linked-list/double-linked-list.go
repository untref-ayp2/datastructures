package double_linked_list

import (
	"cmp"
	"errors"
)

// DoubleLinkedList implementa una lista enlazada doble genérica.
type DoubleLinkedList[T cmp.Ordered] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewDoubleLinkedList[T cmp.Ordered]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

func (l *DoubleLinkedList[T]) InsertAt(index int, data T) error {
	if index < 0 || index > l.size {
		return errors.New("index out of bounds")
	}

	newNode := newNode[T](data)

	if l.size == 0 { // si la lista estaba vacía
		l.head = newNode
		l.tail = newNode
	} else if index == 0 { // si se inserta al principio
		newNode.setNext(l.head)
		l.head.setPrev(newNode)
		l.head = newNode
	} else if index == l.size { // si se inserta al final
		newNode.setPrev(l.tail)
		l.tail.setNext(newNode)
		l.tail = newNode
	} else {
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.getNext()
		}

		newNode.setNext(current.getNext())
		newNode.setPrev(current)
		current.setNext(newNode)
		newNode.getNext().setPrev(newNode)
	}

	l.size++
	return nil
}

func (l *DoubleLinkedList[T]) RemoveAt(index int) (T, error) {
	var data T
	if index < 0 || index >= l.size {
		return data, errors.New("indice fuera de rango")
	}

	var removed *node[T]
	if l.size == 1 { // si la lista tiene un solo elemento
		removed = l.head
		l.head = nil
		l.tail = nil
	} else if index == 0 { // si se elimina el primer elemento
		removed = l.head
		l.head = l.head.getNext()
		l.head.setPrev(nil)
	} else if index == l.size-1 { // si se elimina el último elemento
		removed = l.tail
		l.tail = l.tail.getPrev()
		l.tail.setNext(nil)
	} else {
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.getNext()
		}

		removed = current.getNext()
		current.setNext(removed.getNext())
		removed.getNext().setPrev(current)
	}

	l.size--
	return removed.getData(), nil
}

func (l *DoubleLinkedList[T]) Get(index int) (T, error) {
	var data T
	if index < 0 || index >= l.size {
		return data, errors.New("indice fuera de rango")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.getNext()
	}

	return current.getData(), nil
}

func (l *DoubleLinkedList[T]) IndexOf(data T) int {
	current := l.head
	for i := 0; i < l.size; i++ {
		if current.getData() == data {
			return i
		}
		current = current.getNext()
	}
	return -1
}

func (l *DoubleLinkedList[T]) Set(index int, data T) error {
	if index < 0 || index >= l.size {
		return errors.New("indice fuera de rango")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.getNext()
	}

	current.setData(data)
	return nil
}

func (l *DoubleLinkedList[T]) Size() int {
	return l.size
}

func (l *DoubleLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *DoubleLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *DoubleLinkedList[T]) Iterate() <-chan T {
	ch := make(chan T)
	go func() {
		current := l.head
		for current != nil {
			ch <- current.getData()
			current = current.getNext()
		}
		close(ch)
	}()
	return ch
}
