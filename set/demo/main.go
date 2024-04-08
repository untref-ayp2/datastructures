package main

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/set"
)

func main() {
	// Creo dos conjuntos de números
	n1 := set.NewSet(1, 10, 5)
	n2 := set.NewSet(5, 15, 1)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println("Union: ", set.Union(n1, n2))
	fmt.Println("Intersection: ", set.Intersection(n1, n2))
	fmt.Println("Difference: ", set.Difference(n1, n2))
	fmt.Println("Subset: ", set.Subset(n1, n2))
	fmt.Println("Equal: ", set.Equal(n1, n2))
}
