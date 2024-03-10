# Queue

Una implementación simple de una estructura de datos tipo cola (queue) en Go.
La cola es genérica y puede almacenar cualquier tipo de datos (any) en ella.

## Métodos

### `New() *Queue`

Crea y devuelve una nueva instancia de la cola vacía.

### `Enqueue(any)`

Agrega un elemento al final de la cola.

### `Dequeue() (any, error)`

Elimina y devuelve el elemento más antiguo que se encuentre en la cola.

- Retorna el valor del elemento eliminado.
- Retorna un error si la cola está vacía.


### `Front() (any, error)`

Devuelve el elemento al frente de la cola sin removerlo.
- Retorna el primer elemento de la cola.
- Retorna un error si la cola está vacía.

### `IsEmpty() bool`

Verifica si la cola está vacía.

- Retorna `true` si la cola está vacía, de lo contrario retorna `false`.

## Ejemplo de uso

En la carpeta  [demo](./demo) se encuentra un ejemplo de uso
