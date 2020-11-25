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

package list

type List interface {
	Append(interface{})
	Back() interface{}
	Clear()
	Empty() bool
	Front() interface{}
	Get(int) (interface{}, error)
	Length() int
	Remove(int) error
	Index(interface{}) int
	Insert(int, interface{}) error
	ToSlice() []interface{}
}

type singleNode struct {
	val  interface{}
	next *singleNode
}

type doubleNode struct {
	val   interface{}
	left  *doubleNode
	right *doubleNode
}
