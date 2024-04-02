# Listas Enlazadas
Las listas enlazadas son una estructura de datos que permite almacenar una colección de elementos. Cada elemento de la lista se almacena en un nodo que contiene un campo de datos y uno o dos punteros a otros nodos. Los nodos están enlazados entre sí para formar la lista.
A diferencia de los arreglos, las listas enlazadas no tienen un tamaño fijo y permiten insertar y eliminar elementos en cualquier posición de la lista de manera eficiente.

En general las listas se usan como contenedores de datos para declarar nuevos tipos de datos.

En la carpeta [single-linked-list](./single-linked-list) se encuentra una implementación de una lista enlazada simple, donde en cada nodo se almacena un puntero al siguiente de la lista.

En la carpeta [double-linked-list](./double-linked-list) se encuentra una implementación de una lista enlazada doble, donde en cada nodo se almacena un puntero al siguiente y al anterior de la lista.

En la carpeta [circular-linked-list](./circular-linked-list) se encuentra una implementación de una lista enlazada circular, donde el último nodo apunta al primer nodo.

En [demo](./demo/main.go) se encuentra un ejemplo de uso de las listas enlazadas.