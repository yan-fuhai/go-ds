package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet()
	newElement := 1
	s.Add(newElement)
	if !s.Has(newElement) {
		t.Error("Set add failed.")
	}
	_ = s.Remove(newElement)
	if s.Has(newElement) {
		t.Error("Set remove failed.")
	}
	if err := s.Remove(newElement); err == nil {
		t.Error("error must not be nil when remove element not exists in this Set.")
	}

	s.Clear()
	if s.Size() != 0 {
		t.Error("size of Set must be zero after calling clear function.")
	}
	all := []int{1, 2, 3, 4, 5}
	for _, e := range all {
		s.Add(e)
	}
	if s.Size() != len(all) {
		t.Errorf("size of Set must be %d", len(all))
	}
	if len(s.GetAll()) != len(all) {
		t.Errorf("result that GetAll returns must has %d elements", len(all))
	}
}
