package double_linked_list

import (
	"cmp"
)

// DoubleLinkedList implementa una lista enlazada doble genérica.
type DoubleLinkedList[T cmp.Ordered] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewList[T cmp.Ordered]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

func (dll *DoubleLinkedList[T]) Head() *Node[T] {
	return dll.head
}

func (dll *DoubleLinkedList[T]) Tail() *Node[T] {
	return dll.tail
}

func (dll *DoubleLinkedList[T]) Size() int {
	return dll.size
}

func (dll *DoubleLinkedList[T]) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoubleLinkedList[T]) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

func (dll *DoubleLinkedList[T]) Prepend(data T) {
	newNode := newNode[T](data)

	if dll.size == 0 {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.SetNext(dll.head)
		dll.head.SetPrev(newNode)
		dll.head = newNode
	}

	dll.size++
}

func (dll *DoubleLinkedList[T]) Append(data T) {
	newNode := newNode[T](data)

	if dll.size == 0 {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.SetNext(newNode)
		newNode.SetPrev(dll.tail)
		dll.tail = newNode
	}

	dll.size++
}

func (dll *DoubleLinkedList[T]) Find(data T) *Node[T] {
	current := dll.head

	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}

	return nil
}

func (dll *DoubleLinkedList[T]) RemoveFirst(){
	if dll.size == 0 {
		return
	}

	dll.head = dll.head.next
	dll.size--

	if dll.size == 0 {
		dll.tail = nil
	} else {
		dll.head.SetPrev(nil)
	}
}

func (dll *DoubleLinkedList[T]) RemoveLast(){
	if dll.size == 0 {
		return
	}

	if dll.size == 1 {
		dll.head = nil
		dll.tail = nil
		dll.size = 0
		return
	}

	dll.tail = dll.tail.prev
	dll.tail.SetNext(nil)
	dll.size--
}

func (dll *DoubleLinkedList[T]) Remove(data T){
	node := dll.Find(data)

	if node == nil {
		return
	}

	if node == dll.head {
		dll.RemoveFirst()
		return
	}

	if node == dll.tail {
		dll.RemoveLast()
		return
	}

	node.prev.SetNext(node.next)
	node.next.SetPrev(node.prev)
	dll.size--
}




