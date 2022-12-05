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

	// backpack := []string{
	// 	"vJrwpWtwJgWrhcsFMMfFFhFp",
	// 	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	// 	"PmmdzqPrVvPwwTWBwg",
	// 	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	// 	"ttgJtRGJQctTZtZT",
	// 	"CrZsJsPPZsGzwwsLwLmpwMDw",
	// }

	fmt.Printf("Part One: %v\n", PartOne(backpack))

	fmt.Printf("Part Two: %v\n", PartTwo(backpack))

}

func PartOne(backpack []string) int {
	score := 0
	for _, items := range backpack {
		first, second := items[:len(items)/2], items[len(items)/2:]
		// c := commonItem(first, second)
		c := commonChar([]string{first, second})
		score += getScore(c)
	}
	return score
}

func PartTwo(backpack []string) int {
	score := 0
	for i := 0; i < len(backpack); i = i + 3 {
		c := commonChar([]string{backpack[i], backpack[i+1], backpack[i+2]})
		score += getScore(c)
	}
	return score
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

// [1, 1, 1, 1, 1]

// [1, 1,1, 0, 0]

// [1, 0, 0, 1, 1]

// [1, 0, 0, 1, 1]

// A-Za-z
// 65-90,97-122

func commonChar(l []string) rune {
	primary := make([]bool, 58)
	for i := range primary {
		primary[i] = true
	}

	for _, s := range l {
		secondary := make([]bool, 58)

		for _, c := range s {
			if primary[c-'A'] {
				secondary[c-'A'] = true
			}
		}

		for i, _ := range primary {
			primary[i] = secondary[i]
		}
	}

	for i := range primary {
		if primary[i] {
			return rune(i) + 'A'
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
