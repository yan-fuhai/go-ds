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

package queue

import "fmt"

type Interface interface {
	Back() interface{}
	Clear()
	Empty() bool
	Front() interface{}
	GetAll() []interface{}
	PushBack(interface{})
	PopFront() (interface{}, error)
	Size() int
}

type queueNode struct {
	val  interface{}
	next *queueNode
}

type Queue struct {
	size int // the number of elements in this Queue
	head *queueNode
	tail *queueNode
}

// NewQueue returns a new Queue pointer.
func NewQueue() *Queue {
	head := &queueNode{}
	tail := head
	return &Queue{
		size: 0,
		head: head,
		tail: tail,
	}
}

// Empty returns true if Queue is empty, else false.
func (q *Queue) Empty() bool {
	return q.size == 0
}

// Size returns the number of elements in this Queue.
func (q *Queue) Size() int {
	return q.size
}

// PushBack pushes a new element at the back of this Queue.
func (q *Queue) PushBack(v interface{}) {
	q.tail.next = &queueNode{
		val:  v,
		next: nil,
	}
	q.tail = q.tail.next
	q.size++
}

// PopFront pops the front element of this Queue.
// Error != nil only if this Queue is empty.
func (q *Queue) PopFront() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("can not pop elements from empty Queue")
	}
	ret := q.head.next.val
	q.head.next = q.head.next.next
	q.size--
	return ret, nil
}

// Clear removes all elements in this queue.
func (q *Queue) Clear() {
	q.tail.val = nil
	q.tail.next = nil
	q.head = q.tail
	q.size = 0
}

// GetAll returns a slice that contains all items in this queue with FIFO order.
func (q *Queue) GetAll() []interface{} {
	all := make([]interface{}, 0, q.size)
	p := q.head.next
	for p != q.tail {
		all = append(all, p.val)
		p = p.next
	}
	return all
}

// Back returns the last item of queue.
func (q *Queue) Back() interface{} {
	if q.size != 0 {
		return q.tail.val
	}
	return nil
}

// Front returns the first item of queue.
func (q *Queue) Front() interface{} {
	if q.size != 0 {
		return q.head.next.val
	}
	return nil
}
