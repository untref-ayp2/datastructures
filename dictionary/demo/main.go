package main

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/dictionary"
	"github.com/untref-ayp2/data-structures/list"
	"github.com/untref-ayp2/data-structures/set"
)

func main() {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 60)
	d.Put("Leo", 61)
	d.Put("Fabi", 36)
	d.Put("Leo", 60)
	d.Put("Fede", 36)
	d.Put("Vale", 40)
	d.Put("Leo", 50)

	fmt.Println("Clave: valor en el diccionario (String, Int)")
	fmt.Println(d)
	fmt.Println("--------------------")

	d.Remove("Fede")

	fmt.Println("Borre a Fede")
	fmt.Println("Al pedirle al diccionario el valor para Fede me da como resultado: ", d.Get("Fede"))
	fmt.Println("--------------------")
	fmt.Println("Clave: valor en el diccionario sin Fede (String, Int)")
	fmt.Println(d)
	fmt.Println("--------------------")

	ds := dictionary.NewDictionary[string, set.Set[int]]()
	s1 := set.NewListSet(100, 222, 333)
	s2 := set.NewListSet(1, 2, 3)
	ds.Put("uno", s1)
	ds.Put("dos", s2)
	ds.Put("tres", s2)
	fmt.Println("Clave: valor en el diccionario (String, SetList[Int])")
	fmt.Println(ds)
	fmt.Println("Claves en el diccionario:")
	fmt.Println(ds.Keys())
	fmt.Println("Valores en el diccionario:")
	fmt.Println(ds.Values())
	fmt.Println("--------------------")

	dl := dictionary.NewDictionary[string, list.LinkedList[int]]()
	l1 := list.NewLinkedList[int]()
	l2 := list.NewLinkedList[int]()
	l1.Append(10)
	l1.Append(20)
	l2.Append(30)
	l2.Append(40)
	l2.Append(50)
	dl.Put("uno", *l1)
	dl.Put("dos", *l2)
	fmt.Println("Clave: valor en el diccionario (String, LinkedList[int]]")
	fmt.Println(dl)
	fmt.Println("Claves en el diccionario:")
	fmt.Println(dl.Keys())
	fmt.Println("Valores en el diccionario:")
	fmt.Println(dl.Values())
}
