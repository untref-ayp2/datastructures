package tree

import "github.com/untref-ayp2/data-structures/types"

type Tree[T types.Ordered] interface {
	Insert(k T)
	Remove(k T)
	FindMin() T
	FindMax() T
	IsEmpty() bool
	Clear()
	String() string
	Iterator() types.Iterator[T]
}
