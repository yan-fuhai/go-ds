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
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPQ(t *testing.T) {
	pq := NewEmptyPriorityQueue()
	assert.Equal(t, 0, pq.Len())
	items := []Item{
		{
			value:    "apple",
			priority: 4,
		},
		{
			value:    "banana",
			priority: 3,
		},
		{
			value:    "peach",
			priority: 2,
		},
		{
			value:    "strawberry",
			priority: 1,
		},
	}

	pq = NewPriorityQueue(items)
	for _, item := range items {
		pop := heap.Pop(pq).(*Item)
		assert.Equal(t, pop.value, item.value)
	}
	assert.Equal(t, 0, pq.Len())

	item := Item{
		value:    "apple",
		priority: 10,
	}
	heap.Push(pq, &item)
	assert.Equal(t, 1, pq.Len())
}
