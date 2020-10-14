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

package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Errorf("Size() must be %v", 0)
	}
	e := 1
	s.Push(e)
	if s.Size() != 1 {
		t.Errorf("Size() must be %v", e)
	}
	if s.Top() != 1 {
		t.Errorf("Top() must be %v", e)
	}
	if pop, err := s.Pop(); err == nil {
		if e != pop {
			t.Errorf("Pop() must be %v", e)
		}
	}
	if _, err := s.Pop(); err == nil {
		t.Errorf("Pop() must not be nil")
	}
}

func TestContinuousPushAndPop(t *testing.T) {
	s := NewStack()
	n := 1000000
	for i := 0; i < n; i++ {
		s.Push(i)
	}
	for i := n - 1; i >= 0; i-- {
		if pop, err := s.Pop(); err == nil {
			if i != pop {
				t.Errorf("Pop() must be %v", i)
			}
		}
	}
}
