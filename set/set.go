package set

import "fmt"

// Set Implementation of SET which contains unique elements.
type Set struct {
	m map[interface{}]bool
}

// NewSet returns a new set pointer.
func NewSet() *Set {
	return &Set{
		m: make(map[interface{}]bool),
	}
}

// Add add element into Set.
func (s *Set) Add(v interface{}) {
	s.m[v] = true
}

// Has return true if v already in Set, else false.
func (s *Set) Has(v interface{}) bool {
	if _, has := s.m[v]; has {
		return true
	} else {
		return false
	}
}

// Remove remove v from Set.
// Return error if v doesn't exist in Set.
func (s *Set) Remove(v interface{}) error {
	if _, has := s.m[v]; has {
		delete(s.m, v)
		return nil
	} else {
		return fmt.Errorf("Set has not element: %v", v)
	}
}

// Discard remove v from Set and return nothing no matter v already in Set.
func (s *Set) Discard(v interface{}) {
	_ = s.Remove(v)
}

// GetAll return slice that contains all elements in this Set.
func (s *Set) GetAll() []interface{} {
	all := make([]interface{}, 0, len(s.m))
	for k := range s.m {
		all = append(all, k)
	}
	return all
}

// Clear remove all elements in the Set.
func (s *Set) Clear() {
	s.m = make(map[interface{}]bool)
}

// Size return the number of elements in this Set.
func (s *Set) Size() int {
	return len(s.m)
}
