package main

import (
	"fmt"
)

type Node struct {
	path     string // determines the current path
	size     int
	children []Node
}

func NewNode(path string, size int, children []Node) *Node {
	return &Node{path, size, children}
}

func (n *Node) AddChild(c Node) {
	n.children = append(n.children, c)
	n.size += c.size
}

func (n *Node) ToString() string {
	return fmt.Sprintf("%s, (%d)\n", n.path, n.size)
}
