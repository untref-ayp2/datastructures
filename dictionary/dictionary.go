package dictionary

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/types"
)

type Dictionary[K types.Ordered, V any] struct {
	mapa map[K]V
}

func NewDictionary[K types.Ordered, V any]() Dictionary[K, V] {
	dict := Dictionary[K, V]{mapa: make(map[K]V)}
	return dict
}

func (dict *Dictionary[K, V]) Put(key K, val V) {
	dict.mapa[key] = val
}

func (dict *Dictionary[K, V]) Remove(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	if exists {
		delete(dict.mapa, key)
	}
	return exists
}

func (dict *Dictionary[K, V]) Contains(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	return exists
}

func (dict *Dictionary[K, V]) Get(key K) V {
	return dict.mapa[key]
}

func (dict *Dictionary[K, V]) Size() int {
	return len(dict.mapa)
}

func (dict *Dictionary[K, V]) GetKeys() []K {
	var dictKeys []K
	dictKeys = []K{}
	var key K
	for key = range dict.mapa {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

func (dict *Dictionary[K, V]) GetValues() []V {
	var dictValues []V
	dictValues = []V{}
	var key K
	for key = range dict.mapa {
		dictValues = append(dictValues, dict.mapa[key])
	}
	return dictValues
}

func (dict Dictionary[K, V]) String() string {
	var str string
	str = "Dictionary {\n"
	var key K
	for key = range dict.mapa {
		str += fmt.Sprintf("\t%v: %v\n", key, dict.mapa[key])
	}
	str += "}"
	return str
}
