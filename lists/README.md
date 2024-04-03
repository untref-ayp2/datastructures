# Listas Enlazadas

Las listas enlazadas son una estructura de datos que permite almacenar una
colección de elementos. Cada elemento de la lista se almacena en un nodo que
contiene un campo de datos y uno o dos punteros a otros nodos. Los nodos están
enlazados entre sí para formar la lista.

A diferencia de los arreglos, las listas enlazadas no tienen un tamaño fijo y
permiten insertar y eliminar elementos en cualquier posición de la lista de
manera eficiente.

En general las listas se usan como contenedores de datos para declarar nuevos
tipos de datos.

En la carpeta [single_linked_list](./single_linked_list) se encuentra una
implementación de una lista enlazada simple, donde en cada nodo se almacena un
puntero al siguiente de la lista.

En la carpeta [double_linked_list](./double_linked_list) se encuentra una
implementación de una lista enlazada doble, donde en cada nodo se almacena un
puntero al siguiente y al anterior de la lista.

En la carpeta [circular_linked_list](./circular_linked_list) se encuentra una
implementación de una lista enlazada circular, donde el último nodo apunta al
primer nodo.

En [demo](./demo/main.go) se encuentra un ejemplo de uso de las listas
enlazadas.

## Métodos comunes

### `NewList()`

Crea y devuelve una nueva instancia de la lista vacía.

### `Head() *Node`

Devuelve el primer nodo de la lista.

### `Tail() *Node`

Devuelve el último nodo de la lista.

### `Size() int`

Devuelve la cantidad de nodos en la lista.

### `IsEmpty() bool`

Devuelve `true` si la lista está vacía, `false` en caso contrario.

### `Clear()`

Elimina todos los nodos de la lista.

### `Prepend(data T)`

Agrega un nuevo nodo al principio de la lista.

### `Append(data T)`

Agrega un nuevo nodo al final de la lista.

### `Find(data T) *Node`

Busca un nodo con el valor `data` en la lista y devuelve el nodo encontrado.

### `RemoveFirst()`

Elimina el primer nodo de la lista.

### `RemoveLast()`

Elimina el último nodo de la lista.

### `Remove(data T)`

Elimina el nodo con el valor `data` de la lista.

Para recorrer una lista enlazada se puede hacer uso de un bucle `for` o de un
bucle `while` que recorra los nodos de la lista. A continuación se muestra un
ejemplo de cómo recorrer una lista enlazada simple:

```go
func main() {
    list := NewList()
    list.Append(1)
    list.Append(2)
    list.Append(3)

    for node := list.Head(); node != nil; node = node.Next {
        fmt.Println(node.Data)
    }
}
```
