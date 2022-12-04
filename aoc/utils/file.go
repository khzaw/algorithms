package utils

import (
	"os"
)

// ReadInput reads from AOC input file specified at `path`.
func ReadInput(path string) string {

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}
