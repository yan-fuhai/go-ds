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

package queue

import (
	"testing"
)

var q = NewDeque()

func TestDeque(t *testing.T) {
	q.PushBack(1)
	q.PushBack(2)
	if q.Size() != 2 {
		t.Error("size must be 2")
	}
	if front, err := q.PopFront(); err != nil {
		if front != 1 {
			t.Error("front must be 1")
		}
	}
	if front, err := q.PopFront(); err != nil {
		if front != 2 {
			t.Error("front must be 2")
		}
	}
	if _, err := q.PopFront(); err == nil {
		t.Error("error must not be nil")
	}

	q.PushBack(1)
	q.PushBack(2)
	if q.Size() != 2 {
		t.Error("size must be 2")
	}
	if back, err := q.PopBack(); err != nil {
		if back != 2 {
			t.Error("back must be 2")
		}
	}
	if back, err := q.PopBack(); err != nil {
		if back != 1 {
			t.Error("back must be 1")
		}
	}
	if _, err := q.PopBack(); err == nil {
		t.Error("error must not be nil")
	}

	q.PushFront(2)
	q.PushFront(3)
	if q.Size() != 2 {
		t.Error("size of Queue must be 2")
	}
	if front, err := q.PopFront(); err != nil {
		if front != 3 {
			t.Error("front must be 3")
		}
	}
	if q.Size() != 1 {
		t.Error("size of Queue must be 1")
	}
	if back, err := q.PopBack(); err == nil {
		if back != 2 {
			t.Error("back must be 2")
		}
	}
	if _, err := q.PopBack(); err == nil {
		t.Error("error must not be nil")
	}

	q.PushFront(2)
	q.PushFront(3)
	if q.Size() != 2 {
		t.Error("size of Queue must be 2")
	}
	if back, err := q.PopBack(); err != nil {
		if back != 2 {
			t.Error("back must be 2")
		}
	}
	if q.Size() != 1 {
		t.Error("size of Queue must be 1")
	}
	if front, err := q.PopFront(); err == nil {
		if front != 3 {
			t.Error("front must be 2")
		}
	}
	if _, err := q.PopBack(); err == nil {
		t.Error("error must not be nil")
	}
}

func TestContinuousPush(t *testing.T) {
	var elements []int

	n := 100000
	for i := 0; i < n; i++ {
		elements = append(elements, i)
	}

	for _, e := range elements {
		q.PushBack(e)
	}
	for _, e := range elements {
		if front, err := q.PopFront(); err == nil {
			if e != front {
				t.Errorf("front must be %v", e)
			}
		}
	}
	for _, e := range elements {
		q.PushBack(e)
	}
	for i := len(elements) - 1; i >= 0; i-- {
		if back, err := q.PopBack(); err == nil {
			if elements[i] != back {
				t.Errorf("back must be %v", elements[i])
			}
		}
	}

	for _, e := range elements {
		q.PushFront(e)
	}
	for i := len(elements) - 1; i >= 0; i-- {
		if front, err := q.PopFront(); err == nil {
			if elements[i] != front {
				t.Errorf("front must be %v", elements[i])
			}
		}
	}
	for _, e := range elements {
		q.PushFront(e)
	}
	for _, e := range elements {
		if back, err := q.PopBack(); err == nil {
			if e != back {
				t.Errorf("back must be %v", e)
			}
		}
	}
}
