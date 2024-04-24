// Package bitmap proporciona una implementación de un mapa de bits sobre un entero de 32 bits.
// Expone la estructura BitMap y sus métodos para manipular un mapa de bits.
package bitmap

import "errors"

// BitMap implementa un mapa de bits sobre un entero de 32 bits.
type BitMap struct {
	bits uint32
}

// NewBitMap crea un mapa de bits y lo inicializa con todos los bits apagados.
//
// Uso:
//
//	bm := bitmap.NewBitMap() // Crea un nuevo mapa de bits.
func NewBitMap() *BitMap {
	return &BitMap{bits: 0x0}
}

// On enciende el bit de la posición indicada.
//
// Uso:
//
//	bm.On(5) // Enciende el bit en la posición 5.
//
// Parámetros:
//   - `pos`: la posición del bit a encender.
//
// Retorna:
//   - `error` si la posición es inválida.
func (bm *BitMap) On(pos uint8) error {
	if isOutOfRange(pos) {
		return errors.New("posición no válida")
	}

	bm.bits |= 0x1 << pos

	return nil
}

// Off apaga el bit de la posición indicada.
//
// Uso:
//
//	bm.Off(5) // Apaga el bit en la posición 5.
//
// Parámetros:
//   - `pos`: la posición del bit a apagar.
//
// Retorna:
//   - `error` si la posición es inválida.
func (bm *BitMap) Off(pos uint8) error {
	if isOutOfRange(pos) {
		return errors.New("posición no válida")
	}

	bm.bits &= ^(0x1 << pos)

	return nil
}

// IsOn obtiene el estado del bit en la posición indicada.
//
// Uso:
//
//	on, _ := bm.IsOn(5) // Verifica si el bit en la posición 5 está encendido.
//
// Parámetros:
//   - `pos`: la posición del bit a verificar.
//
// Retorna:
//   - `true` si el bit está encendido; `false` si está apagado.
func (bm *BitMap) IsOn(pos uint8) (bool, error) {
	if isOutOfRange(pos) {
		return false, errors.New("posición no válida")
	}

	return bm.bits&(1<<pos) != 0x0, nil
}

// GetMap obtiene la representación interna del mapa de bits
//
// Uso:
//
//	bits := bm.GetMap() // Obtiene la representación interna del mapa de bits.
//
// Retorna:
//   - La representación interna del mapa de bits como un uint32.
func (bm *BitMap) GetMap() uint32 {
	return bm.bits
}

// isOutOfRange verifica si la posición está fuera de rango.
//
// Uso:
//
//	if isOutOfRange(32) {
//		fmt.Println("Posición no válida")
//	}
//
// Retorna:
//   - `true` si la posición está fuera de rango; `false` en caso contrario.
func isOutOfRange(pos uint8) bool {
	return 32 <= pos
}
