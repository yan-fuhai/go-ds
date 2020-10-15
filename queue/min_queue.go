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

type MinQueueItemInterface interface {
	Val() interface{}
	Less(MinQueueItemInterface) bool
	Equal(item MinQueueItemInterface) bool
}

type MinQueue struct {
	Queue
	monoQueue *Deque
}

func NewMinQueue() *MinQueue {
	return &MinQueue{
		Queue:     *NewQueue(),
		monoQueue: NewDeque(),
	}
}

func (q *MinQueue) PushBack(v interface{}) {
	var back MinQueueItemInterface

	item := v.(MinQueueItemInterface)
	q.Queue.PushBack(item.Val())

	for !q.monoQueue.Empty() {
		back = q.monoQueue.Back().(MinQueueItemInterface)
		if item.Less(back) {
			_, _ = q.monoQueue.PopBack()
		} else {
			break
		}
	}

	q.monoQueue.PushBack(item)
}

func (q *MinQueue) PopFront() (interface{}, error) {
	ret, err := q.Queue.PopFront()
	if err == nil {
		if !q.monoQueue.Empty() {
			if front := q.monoQueue.Front().(MinQueueItemInterface); front.Val() == ret {
				_, _ = q.monoQueue.PopFront()
			}
		}
	}
	return ret, err
}

func (q *MinQueue) Clear() {
	q.Queue.Clear()
	q.monoQueue.Clear()
}

func (q *MinQueue) Min() interface{} {
	min := q.monoQueue.Front()
	if min != nil {
		return min.(MinQueueItemInterface).Val()
	}
	return nil
}
