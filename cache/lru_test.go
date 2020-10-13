package cache

import "testing"

type kv struct {
	key   string
	value interface{}
}

var puts = []kv{
	{
		key:   "a",
		value: 1,
	},
	{
		key:   "a",
		value: 2,
	},
	{
		key:   "c",
		value: 3,
	},
	{
		key:   "d",
		value: 4,
	},
	{
		key:   "e",
		value: 5,
	},
	{
		key:   "f",
		value: 6,
	},
	{
		key:   "g",
		value: 7,
	},
	{
		key:   "h",
		value: 8,
	},
	{
		key:   "h",
		value: 9,
	},
}

func TestLRUCache(t *testing.T) {
	capacity := 2
	c := NewLRUCache(capacity)
	for _, p := range puts {
		c.Put(p.key, p.value)
	}

	gets := []kv{
		{
			key:   "a",
			value: nil,
		},
		{
			key:   "b",
			value: nil,
		},
		{
			key:   "h",
			value: 9,
		},
	}
	for _, g := range gets {
		got := c.Get(g.key)
		if g.value != got {
			t.Errorf("Get() must be %v, got %v", g.value, got)
		}
	}
}

func TestClear(t *testing.T) {
	c := NewLRUCache(5)
	for _, p := range puts {
		c.Put(p.key, p.value)
	}
	c.Clear()
	for _, p := range puts {
		got := c.Get(p.key)
		if got != nil {
			t.Errorf("Get() must be nil, got %v", got)
		}
	}
}
