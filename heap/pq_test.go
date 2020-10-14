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
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	n := 100000
	nums := make([]int, n)
	pq := NewPriorityQueue()

	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}

	for _, n := range nums {
		pq.Push(n, n)
	}

	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		val, _, err := pq.Pop()
		assert.Equal(t, val, nums[i])
		assert.NoError(t, err)
		assert.Equal(t, pq.Len(), i)
	}
}

func TestPopFromEmptyPQ(t *testing.T) {
	pq := NewPriorityQueue()
	_, _, err := pq.Pop()
	assert.Error(t, err)
}

func BenchmarkPush(b *testing.B) {
	pq := NewPriorityQueue()
	for i := 0; i < b.N; i++ {
		pq.Push(0, i)
	}
}

func BenchmarkPop(b *testing.B) {
	pq := NewPriorityQueue()
	for i := 0; i < b.N; i++ {
		_, _, _ = pq.Pop()
	}
}
