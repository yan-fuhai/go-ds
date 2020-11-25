// Copyright (c) 2020 Fuhai Yan.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet()
	newElement := 1
	s.Add(newElement)
	if !s.Has(newElement) {
		t.Error("set add failed.")
	}
	_ = s.Remove(newElement)
	if s.Has(newElement) {
		t.Error("set remove failed.")
	}
	if err := s.Remove(newElement); err == nil {
		t.Error("error must not be nil when remove element not exists in this set.")
	}

	s.Clear()
	if s.Size() != 0 {
		t.Error("size of set must be zero after calling clear function.")
	}
	all := []int{1, 2, 3, 4, 5}
	for _, e := range all {
		s.Add(e)
	}
	if s.Size() != len(all) {
		t.Errorf("size of set must be %d", len(all))
	}
	if len(s.GetAll()) != len(all) {
		t.Errorf("result that GetAll returns must has %d elements", len(all))
	}
}
