package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// readFile("./input.txt")
	test := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	}

	for l := range test {
		parseBlockLine(test[l])
	}
}

func readFile(p string) {
	f, err := os.Open(p)
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("read file line error: %v", err)
			return
		}

		parseBlockLine(line)
	}
}

func parseBlockLine(line string) {
	fmt.Println(line, len(line))
}
