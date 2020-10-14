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

package heap

import (
	"sort"
)

type Interface interface {
	sort.Interface
	Push(interface{}, int)
	Pop() (interface{}, int, error)
}

// up shifts up h[j] until it's less than its parent (according to Less()).
// The code below was copied from heap, the official golang package
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			// i == j means that i is the top item of this heap
			// !h.Less(j, i) means that h[j] >= [i]
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// down shifts down h[j] until it's NOT less than left child and right child (according to Less()).
// The code below was copied from heap, the official golang package.
func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			// j1 >= n means that j1 is out of range
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			// h.Less(j2, j1) means that item[j2] < item[j1].
			// In each time of down operation, it will swap current item h[i] with the smaller child
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			// current item h[i] is less than all children
			break
		}
		h.Swap(i, j)
		i = j
	}

	// i == i0 if h[i] not need to down and this function would return true in this situation
	return i > i0
}
