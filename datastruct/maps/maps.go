package maps

import "github.com/Pilipupu/go-kit/datastruct/containers"

// Map interface that all maps implement
type Map[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	containers.Container[V]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
