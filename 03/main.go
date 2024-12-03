package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseOps(input string) [][]int {
	var out [][]int
	enabled := true
	prog := []string{}
	for i := 0; i < len(input); {
		cur := input[i:]
		fmt.Printf("\ninput: %v\n", cur)
		fmt.Printf("i: %v\n", i)
		nextDisable := strings.Index(cur, "don't()")
		nextEnable := strings.Index(cur, "do()")
		fmt.Printf("nextDisable: %v\n", nextDisable)
		fmt.Printf("nextEnable: %v\n", nextEnable)
		if enabled {
			if nextDisable == -1 || i+nextDisable > len(input) {
				prog = append(prog, cur)
				break
			} else {
				fmt.Printf("appending: %v\n", cur[:nextDisable])
				prog = append(prog, cur[:nextDisable])
				i = i + nextDisable + 7
				enabled = false
			}
		} else {
			if nextEnable == -1 {
				break
			} else {
				i = i + nextEnable + 4
				enabled = true
			}
		}
	}
	fmt.Printf("\n\n%+v\n\n", prog)
	for i := 0; i < len(prog); i++ {
		r := regexp.MustCompile(`mul\([0-9]+\,[0-9]+\)`)
		results := r.FindAllString(prog[i], -1)
		for _, result := range results {
			parts := strings.Split(result, ",")
			num1, _ := strconv.Atoi(parts[0][4:])
			num2, _ := strconv.Atoi(parts[1][:len(parts[1])-1])
			out = append(out, []int{num1, num2})
		}
	}
	return out
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	result := parseOps(string(content))

	fmt.Println(result)

	sum := 0
	for _, op := range result {
		sum += op[0] * op[1]
	}
	fmt.Println(sum)
}
