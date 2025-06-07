package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x int
	y int
}

var START Point = Point{0, 0}
var HEAD, TAIL Point = START, START
var VISITED = map[Point]bool{
	START: true,
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func init() {
	input = strings.TrimRight(input, "\n")
}

func main() {

	moves := strings.Split(input, "\n")
	for _, m := range moves {
		m := strings.Split(m, " ")
		MoveHead(m[0], toInt(m[1]))

	}
	fmt.Printf("%v\n", HEAD)
}

func MoveHead(m string, len int) {
	switch m {
	case "R":
		HEAD.x += len
	case "L":
		HEAD.x -= len
	case "U":
		HEAD.y += len
	case "D":
		HEAD.y -= len
	}
	MoveTail()
}

func MoveTail() {

}
