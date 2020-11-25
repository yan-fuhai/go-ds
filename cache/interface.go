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

type Cache interface {
	Size() int
	Cap() int
	Resize(int)
	Clear()
}

type LRU interface {
	Cache
	Delete(k interface{})
	Keys() []interface{}
	Put(k interface{}, v interface{})
	Get(k interface{}) interface{}
}

type Set interface {
	Cache
	Has(k interface{}) bool
	Add(k interface{}) interface{}
	Delete(k interface{})
	Keys() []interface{}
	MostRU() interface{}
	LeastRU() interface{}
}

type doubleListNode struct {
	key   interface{}
	val   interface{}
	left  *doubleListNode
	right *doubleListNode
}

// moveToHead moves a node to the head of double linked-list.
func moveToHead(head, node *doubleListNode) {
	if node != nil {
		removeNode(node)
		addToHead(head, node)
	}
}

// removeNode removes a node from double linked-list.
func removeNode(node *doubleListNode) {
	node.right.left, node.left.right = node.left, node.right
}

// addToHead adds a node to the head of double linked-list.
func addToHead(head, node *doubleListNode) {
	node.left, node.right = head, head.right
	head.right.left, head.right = node, node
}

// removeTail removes the tail node of double linked-list.
func removeTail(head, tail *doubleListNode) *doubleListNode {
	if head.right != tail {
		ret := tail.left
		removeNode(tail.left)
		return ret
	}
	return nil
}
