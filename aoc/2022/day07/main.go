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
	tree, _ := ReadFileTree("./input.txt")

	var MAX_DISK_SPACE int = 70000000
	var UPDATE int = 30000000

	total := 0
	m := tree.GetMap()
	for _, v := range m {
		if v <= 100000 {
			total += v
		}
	}
	fmt.Println("Part One: ", total)

	totalSpace := m["/"]
	// fmt.Println("root space: ", totalSpace)

	remaining := MAX_DISK_SPACE - totalSpace
	// fmt.Println("remaining: ", remaining)

	needForUpdate := UPDATE - remaining
	// fmt.Println("needForUpdate: ", needForUpdate)

	smallest := totalSpace
	for _, v := range m {
		if v >= needForUpdate && v < smallest {
			smallest = v
		}
	}

	fmt.Println("Part Two: ", smallest)
}

func ReadFileTree(p string) (*Tree, error) {
	f, err := os.Open(p)
	defer f.Close()

	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}

	rd := bufio.NewReader(f)
	tree := NewTree()

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read file line error: %w", err)
		}
		ParseTree(strings.TrimRight(line, "\n"), tree)
	}
	return tree, nil
}

func ParseTree(line string, tree *Tree) {
	if line[0] == '$' && line[2:4] == "cd" {
		tree.ProcessCD(line[4:])
	} else if line[0] == '$' && line[2:4] == "ls" {
		// continue
	} else if line[0:3] == "dir" {
		tree.ProcessDir(line[4:])
	} else {
		l := strings.Split(line, " ")
		tree.ProcessFile(toInt(l[0]))
	}

}

func toInt(s string) int {
	d, _ := strconv.Atoi(s)
	return d
}
