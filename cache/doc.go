package cache

type Cache interface {
	Put(k interface{}, v interface{})
	Get(k interface{}) interface{}
	Clear()
}
