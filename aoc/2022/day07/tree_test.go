package main

import (
	"fmt"
	"testing"
)

func TestTree_ProcessCD(t *testing.T) {
	tests := []struct {
		p   string
		exp string
	}{
		{
			p:   "/",
			exp: "/",
		},
		{
			p:   "btsgrbd",
			exp: "/btsgrbd",
		},
		{
			p:   "dvght",
			exp: "/btsgrbd/dvght",
		},
		{
			p:   "..",
			exp: "/btsgrbd",
		},
	}

	tree := NewTree()
	for _, tt := range tests {
		tree.ProcessCD(tt.p)
		p := tree.GetPath()
		fmt.Println("path", p)
		if p != tt.exp {
			t.Errorf("path not equal, expected: %s, got: %s\n", tt.exp, p)
		}

	}
}
