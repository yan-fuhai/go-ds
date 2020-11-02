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

// KV stores key-value pairs with fixed capacity.
// It would remove the least recently used (LRU) key-value pair as it exceeds the capacity.
type KV struct {
	capacity int
	head     *doubleListNode
	tail     *doubleListNode
	keyMap   map[interface{}]*doubleListNode
}

// NewLRUCache returns a new KV pointer
func NewLRUCache(capacity int) *KV {
	head, tail := &doubleListNode{}, &doubleListNode{}
	head.right, tail.left = tail, head
	return &KV{
		capacity: capacity,
		head:     head,
		tail:     tail,
		keyMap:   make(map[interface{}]*doubleListNode),
	}
}

// Resize set a new capacity for LRU cache.
func (c *KV) Resize(capacity int) {
	if c.capacity > capacity {
		for ; c.capacity != capacity; c.capacity-- {
			tail := removeTail(c.head, c.tail)
			delete(c.keyMap, tail.key)
		}
	} else {
		c.capacity = capacity
	}
}

// Delete deletes a key-value pair in this LRU cache.
func (c *KV) Delete(k interface{}) {
	if nPtr, ok := c.keyMap[k]; ok {
		delete(c.keyMap, k)
		removeNode(nPtr)
	}
}

// Keys returns a slice which contains all unique keys in this LRU cache.
func (c *KV) Keys() []interface{} {
	keys := make([]interface{}, 0, len(c.keyMap))
	for k := range c.keyMap {
		keys = append(keys, k)
	}
	return keys
}

// Put puts a new key-value pair in this cache.
// It will update the value if the key already exist in cache.
func (c *KV) Put(k interface{}, v interface{}) {
	if nPtr, has := c.keyMap[k]; has {
		moveToHead(c.head, nPtr)
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
		addToHead(c.head, newNode)
		c.keyMap[k] = newNode
		if len(c.keyMap) > c.capacity {
			delete(c.keyMap, c.tail.left.key)
			removeTail(c.head, c.tail)
		}
	}
}

// Get returns the value corresponding to the key k.
func (c *KV) Get(k interface{}) interface{} {
	if nPtr, has := c.keyMap[k]; has {
		moveToHead(c.head, nPtr)
		return nPtr.val
	}
	return nil
}

// Clear removes all key-value pairs in this cache.
func (c *KV) Clear() {
	c.head.right = c.tail
	c.tail.left = c.head
	c.keyMap = make(map[interface{}]*doubleListNode)
}

// Size returns the size of cache.
func (c *KV) Size() int {
	return len(c.keyMap)
}

// Cap returns the capacity of cache.
func (c *KV) Cap() int {
	return c.capacity
}
