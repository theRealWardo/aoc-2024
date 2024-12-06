package main

import (
	"fmt"
	"os"
)

type Node struct {
	Val     rune
	Visited bool
	Next    []*Node // [0]=top, [1]=right, [2]=bottom, [3]=left
}

// Rest of imports and constants remain the same
func BuildGraph(content string) [][]*Node {
	rows := [][]*Node{}
	row := 0
	rows = append(rows, []*Node{})
	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			row++
			rows = append(rows, []*Node{})
		} else {
			rows[row] = append(rows[row], &Node{Val: rune(content[i]), Next: make([]*Node, 4), Visited: false})
		}
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			// Top
			if i > 0 {
				rows[i][j].Next[0] = rows[i-1][j]
			}
			// Right
			if j < len(rows[i])-1 {
				rows[i][j].Next[1] = rows[i][j+1]
			}
			// Bottom
			if i < len(rows)-1 {
				rows[i][j].Next[2] = rows[i+1][j]
			}
			// Left
			if j > 0 {
				rows[i][j].Next[3] = rows[i][j-1]
			}
		}
	}
	return rows
}

func Walk(node *Node, dir int) bool {
	node.Visited = true
	if node.Next[dir] == nil {
		return true
	}
	if node.Next[dir].Val == '#' {
		turn := dir + 1
		if turn > 3 {
			turn = 0
		}
		return Walk(node, turn)
	}
	return Walk(node.Next[dir], dir)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := BuildGraph(string(content))

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			// Starting node.
			if rows[i][j].Val == '^' {
				// 0 = up, 1 = right, 2 = down, 3 = left
				Walk(rows[i][j], 0)
			}
		}
	}

	total := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j].Visited {
				fmt.Printf("X")
				total++
			} else {
				fmt.Printf("%s", string(rows[i][j].Val))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nTotal Visisted: %d\n", total)
}
