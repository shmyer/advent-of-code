package utils

import "strconv"

func Atoi(input string) int {
	atoi, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return atoi
}
