package main

import (
	"fmt"

	sl "github.com/untref-ayp2/data-structures/set"
)

func main() {
	set := sl.NewListSet(1, 10, 5)
	fmt.Println(set)
	set.Add(7)
	fmt.Println(set)
	fmt.Println(set.Contains(7))
	set.Remove(10)
	fmt.Println(set.Contains(10))
	fmt.Println(set)
	fmt.Println(set.Values())
	set.Add(4, 12)
	fmt.Println(set)
}
