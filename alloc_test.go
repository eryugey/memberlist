/*
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

import "testing"

func TestPowerOf2(t *testing.T) {
	tests := []struct {
		input    int
		expected int
		desc     string
	}{
		// Edge cases
		{-1, 0, "negative value returns 0"},
		{0, 0, "zero returns 0"},
		// Power of 2 values
		{1, 0, "1 -> 2^0"},
		{2, 1, "2 -> 2^1"},
		{4, 2, "4 -> 2^2"},
		{8, 3, "8 -> 2^3"},
		{16, 4, "16 -> 2^4"},
		{1024, 10, "1024 -> 2^10"},
		{1 << 20, 20, "1MB -> 2^20"},
		// Non-power of 2 values
		{3, 2, "3 -> 2^2 = 4"},
		{5, 3, "5 -> 2^3 = 8"},
		{7, 3, "7 -> 2^3 = 8"},
		{9, 4, "9 -> 2^4 = 16"},
		{15, 4, "15 -> 2^4 = 16"},
		{1023, 10, "1023 -> 2^10 = 1024"},
		{1025, 11, "1025 -> 2^11 = 2048"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			result := powerOf2(tt.input)
			if result != tt.expected {
				t.Errorf("powerOf2(%d) = %d, expected %d", tt.input, result, tt.expected)
				return
			}
		})
	}
}

func TestPowerOf2Property(t *testing.T) {
	// Property-based test: verify that 2^result >= input for all inputs and 2^(result-1) <
	// input for inputs > 1
	for input := 1; input <= 10000; input++ {
		result := powerOf2(input)
		powerValue := 1 << result
		if powerValue < input {
			t.Errorf("powerOf2(%d) = %d, but 2^%d = %d < %d", input, result, result, powerValue, input)
			return
		}
		if input > 1 && result > 0 {
			prevPower := 1 << (result - 1)
			if prevPower >= input {
				t.Errorf("powerOf2(%d) = %d, but 2^%d = %d >= %d (should be smaller)", input, result, result-1, prevPower, input)
				return
			}
		}
	}
}
