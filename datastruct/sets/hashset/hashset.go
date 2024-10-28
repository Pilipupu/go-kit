package hashset

import (
	"fmt"
	"github.com/Pilipupu/go-kit/datastruct/sets"
	"strings"
)

// Assert Set implementation
var _ sets.Set[int] = (*Set[int])(nil)

var itemExists = struct{}{}

// Set holds elements in go's native map
type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable](items ...T) *Set[T] {
	set := &Set[T]{items: make(map[T]struct{})}
	if len(items) > 0 {
		set.Add(items...)
	}
	return set
}

func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(set.items, item)
	}
}

func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

func (set *Set[T]) Empty() bool {
	return set.Size() == 0
}

func (set *Set[T]) Size() int {
	return len(set.items)
}

func (set *Set[T]) Clear() {
	set.items = make(map[T]struct{})
}

func (set *Set[T]) Values() []T {
	values := make([]T, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Set[T]) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

func (set *Set[T]) Intersection(another *Set[T]) *Set[T] {
	result := New[T]()
	// Iterate over smaller set (optimization)
	if set.Size() <= another.Size() {
		for item := range set.items {
			if _, contains := another.items[item]; contains {
				result.Add(item)
			}
		}
	} else {
		for item := range another.items {
			if _, contains := set.items[item]; contains {
				result.Add(item)
			}
		}
	}

	return result
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "another" (possibly both).
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set[T]) Union(another *Set[T]) *Set[T] {
	result := New[T]()

	for item := range set.items {
		result.Add(item)
	}
	for item := range another.items {
		result.Add(item)
	}

	return result
}

// Difference returns the difference between two sets.
// The new set consists of all elements that are in "set" but not in "another".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set[T]) Difference(another *Set[T]) *Set[T] {
	result := New[T]()

	for item := range set.items {
		if _, contains := another.items[item]; !contains {
			result.Add(item)
		}
	}

	return result
}
