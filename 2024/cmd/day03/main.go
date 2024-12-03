package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := utils.ReadFile("inputs/day03.txt")
	result := solve(input)
	fmt.Println("Result for Day 1:", result)
}

func solve(input string) int {
	regex, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")

	found := regex.FindAllString(input, -1)

	var sum int
	for _, value := range found {
		numsRaw := value[4 : len(value)-1]

		nums := strings.Split(numsRaw, ",")

		a := utils.Atoi(nums[0])
		b := utils.Atoi(nums[1])

		sum += a * b
	}

	return sum
}
