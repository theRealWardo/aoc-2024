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

func countXMAS(node *Node, dir int, seen string, debug bool) int {
	if debug {
		fmt.Printf("\tseen: %v\n", seen)
		fmt.Printf("\tval: %v\n", string(node.Val))
		fmt.Printf("\tdir: %v\n", dir)
	}
	if seen+string(node.Val) == "XMAS" {
		fmt.Printf("\t\tDING!\n")
		fmt.Printf("dir: %v\n", dir)
		R := node.Next[4]
		L := node.Next[3]
		if L != nil {
			fmt.Printf("L: %s\t", string(L.Val))
		}
		if R != nil {
			fmt.Printf("R: %s\n", string(R.Val))
		}
		fmt.Println("")
		return 1
	}
	next := seen + string(node.Val)
	if strings.Index(next, XMAS[0:len(next)]) == -1 {
		if debug {
			fmt.Println("\tnope.")
		}
		return 0
	}
	count := 0
	nextLetter := XMAS[len(next)]
	if debug {
		fmt.Printf("\tnextLetter: %v\n", string(nextLetter))
	}
	opt := node.Next[dir]
	if opt == nil {
		return 0
	}
	if debug {
		fmt.Printf("\topt: %s\n", string(opt.Val))
	}
	if opt.Val == rune(nextLetter) {
		if debug {
			fmt.Printf("\t\tRECURSE: %s\n", seen+string(node.Val))
		}
		count += countXMAS(opt, dir, seen+string(node.Val), debug)
	}
	return count
}

func BuildGraph(content string) [][]*Node {
	rows := [][]*Node{}
	row := 0
	rows = append(rows, []*Node{})
	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			row++
			rows = append(rows, []*Node{})
		} else {
			rows[row] = append(rows[row], &Node{Val: rune(content[i]), Next: make([]*Node, 8)})
		}
	}

	//fmt.Printf("len rows: %d\n", len(rows))
	for i := 0; i < len(rows); i++ {
		//fmt.Printf("len row: %d\n", len(rows[i]))
		for j := 0; j < len(rows[i]); j++ {
			//fmt.Printf("i: %d j: %d\n", i, j)
			if i > 0 && j > 0 {
				rows[i][j].Next[0] = rows[i-1][j-1]
			}
			if i > 0 {
				rows[i][j].Next[1] = rows[i-1][j]
			}
			if i > 0 && j < len(rows[i])-1 {
				rows[i][j].Next[2] = rows[i-1][j+1]
			}
			if j > 0 {
				rows[i][j].Next[3] = rows[i][j-1]
			}
			if j < len(rows[i])-1 {
				rows[i][j].Next[4] = rows[i][j+1]
			}
			if i < len(rows)-1 && j > 0 {
				rows[i][j].Next[5] = rows[i+1][j-1]
			}
			if i < len(rows)-1 {
				rows[i][j].Next[6] = rows[i+1][j]
			}
			if i < len(rows)-1 && j < len(rows[i])-1 {
				rows[i][j].Next[7] = rows[i+1][j+1]
			}
		}
	}
	return rows
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := BuildGraph(string(content))

	count := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			fmt.Printf("i: %d j: %d\t\t%s\n", i, j, string(rows[i][j].Val))
			for dir := 0; dir < 8; dir++ {
				debug := false
				if string(rows[i][j].Val) == "X" {
					debug = false
				}
				count += countXMAS(rows[i][j], dir, "", debug)
			}
		}
		fmt.Println("")
	}
	fmt.Println(count)
}
