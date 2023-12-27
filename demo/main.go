package main

import (
	"fmt"
	"github.com/untref-ayp2/queue"
)

func main() {
	q := queue.Queue{}
	q.Enqueue(1)
	q.Enqueue("hola")
	q.Enqueue("Mundo")

	v, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	q.Enqueue(1)
	q.Enqueue("hola")
	q.Enqueue("Mundo")

	v, err = q.Dequeue()
	for err == nil {
		fmt.Println(v)
		v, err = q.Dequeue()
	}

	q.Enqueue(1)
	q.Enqueue(2)

	//Modificamos la cola sin usar los métodos definidos
	q[0] = "Hola"
	q[1] = "Mundo"

	v, err = q.Dequeue()
	for err == nil {
		fmt.Printf("%s", v)
		v, err = q.Dequeue()
	}
}