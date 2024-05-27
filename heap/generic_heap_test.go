package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Persona struct {
	nombre string
	edad   int
}

func personasDeMayorAMenorEdad(a Persona, b Persona) int {
	if a.edad < b.edad {
		return 1
	} else if a.edad > b.edad {
		return -1
	}

	return 0
}

func TestHeapCrearVacio(t *testing.T) {
	m := NewGenericHeap[Persona](personasDeMayorAMenorEdad)
	assert.Equal(t, 0, m.Size())
}

func TestHeapRemoveVacio(t *testing.T) {
	m := NewGenericHeap[Persona](personasDeMayorAMenorEdad)
	_, err := m.Remove()
	assert.NotNil(t, err)
}

func TestHeapCrearInsertarYExtraer(t *testing.T) {
	secuenciaDeInsercion := []Persona{
		{"Ana", 44},
		{"Juan", 29},
		{"Pedro", 58},
		{"Maria", 2},
		{"Jose", 98},
		{"Lucia", 11},
		{"Carlos", 65},
		{"Sofia", 3},
		{"Miguel", 68},
		{"Laura", 99},
	}

	ordenEsperadoDespuesDeInsertar := [][]Persona{
		{{"Ana", 44}},
		{{"Ana", 44}, {"Juan", 29}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}, {"Maria", 2}},
		{{"Jose", 98}, {"Pedro", 58}, {"Ana", 44}, {"Maria", 2}, {"Juan", 29}},
		{{"Jose", 98}, {"Pedro", 58}, {"Ana", 44}, {"Maria", 2}, {"Juan", 29}, {"Lucia", 11}},
		{{"Jose", 98}, {"Pedro", 58}, {"Carlos", 65}, {"Maria", 2}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}},
		{{"Jose", 98}, {"Pedro", 58}, {"Carlos", 65}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}},
		{{"Jose", 98}, {"Miguel", 68}, {"Carlos", 65}, {"Pedro", 58}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}},
		{{"Laura", 99}, {"Jose", 98}, {"Carlos", 65}, {"Pedro", 58}, {"Miguel", 68}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}, {"Juan", 29}},
	}

	// Verificaciones iniciales
	m := NewGenericHeap[Persona](personasDeMayorAMenorEdad)
	assert.Equal(t, 0, m.Size())

	// Verificaciones a medida que vamos insertando
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.Insert(secuenciaDeInsercion[i])
		assert.Equal(t, ordenEsperadoDespuesDeInsertar[i], m.elements)
	}

	ordenEsperadoDespuesDeEliminar := [][]Persona{
		{{"Jose", 98}, {"Miguel", 68}, {"Carlos", 65}, {"Pedro", 58}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}},
		{{"Miguel", 68}, {"Pedro", 58}, {"Carlos", 65}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}},
		{{"Carlos", 65}, {"Pedro", 58}, {"Ana", 44}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Maria", 2}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}, {"Sofia", 3}, {"Maria", 2}, {"Lucia", 11}},
		{{"Ana", 44}, {"Juan", 29}, {"Lucia", 11}, {"Sofia", 3}, {"Maria", 2}},
		{{"Juan", 29}, {"Sofia", 3}, {"Lucia", 11}, {"Maria", 2}},
		{{"Lucia", 11}, {"Sofia", 3}, {"Maria", 2}},
		{{"Sofia", 3}, {"Maria", 2}},
		{{"Maria", 2}},
		{},
	}

	for i := 0; i < len(secuenciaDeInsercion); i++ {
		_, err := m.Remove()
		assert.Equal(t, ordenEsperadoDespuesDeEliminar[i], m.elements)
		assert.NoError(t, err)
	}
}
