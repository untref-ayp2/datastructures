package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/untref-ayp2/structures/stack"
)

func InvertirCadena(cadena string) string {
	s := *stack.New[string]()
	salida := ""
	for _, c := range cadena {
		s.Push(string(c))
	}

	for !s.IsEmpty() {
		c, _ := s.Pop()
		salida += c
	}
	return salida
}

func Demo() {
	s := *stack.New[string]()
	fmt.Println("Encolando Hola")
	s.Push("Hola")
	fmt.Println("Encolando 2")
	s.Push("2")
	fmt.Println("Encolando 3")
	s.Push("3")
	fmt.Println("Desencolando todos los elementos:")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println("La pila está vacía:", s.IsEmpty())
	fmt.Println("Desencolando de una pila vacía:")
	fmt.Println(s.Pop())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Demo")
		fmt.Println("2. Invertir cadena")
		fmt.Println("Otra opción. Salir")
		fmt.Print("Elija una opcion: ")

		scanner.Scan()
		choice, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		switch choice {
		case 1:
			Demo()
		case 2:
			fmt.Print("Ingrese una cadena: ")
			scanner.Scan()
			cadena := scanner.Text()
			fmt.Println(InvertirCadena(cadena))
		default:
			fmt.Println("Fin...")
			return
		}
	}
}
