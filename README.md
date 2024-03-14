# Stack

Una implementación simple de una estructura de datos tipo pila (stack) en Go.
La pila está parametrizada, por lo que se puede instanciar para obtener una
pila de enteros, o una pila de strings, etc.

## Métodos

### `New[T any]() *Stack[T]`

Crea y devuelve una nueva instancia de la pila vacía.

### `Push[T any](element T)`

Agrega un elemento al tope de la pila.

### `Pop[T any]() (T, error)`

Elimina y devuelve el elemento del tope de la pila.

- Retorna el valor del elemento eliminado.
- Retorna un error si la pila está vacía.

### `Top[T any]() (T, error)`

Devuelve el elemento del tope de la pila sin eliminarlo.

- Retorna el valor del elemento del tope.
- Retorna un error si la pila está vacía.

### `IsEmpty() bool`

Verifica si la pila está vacía.

## Ejemplo de uso

En la carpeta  [demo](./demo/main.go) se encuentra un ejemplo de uso.