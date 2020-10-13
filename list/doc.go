package list

type List interface {
	Append(interface{})
	Clear()
	Empty() bool
	Get(int) (interface{}, error)
	Length() int
	Remove(int) error
	Index(interface{}) int
	Insert(int, interface{}) error
	ToSlice() []interface{}
}
