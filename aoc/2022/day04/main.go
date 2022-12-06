package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var data string

type Range struct {
	start int
	end   int
}

func sliceAtoi(l []string) []int {
	var ls = make([]int, len(l))
	var err error
	for i := range l {
		if ls[i], err = strconv.Atoi(l[i]); err != nil {
			panic("invalid input")
		}
	}
	return ls
}

func parse(r string) Range {
	n := sliceAtoi(strings.Split(r, "-"))
	return Range{n[0], n[1]}
}

func init() {
	data = strings.TrimRight(data, "\n")
}

func main() {
	d := strings.Split(data, "\n")

	//d := []string{
	//	"2-10,1-9",
	//	"2-3,4-5",
	//	"5-7,7-9",
	//	"2-8,3-7",
	//}

	count, overlap := 0, 0
	for _, row := range d {
		ranges := strings.Split(row, ",")
		a, b := parse(ranges[0]), parse(ranges[1])

		m := make(map[int]bool)
		for i := a.start; i <= a.end; i++ {
			m[i] = true
		}

		var inter []int
		for i := b.start; i <= b.end; i++ {
			if m[i] {
				inter = append(inter, i)
			}
		}

		if len(inter) == 0 {
			continue
		}

		if len(inter) > 0 {
			overlap++
		}

		if (inter[0] == a.start && inter[len(inter)-1] == a.end) ||
			(inter[0] == b.start && inter[len(inter)-1] == b.end) {
			count++
		}
	}
	fmt.Printf("Part One: Count: %d\n", count)

	fmt.Printf("Part Two: Overlap: %d\n", overlap)
}
