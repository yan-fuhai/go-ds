package set

type GeneralSet interface {
	Add(interface{})
	Has(interface{}) bool
	Empty() bool
	Remove(interface{}) error
	Discard(interface{})
	GetAll() []interface{}
	Size() int
	Clear()
}
