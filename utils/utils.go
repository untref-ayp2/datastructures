// Package utils provee funciones de utilidad para el resto de los paquetes.
package utils

import "github.com/untref-ayp2/data-structures/types"

// Compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
func Compare[T types.Ordered](x, y T) int {
	xNaN := isNaN(x)
	yNaN := isNaN(y)
	if xNaN && yNaN {
		return 0
	}
	if xNaN || x < y {
		return -1
	}
	if yNaN || x > y {
		return +1
	}

	return 0
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T types.Ordered](x T) bool {
	return x != x
}

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
