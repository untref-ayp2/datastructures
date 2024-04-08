# Conjuntos

El TAD conjunto representa matemáticamente un conjunto finito de elementos que se pueden comparar por igual o distinto.
Un conjunto no admite elementos repetidos ni garabtiza un orden en los elementos

En [set](./set.go) se encuentra una
implementación de un conjunto utilizando como estructura auxiliar una lista enlazada simple [single_linked_list](../lists/single_linked_list)

En [demo](./demo/main.go) se encuentra un ejemplo de uso de las listas
enlazadas.

## Métodos comunes

### `NewSet[T types.Ordered](elements ...T) *Set[T]`

Crea y devuelve una nueva instancia de un conjunto.

### `Contains(element T) bool`

Devuelve si el conjunto contiene el elemento indicado.

### `Add(element T)`
Agrega un elemento al conjunto si este ya no lo contoene.

### `Remove(element T)`
Elimina el elemento indicado del conjunto.

### `Size() int`
Devuelve el tamaño del conjunto.

### `Values() []T`
Devuelve un array con los elementos del conjunto.

### `String() string`
Devuelve un string con la respresentación del conjunto.

### `Union[T types.Ordered](s1, s2 *Set[T]) *Set[T]`
Devuelve un conjunto que es la union entre dos conjuntos(los elementos de que estan en s1 mas los que estan en s2).

### `Intersection[T types.Ordered](s1, s2 *Set[T]) *Set[T]`
Devuelve un conjunto que es la intersección entre dos conjuntos(los elementos de que estan en s1 y tambien en s2).

### `Difference[T types.Ordered](s1, s2 *Set[T]) *Set[T]`
Devuelve un conjunto que es la diferencia entre dos conjuntos(los elementos de s1 menos los elementos que estan en s2).

### `Subset[T types.Ordered](s1, s2 *Set[T]) bool`
Devuelve si s2 es un sunconjunto de s1

### `Equal[T types.Ordered](s1, s2 *Set[T]) bool`
Devuelve si dos conjunros tienen los mismos elementos
