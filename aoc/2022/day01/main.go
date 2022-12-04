package main

import (
	"fmt"
	"strconv"
	"strings"

	"sort"

	"github.com/khzaw/algorithms/aoc/utils"
)

func GetMax(l []int) int {
	max := 0
	for _, c := range l {
		if c > max {
			max = c
		}
	}
	return max
}

func main() {
	data := utils.ReadInput("./input.txt")
	elfs := strings.Split(data, "\n")
	calories := []int{}

	total := 0
	for _, i := range elfs {
		c, err := strconv.Atoi(i)
		if err != nil {
			c = 0
			calories = append(calories, total)
			total = 0
		}
		total += c
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	fmt.Println("Part One: Max Calories")
	fmt.Println(calories[0])

	fmt.Println("Part Two: Sum Top 3 Max Calories")
	fmt.Println(calories[0] + calories[1] + calories[2])
}
