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
	"github.com/stretchr/testify/assert"
	"testing"
)

type mqItem struct {
	val int
}

func (i mqItem) Val() interface{} {
	return i.val
}

func (i mqItem) Equal(item MinQueueItem) bool {
	cmpVal := item.(mqItem).val
	return i.val == cmpVal
}

func (i mqItem) Less(item MinQueueItem) bool {
	cmpVal := item.(mqItem).val
	return i.val < cmpVal
}

func TestMinQueue_Min(t *testing.T) {
	mq := NewMinQueue()

	mq.PushBack(mqItem{1})
	assert.Equal(t, 1, mq.Size())
	assert.Equal(t, 1, mq.Front())
	assert.Equal(t, 1, mq.Back())
	assert.Equal(t, 1, mq.Min())

	mq.PushBack(mqItem{-1})
	assert.Equal(t, 2, mq.Size())
	assert.Equal(t, 1, mq.Front())
	assert.Equal(t, -1, mq.Back())
	assert.Equal(t, -1, mq.Min())

	mq.PushBack(mqItem{-3})
	assert.Equal(t, 3, mq.Size())
	assert.Equal(t, 1, mq.Front())
	assert.Equal(t, -3, mq.Back())
	assert.Equal(t, -3, mq.Min())

	mq.PushBack(mqItem{5})
	assert.Equal(t, 4, mq.Size())
	assert.Equal(t, 1, mq.Front())
	assert.Equal(t, 5, mq.Back())
	assert.Equal(t, -3, mq.Min())

	_, _ = mq.PopFront()
	assert.Equal(t, 3, mq.Size())
	assert.Equal(t, -1, mq.Front())
	assert.Equal(t, 5, mq.Back())
	assert.Equal(t, -3, mq.Min())

	mq.Clear()
	assert.Equal(t, 0, mq.Size())
	assert.Equal(t, nil, mq.Front())
	assert.Equal(t, nil, mq.Back())
	assert.Equal(t, nil, mq.Min())
}

func BenchmarkMinQueue_PushBack(b *testing.B) {
	mq := NewMinQueue()
	for i := 0; i < b.N; i++ {
		mq.PushBack(mqItem{1})
	}
}

func BenchmarkMinQueue_Min(b *testing.B) {
	mq := NewMinQueue()
	mq.PushBack(mqItem{1})
	for i := 0; i < b.N; i++ {
		mq.Min()
	}
}
