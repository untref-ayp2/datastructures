package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeapCrearVacio(t *testing.T) {
	m := NewMinHeap[int]()
	assert.Equal(t, 0, m.Len())
}

func TestMinHeapRemoveMaxVacio(t *testing.T) {
	m := NewMinHeap[int]()
	_, err := m.RemoveMin()
	assert.NotNil(t, err)
}

func TestMinHeapCrearInsertarYExtraer(t *testing.T) {
	/*
		 Gracias a visualgo.net/en/heap
		 por la ayuda para preparar este caso de prueba.

		 Insertando los siguientes elementos en orden:
		 44, 29, 58, 2, 98, 11, 65, 3, 68, 99

		 El arbol resultante debería ser:
		      		       [99]
			       [98]            [65]
		   [58]        [68]   [11]  [44]
		 [2]  [3]   [29]

		 Como arreglo:
		 [99, 98, 65, 58, 68, 11, 44, 2, 3, 29]
	*/
	secuenciaDeInsercion := []int{44, 29, 58, 2, 98, 11, 65, 3, 68, 99}

	ordenEsperadoDespuesDeInsertar := [][]int{
		{44},
		{29, 44},
		{29, 44, 58},
		{2, 29, 58, 44},
		{2, 29, 58, 44, 98},
		{2, 29, 11, 44, 98, 58},
		{2, 29, 11, 44, 98, 58, 65},
		{2, 3, 11, 29, 98, 58, 65, 44},
		{2, 3, 11, 29, 98, 58, 65, 44, 68},
		{2, 3, 11, 29, 98, 58, 65, 44, 68, 99},
	}

	// Verificaciones iniciales
	m := NewMinHeap[int]()
	assert.Equal(t, 0, m.Len())

	// Verificaciones a medida que vamos insertando
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.Insert(secuenciaDeInsercion[i])
		assert.Equal(t, ordenEsperadoDespuesDeInsertar[i], m.elements)
	}

	ordenEsperadoDespuesDeEliminar := [][]int{
		{3, 29, 11, 44, 98, 58, 65, 99, 68},
		{11, 29, 58, 44, 98, 68, 65, 99},
		{29, 44, 58, 99, 98, 68, 65},
		{44, 65, 58, 99, 98, 68},
		{58, 65, 68, 99, 98},
		{65, 98, 68, 99},
		{68, 98, 99},
		{98, 99},
		{99},
		{},
	}

	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.RemoveMin()
		assert.Equal(t, ordenEsperadoDespuesDeEliminar[i], m.elements)
	}
}
