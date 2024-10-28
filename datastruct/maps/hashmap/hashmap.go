package hashmap

import "fmt"

// Map holds the elements in go's native map
type Map[K comparable, V any] struct {
	m map[K]any
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: make(map[K]any)}
}

func (m *Map[K, V]) Put(key K, value V) {
	m.m[key] = value
}

func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.m[key]
	return value, ok
}

func (m *Map[K, V]) Remove(key K) {
	delete(m.m, key)
}

func (m *Map[K, V]) Size() int {
	return len(m.m)
}

func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	clear(m.m)
}

// String returns a string representation of containers
func (m *Map[K, V]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}

func (m *Map[K, V]) PutIfNotExist() {

}
