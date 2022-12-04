package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var data string

func init() {
	data = strings.TrimRight(data, "\n")
}

func main() {
	backpack := strings.Split(data, "\n")

	score := 0
	for _, items := range backpack {
		first, second := items[:len(items)/2], items[len(items)/2:]
		c := commonItem(first, second)
		score += getScore(c)
	}

	fmt.Printf("Score: %v\n", score)

}

func commonItem(first, second string) rune {
	d := map[rune]struct{}{}
	for _, i := range first {
		d[i] = struct{}{}
	}

	for _, i := range second {
		if _, ok := d[i]; ok {
			return i
		}
	}
	return 0
}

func getScore(c rune) int {
	if c >= 'a' {
		return int(c) - 96
	}
	return int(c) - 38
}
