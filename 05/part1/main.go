package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInputToMap(input string) map[int][]int {
	result := make(map[int][]int)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}

		x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue
		}

		y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}

		result[y] = append(result[y], x)
	}

	return result
}

func parseLineToInts(line string) []int {
	numbers := strings.Split(strings.TrimSpace(line), ",")
	result := make([]int, 0, len(numbers))

	for _, num := range numbers {
		n, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			continue
		}
		result = append(result, n)
	}

	return result
}

func depsMet(depsMap map[int][]int, in []int) bool {
	allNumbers := make(map[int]bool)
	for _, n := range in {
		allNumbers[n] = true
	}
	seen := make(map[int]bool)
	for j := 0; j < len(in); j++ {
		dep := in[j]
		if needs, ok := depsMap[dep]; ok {
			for _, need := range needs {
				if !seen[need] {
					if !allNumbers[need] {
						continue
					}
					return false
				}
			}
		}
		seen[dep] = true
	}
	return true
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(content), "\n\n")

	depsMap := parseInputToMap(parts[0])
	fmt.Printf("%+v\n", depsMap)

	lines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	okLists := [][]int{}
	for i := 0; i < len(lines); i++ {
		deps := parseLineToInts(lines[i])
		fmt.Printf("checking: %+v\n", deps)
		if depsMet(depsMap, deps) {
			okLists = append(okLists, deps)
			fmt.Printf("\tok.\n")
		}
	}

	sum := 0
	for _, ok := range okLists {
		fmt.Printf("%d + ", ok[len(ok)/2])
		sum += ok[len(ok)/2]
	}
	fmt.Printf("= %d\n", sum)
}
