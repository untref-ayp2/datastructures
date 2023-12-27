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
		t.Error("queue should not be empty")
	}

	if v, _ := q.Dequeue(); v != 1 {
		t.Error("first value should be 1")
	}

	if v, _ := q.Dequeue(); v != 2 {
		t.Error("second value should be 2")
	}

	if v, _ := q.Dequeue(); v != 3 {
		t.Error("third value should be 3")
	}

	if _, err := q.Dequeue(); err == nil {
		t.Error("should return error when queue is empty")
	}
}