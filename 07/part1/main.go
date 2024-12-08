package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLines(input string) [][][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][][]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ":")
		left := make([]int, 1)
		num, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		left[0] = num

		rightStr := strings.Fields(strings.TrimSpace(parts[1]))
		right := make([]int, len(rightStr))
		for j, numStr := range rightStr {
			right[j], _ = strconv.Atoi(numStr)
		}

		result[i] = [][]int{left, right}
	}
	return result
}

type NodeType int

const (
	NumberNode NodeType = iota
	OperatorNode
)

type Node struct {
	Type   NodeType
	IntVal int
	OpVal  string
	Next   []*Node
}

// Helper functions to create nodes
func NewNumberNode(val int) *Node {
	return &Node{
		Type:   NumberNode,
		IntVal: val,
	}
}

func NewOperatorNode(op string) *Node {
	return &Node{
		Type:  OperatorNode,
		OpVal: op,
	}
}

func intSliceToNodes(numbers []int) []*Node {
	nodes := make([]*Node, len(numbers))
	for i, num := range numbers {
		nodes[i] = NewNumberNode(num)
	}
	return nodes
}

func possibleResult(total int, target int, start *Node) bool {
	if start == nil || len(start.Next) == 0 {
		return total == target
	}
	for _, next := range start.Next {
		if next.Type == OperatorNode {
			switch next.OpVal {
			case "+":
				if possibleResult(total+next.Next[0].IntVal, target, next.Next[0]) {
					return true
				}
			case "*":
				if possibleResult(total*next.Next[0].IntVal, target, next.Next[0]) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	score := 0
	result := parseLines(string(input))
	for _, line := range result {
		nodes := intSliceToNodes(line[1])
		first := nodes[0]
		next := first
		for _, node := range nodes[1:] {
			mul := NewOperatorNode(("*"))
			mul.Next = append(mul.Next, node)
			add := NewOperatorNode(("+"))
			add.Next = append(add.Next, node)
			concat := NewOperatorNode(("."))
			add.Next = append(concat.Next, node)
			next.Next = append(next.Next, mul, add, concat)
			next = node
		}
		if possibleResult(first.IntVal, line[0][0], first) {
			fmt.Println(line[0], "  = ", line[1])
			score += line[0][0]
		} else {
			fmt.Println(line[0], " != ", line[1])
		}
	}
	fmt.Println("Score = ", score)
}
