package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		var row []int
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

// 0 = unknown
// -1 = decreasing
// 1 = increasing
func safeRow(row []int, buffer, state int) bool {
	if len(row) < 2 {
		return true
	}
	val := row[0] - row[1]
	if state == 0 {
		if val < 0 {
			// increasing
			state = 1
		} else if val > 0 {
			// decreasing
			state = -1
		} else {
			// UNSAFE
			return false
		}
	}
	// increasing val must be negative
	if state == 1 && val < 0 && val > -4 {
		return safeRow(row[1:], buffer, state)
		// decreasing val must be positive
	} else if state == -1 && val > 0 && val < 4 {
		return safeRow(row[1:], buffer, state)
	}
	return false
}

func part1(rows [][]int) int {
	sum := 0
	for _, row := range rows {
		if safeRow(row, 0, 0) {
			sum++
		}
	}
	return sum
}

func main() {
	rows, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", part1(rows))
}
