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

package cache

type set struct {
	capacity int
	head     *doubleListNode
	tail     *doubleListNode
	keyMap   map[interface{}]*doubleListNode
}

// NewSet returns a new set pointer.
func NewSet(capacity int) Set {
	head, tail := &doubleListNode{}, &doubleListNode{}
	head.right, tail.left = tail, head
	return &set{
		capacity: capacity,
		head:     head,
		tail:     tail,
		keyMap:   make(map[interface{}]*doubleListNode),
	}
}

// Size returns the size of this cache. The size would not exceed the capacity of this cache.
func (s *set) Size() int {
	return len(s.keyMap)
}

// Cap returns the capacity of this cache.
func (s *set) Cap() int {
	return s.capacity
}

// Resize sets a new capacity for this cache.
func (s *set) Resize(capacity int) {
	for len(s.keyMap) > capacity {
		delete(s.keyMap, removeTail(s.head, s.tail).key)
	}
	s.capacity = capacity
}

// Clear clears this cache.
func (s *set) Clear() {
	s.head.right = s.tail
	s.tail.left = s.head
	s.keyMap = make(map[interface{}]*doubleListNode)
}

// Has returns true if k already exist in this cache, else false.
func (s *set) Has(k interface{}) bool {
	nPtr, has := s.keyMap[k]
	if has {
		moveToHead(s.head, nPtr)
	}
	return has
}

// Add adds a new item into this set.
// If the size of this set exceed capacity after this operation, the Cache item would be removed and be returned.
func (s *set) Add(k interface{}) interface{} {
	if nPtr, has := s.keyMap[k]; has {
		moveToHead(s.head, nPtr)
	} else {
		newNode := &doubleListNode{
			key:   k,
			val:   nil,
			left:  nil,
			right: nil,
		}
		s.keyMap[k] = newNode
		addToHead(s.head, newNode)
	}

	if len(s.keyMap) > s.capacity {
		tail := removeTail(s.head, s.tail)
		delete(s.keyMap, tail.key)
		return tail.key
	}
	return nil
}

// Delete deletes an item from this set and returns nothing, no matter whether or not it already exist in this set.
func (s *set) Delete(k interface{}) {
	if nPtr, has := s.keyMap[k]; has {
		removeNode(nPtr)
		delete(s.keyMap, k)
	}
}

// Keys return a slice which contains all unique items in this set.
func (s *set) Keys() []interface{} {
	all := make([]interface{}, 0, len(s.keyMap))
	p := s.head.right
	for p != s.tail {
		all = append(all, p.key)
		p = p.right
	}
	return all
}

// MostRU returns the most recently used items in this set.
func (s *set) MostRU() interface{} {
	if s.head.right != s.tail && s.tail.left != s.head {
		return s.head.right.key
	} else {
		return nil
	}
}

// LeastRU returns the least recently used items in this set.
func (s *set) LeastRU() interface{} {
	if s.head.right != s.tail && s.tail.left != s.head {
		return s.tail.left.key
	} else {
		return nil
	}
}
