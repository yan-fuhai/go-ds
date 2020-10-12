package queue

type GeneralQueue interface {
	Empty() bool
	Size() int
	PushBack(interface{})
	PopFront() (interface{}, error)
	Clear()
}

type DoubleEndedQueue interface {
	GeneralQueue
	PushFront(interface{})
	PopBack() (interface{}, error)
}
