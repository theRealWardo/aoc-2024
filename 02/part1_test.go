package main

import (
	"testing"
)

func TestSafeRow(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected bool
	}{
		{
			numbers:  []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			numbers:  []int{1, 2, 7, 8, 9},
			expected: false,
		},

		{
			numbers:  []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			numbers:  []int{1, 3, 2, 4, 5},
			expected: true,
		},
		{
			numbers:  []int{8, 6, 4, 4, 1},
			expected: true,
		},
		{
			numbers:  []int{1, 3, 6, 7, 9},
			expected: true,
		},
	}
	for _, test := range tests {
		sum := part1([][]int{test.numbers})
		result := sum == 1
		if result != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, result)
		}
	}
}
