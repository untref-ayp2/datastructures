package single_linked_list

import (
	"cmp"
	"errors"
)


type SingleLinkedList[T cmp.Ordered] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewSingleLinkedList[T cmp.Ordered]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (sll *SingleLinkedList[T]) Size() int {
	return sll.size
}

func (sll *SingleLinkedList[T]) IsEmpty() bool {
	return sll.size == 0
}

func (ssl *SingleLinkedList[T]) InsertAt(index int, data T) error {
	if index < 0 || index > ssl.size {
		return errors.New("índice fuera de rango")
	}
	newNode := newNode(data)
	// si la lista está vacía O(1)
	if ssl.IsEmpty() {
		ssl.head = newNode
		ssl.tail = newNode
	// si se inserta al inicio O(1)
	} else if index == 0 {
		newNode.setNext(ssl.head)
		ssl.head = newNode
	// si se inserta al final O(1)
	} else if index == ssl.size {
		ssl.tail.setNext(newNode)
		ssl.tail = newNode
	// si se inserta en cualquier otro lugar O(n)
	} else {
		current := ssl.head
		for i := 0; i < index-1; i++ {
			current = current.getNext()
		}
		newNode.setNext(current.getNext())
		current.setNext(newNode)
	}
	ssl.size++
	return nil
}

func (ssl* SingleLinkedList[T]) RemoveAt(index int) (T, error) {
	var data T
	if index < 0 || index >= ssl.size {
		return data, errors.New("índice fuera de rango")
	}
	// si se elimina al inicio O(1)
	if index == 0 {
		data = ssl.head.getData()
		ssl.head = ssl.head.getNext()
		// si la lista queda vacía
		if ssl.head == nil {
			ssl.tail = nil
		}
	// si se elimina en cualquier otro lugar O(n)
	} else {
		current := ssl.head
		for i := 0; i < index-1; i++ {
			current = current.getNext()
		}
		data = current.getNext().getData()
		current.setNext(current.getNext().getNext())
		// si se elimina el último nodo
		if index == ssl.size-1 {
			ssl.tail = current
		}
	}
	ssl.size--
	return data, nil
}

func (ssl *SingleLinkedList[T]) Get(index int) (T, error) {
	var data T
	if index < 0 || index >= ssl.size {
		return data, errors.New("índice fuera de rango")
	}
	current := ssl.head
	for i := 0; i < index; i++ {
		current = current.getNext()
	}
	return current.getData(), nil
}

func (ssl *SingleLinkedList[T]) IndexOf(data T) int {
	current := ssl.head
	for i := 0; i < ssl.size; i++ {
		if current.getData() == data { 
			return i
		}
		current = current.getNext()
	}
	return -1
}

func (ssl *SingleLinkedList[T]) Set(index int, data T) error {
	if index < 0 || index >= ssl.size {
		return errors.New("índice fuera de rango")
	}
	current := ssl.head
	for i := 0; i < index; i++ {
		current = current.getNext()
	}
	current.setData(data)
	return nil
}

func (ssl *SingleLinkedList[T]) Clear() {
	ssl.head = nil
	ssl.tail = nil
	ssl.size = 0
}

func (ssl *SingleLinkedList[T]) Iterate() <-chan T {
	ch := make(chan T)
	go func() {
		current := ssl.head
		for current != nil {
			ch <- current.getData()
			current = current.getNext()
		}
		close(ch)
	}()
	return ch
}