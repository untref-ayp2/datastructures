package dictionary

import "fmt"

// Dictionary implementa un diccionario sobre un map de Go.
type Dictionary[K comparable, V any] struct {
	dict map[K]V
}

// NewDictionary crea y devuelve una nueva instancia de un diccionario.
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{dict: make(map[K]V)}
}

// Put inserta el par (key, val) en el Diccionario, si la clave existe reemplaza
// el par existente, si no existe, el par es agregado al diccionario.
func (d *Dictionary[K, V]) Put(key K, val V) {
	d.dict[key] = val
}

// Remove remueve el par del diccionario con clave `key` y retorna true si la
// clave existe o false si no.
func (d *Dictionary[K, V]) Remove(key K) bool {
	_, exists := d.dict[key]
	if exists {
		delete(d.dict, key)
	}
	return exists
}

// Contains devuelve true si la clave especificada existe en el diccionario.
func (d *Dictionary[K, V]) Contains(key K) bool {
	_, exists := d.dict[key]
	return exists
}

// Get devuelve el valor asociado a la clave específicada o el nil type del tipo
// V si la clave no existe.
func (d *Dictionary[K, V]) Get(key K) V {
	return d.dict[key]
}

// Size devuelve la cantidad de entradar hay en el diccionario.
func (d *Dictionary[K, V]) Size() int {
	return len(d.dict)
}

// Keys devuelve una lista de todas las claves del diccionario.
func (d *Dictionary[K, V]) Keys() []K {
	dictKeys := make([]K, 0, d.Size())
	for key := range d.dict {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

// Values devuelve una lista de todas los valores del diccionario.
func (d *Dictionary[K, V]) Values() []V {
	dictValues := make([]V, 0, d.Size())
	for _, value := range d.dict {
		dictValues = append(dictValues, value)
	}
	return dictValues
}

// String devuelve la presentación en string del diccionario.
func (d *Dictionary[K, V]) String() string {
	str := "Dictionary: {\n"
	for key, value := range d.dict {
		str += fmt.Sprintf("  %v: %v\n", key, value)
	}
	str += "}"
	return str
}
