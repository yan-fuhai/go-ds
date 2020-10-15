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

type doubleListNode struct {
	key   interface{}
	val   interface{}
	left  *doubleListNode
	right *doubleListNode
}

// LRUCache store key-value pairs with fixed capacity.
// It would remove the least recently used (LRU) key-value pair as it exceeds the capacity.
type LRUCache struct {
	capacity int
	head     *doubleListNode
	tail     *doubleListNode
	keyMap   map[interface{}]*doubleListNode
}

// Delete deletes a key-value pair in this LRU cache.
func (c *LRUCache) Delete(k interface{}) {
	if nPtr, ok := c.keyMap[k]; ok {
		delete(c.keyMap, k)
		c.removeNode(nPtr)
	}
}

// Keys returns a slice which contains all unique keys in this LRU cache.
func (c *LRUCache) Keys() []interface{} {
	keys := make([]interface{}, 0, len(c.keyMap))
	for k := range c.keyMap {
		keys = append(keys, k)
	}
	return keys
}

// NewLRUCache returns a new LRUCache pointer
func NewLRUCache(capacity int) *LRUCache {
	head, tail := &doubleListNode{}, &doubleListNode{}
	head.right, tail.left = tail, head
	return &LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		keyMap:   make(map[interface{}]*doubleListNode),
	}
}

// Put puts a new key-value pair in this cache.
// It will update the value if the key already exist in cache.
func (c *LRUCache) Put(k interface{}, v interface{}) {
	if nPtr, has := c.keyMap[k]; has {
		c.moveToHead(nPtr)
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
		c.addToHead(newNode)
		c.keyMap[k] = newNode
		if len(c.keyMap) > c.capacity {
			delete(c.keyMap, c.tail.left.key)
			c.removeTail()
		}
	}
}

// Get returns the value corresponding with the key k.
func (c *LRUCache) Get(k interface{}) interface{} {
	if nPtr, has := c.keyMap[k]; has {
		c.moveToHead(nPtr)
		return nPtr.val
	}
	return nil
}

// Clear removes all key-value pairs in this cache.
func (c *LRUCache) Clear() {
	c.head.right = c.tail
	c.tail.left = c.head
	c.keyMap = make(map[interface{}]*doubleListNode)
}

// moveToHead moves a node to the head of double linked-list.
func (c *LRUCache) moveToHead(node *doubleListNode) {
	if node != nil {
		c.removeNode(node)
		c.addToHead(node)
	}
}

// removeNode removes a node from double linked-list.
func (c *LRUCache) removeNode(node *doubleListNode) {
	node.right.left, node.left.right = node.left, node.right
}

// addToHead adds a node to the head of double linked-list.
func (c *LRUCache) addToHead(node *doubleListNode) {
	node.left, node.right = c.head, c.head.right
	c.head.right.left, c.head.right = node, node
}

// removeTail remove the tail node of double linked-list.
func (c *LRUCache) removeTail() {
	if c.head.right != c.tail {
		c.removeNode(c.tail.left)
	}
}
