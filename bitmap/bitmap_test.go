package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncenderTodosDeAUno(t *testing.T) {
	m := NewBitMap()
	for i := uint8(0); i < 32; i++ {
		m.On(i)
	}
	var max uint32 = 0xffffffff
	assert.Equal(t, max, m.GetMap())

}

func TestQueSigaEncendido(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	m.On(1)
	is1On, _ := m.IsOn(1)
	assert.True(t, is1On)
}

func TestPosicionExcedida(t *testing.T) {
	m := NewBitMap()
	err := m.On(32)
	assert.Error(t, err, "posición inválida")
}
