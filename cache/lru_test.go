package cache

import "testing"

func TestLRUCache(t *testing.T) {
	capacity := 5
	c := NewLRUCache(capacity)

	k, v := 5, 1
	c.Put(k, v)
	if c.Get(k) != 1 {
		t.Errorf("Get() must be %v", v)
	}
	k, v = 5, 2

}
