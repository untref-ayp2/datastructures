# Queue

Una implementación simple de una estructura de datos tipo cola (queue) en Go.

La cola por lo que se puede instanciar para obtener una cola de enteros, o una
cola de strings, etc.

## Métodos

### `New[T any]() *Queue[T]`

Crea y devuelve una nueva instancia de la cola vacía.

### `Enqueue[T any](element T)`

Agrega un elemento al final de la cola.

### `Dequeue[T any]() (T, error)`

Elimina y devuelve el elemento más antiguo que se encuentre en la cola.

- Retorna el valor del elemento eliminado.
- Retorna un error si la cola está vacía.

### `Front[T any]() (T, error)`

Devuelve el elemento al frente de la cola sin removerlo.

- Retorna el primer elemento de la cola.
- Retorna un error si la cola está vacía.

### `IsEmpty() bool`

Verifica si la cola está vacía.

## Ejemplo de uso

En la carpeta [demo](./demo/main.go) se encuentra un ejemplo de uso.
