package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// forest := parseInput(strings.TrimRight(input, "\n"))
	forest := [][]int{
		[]int{3, 0, 3, 7, 3},
		[]int{2, 5, 5, 1, 2},
		[]int{6, 5, 3, 3, 2},
		[]int{3, 3, 5, 4, 9},
		[]int{3, 5, 3, 9, 0},
	}

	vt, score := IterateForest(forest)

	fmt.Println("Part One: ", vt)
	fmt.Println("Part Two: ", score)

	fmt.Println(GetScore(forest, 1, 2))

}

func IterateForest(forest [][]int) (int, int) {
	inner := 0

	scenicScore := 0

	for i := 1; i < len(forest)-1; i += 1 {
		for j := 1; j < len(forest[i])-1; j += 1 {

			score := GetScore(forest, i, j)
			if score > scenicScore {
				scenicScore = score
			}

			if IsVisible(forest, i, j) {
				inner += 1
			}
		}
	}
	outer := 2*len(forest[0]) + 2*len(forest) - 4
	return inner + outer, scenicScore
}

func GetScore(forest [][]int, row, col int) int {
	var top, bottom, left, right int

	// top
	for i := row - 1; i >= 0; i-- {
		if forest[i][col] <= forest[row][col] {
			top += 1
		}
	}

	// bottom
	for i := row + 1; i < len(forest); i++ {
		if forest[i][col] <= forest[row][col] {
			bottom += 1
		}
	}

	// left
	for j := col - 1; j >= 0 && forest[row][j] < forest[row][col]; j-- {
		if forest[row][j] <= forest[row][col] {
			left++
		}
	}

	// right
	for j := col + 1; j < len(forest[0]); j++ {
		if forest[row][j] <= forest[row][col] {
			right++
		}
	}

	fmt.Println("top", top, "left", left, "right", right, "down", bottom)

	return top * bottom * left * right
}

func IsVisible(forest [][]int, row, col int) bool {
	var top, bottom, left, right bool = true, true, true, true

	// top
	for i := row - 1; i >= 0; i-- {
		if forest[i][col] >= forest[row][col] {
			top = false
			break
		}
	}

	// bottom
	for i := row + 1; i < len(forest); i++ {
		if forest[i][col] >= forest[row][col] {
			bottom = false
			break
		}
	}

	// left
	for j := col - 1; j >= 0; j-- {
		if forest[row][j] >= forest[row][col] {
			left = false
			break
		}
	}

	// right
	for j := col + 1; j < len(forest[0]); j++ {
		if forest[row][j] >= forest[row][col] {
			right = false
			break
		}
	}

	return top || bottom || left || right
}

func parseInput(input string) (forest [][]int) {
	lines := strings.Split(input, "\n")
	forest = make([][]int, len(lines))
	for i, line := range lines {
		forest[i] = make([]int, len(line))

		for j, char := range line {
			forest[i][j] = int(char - '0')
		}
	}
	return forest
}
