package main

import (
	"fmt"

	set "github.com/untref-ayp2/data-structures/set/set_list"
)

func main() {
	set := set.NewSet(1, 10, 5)
	fmt.Println(set)
	set.Add(7)
	fmt.Println(set)
	fmt.Println(set.Contains(7))
	set.Remove(10)
	fmt.Println(set.Contains(10))
	fmt.Println(set.Values())

}
