package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	stacks, _ := ReadFile("./input.txt")
	ans := ""
	for _, s := range *stacks {
		if len(s.Items) > 0 {
			ans += s.Items[len(s.Items)-1]
		}
	}
	fmt.Println("Part Two: ", ans)
}

func ReadFile(p string) (*[]Block[string], error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}
	defer f.Close()

	rows, cols, stacks := 0, 0, []Block[string]{}

	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read file line error: %w", err)
		}

		if rows == 0 {
			cols = (len(line) + 1) / 4 // line length is the same for all rows
			stacks = make([]Block[string], cols)
		}

		if string(strings.TrimLeft(line, " ")[0]) == "1" || string(line[0]) == "\n" {
			continue
		} else if string(line[0]) == "m" {
			ins := ParseInstruction(line)
			// ProcessPartOne(&stacks, ins)
			ProcessPartTwo(&stacks, ins)
		} else {
			ParseBlockLine(&stacks, line, rows, cols)
		}
		rows++

	}
	// printStacks(&stacks)
	return &stacks, nil
}

// ParseBlockLine will parse a string line and place the character in the correct column
func ParseBlockLine(stacks *[]Block[string], line string, row int, cols int) {
	// 0 -> 1
	// 1 -> 5
	// 2 -> 9
	// (x - 1)/4   -> 4n + 1
	// 11, 3 -> 1 5 9

	for i := 1; i <= len(line); i = i + 4 {
		c := (i - 1) / 4            // get column from index
		if string(line[i]) != " " { // if found a character
			(*stacks)[c].Items = append([]string{string(line[i])}, (*stacks)[c].Items...)
		}
	}
}

type Instruction struct {
	Blocks int
	From   int
	To     int
}

func ParseInstruction(ins string) *Instruction {
	s := strings.Split(strings.TrimRight(ins, "\n"), " ")
	i := &Instruction{toInt(s[1]), toInt(s[3]) - 1, toInt(s[5]) - 1}
	return i
}

func ProcessPartOne(stacks *[]Block[string], in *Instruction) {
	for i := 0; i < in.Blocks; i++ {
		(*stacks)[in.To].Add((*stacks)[in.From].Pop())
	}
}

func ProcessPartTwo(stacks *[]Block[string], in *Instruction) {
	top := (*stacks)[in.From].Items[len((*stacks)[in.From].Items)-in.Blocks:]
	(*stacks)[in.From].Items = (*stacks)[in.From].Items[:len((*stacks)[in.From].Items)-in.Blocks]
	(*stacks)[in.To].Items = append((*stacks)[in.To].Items, top...)
}

func printStacks(stacks *[]Block[string]) {
	for s := range *stacks {
		fmt.Printf("%#v\n", (*stacks)[s])
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
