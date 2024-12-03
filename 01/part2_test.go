package main

import (
	"sort"
	"testing"
)

func TestPart2(t *testing.T) {
	tests := []struct {
		firstNumbers  []int
		secondNumbers []int
		expected      int
	}{
		{
			firstNumbers:  []int{1},
			secondNumbers: []int{1},
			expected:      1,
		},
		{
			firstNumbers:  []int{2, 2},
			secondNumbers: []int{1, 2},
			expected:      4,
		},
		{
			firstNumbers:  []int{1, 1, 3},
			secondNumbers: []int{1, 2, 4},
			expected:      2,
		},
		{
			firstNumbers:  []int{3, 4, 2, 1, 3, 3},
			secondNumbers: []int{4, 3, 5, 3, 9, 3},
			expected:      31,
		},
	}
	for _, test := range tests {
		// assume inputs are sorted.
		sort.Ints(test.firstNumbers)
		sort.Ints(test.secondNumbers)

		result := part2(test.firstNumbers, test.secondNumbers, 0)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, result)
		}
	}
}
