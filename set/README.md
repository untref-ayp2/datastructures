# Conjuntos

El TAD conjunto representa matemáticamente un conjunto finito de elementos que se pueden comparar por igual o distinto.
Un conjunto no admite elementos repetidos ni garabtiza un orden en los elementos

En [SetList](./set_list.go) se encuentra una
implementación de un conjunto utilizando como estructura auxiliar una lista enlazada simple [single_linked_list](../lists/single_linked_list)

En [Set](./set.go) se encuentra la interfaz con los metodos que un conjunto debe tener

En [demo](./demo/main.go) se encuentra un ejemplo de uso de conjuntos.

## Métodos comunes

### `NewSetList[T types.Ordered](elements ...T) *Set[T]`

Crea y devuelve una nueva instancia de un conjunto.

### `Contains(element T) bool`

Devuelve si el conjunto contiene el elemento indicado.

### `(s *SetList[T]) Add(elements ...T)`
Agrega uno o mas ementos al conjunto, si estos ya no estan en el conjunto.

### `Remove(element T)`
Elimina el elemento indicado del conjunto.

### `Size() int`
Devuelve el tamaño del conjunto.

### `Values() []T`
Devuelve un array con los elementos del conjunto.

### `String() string`
Devuelve un string con la respresentación del conjunto.
