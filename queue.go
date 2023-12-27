package queue

import "errors"

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue []any

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	*q = append(*q, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (*q)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
