/*
 * JuiceFS, Copyright 2020 Juicedata, Inc.
 * Copyright 2026 Alibaba Cloud, Inc. or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package memberlist

import (
	"fmt"
	"math/bits"
	"sync"
)

// Alloc returns size bytes memory from Go heap.
func Alloc(size int) []byte {
	zeros := powerOf2(size)
	b := *pools[zeros].Get().(*[]byte)
	if cap(b) < size {
		panic(fmt.Sprintf("%d < %d", cap(b), size))
	}
	return b[:size]
}

// Free returns memory to Go heap.
func Free(b []byte) {
	// buf could be zero length
	pools[powerOf2(cap(b))].Put(&b)
}

var pools []*sync.Pool

// PowerOf2 returns the smallest power of 2 that is >= s
func powerOf2(s int) int {
	if s <= 0 {
		return 0
	}
	// Find position of the most significant bit (MSB)
	return bits.Len(uint(s - 1))
}

func init() {
	pools = make([]*sync.Pool, 30) // 1 - 1G
	for i := 0; i < 30; i++ {
		func(bits int) {
			pools[i] = &sync.Pool{
				New: func() interface{} {
					b := make([]byte, 1<<bits)
					return &b
				},
			}
		}(i)
	}
}
