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

import "sync"

// lru stores key-value pairs with fixed capacity.
// It would remove the least recently used (Cache) key-value pair as it exceeds the capacity.
type lru struct {
	capacity int
	head     *doubleListNode
	tail     *doubleListNode
	keyMap   map[interface{}]*doubleListNode
	mux      *sync.RWMutex
}

// NewLRU returns a new keyValue pointer
func NewLRU(capacity int) LRU {
	head, tail := &doubleListNode{}, &doubleListNode{}
	head.right, tail.left = tail, head
	return &lru{
		capacity: capacity,
		head:     head,
		tail:     tail,
		keyMap:   make(map[interface{}]*doubleListNode),
		mux:      &sync.RWMutex{},
	}
}

// Resize set a new capacity for Cache cache.
func (l *lru) Resize(capacity int) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if l.capacity > capacity {
		for ; l.capacity != capacity; l.capacity-- {
			tail := removeTail(l.head, l.tail)
			delete(l.keyMap, tail.key)
		}
	} else {
		l.capacity = capacity
	}
}

// Delete deletes a key-value pair in this Cache cache.
func (l *lru) Delete(k interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if nPtr, ok := l.keyMap[k]; ok {
		delete(l.keyMap, k)
		removeNode(nPtr)
	}
}

// Keys returns a slice which contains all unique keys in this Cache cache.
func (l *lru) Keys() []interface{} {
	l.mux.RLock()
	defer l.mux.RUnlock()

	keys := make([]interface{}, 0, len(l.keyMap))
	for k := range l.keyMap {
		keys = append(keys, k)
	}
	return keys
}

// Put puts a new key-value pair in this cache.
// It will update the value if the key already exist in cache.
func (l *lru) Put(k interface{}, v interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if nPtr, has := l.keyMap[k]; has {
		moveToHead(l.head, nPtr)
		if v != nPtr.val {
			nPtr.val = v
		}
	} else {
		newNode := &doubleListNode{
			key:   k,
			val:   v,
			left:  nil,
			right: nil,
		}
		addToHead(l.head, newNode)
		l.keyMap[k] = newNode
		if len(l.keyMap) > l.capacity {
			delete(l.keyMap, l.tail.left.key)
			removeTail(l.head, l.tail)
		}
	}
}

// Get returns the value corresponding to the key k.
func (l *lru) Get(k interface{}) interface{} {
	l.mux.Lock()
	defer l.mux.Unlock()

	if nPtr, has := l.keyMap[k]; has {
		moveToHead(l.head, nPtr)
		return nPtr.val
	}
	return nil
}

// Clear removes all key-value pairs in this cache.
func (l *lru) Clear() {
	l.mux.Lock()
	defer l.mux.Unlock()

	l.head.right = l.tail
	l.tail.left = l.head
	l.keyMap = make(map[interface{}]*doubleListNode)
}

// Size returns the size of cache.
func (l *lru) Size() int {
	l.mux.RLock()
	defer l.mux.RUnlock()

	return len(l.keyMap)
}

// Cap returns the capacity of cache.
func (l *lru) Cap() int {
	l.mux.RLock()
	defer l.mux.RUnlock()

	return l.capacity
}
