package tests

import (
	"testing"

	"github.com/untref-ayp2/queue"
)

func TestQueue(t *testing.T) {
	q := queue.Queue{}

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
