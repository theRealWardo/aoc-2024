package main

import (
	"testing"
)

func TestGraph(t *testing.T) {
	input := `ABC
DEF
HIJ`
	graph := BuildGraph(input)

	tests := []struct {
		val  rune
		next []rune
	}{
		{
			'A',
			[]rune{'B', 'D', 'E'},
		},
		{
			'B',
			[]rune{'A', 'C', 'D', 'E', 'F'},
		},
		{
			'C',
			[]rune{'B', 'E', 'F'},
		},
		{
			'D',
			[]rune{'A', 'B', 'E', 'H', 'I'},
		},
		{
			'E',
			[]rune{'A', 'B', 'C', 'D', 'F', 'H', 'I', 'J'},
		},
		{
			'F',
			[]rune{'B', 'C', 'E', 'I', 'J'},
		},
		{
			'H',
			[]rune{'D', 'E', 'I'},
		},
		{
			'I',
			[]rune{'D', 'E', 'F', 'H', 'J'},
		},
		{
			'J',
			[]rune{'E', 'F', 'I'},
		},
	}
	for _, test := range tests {
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph[i]); j++ {
				if graph[i][j].Val == test.val {

					for k := 0; k < len(test.next); k++ {
						found := false
						for l := 0; l < len(graph[i][j].Next); l++ {
							if graph[i][j].Next[l] != nil && graph[i][j].Next[l].Val == test.next[k] {
								if found {
									t.Errorf("Found %v twice", test.next[k])
								}
								found = true
							}
						}
						if !found {
							t.Errorf("Didn't find %q for node %q", string(test.next[k]), string(test.val))
						}
					}
				}
			}
		}
	}
}
