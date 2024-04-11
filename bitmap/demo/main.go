package main

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/bitmap"
)

func main() {
	m := bitmap.NewBitMap()

	m.On(3)
	m.On(0)
	fmt.Printf("%032b\n", m.GetMap())

	m.On(5)
	m.On(6)
	fmt.Printf("%032b\n", m.GetMap())

	m.Off(3)
	fmt.Printf("%032b\n", m.GetMap())

	fmt.Println("=== Ejemplos de operadores a nivel de bits ===")
	var x int32 = 0x2f // 63 en decimal
	fmt.Printf("%032b = x, equivale al conjunto {0, 1, 2, 3, 4, 5}\n", x)
	var y int32 = 0xdc19 // 56_345 en decimal
	fmt.Printf("%032b = y, equivale al conjunto {0, 3, 4, 10, 11, 12, 14, 15}\n", y)

	fmt.Printf("Intersección\n%032b = x&y,  equivale al conjunto {0, 3, 4}\n", x&y)
	fmt.Printf("Unión\n%032b = x|y,  equivale al conjunto {0, 1, 2, 3, 4, 5, 10, 11, 12, 14, 15}\n", x|y)
	fmt.Printf("Diferencia simétrica\n%032b = x^y,  equivale al conjunto {1, 2, 5, 10, 11, 12, 14, 15}\n", x^y)
	fmt.Printf("Diferencia\n%032b = x&^y, equivale al conjunto {1, 3, 5}\n", x&^y)
}
