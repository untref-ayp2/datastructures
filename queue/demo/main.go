package main

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/queue"
)

func main() {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for !q.IsEmpty() {
		item, _ := q.Dequeue()
		fmt.Println(item)
	}
}
