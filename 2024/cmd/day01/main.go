package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"math"
	"regexp"
	"sort"
)

func main() {
	input := utils.ReadFile("inputs/day01.txt")
	result := solveDay01(input)
	fmt.Println("Result for Day 1:", result)
}

func solveDay01(input string) int {
	list1, list2 := parseAndSortLists(input)

	distances := 0
	for i := 0; i < len(list1); i++ {
		distances += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return distances
}

func parseAndSortLists(input string) ([]int, []int) {
	regex, _ := regexp.Compile("\\s+")
	result := regex.Split(input, -1)

	// minus one because of trailing newline
	l := (len(result) - 1) / 2

	list1 := make([]int, l)
	list2 := make([]int, l)

	for i := 0; i < l; i++ {
		ri := i * 2
		list1[i] = utils.Atoi(result[ri])
		list2[i] = utils.Atoi(result[ri+1])
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return list1, list2
}
