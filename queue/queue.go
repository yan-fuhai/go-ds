package queue

import "fmt"

type queueNode struct {
	val   interface{}
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
