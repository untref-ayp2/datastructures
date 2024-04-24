package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitMapNuevoTodoEnCero(t *testing.T) {
	m := NewBitMap()
	assert.Equal(t, uint32(0), m.GetMap())
}

func TestBitMapEncenderUnBit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

func TestBitMapEncenderUnBitEnPosicionNoValida(t *testing.T) {
	m := NewBitMap()
	err := m.On(32)
	assert.EqualError(t, err, "posición no válida")
}

func TestBitMapApagarUnBit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	m.Off(1)
	isOn, _ := m.IsOn(1)
	assert.False(t, isOn)
}

func TestBitMapApagarUnBitEnPosicionNoValida(t *testing.T) {
	m := NewBitMap()
	err := m.Off(32)
	assert.EqualError(t, err, "posición no válida")
}

func TestBitMapEncenderTodosLosBits(t *testing.T) {
	m := NewBitMap()
	for i := uint8(0); i < 32; i++ {
		m.On(i)
	}
	var max uint32 = 0xffffffff
	assert.Equal(t, max, m.GetMap())
}

func TestBitMapApagarTodosLosBits(t *testing.T) {
	m := NewBitMap()
	for i := uint8(0); i < 32; i++ {
		m.On(i)
	}
	for i := uint8(0); i < 32; i++ {
		m.Off(i)
	}
	assert.Equal(t, uint32(0), m.GetMap())
}

func TestBitMapEstadoDeUnBit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

func TestBitMapEstadoDeUnBitEnPosicionNoValida(t *testing.T) {
	m := NewBitMap()
	_, err := m.IsOn(32)
	assert.EqualError(t, err, "posición no válida")
}

func TestBitMapEncenderVariasVecesUnMismoBitNoLoApaga(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

func TestBitMapGetMap(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	assert.Equal(t, uint32(2), m.GetMap())
}
