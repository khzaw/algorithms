package main

import (
	"fmt"

	"github.com/khzaw/algorithms/aoc/utils"
	"strings"
)

func main() {
	data := strings.TrimRight(utils.ReadInput("./input.txt"), "\n")
	fmt.Println("Part One: ", GetMarker(data, 4))
	fmt.Println("Part Two: ", GetMarker(data, 14))
}

func GetMarker(s string, l int) int {
	if len(s) < l {
		return 0
	}

	for i := l; i < len(s); i += 1 {
		if IsMarker(s[i-l : i]) {
			return i
		}
	}

	return 0
}

func IsMarker(s string) bool {
	m := make(map[rune]int, len(s))
	for _, c := range s {
		if _, exists := m[c]; exists {
			return false
		} else {
			m[c] = 1
		}
	}
	return true
}
