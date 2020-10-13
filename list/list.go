package list

import "fmt"

type listNode struct {
	val   interface{}
	left  *listNode
	right *listNode
}

type LinkedList struct {
	length int
	head   *listNode
	tail   *listNode
}

func NewLinkedList() *LinkedList {
	head, tail := &listNode{}, &listNode{}
	head.right, tail.left = tail, head
	return &LinkedList{
		length: 0,
		head:   head,
		tail:   tail,
	}
}

func (l *LinkedList) Append(v interface{}) {
	l.addToTail(&listNode{
		val: v,
	})
	l.length++
}

func (l *LinkedList) Clear() {
	l.head.right, l.tail.left = l.tail, l.head
	l.length = 0
}

func (l *LinkedList) Empty() bool {
	return l.length == 0
}

func (l *LinkedList) Get(idx int) (interface{}, error) {
	if idx >= l.length {
		return nil, l.outOfRangeError(idx)
	}
	return l.getNodeByIndex(idx).val, nil
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) Remove(idx int) error {
	if idx >= l.length {
		return l.outOfRangeError(idx)
	}
	l.removeNode(l.getNodeByIndex(idx))
	l.length--
	return nil
}

func (l *LinkedList) Index(v interface{}) int {
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

func (l *LinkedList) Insert(idx int, v interface{}) error {
	if idx > l.length {
		return fmt.Errorf("index must be in [0, %v], got index %v", l.length, idx)
	}
	replacedNode := l.getNodeByIndex(idx)
	newNode := &listNode{
		val:   v,
		left:  replacedNode.left,
		right: replacedNode,
	}
	replacedNode.left.right, replacedNode.left = newNode, newNode
	l.length++
	return nil
}

func (l *LinkedList) getNodeByIndex(idx int) *listNode {
	n := l.head.right
	for i := 0; i < idx && n != nil; i++ {
		n = n.right
	}
	return n
}

func (l *LinkedList) addToTail(node *listNode) {
	node.left, node.right = l.tail.left, l.tail
	l.tail.left.right, l.tail.left = node, node
}

func (l *LinkedList) removeNode(node *listNode) {
	node.right.left, node.left.right = node.left, node.right
}

func (l *LinkedList) outOfRangeError(index int) error {
	return fmt.Errorf("index out of range [%v] with length %v", index, l.length)
}
