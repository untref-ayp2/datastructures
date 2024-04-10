package mapadebits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Prueba que se enciendan todos los bits
func TestEncenderTodosDeAUno(t *testing.T) {
	m := NewMapaDeBits()
	for i := uint8(0); i < 32; i++ {
		m.Encender(i)
	}
	var max uint32 = 0
	max--
	assert.Equal(t, max, m.GetMapa(), "Error en Encender todos")

}

// Prueba que al intentar encender un bit encendido este no se apague
func TestQueSigaEncendido(t *testing.T) {
	m := NewMapaDeBits()
	m.Encender(1)
	m.Encender(1)
	aux, _ := m.EstaEncendido(1)
	fmt.Println(aux)
	assert.False(t, !aux, "Error en Encender multiple")
}

// Prueba que no se pyeda encender un bit de una posicion invalida
func TestPosicionExcedida(t *testing.T) {
	m := NewMapaDeBits()
	err := m.Encender(32)
	assert.NotEqual(t, nil, err, "Error en Encender multiple")
}
