package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	_ = ReadFileTree("./input.txt")
}

func ReadFileTree(p string) error {
	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("read file line error: %w", err)
		}
		ParseTree(line, NewNode("/", 0, []Node{}))
	}
	return nil
}

func ParseTree(line string, node *Node) {
	if line[0] == '$' {
		ParseCommand(line[2:], node)
	}
}

func ParseCommand(command string, node *Node) {
	c := strings.Split(command, " ")
}
