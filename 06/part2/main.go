package main

import (
	"fmt"
	"os"
)

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
			rows[row] = append(rows[row], &Node{
				Val:       rune(content[i]),
				Next:      make([]*Node, 4),
				Visited:   false,
				LoopCheck: make([]bool, 4),
			})
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

type Node struct {
	Val          rune
	Visited      bool
	Loopable     bool
	FakeObstacle bool
	LoopCheck    []bool // Track all directions
	Next         []*Node
}

func Walk(node *Node, dir int, visit bool) bool {
	if visit {
		node.Visited = true
	}
	if node.Next[dir] == nil {
		return true
	}
	if node.LoopCheck[dir] {
		return false
	}
	if node.Next[dir].Val == '#' {
		turn := dir + 1
		if turn > 3 {
			turn = 0
		}
		node.LoopCheck[dir] = true
		result := Walk(node, turn, visit)
		node.LoopCheck[dir] = false
		return result
	}
	node.LoopCheck[dir] = true
	result := Walk(node.Next[dir], dir, visit)
	node.LoopCheck[dir] = false
	return result
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := BuildGraph(string(content))

	obstacleOpts := 0
	startI := 0
	startJ := 0

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j].Val == '^' {
				// 0 = up, 1 = right, 2 = down, 3 = left
				Walk(rows[i][j], 0, true)
				startI = i
				startJ = j
				break
			}
		}
	}

	for k := 0; k < len(rows); k++ {
		for l := 0; l < len(rows[k]); l++ {
			// if [k][l] was visited, make it a fake obstacle.
			node := rows[k][l]
			oldVal := node.Val
			if node.Visited && node.Val == '.' {
				node.Val = '#'
			}
			if !Walk(rows[startI][startJ], 0, false) {
				node.Loopable = true
				obstacleOpts++
			}
			node.Val = oldVal
		}
	}

	total := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j].Loopable {
				fmt.Printf("O")
			} else if rows[i][j].Visited && rows[i][j].Val != '^' {
				fmt.Printf("X")
				total++
			} else {
				fmt.Printf("%s", string(rows[i][j].Val))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nTotal Visisted: %d\n", total)
	fmt.Printf("\nObstacle Options: %d\n", obstacleOpts)
}
