package main

import (
	"fmt"
	"os"
)

type Node struct {
	Val  rune
	Next []*Node
}

const (
	XMASA string = "M.S..M.S"
	XMASB string = "S.M..S.M"
	XMASC string = "M.M..S.S"
	XMASD string = "S.S..M.M"
)

func countXMAS(node *Node, debug bool) int {
	if debug {
		fmt.Printf("\tval: %v\n\t\"", string(node.Val))
	}
	matchA := true
	matchB := true
	matchC := true
	matchD := true
	for i := 0; i < len(node.Next); i++ {
		if XMASA[i] == '.' {
			continue
		}
		if node.Next[i] == nil {
			if debug {
				fmt.Printf(" ")
			}
			return 0
		}
		if debug {
			fmt.Printf("%s", string(node.Next[i].Val))
		}
		if rune(XMASA[i]) != node.Next[i].Val {
			if debug && matchA {
				fmt.Printf("*")
			}
			matchA = false
		}
		if rune(XMASB[i]) != node.Next[i].Val {
			if debug && matchB {
				fmt.Printf("~")
			}
			matchB = false
		}
		if rune(XMASC[i]) != node.Next[i].Val {
			if debug && matchC {
				fmt.Printf("@")
			}
			matchC = false
		}
		if rune(XMASD[i]) != node.Next[i].Val {
			if debug && matchD {
				fmt.Printf("$")
			}
			matchD = false
		}
	}
	if debug {
		fmt.Printf("\"")
		if matchA || matchB || matchC || matchD {
			fmt.Printf(" MATCH")
		}
		fmt.Printf("\n\n")
	}
	if matchA || matchB || matchC || matchD {
		return 1
	}
	return 0
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
			debug := false
			if string(rows[i][j].Val) == "A" {
				debug = true
				count += countXMAS(rows[i][j], debug)
			}
		}
		fmt.Println("")
	}
	fmt.Println(count)
}
