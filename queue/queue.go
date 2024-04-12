package queue

import "errors"

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue[T any] struct {
	data []T
}

// New crea una nueva cola vacía. O(1)
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	var head T
	if len(q.data) == 0 {
		return head, errors.New("cola vacía")
	}
	head = q.data[0]
	q.data = q.data[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue[T]) Front() (T, error) {
	var head T
	if len(q.data) == 0 {
		return head, errors.New("cola vacía")
	}
	head = q.data[0]
	return head, nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}
