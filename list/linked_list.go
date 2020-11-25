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

package list

import "fmt"

// linkedList was implemented by double linked nodes.
type linkedList struct {
	length int
	head   *doubleNode
	tail   *doubleNode
}

// NewLinkedList returns a new linked-list pointer.
func NewLinkedList() List {
	head, tail := &doubleNode{}, &doubleNode{}
	head.right, tail.left = tail, head
	return &linkedList{
		length: 0,
		head:   head,
		tail:   tail,
	}
}

// Back returns the last element of list or nil if the list is empty.
func (l *linkedList) Back() interface{} {
	if l.Empty() {
		return nil
	}
	return l.tail.left.val
}

// Front returns the first element of list of nil if the list if empty.
func (l *linkedList) Front() interface{} {
	if l.Empty() {
		return nil
	}
	return l.head.right.val
}

// Append appends new element at the tail of list.
func (l *linkedList) Append(v interface{}) {
	l.addToTail(&doubleNode{
		val: v,
	})
	l.length++
}

// Clear removes all element in this list.
func (l *linkedList) Clear() {
	l.head.right, l.tail.left = l.tail, l.head
	l.length = 0
}

// Empty return true if this list is empty, else false.
func (l *linkedList) Empty() bool {
	return l.length == 0
}

// Get returns a element corresponding with index idx.
func (l *linkedList) Get(idx int) (interface{}, error) {
	if idx >= l.length {
		return nil, l.outOfRangeError(idx)
	}
	return l.getNodeByIndex(idx).val, nil
}

// Length returns the length of this list.
func (l *linkedList) Length() int {
	return l.length
}

// Remove removes the element corresponding with index idx.
func (l *linkedList) Remove(idx int) error {
	if idx >= l.length {
		return l.outOfRangeError(idx)
	}
	l.removeNode(l.getNodeByIndex(idx))
	l.length--
	return nil
}

// Index returns the first-occur index of v.
// It would return -1 when v was not found.
func (l *linkedList) Index(v interface{}) int {
	i, p := 0, l.head.right
	for i < l.length && p != nil {
		if p.val == v {
			return i
		}
		i++
		p = p.right
	}
	return -1
}

// Insert inserts new element v with index idx.
// The insert range is [0, list.length].
// After calling Insert(), the length of list would increment by one.
func (l *linkedList) Insert(idx int, v interface{}) error {
	if idx > l.length {
		return fmt.Errorf("index must be in [0, %v], got index %v", l.length, idx)
	}
	l.addBefore(&doubleNode{val: v}, l.getNodeByIndex(idx))
	l.length++
	return nil
}

func (l *linkedList) ToSlice() []interface{} {
	s := make([]interface{}, 0, l.length)
	p := l.head.right
	for p != l.tail {
		s = append(s, p.val)
		p = p.right
	}
	return s
}

// addBefore add newNode at the place before curNode
// ... <-> newNode <-> curNode <-> ...
func (l *linkedList) addBefore(newNode, curNode *doubleNode) {
	newNode.left, newNode.right = curNode.left, curNode
	curNode.left.right, curNode.left = newNode, newNode
}

func (l *linkedList) getNodeByIndex(idx int) *doubleNode {
	n := l.head.right
	for i := 0; i < idx && n != l.tail; i++ {
		n = n.right
	}
	return n
}

func (l *linkedList) addToTail(node *doubleNode) {
	node.left, node.right = l.tail.left, l.tail
	l.tail.left.right, l.tail.left = node, node
}

func (l *linkedList) removeNode(node *doubleNode) {
	node.right.left, node.left.right = node.left, node.right
}

func (l *linkedList) outOfRangeError(index int) error {
	return fmt.Errorf("index out of range [%v] with length %v", index, l.length)
}
