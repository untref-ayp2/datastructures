package set

import "github.com/untref-ayp2/data-structures/types"

type Set[T types.Ordered] interface {
	Contains(element T) bool
	Add(element T)
	Remove(element T)
	Size() int
	Values() []T
}
