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

package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUSet(t *testing.T) {
	s := NewSet(5)
	for i := 0; i < 200; i++ {
		s.Add(i)
	}
	assert.Equal(t, 5, s.Size())
	fmt.Println(s.Keys())
}
