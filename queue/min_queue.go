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

type minQueue struct {
	*queue
	monoQueue Deque
}

func NewMinQueue() MinQueue {
	return &minQueue{
		queue:     NewQueue().(*queue),
		monoQueue: NewDeque(),
	}
}

func (q *minQueue) PushBack(v interface{}) {
	var back MinQueueItem

	item := v.(MinQueueItem)
	q.queue.PushBack(item.Val())

	for !q.monoQueue.Empty() {
		back = q.monoQueue.Back().(MinQueueItem)
		if item.Less(back) {
			_, _ = q.monoQueue.PopBack()
		} else {
			break
		}
	}

	q.monoQueue.PushBack(item)
}

func (q *minQueue) PopFront() (interface{}, error) {
	ret, err := q.queue.PopFront()
	if err == nil {
		if !q.monoQueue.Empty() {
			if front := q.monoQueue.Front().(MinQueueItem); front.Val() == ret {
				_, _ = q.monoQueue.PopFront()
			}
		}
	}
	return ret, err
}

func (q *minQueue) Clear() {
	q.queue.Clear()
	q.monoQueue.Clear()
}

func (q *minQueue) Min() interface{} {
	min := q.monoQueue.Front()
	if min != nil {
		return min.(MinQueueItem).Val()
	}
	return nil
}
