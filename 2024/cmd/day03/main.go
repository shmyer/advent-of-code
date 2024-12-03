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
	fmt.Println("Result for Day 3:", result)
}

func solve(input string) int {
	regex, _ := regexp.Compile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")

	//At the beginning of the program, mul instructions are enabled.
	enabled := true

	found := regex.FindAllString(input, -1)

	var sum int
	for _, value := range found {
		if strings.HasPrefix(value, "mul") {
			// skip mul if not enabled
			if !enabled {
				continue
			}

			numsRaw := value[4 : len(value)-1]

			nums := strings.Split(numsRaw, ",")

			a := utils.Atoi(nums[0])
			b := utils.Atoi(nums[1])

			sum += a * b
		} else if value == "do()" {
			enabled = true
		} else if value == "don't()" {
			enabled = false
		} else {
			panic("WAT?")
		}
	}

	return sum
}
