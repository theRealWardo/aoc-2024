package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var firstNumbers []int
	var secondNumbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")

		if len(numbers) == 2 {
			first, err := strconv.Atoi(numbers[0])
			if err != nil {
				return nil, nil, err
			}

			second, err := strconv.Atoi(numbers[1])
			if err != nil {
				return nil, nil, err
			}

			firstNumbers = append(firstNumbers, first)
			secondNumbers = append(secondNumbers, second)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return firstNumbers, secondNumbers, nil
}

func main() {
	first, second, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	if len(first) != len(second) {
		panic("Input files have different lengths")
	}

	sort.Ints(first)
	sort.Ints(second)

	fmt.Printf("part 1: %d\n", part1(first, second))
	fmt.Printf("part 2: %d\n", part2(first, second, 0))
}

func part1(first, second []int) int {
	sum := 0
	for i := 0; i < len(first); i++ {
		if first[i] > second[i] {
			sum += first[i] - second[i]
		} else {
			sum += second[i] - first[i]
		}
	}
	return sum
}

func part2(first, second []int, score int) int {
	if len(first) == 0 || len(second) == 0 {
		return score
	}
	if first[0] < second[0] {
		return part2(first[1:], second, score)
	} else if first[0] > second[0] {
		return part2(first, second[1:], score)
	}
	if first[0] == second[0] {
		i := 0
		for i < len(second) && first[0] == second[i] {
			i++
		}
		return part2(first[1:], second, score+(first[0]*i))
	}
	return score
}
