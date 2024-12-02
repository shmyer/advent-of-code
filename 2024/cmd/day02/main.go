package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := utils.ReadFile("tests/day02.txt")
	result := solve(input)
	fmt.Println("Result for Day 1:", result)
}

func solve(input string) int {

	lines := strings.Split(input, "\r\n")

	count := 0
	for _, line := range lines {

		if isSafe(line) {
			count++
			fmt.Printf("%s: Safe\n", line)
		} else {
			fmt.Printf("%s: Unsafe\n", line)
		}
	}

	return count
}

func isSafe(line string) bool {

	values := strings.Split(line, " ")

	var first int
	previous := -1
	for _, value := range values {
		current := utils.Atoi(value)

		if previous == -1 {
			first = current
			previous = current
			continue
		}

		// check for monotonic increase
		if previous > first && current < previous || previous < first && current > previous {
			return false
		}

		diff := math.Abs(float64(previous - current))
		previous = current

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
