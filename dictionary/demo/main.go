package main

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/dictionary"
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
	fmt.Println("Borre a fede")
	fmt.Println("Al pedirle al diccionario el valor para fede me da como resultado: ", d.Get("Fede"))
	fmt.Println("--------------------")
	fmt.Println("Clave: valor en el diccionario sin fede(String, Int)")
	fmt.Println(d)
	fmt.Println("--------------------")
	ds := dictionary.NewDictionary[string, set.Set[int]]()
	s := set.NewSetList[int]()
	s.Add(100)
	s.Add(222)
	s.Add(333)
	ss := set.NewSetList[int]()
	ss.Add(1)
	ss.Add(2)
	ss.Add(3)
	ds.Put("uno", s)
	ds.Put("dos", ss)
	ds.Put("tres", ss)
	fmt.Println("Clave: valor en el diccionario (String, SetList[Int])")
	fmt.Println(ds)
	fmt.Println("Claves en el diccionario:")
	fmt.Println(ds.GetKeys())
	fmt.Println("Valores en el diccionario:")
	fmt.Println(ds.GetValues())
}
