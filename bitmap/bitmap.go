package bitmap

import "errors"

// BitMap implementa un mapa de bits sobre un entero de 32 bits.
type BitMap struct {
	bits uint32
}

// NewBitMap crea un mapa de bits y lo inicializa con todos los bits apagados.
func NewBitMap() *BitMap {
	return &BitMap{bits: 0x0}
}

// On enciende el bit de la posición indicada.
func (bm *BitMap) On(pos uint8) error {
	if isOutOfRange(pos) {
		return errors.New("posición inválida")
	}

	bm.bits |= 0x1 << pos
	return nil
}

// Off apaga el bit de la posición indicada.
func (bm *BitMap) Off(pos uint8) error {
	if isOutOfRange(pos) {
		return errors.New("posición inválida")
	}

	bm.bits &= ^(0x1 << pos)
	return nil
}

// IsOn devuelve true si el bit está encendido o false si está apagado.
func (bm *BitMap) IsOn(pos uint8) (bool, error) {
	if isOutOfRange(pos) {
		return false, errors.New("posición inválida")
	}

	return bm.bits&(1<<pos) != 0x0, nil
}

// GetMap devuelve la representación interna del mapa de bits como un uint32.
func (bm *BitMap) GetMap() uint32 {
	return bm.bits
}

func isOutOfRange(pos uint8) bool {
	return 32 <= pos
}
