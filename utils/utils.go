// Package utils provee funciones de utilidad para el resto de los paquetes.
package utils

// Min devuelve el menor de dos enteros.
//
// Uso:
//
//	a := 5
//	b := 3
//	utils.Min(a, b)
//
// Parámetros:
//   - `a` un entero.
//   - `b` un entero.
//
// Retorna:
//   - el menor de los dos enteros.
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Max devuelve el mayor de dos enteros.
//
// Uso:
//
//	a := 5
//	b := 3
//	utils.Max(a, b)
//
// Parámetros:
//   - `a` un entero.
//   - `b` un entero.
//
// Retorna:
//   - el mayor de los dos enteros.
func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
