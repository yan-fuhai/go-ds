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

package list

import (
	"reflect"
	"testing"
)

var list = []int{1, 2, 3, 4, 5}

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()
	if !l.Empty() {
		t.Errorf("Empty() must be %v, got %v", true, false)
	}
	for _, v := range list {
		l.Append(v)
	}
	if _, err := l.Get(len(list)); err == nil {
		t.Errorf("Get() must be error")
	}

	for i := 0; i < len(list); i++ {
		got, _ := l.Get(i)
		if got != list[i] {
			t.Errorf("Get() must be %v, got %v", list[i], got)
		}
	}
	if err := l.Remove(0); err != nil {
		if l.Length() != len(list)-1 {
			t.Errorf("Length() must be %v, got %v", len(list)-1, l.Length())
		}
	}
	if err := l.Insert(len(list), 100); err == nil {
		t.Errorf("Insert() must be error")
	}
	if err := l.Insert(0, 100); err == nil {
		if val, _ := l.Get(0); val != 100 {
			t.Errorf("Get() must be %v, got %v", 100, val)
		}
	}
	if l.Length() != len(list) {
		t.Errorf("Length() must be %v, got %v", len(list), l.Length())
	}
	if got := l.Index(100); got != 0 {
		t.Errorf("Index() must be %v, got %v", 0, got)
	}
	if got := l.Index(5); got != 4 {
		t.Errorf("Index() must be %v, got %v", 4, got)
	}
	l.Clear()
	if l.Length() != 0 {
		t.Errorf("Length() must be 0, got %v", l.Length())
	}
}

func TestInsertAndRemove(t *testing.T) {
	l := NewLinkedList()
	for i, v := range list {
		_ = l.Insert(i, v)
	}
	if l.Length() != len(list) {
		t.Errorf("Length() must be %v, got %v", len(list), l.Length())
	}

	for i := 0; i < len(list); i++ {
		_ = l.Remove(0)
	}
	if l.Length() != 0 {
		t.Errorf("Length() must be 0, got %v", l.Length())
	}
}

func TestToSlice(t *testing.T) {
	l := NewLinkedList()
	for i, v := range list {
		_ = l.Insert(i, v)
	}
	if reflect.DeepEqual(l.ToSlice(), list) {
		t.Errorf("ToSlice() must deeply equal to list")
	}
}
