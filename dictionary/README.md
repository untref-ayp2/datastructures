# Diccionarios

El Diccionario es una estructura de datos que almacena colecciones de objetos representados por pares únicos (clave: valor).

Tiene una clave (key), y cada clave tiene un elemento asociado a ella.

En [Dictionary](./dictionary.go) se encuentra una implementación de un conjunto,
utilizando como estructura auxiliar un mapa ya proporcionado por go.

En [demo](./demo/main.go) se encuentra un ejemplo de uso de diccionarios.


## Métodos comunes

### `NewDictionary[K types.Ordered, V any]() Dictionary[K, V]`

Crea y devuelve una nueva instancia de un diccionario.

### `Contains(key K) bool`

Sí la clave existe en el diccionario retorna true, si no existe, retorna false.

### `Put(key K, val V)`

Inserta el par (key, val) en el Diccionario, si la clave existe, reemplaza el par existente, si no existe, el par es agregado al diccionario.

### `Remove(key K) bool `

Sí la clave existe, remueve el par del diccionario y retorna truem si no existe, retorna false.

### `Get(key K) V`

Sí la clave existe, devuelve el valor asociado a la clave.

### `Size() int`

Devuelve la cantidad de elementos en el diccionario.

### `GetKeys() []K`

Devuelve un array no ordenado con las keys del diccionario.

### `GetValues() []K`

Devuelve un array no ordenado con los values del diccionario.

### `String() string`

Devuelve un string con la respresentación del diccionario.
