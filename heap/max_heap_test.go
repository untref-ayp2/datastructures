package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearMonticuloDeMaxVacio(t *testing.T) {
	m := NewMaxHeap[int]()
	assert.Equal(t, 0, m.Len())
}

func TestRemoveMaxDeMonticuloDeMaxVacio(t *testing.T) {
	m := NewMaxHeap[int]()
	_, err := m.RemoveMax()
	assert.NotNil(t, err)
}

func TestCrearMonticuloInsertarYExtraer(t *testing.T) {
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
		{44, 29},
		{58, 29, 44},
		{58, 29, 44, 2},
		{98, 58, 44, 2, 29},
		{98, 58, 44, 2, 29, 11},
		{98, 58, 65, 2, 29, 11, 44},
		{98, 58, 65, 3, 29, 11, 44, 2},
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{99, 98, 65, 58, 68, 11, 44, 2, 3, 29},
	}

	// Verificaciones iniciales
	m := NewMaxHeap[int]()
	assert.Equal(t, 0, m.Len())

	// Verificaciones a medida que vamos insertando
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.Insert(secuenciaDeInsercion[i])
		assert.Equal(t, ordenEsperadoDespuesDeInsertar[i], m.elements)
	}

	ordenEsperadoDespuesDeEliminar := [][]int{
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{68, 58, 65, 3, 29, 11, 44, 2},
		{65, 58, 44, 3, 29, 11, 2},
		{58, 29, 44, 3, 2, 11},
		{44, 29, 11, 3, 2},
		{29, 3, 11, 2},
		{11, 3, 2},
		{3, 2},
		{2},
		{},
	}

	for i := 0; i < m.Len(); i++ {
		m.RemoveMax()
		assert.Equal(t, ordenEsperadoDespuesDeEliminar[i], m.elements)
	}
}
