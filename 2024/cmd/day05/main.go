package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	input := utils.ReadFile("inputs/day05.txt")
	result := solvePart1(input)
	fmt.Println("Result for Day 5 Part 1:", result)

	result = solvePart2(input)
	fmt.Println("Result for Day 5 Part 2:", result)
}

// part 1
func solvePart1(input string) int {
	rulesAndUpdates := strings.Split(input, "\r\n\r\n")

	allowedAfter, allowedBefore := parseRules(rulesAndUpdates[0])

	return validUpdates(rulesAndUpdates[1], allowedAfter, allowedBefore)
}

func validUpdates(updates string, allowedAfter map[int][]int, allowedBefore map[int][]int) int {
	lines := strings.Split(updates, "\r\n")

	count := 0
	for _, line := range lines {
		pages := atoiSlice(strings.Split(line, ","))

		valid, _, _ := validUpdate(pages, allowedAfter, allowedBefore)
		if valid {
			count += pages[len(pages)/2]
		}
	}
	return count
}

func validUpdate(pages []int, allowedAfter map[int][]int, allowedBefore map[int][]int) (bool, int, int) {
	for i := range pages {
		valid, currIdx, nextIdx := validPage(pages, i, allowedAfter, allowedBefore)
		if !valid {
			return false, currIdx, nextIdx
		}
	}
	return true, -1, -1
}

func validPage(pages []int, currIdx int, allowedAfter map[int][]int, allowedBefore map[int][]int) (bool, int, int) {
	current := pages[currIdx]

	for nextIdx, next := range pages {
		if nextIdx > currIdx && !slices.Contains(allowedAfter[current], next) {
			return false, currIdx, nextIdx
		}
		if nextIdx < currIdx && !slices.Contains(allowedBefore[current], next) {
			return false, currIdx, nextIdx
		}
	}
	return true, -1, -1
}

func parseRules(rawRules string) (map[int][]int, map[int][]int) {
	lines := strings.Split(rawRules, "\r\n")

	allowedAfter := make(map[int][]int)
	allowedBefore := make(map[int][]int)

	for _, line := range lines {
		parts := strings.Split(line, "|")
		left := utils.Atoi(parts[0])
		right := utils.Atoi(parts[1])

		leftAllowedAfter, found := allowedAfter[left]
		if found {
			allowedAfter[left] = append(leftAllowedAfter, right)
		} else {
			allowedAfter[left] = []int{right}
		}

		rightAllowedBefore, found := allowedBefore[right]
		if found {
			allowedBefore[right] = append(rightAllowedBefore, left)
		} else {
			allowedBefore[right] = []int{left}
		}
	}

	return allowedAfter, allowedBefore
}

func atoiSlice(pages []string) []int {
	converted := make([]int, len(pages))
	for i, page := range pages {
		converted[i] = utils.Atoi(page)
	}
	return converted
}

// part 2
func solvePart2(input string) int {
	rulesAndUpdates := strings.Split(input, "\r\n\r\n")

	allowedAfter, allowedBefore := parseRules(rulesAndUpdates[0])

	return validUpdatesPart2(rulesAndUpdates[1], allowedAfter, allowedBefore)
}

func validUpdatesPart2(updates string, allowedAfter map[int][]int, allowedBefore map[int][]int) int {
	lines := strings.Split(updates, "\r\n")

	count := 0
	for _, line := range lines {
		pages := atoiSlice(strings.Split(line, ","))

		valid, currIdx, nextIdx := validUpdate(pages, allowedAfter, allowedBefore)
		if valid {
			continue
		}
		for !valid {
			// switch the incorrect indexes until it is valid
			pages[currIdx], pages[nextIdx] = pages[nextIdx], pages[currIdx]
			valid, currIdx, nextIdx = validUpdate(pages, allowedAfter, allowedBefore)
		}

		if valid {
			count += pages[len(pages)/2]
		} else {
			fmt.Printf("Invalid line: %s\n", line)
		}
	}
	return count
}
