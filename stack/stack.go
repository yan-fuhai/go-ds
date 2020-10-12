package stack

import "fmt"

type stackNode struct {
	val  interface{}
	next *stackNode
}

type Stack struct {
	size int
	top  *stackNode
}

// NewStack returns a new Stack pointer
func NewStack() *Stack {
	return &Stack{
		size: 0,
		top: &stackNode{},
	}
}

// Clear removes all elements in stack.
func (s *Stack) Clear() {
	// GC would automatically free memory.
	s.top.next = nil
	s.size = 0
}

// Empty return true if stack is empty, else false
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Push pushed an element into stack and add it to the top of stack.
func (s *Stack) Push(v interface{}) {
	s.top.next = &stackNode{
		val:  v,
		next: s.top.next,
	}
	s.size++
}

// Pop pops the top element of stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("can not pop elements from empty stack")
	}
	ret := s.top.next.val
	s.top = s.top.next
	s.size--
	return ret, nil
}

// Top returns the top element of stack.
func (s *Stack) Top() interface{} {
	if !s.Empty() {
		return s.top.next.val
	}
	return nil
}

// Size returns the size of stack.
func (s *Stack) Size() int {
	return s.size
}
