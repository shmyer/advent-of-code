package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"regexp"
	"sort"
)

func main() {
	input := utils.ReadFile("inputs/day01.txt")
	result := solveDay01(input)
	fmt.Println("Result for Day 1:", result)
}

func solveDay01(input string) int {
	left, right := parseAndSortLists(input)

	similarityScore := 0
	// as the lists are sorted, keep a pointer on the right list
	ri := 0
	for i := 0; i < len(left); i++ {
		// left value
		lv := left[i]

		// skip elements in right if smaller
		for ri+1 < len(right) && right[ri] < lv {
			ri++
		}

		// count appearance in right
		count := 0
		for lv == right[ri] {
			count++
			ri++
		}

		similarityScore += lv * count

		// look ahead in left list if the value appears again and revert ri if it is the case
		if i+1 < len(left) && left[i+1] == lv {
			ri -= count
		}
	}

	return similarityScore
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
