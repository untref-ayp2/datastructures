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

En el archivo [`linked_list.go`](./linked_list.go) se encuentra una
implementación de una lista enlazada simple, donde en cada nodo se almacena un
puntero al siguiente de la lista.

En el archivo [`double_linked_list.go`](./double_linked_list.go) se encuentra
una implementación de una lista enlazada doble, donde en cada nodo se almacena
un puntero al siguiente y al anterior de la lista.

En el archivo [`circular_list.go`](./circular_list.go) se encuentra una
implementación de una lista enlazada circular, donde el último nodo apunta al
primer nodo.

## Instalación

Para utilizar estas estructuras de datos en tu proyecto Go, simplemente debe importarse el paquete `github.com/untref-ayp2/data-structures/list`

```shell
go get github.com/untref-ayp2/data-structures/list
```

## Ejemplos de uso
Se proporcionan algunos ejemplos de uso en el [repositorio de snippets de la cátedra](https://github.com/untref-ayp2/snippets)
