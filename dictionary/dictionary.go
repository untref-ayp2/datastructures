// Package dictionary proporciona una implementación de un diccionario genérico basado en un map de Go.
// Expone la estructura Dictionary y sus métodos para manipular un diccionario
package dictionary

import (
	"errors"
	"fmt"
)

// Dictionary implementa un diccionario genérico basado en un map de Go.
// Las claves deben ser de un tipo comparable y los valores pueden ser de cualquier tipo.
type Dictionary[K comparable, V any] struct {
	dict map[K]V
}

// NewDictionary crea un nuevo diccionario vacío.
//
// Uso:
//
//	d := dictionary.NewDictionary[int, string]() // Crea un diccionario de enteros a cadenas.
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{dict: make(map[K]V)}
}

// Put inserta el par (key, value) en el Diccionario. Si la `key` existe, reemplaza
// el par existente; si no existe, el par es agregado al diccionario.
//
// Uso:
//
//	d.Put(10, "ten") // Agrega el par (10, "ten") al diccionario.
//
// Parámetros:
//
//	`key`: la clave del par a insertar.
//	`value`: el valor del par a insertar.
func (d *Dictionary[K, V]) Put(key K, value V) {
	d.dict[key] = value
}

// Contains verifica si la `key` especificada existe en el diccionario
//
// Uso:
//
//	if d.Contains(10) {
//		fmt.Println("Clave 10 existe.")
//	}
//
// Parámetros:
//
//	`key`: la clave a buscar en el diccionario.
//
// Retorna:
//
//	`true` si la clave existe en el diccionario; `false` en caso contrario.
func (d *Dictionary[K, V]) Contains(key K) bool {
	_, exists := d.dict[key]

	return exists
}

// Get devuelve el `value` asociado a la `key` específicada.
//
// Uso:
//
//	if value, err := d.Get(10); err != nil {
//		fmt.Println("Valor asociado a la clave 10:", value)
//	}
//
// Parámetros:
//
//	`key`: la clave cuyo valor asociado se desea obtener.
//
// Retorna:
//
//	El valor asociado a la clave especificada  y nil como error, o el nil type del tipo V
//	y un error si la clave no existe.
func (d *Dictionary[K, V]) Get(key K) (V, error) {
	value, exists := d.dict[key]
	if !exists {
		return value, errors.New("clave inexistente")
	}

	return value, nil
}

// Remove remueve el par asociado a la `key` especificada del diccionario.
//
// Uso:
//
//	d.Remove(10)
//
// Parámetros:
//
//	`key`: la clave a remover del diccionario.
func (d *Dictionary[K, V]) Remove(key K) {
	delete(d.dict, key)
}

// Size devuelve el tamaño del diccionario.
//
// Uso:
//
//	fmt.Println("Tamaño del diccionario:", d.Size())
//
// Retorna:
//
//	La cantidad de entradas presentes en el diccionario.
func (d *Dictionary[K, V]) Size() int {
	return len(d.dict)
}

// Keys devuelve todas las claves del diccionario. Por la naturaleza de los map de Go,
// las claves no se devuelven en un orden específico. No puede contener duplicados.
//
// Uso:
//
//	keys := d.Keys()
//
// Retorna:
//
//	Un slice de todas las claves presentes en el diccionario.
func (d *Dictionary[K, V]) Keys() []K {
	dictKeys := make([]K, 0, d.Size())
	for key := range d.dict {
		dictKeys = append(dictKeys, key)
	}

	return dictKeys
}

// Values devuelve todos los valores del diccionario. Por la naturaleza de los map de Go,
// los valores no se devuelven en un orden específico. Puede contener duplicados.
//
// Uso:
//
//	values := d.Values()
//
// Retorna:
//
//	Un slice de todos los valores presentes en el diccionario.
func (d *Dictionary[K, V]) Values() []V {
	dictValues := make([]V, 0, d.Size())
	for _, value := range d.dict {
		dictValues = append(dictValues, value)
	}

	return dictValues
}

// String devuelve la representación en `string` del diccionario.
//
// Uso:
//
//	fmt.Println(d)
//
// Retorna:
//
//	Una representación en `string` del diccionario.
func (d *Dictionary[K, V]) String() string {
	if d.Size() == 0 {
		return "Dictionary: {}"
	}

	str := "Dictionary: {\n"
	for key, value := range d.dict {
		str += fmt.Sprintf("  %v: %v\n", key, value)
	}
	str += "}"

	return str
}
