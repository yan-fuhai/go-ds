package queue

import "fmt"

type dequeNode struct {
	val   interface{}
	left  *dequeNode
	right *dequeNode
}

type Deque struct {
	size int
	head *dequeNode
	tail *dequeNode
}

// NewDeque returns a new Deque pointer
func NewDeque() *Deque {
	head, tail := &dequeNode{}, &dequeNode{}
	head.right, tail.left = tail, head
	return &Deque{
		size: 0,
		head: head,
		tail: tail,
	}
}

// Empty returns true if deque is empty, else false.
func (q *Deque) Empty() bool {
	return q.size == 0
}

// Size returns the number of elements in this deque.
func (q *Deque) Size() int {
	return q.size
}

// PushBack pushes a new element at the back of this deque.
func (q *Deque) PushBack(v interface{}) {
	newNode := &dequeNode{
		val:   v,
		left:  q.tail.left,
		right: q.tail,
	}
	q.tail.left.right, q.tail.left = newNode, newNode
	q.size++
}

// PushBack pushes a new element at the front of this deque.
func (q *Deque) PushFront(v interface{}) {
	newNode := &dequeNode{
		val:   v,
		left:  q.head,
		right: q.head.right,
	}
	q.head.right.left, q.head.right = newNode, newNode
	q.size++
}

// PopFront pops the front element of this deque.
// Error != nil only if this deque is empty.
func (q *Deque) PopFront() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("can not pop element from empty Queue")
	}
	ret := q.head.right.val
	q.head.right.right.left, q.head.right = q.head, q.head.right.right
	q.size--
	return ret, nil
}

// PopBack pops the back element of this deque.
// Error != nil only if this deque is empty.
func (q *Deque) PopBack() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("can not pop element from empty Queue")
	}
	ret := q.tail.left.val
	q.tail.left.left.right, q.tail.left = q.tail, q.tail.left.left
	q.size--
	return ret, nil
}

// Clear removes all elements in this deque.
func (q *Deque) Clear() {
	q.head.left = q.tail
	q.tail.left = q.head
	q.size = 0
}
