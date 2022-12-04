package main

import (
	"fmt"
	"strings"

	"github.com/khzaw/algorithms/aoc/utils"
)

const (
	Rock     string = "Rock"
	Paper           = "Paper"
	Scissors        = "Scissors"
)

var RPS = map[string]string{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var Condition = map[string]int{
	"X": -1,
	"Y": 0,
	"Z": 1,
}

func GetShape(move string) int {
	switch move {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

func GetOutcome(m int) int {
	if m > 0 {
		return 6
	} else if m < 0 {
		return 0
	} else {
		return 3
	}
}

func Move(a string, b string) int {
	if a == b {
		return 0
	} else if a == Rock && b == Scissors {
		return 1
	} else if a == Rock && b == Paper {
		return -1
	} else if a == Paper && b == Rock {
		return 1
	} else if a == Paper && b == Scissors {
		return -1
	} else if a == Scissors && b == Paper {
		return 1
	} else {
		return -1
	}
}

// GetPlayerShape returns shape for a player given a win/lose/draw integer
func GetPlayerShape(oppo string, m int) string {
	if oppo == Rock && m > 0 {
		return Paper
	} else if oppo == Rock && m < 0 {
		return Scissors
	} else if oppo == Rock && m == 0 {
		return Rock
	} else if oppo == Paper && m > 0 {
		return Scissors
	} else if oppo == Paper && m < 0 {
		return Rock
	} else if oppo == Paper && m == 0 {
		return Paper
	} else if oppo == Scissors && m > 0 {
		return Rock
	} else if oppo == Scissors && m < 0 {
		return Paper
	} else {
		return Scissors
	}
}

func PartOne(l []string) int {
	score := 0
	for _, game := range l {
		g := strings.Split(game, " ")
		oppo, player := g[0], g[1]
		score += GetShape(RPS[player]) + GetOutcome(Move(RPS[player], RPS[oppo]))
	}
	return score
}

func PartTwo(l []string) int {
	score := 0
	for _, game := range l {
		g := strings.Split(game, " ")
		oppo, player := RPS[g[0]], GetPlayerShape(RPS[g[0]], Condition[g[1]])
		score += GetShape(player) + GetOutcome(Move(player, oppo))
	}
	return score
}

func main() {
	data := utils.ReadInput("./input.txt")
	data = strings.TrimRight(data, "\n")
	strategy := strings.Split(data, "\n")

	fmt.Printf("Part One: %d\n", PartOne(strategy))
	fmt.Printf("Part Two: %d\n", PartTwo(strategy))
}
