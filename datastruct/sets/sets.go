package sets

import "github.com/Pilipupu/go-kit/datastruct/containers"

// Set interface that all sets implement
type Set[T comparable] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
