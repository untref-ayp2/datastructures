package main

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/mapadebits"
)

func main() {
	m := mapadebits.NewMapaDeBits()

	m.Encender(3)
	m.Encender(0)

	fmt.Printf("%032b\n", *m)

	m.Encender(5)
	m.Encender(6)
	fmt.Printf("%032b\n", *m)

	m.Apagar(3)
	fmt.Printf("%032b\n", *m)

	var x int32 = 63
	var y int32 = 56345
	fmt.Printf("%032b = x, representa el conjunto {1,2,3,4,5,6}\n", x)
	fmt.Printf("%032b = x, representa el conjunto {1,4,5,11,12,13,15,16}\n", y)
	fmt.Printf("%032b = x&y, representa el conjunto {1,4,5}\n", x&y)                      // Interseccion
	fmt.Printf("%032b = x|y, representa el conjunto {1,2,3,4,5,6,11,12,13,15,16}\n", x|y) //Union
	fmt.Printf("%032b = x^y, representa el conjunto {2,3,6,11,12,13,15,16}\n", x^y)
	fmt.Printf("%032b = x&^y,representa el conjunto {}\n", x&^y) //diferencia
}
