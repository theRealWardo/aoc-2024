package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Val  rune
	Next []*Node
}

const XMAS string = "XMAS"

func countXMAS(node *Node, last *Node, seen string) int {
	fmt.Printf("\tseen: %v\n", seen)
	fmt.Printf("\tval: %v\n", string(node.Val))
	if seen+string(node.Val) == "XMAS" {
		fmt.Printf("\t\tDING!\n")
		return 1
	}
	next := seen + string(node.Val)
	if strings.Index(next, XMAS[0:len(next)]) == -1 {
		fmt.Println("\tnope.\n")
		return 0
	}
	count := 0
	nextLetter := XMAS[len(next)]
	fmt.Printf("\tnextLetter: %v\n\n", string(nextLetter))
	for _, n := range node.Next {
		if n == last {
			continue
		}
		fmt.Printf("\topt: %s\n", string(n.Val))
		if n.Val == rune(nextLetter) {
			fmt.Printf("\t\tRECURSE: %s\n", seen+string(node.Val))
			count += countXMAS(n, node, seen+string(node.Val))
		}
	}
	return count
}

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	rows := [][]*Node{}
	row := 0
	rows = append(rows, []*Node{})
	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			row++
			rows = append(rows, []*Node{})
		} else {
			rows[row] = append(rows[row], &Node{Val: rune(content[i])})
		}
	}

	fmt.Printf("len rows: %d\n", len(rows))
	for i := 0; i < len(rows); i++ {
		fmt.Printf("len row: %d\n", len(rows[i]))
		for j := 0; j < len(rows[i]); j++ {
			fmt.Printf("i: %d j: %d\n", i, j)
			if j > 0 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i][j-1])
			}
			if j < len(rows[i])-2 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i][j+1])
			}
			if i > 0 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i-1][j])
			}
			if i < len(rows)-2 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i+1][j])
			}
			if i > 0 && j > 0 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i-1][j-1])
			}
			if i > 0 && j < len(rows[i])-2 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i-1][j+1])
			}
			if i < len(rows)-2 && j > 0 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i+1][j-1])
			}
			if i < len(rows)-2 && j < len(rows[i])-2 {
				rows[i][j].Next = append(rows[i][j].Next, rows[i+1][j+1])
			}
		}
	}

	count := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			fmt.Printf("i: %d j: %d\n", i, j)
			count += countXMAS(rows[i][j], nil, "")
		}
	}
	fmt.Println(count)
}
