package demo
import (
	"fmt"
	"github.com/untre-ayp2/stack"
)

func main() {
	var s stack.Stack
	cadena := "Hola Mundo"
	for _, v := range cadena {
		s.Push(string(v))
	}

	c, err := s.Pop()
	for err == nil {
		fmt.Printf("%s", c)
		c, err = s.Pop()
	}
	fmt.Println()

}