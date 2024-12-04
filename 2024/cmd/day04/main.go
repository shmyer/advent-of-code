package main

import (
	"advent-of-code/pkg/utils"
	"fmt"
	"strings"
)

const search = "XMAS"

func main() {
	input := utils.ReadFile("inputs/day04.txt")
	result := solvePart1(input)
	fmt.Println("Result for Day 4 Part 1:", result)

	result = solvePart2(input)
	fmt.Println("Result for Day 4 Part 2:", result)
}

// part 1
func solvePart1(input string) int {

	lines := strings.Split(input, "\r\n")

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

func solvePart2(input string) int {

	lines := strings.Split(input, "\r\n")

	count := 0
	for y, line := range lines {

		for x, char := range line {
			// skip all non-A characters
			if char != 'A' {
				continue
			}

			if isXMas(lines, x, y) {
				count++
			}
		}
	}

	return count
}

func isXMas(lines []string, x int, y int) bool {
	lengthX := len(lines[y])
	lengthY := len(lines)

	// check boundaries first as they can't be true.
	if x == 0 || x == lengthX-1 || y == 0 || y == lengthY-1 {
		return false
	}

	return isSAndM(lines, x-1, y-1, x+1, y+1) && isSAndM(lines, x-1, y+1, x+1, y-1)
}

func isSAndM(lines []string, x1 int, y1 int, x2 int, y2 int) bool {
	return lines[y1][x1] == 'S' && lines[y2][x2] == 'M' || lines[y1][x1] == 'M' && lines[y2][x2] == 'S'
}
