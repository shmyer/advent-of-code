package utils

import "os"

func ReadFile(name string) string {

	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(data)
}
