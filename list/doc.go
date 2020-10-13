package list

type List interface {
	Append(interface{})
	Back() interface{}
	Clear()
	Empty() bool
	Front() interface{}
	Get(int) (interface{}, error)
	Length() int
	Remove(int) error
	Index(interface{}) int
	Insert(int, interface{}) error
	ToSlice() []interface{}
}
