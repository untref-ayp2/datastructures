# Stack

Una implementación simple de una estructura de datos tipo pila (stack) en Go.

## Métodos

### `New() *Stack`

Crea y devuelve una nueva instancia de la pila vacía.

### `Push(value any)`

Agrega un elemento al tope de la pila.

- `value`: El valor a agregar a la pila.

### `Pop() (any, error)`

Elimina y devuelve el elemento del tope de la pila.

- Retorna el valor del elemento eliminado.
- Retorna un error si la pila está vacía.

### `Top() (any, error)`

Devuelve el elemento del tope de la pila sin eliminarlo.

- Retorna el valor del elemento del tope.
- Retorna un error si la pila está vacía.

### `IsEmpty() bool`

Verifica si la pila está vacía.

- Retorna `true` si la pila está vacía, de lo contrario retorna `false`.

## Ejemplo de uso

En la carpeta  [demo](./demo) se encuentra un ejemplo de uso
