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
	"fmt"
	"math"
)

// An item was similar with heap.heap_test.item. The difference is that
// this one set the type of value with empty interface, which makes it more general.
type item struct {
	value    interface{}
	priority int
}

// An PriorityQueue is an implementation of heap.
//
// ATTENTION: The top of non-empty priority queue would be the item with MAXIMUM priority.
type PriorityQueue []*item

// NewEmptyPriorityQueue returns a new empty PriorityQueue pointer.
func NewPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	return &pq
}

// Len returns length of priority queue.
func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

// Less return true if item with index i is less than item with index j, else false.
func (pq *PriorityQueue) Less(i, j int) bool {
	// In this case, the item with maximum priority would be the top of heap.
	// Reverse the '<' to '>' in below statement and then the top of heap would be the item with minimum priority.
	return (*pq)[i].priority > (*pq)[j].priority
}

// Swap swaps items locate at index i and j.
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

// Push pushes a new item at the end of priority queue with given priority.
func (pq *PriorityQueue) Push(v interface{}, priority int) {
	item := &item{
		value:    v,
		priority: priority,
	}
	*pq = append(*pq, item)
	up(pq, pq.Len()-1)
}

// Pop pops the item with maximum priority in priority queue, both the value and priority would be returned.
func (pq *PriorityQueue) Pop() (interface{}, int, error) {
	n := pq.Len()
	if n == 0 {
		return nil, math.MinInt64, fmt.Errorf("can not pop from empty heap")
	}
	top := (*pq)[0]
	pq.Swap(0, n-1)
	down(pq, 0, n-1)
	(*pq)[n-1] = nil // avoid memory leak, that is, remove pointer to last element for garbage collection
	*pq = (*pq)[:n-1]
	return top.value, top.priority, nil
}
