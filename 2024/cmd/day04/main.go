package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"strings"
)

const search = "XMAS"

func main() {
	input := utils.ReadFile("inputs/day04.txt")
	result := solve(input)
	fmt.Println("Result for Day 4:", result)
}

func solve(input string) int {

	lines := strings.Split(input, "\n")

	count := 0
	for y, line := range lines {

		for x, char := range line {
			// skip all non-X characters
			if char != 'X' {
				continue
			}

			count += searchHorizontal(line, x)
			count += searchVertical(lines, x, y)
			count += searchDiagonal(lines, x, y)
		}
	}

	return count
}

func searchHorizontal(line string, x int) int {
	length := len(line)
	west := x >= 3       // west can only be found if there are at least 3 chars west of the current one
	east := length-x > 3 // east can only be found if there are at least 3 chars east of the current one

	for i := 1; i < 4 && (west || east); i++ {
		west = west && line[x-i] == search[i]
		east = east && line[x+i] == search[i]
	}

	return countTrue(west, east)
}

func searchVertical(lines []string, x int, y int) int {
	length := len(lines)
	north := y >= 3       // north can only be found if there are at least 3 lines above the current one
	south := length-y > 3 // north can only be found if there are at least 3 lines below the current one

	for i := 1; i < 4 && (north || south); i++ {
		north = north && lines[y-i][x] == search[i]
		south = south && lines[y+i][x] == search[i]
	}

	return countTrue(north, south)
}

func searchDiagonal(lines []string, x int, y int) int {
	lengthX := len(lines[y])
	lengthY := len(lines)

	northwest := y >= 3 && x >= 3
	northeast := y >= 3 && lengthX-x > 3
	southwest := lengthY-y > 3 && x >= 3
	southeast := lengthY-y > 3 && lengthX-x > 3

	for i := 1; i < 4; i++ {
		northwest = northwest && lines[y-i][x-i] == search[i]
		northeast = northeast && lines[y-i][x+i] == search[i]
		southwest = southwest && lines[y+i][x-i] == search[i]
		southeast = southeast && lines[y+i][x+i] == search[i]
	}

	return countTrue(northeast, northwest, southeast, southwest)
}

func countTrue(flags ...bool) int {
	count := 0
	for _, val := range flags {
		if val {
			count++
		}
	}
	return count
}
