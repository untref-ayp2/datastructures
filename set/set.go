package set

type Set[T comparable] interface {
	Contains(element T) bool
	Add(elements ...T)
	Remove(element T)
	Size() int
	Values() []T
}
