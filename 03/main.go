package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseOps(input string) [][]int {
	r := regexp.MustCompile(`mul\([0-9]+\,[0-9]+\)`)
	results := r.FindAllString(input, -1)
	var out [][]int
	for _, result := range results {
		parts := strings.Split(result, ",")
		num1, _ := strconv.Atoi(parts[0][4:])
		num2, _ := strconv.Atoi(parts[1][:len(parts[1])-1])
		out = append(out, []int{num1, num2})
	}
	return out
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(parseOps(string(content)))

	sum := 0
	for _, op := range parseOps(string(content)) {
		sum += op[0] * op[1]
	}
	fmt.Println(sum)
}
