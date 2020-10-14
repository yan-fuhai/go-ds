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

import "container/heap"

// An Item was similar with heap.heap_test.Item. The difference is that
// this one set the type of value with empty interface, which makes it more general.
type Item struct {
	value    interface{}
	priority int
	index    int
}

// An PriorityQueue is a copy of heap.heap_test.PriorityQueue.
type PriorityQueue []*Item

// NewEmptyPriorityQueue returns a new empty PriorityQueue pointer.
func NewEmptyPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	return &pq
}

// NewPriorityQueue returns a new PriorityQueue and initializes it with items.
func NewPriorityQueue(items []Item) *PriorityQueue {
	pq := make(PriorityQueue, 0, len(items))
	heap.Init(&pq)
	for i := 0; i < len(items); i++ {
		items[i].index = i
		heap.Push(&pq, &items[i])
	}
	return &pq
}

// Len returns length of priority queue.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less return true if item with index i is less than item with index j, else false.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

// Swap swaps items locate at index i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

// Push simply pushes a new item at the end of priority queue.
// ATTENTION: This operation doesn't apply shift-down or shift-up for priority queue,
// which means that it would destroy the property of heap.
// For pushing a new item into priority queue and remaining property of heap, use heap.Push.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop simply pops the last item of priority queue.
// ATTENTION: This operation doesn't apply shift-down or shift-up for priority queue,
// which means that it would destroy the property of heap.
// For pushing a new item into priority queue and remaining property of heap, use heap.Push.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
