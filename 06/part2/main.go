package main

import (
	"fmt"
	"os"
)

type Node struct {
	Val          rune
	Visited      bool
	Loopable     bool
	FakeObstacle bool
	LoopCheck    bool
	Next         []*Node // [0]=top, [1]=right, [2]=bottom, [3]=left
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

func IsLoop(start *Node, node *Node, dir int) bool {
	if node == start {
		return true
	}
	if node.Next[dir] == nil {
		return false
	}
	if node.LoopCheck && node.Next[dir].LoopCheck {
		return true
	}
	if node.Next[dir].Val == '#' || node.Next[dir].FakeObstacle {
		turn := dir + 1
		if turn > 3 {
			turn = 0
		}
		node.LoopCheck = true
		result := IsLoop(start, node, turn)
		node.LoopCheck = false
		return result
	}
	node.LoopCheck = true
	result := IsLoop(start, node.Next[dir], dir)
	node.LoopCheck = false
	return result
}

func Walk(node *Node, dir int) int {
	node.Visited = true
	if node.Next[dir] == nil {
		return 0
	}
	turn := dir + 1
	if turn > 3 {
		turn = 0
	}
	crossing := 0
	node.Next[dir].FakeObstacle = true
	if node.Next[turn] != nil && IsLoop(node, node.Next[turn], turn) {
		// You can put an obstacle in front of me and I shall get stuck.
		node.Next[dir].Loopable = true
		crossing = 1
	}
	node.Next[dir].FakeObstacle = false
	if node.Next[dir].Val == '#' {
		return crossing + Walk(node, turn)
	}
	return crossing + Walk(node.Next[dir], dir)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := BuildGraph(string(content))

	obstacleOpts := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			// Starting node.
			if rows[i][j].Val == '^' {
				// 0 = up, 1 = right, 2 = down, 3 = left
				obstacleOpts = Walk(rows[i][j], 0)
				break
			}
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
