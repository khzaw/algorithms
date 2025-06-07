package parser

type Token int

const (
	ILLEGAL Token = iota
	IDENT

	EOF
	WS

	// keywords
	CD
	LS
	DIR

	// Misc characters
	DOTDOT // ..
)
