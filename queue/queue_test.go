package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.IsEmpty() {
		t.Error("La cola no debería estar vacía")
	}

	if v, _ := q.Dequeue(); v != 1 {
		t.Error("El primer valor debería ser 1")
	}

	if v, _ := q.Dequeue(); v != 2 {
		t.Error("El segundo valor debería ser 2")
	}

	if v, _ := q.Dequeue(); v != 3 {
		t.Error("El tercer valor debería ser 3")
	}

	if _, err := q.Dequeue(); err == nil {
		t.Error("La cola debería estar vacía")
	}
}

func TestEmptyQueue(t *testing.T) {
	q := New[int]()

	if !q.IsEmpty() {
		t.Error("La cola debería estar vacía")
	}

	if _, err := q.Dequeue(); err == nil {
		t.Error("La cola debería estar vacía")
	}

	if _, err := q.Front(); err == nil {
		t.Error("La cola debería estar vacía")
	}
}

func TestFront(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if v, _ := q.Front(); v != 1 {
		t.Error("El frente de la cola debería ser 1")
	}

	q.Dequeue()

	if v, _ := q.Front(); v != 2 {
		t.Error("El frente de la cola debería ser 2")
	}

	q.Dequeue()

	if v, _ := q.Front(); v != 3 {
		t.Error("El frente de la cola debería ser 3")
	}
}
