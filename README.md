# Tipos de Datos Abstractos

[![codecov](https://codecov.io/gh/untref-ayp2/data-structures/graph/badge.svg?token=LIV0D9PK0X)](https://codecov.io/gh/untref-ayp2/data-structures)

En este repositorio encontrarán una serie de estructuras de datos básicas que
utilizaremos en el curso.

## Cómo usar este módulo como dependencia de un proyecto

Para usar este módulo se debe agregar la dependencia en el archivo `go.mod` de
ese proyecto.

1. Crear una nueva carpeta e inicializar un módulo de Go.

   ```bash
   mkdir mi-proyecto
   cd mi-proyecto
   go mod init github.com/usuario/mi-proyecto # o cualquier otro nombre
   ```

2. Agregar la dependencia en el archivo `go.mod` del proyecto.

   ```bash
   go get github.com/untref-ayp2/data-structures
   ```

   Nota: este paso podría saltearse y comenzar a utilizar un módulo, para luego
   ejecutar `go mod tidy` y que se descarguen las dependencias.

3. Utilizar el módulo en el código del proyecto (por ejemplo
   `my_proyecto/main.go`).

   ```go
   package main

   import (
        "fmt"

        "github.com/untref-ayp2/data-structures/stack"
   )

   func main() {
       palabra := "omtirogla"
       s := stack.New[rune]()
       for _, letra := range palabra {
           s.Push(letra)
       }
       for !s.IsEmpty() {
           letra, _ := s.Pop()
           fmt.Printf("%c ", letra)
       }
   }
   ```

   Es importante notar que se importan "paquetes" y no módulos. En este caso, el
   paquete `stack`.

4. Se recomienda ejecutar el comando `go mod tidy` para que se descarguen
   las dependencias. Y se corrija cualquier tipo de inconsistencia entre las
   dependencias usadas en el código y las que se hayan declarado en el archivo
   `go.mod`.

   ```bash
   go mod tidy
   ```

   Este comando también se encarga de agregar las dependencias que se usan en el
   código al archivo `go.mod`, y que no hayamos descargado previamente con
   `go get`.

5. Finalmente podemos correr el proyecto haciendo:

   ```bash
   go run .
   ```

## Cómo trabajar en forma local con esta dependencia

Si fuera necesario modificar el código de este módulo, es posible realizar y
utilizar dichas modificaciones de forma local. Para ello debemos seguir los
siguientes pasos.

1. Crear un "espacio de trabajo" (_workspace_), por ejemplo:

   ```bash
   mkdir untref-ayp2
   cd untref-ayp2
   ```

2. Dentro de esta nueva carpeta, se clona el repositorio de este módulo.

   ```bash
   git clone git@github.com:untref-ayp2/data-structures.git
   ```

3. La carpeta de nuestro proyecto debería estar dentro de este mismo
   _workspace_:

   ```bash
   untref-ayp2/
   ├── data-structures/
   └── mi-proyecto/
   ```

4. Se crea el archivo `go.work` que servirá para que Go pueda encontrar el
   módulo `data-structures` de forma local.

   ```bash
   go work init
   go work use -r . # notar el punto al final
   ```

   el archivo `go.work` debería quedar de la siguiente forma:

   ```bash
   go 1.20.0

    use (
        ./data-structures
        ./mi-proyecto
    )
   ```

5. Para verificar que todo está funcionando correctamente podemos modificar el
   código de algún paquete de `data-structures`. Por ejemplo, si agregámos un
   mensaje de log en la función `stack.New`:

   ```go
   // New crea una nueva pila vacía.
    func New[T any]() *Stack[T] {
        fmt.Println("Usando la pila de forma local")
        return &Stack[T]{}
    }
   ```

   Ahora podemos correr el proyecto `mi-proyecto` y ver si se imprime el mensaje
   que acabamos de agregar. Cabe destacar que dado que estamos trabajando dentro
   de un _workspace_, tenemos varias formas de ejecutar el código de nuestro
   módulo. Por ejemplo:

   Accediendo a la carpeta del proyecto y corriendo el comando `go run .`:

   ```bash
   cd mi-proyecto
   go run .
   ```

   Desde la raíz del _workspace_ y pasando como argumento la carpeta de nuestro
   módulo:

   ```bash
   go run ./mi-proyecto
   ```

   Si todo está funcionando correctamente, deberíamos ver el mensaje que
   acabamos de agregar.

   ```bash
   untref-ayp2 ❯ go run ./mi-proyecto
   Usando la pila de forma local
   a l g o r i t m o
   ```

## Algunas notas sobre la estructura de este repositorio

Siguiendo las guías de Go, la estructura de este repositorio es la siguiente:

```bash
data-structures/ # módulo github.com/untref-ayp2/data-structures
├── bitmap       # paquete github.com/untref-ayp2/data-structures/bitmap
├── dictionary   # paquete github.com/untref-ayp2/data-structures/dictionary
├── list         # paquete github.com/untref-ayp2/data-structures/list
├── queue        # paquete github.com/untref-ayp2/data-structures/queue
├── set          # paquete github.com/untref-ayp2/data-structures/set
└── stack        # paquete github.com/untref-ayp2/data-structures/stack
```

El objetivo de publicar todas la estructuras de datos en un unico repositorio es
para facilitar su mantenimiento y distribución. Además, se busca que sea
sencillo de utilizar como dependencia en otros proyectos.

Esta estructura está basada en la organización de los módulos estandar de Go.

Algunas de las guías y tutoriales oficiales de Go que sirvieron de referencia
para la creación de este repositorio son:

- [Tutorial: Create a Go module](https://golang.org/doc/tutorial/create-module)
- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- [Developing and publishing modules](https://go.dev/doc/modules/developing)
- [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow)
- [About pkgsite > Adding a package](https://pkg.go.dev/about#adding-a-package)

Los nombres y convenciones del código están tomadas de esos recursos y además de
la guía [Effective Go](https://golang.org/doc/effective_go).

### Enfoques descartados

- **Un módulo por paquete**: Si bien es posible tener un módulo por paquete, no
  resulta práctico para el mantenimiento de este módulo.

- **Varios módulos en un solo repositorio**: Esta organización dificultó el
  descubrimiento de paquetes, cuando se instalan con `go get`.
