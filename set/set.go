// Copyright (c) 2020 Fuhai Yan.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set

import (
	"fmt"
)

// set Implementation of SET which contains unique elements.
type set struct {
	m map[interface{}]bool
}

// NewSet returns a new set pointer.
func NewSet() Set {
	return &set{
		m: make(map[interface{}]bool),
	}
}

// Add add element into set.
func (s *set) Add(k interface{}) {
	s.m[k] = true
}

// Has return true if k already in set, else false.
func (s *set) Has(k interface{}) bool {
	_, has := s.m[k]
	return has
}

// Empty returns true if set is empty, else false.
func (s *set) Empty() bool {
	return len(s.m) == 0
}

// Remove remove k from set.
// Return error if k doesn't exist in set.
func (s *set) Remove(k interface{}) error {
	if _, has := s.m[k]; has {
		delete(s.m, k)
		return nil
	} else {
		return fmt.Errorf("set has not element: %v", k)
	}
}

// Discard remove k from set and return nothing no matter k already in set.
func (s *set) Discard(k interface{}) {
	_ = s.Remove(k)
}

// GetAll return slice that contains all elements in this set.
func (s *set) GetAll() []interface{} {
	all := make([]interface{}, 0, len(s.m))
	for k := range s.m {
		all = append(all, k)
	}
	return all
}

// Clear remove all elements in the set.
func (s *set) Clear() {
	s.m = make(map[interface{}]bool)
}

// Size return the number of elements in this set.
func (s *set) Size() int {
	return len(s.m)
}
