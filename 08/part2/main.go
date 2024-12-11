package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value rune
}

type Location struct {
	Row    int
	Column int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	var grid [][]*Node
	for _, line := range lines {
		var row []*Node
		for _, char := range line {
			row = append(row, &Node{Value: char})
		}
		grid = append(grid, row)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%c", grid[i][j].Value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n")

	antenas := make(map[rune][]*Location)
	results := make(map[string]bool)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].Value == '#' {
				//results[grid[i][j].Value] = append(results[grid[i][j].Value], &Location{Row: i, Column: j})
				continue
			}
			if grid[i][j].Value != '.' {
				antenas[grid[i][j].Value] = append(antenas[grid[i][j].Value], &Location{Row: i, Column: j})
			}
		}
	}
	for _, antena := range antenas {
		for _, first := range antena {
			for _, second := range antena {
				distanceX := second.Column - first.Column
				distanceY := second.Row - first.Row
				if distanceX == 0 || distanceY == 0 {
					continue
				}
				for w := 1; w < len(grid)*10; w++ {
					results[fmt.Sprintf("%d,%d", second.Row+(distanceY*w), second.Column+(w*distanceX))] = true
				}
			}
		}
	}

	locations := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if _, ok := results[fmt.Sprintf("%d,%d", i, j)]; ok {
				fmt.Printf("*")
				locations++
			} else {
				if grid[i][j].Value != '.' {
					locations++
				}
				fmt.Printf("%c", grid[i][j].Value)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n\n\n")
	fmt.Printf("locations: %d\n", locations)
}
