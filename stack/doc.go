package stack

type GeneralStack interface {
	Clear()
	Empty() bool
	Push(interface{})
	Pop() (interface{}, error)
	Size() int
	Top() interface{}
}
