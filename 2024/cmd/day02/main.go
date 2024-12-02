package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile("inputs/day02.txt")
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
	report := make([]int, len(values))

	for i, value := range values {
		report[i] = utils.Atoi(value)
	}

	if isSafeReport(report) {
		return true
	}

	for i := 0; i < len(values); i++ {
		reportCopy := make([]int, len(report))
		copy(reportCopy, report)
		if isSafeReport(slice(reportCopy, i)) {
			return true
		}
	}

	return false
}

func isSafeReport(report []int) bool {

	previousDiff := report[1] - report[0]
	if !isSafeDiff(previousDiff) {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		nextDiff := report[i+1] - report[i]

		if !isSafeDiff(nextDiff) {
			return false
		}

		if previousDiff < 0 && nextDiff > 0 || previousDiff > 0 && nextDiff < 0 {
			return false
		}
	}

	return true
}

func isSafeDiff(diff int) bool {
	return diff >= -3 && diff <= -1 || diff >= 1 && diff <= 3
}

func slice(report []int, remove int) []int {
	return append(report[:remove], report[remove+1:]...)
}
