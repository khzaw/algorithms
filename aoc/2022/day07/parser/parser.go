package parser

import (
	"io"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read string literal
		n   int    // buffer size (max = 1)
	}
}

type Command[C any] struct {
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}
