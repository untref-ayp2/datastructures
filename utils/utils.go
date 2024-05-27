// Package utils provee funciones de utilidad para el resto de los paquetes.
package utils

import "github.com/untref-ayp2/data-structures/types"

// Compare devuelve
//
//	-1 si x es menor que y,
//	 0 si x es igual a y,
//	+1 si x es mayor que y.
//
// Para tipos de punto flotante, un NaN se considera menor que cualquier valor no-NaN,
// un NaN se considera igual a un NaN, y -0.0 es igual a 0.0.
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

// isNaN informa si x es un NaN sin requerir el paquete math.
// Esto siempre devolverá falso si T no es de punto flotante.
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
