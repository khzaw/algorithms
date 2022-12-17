package main

import (
	"fmt"
	"strings"
)

type Tree struct {
	path string // determines the current path
	m    map[string]int
}

func NewTree() *Tree {
	return &Tree{m: map[string]int{}}
}

func (t *Tree) GetPath() string {
	return t.path
}

func (t *Tree) getNewPath(child string) string {
	if len(t.path) >= 1 && t.path[len(t.path)-1] != '/' {
		return t.path + "/" + child
	}
	return t.path + child
}

func (t *Tree) getParentPath(curr string) string {
	g := strings.Split(curr, "/")
	j := strings.Join(g[:len(g)-1], "/")
	if j == "" {
		return "/"
	}
	return j
}

func (t *Tree) ProcessCD(p string) {
	p = strings.Trim(p, " ")
	if p == ".." {
		t.path = t.getParentPath(t.path)
	} else {
		t.path = t.getNewPath(p)
	}

	if _, exists := t.m[t.path]; !exists {
		t.m[t.path] = 0
	}
}

func (t *Tree) GetMap() map[string]int {
	return t.m
}

// ProcessDir add new path  key to the map
func (t *Tree) ProcessDir(d string) {
	newPath := t.getNewPath(d)
	if _, exists := t.m[newPath]; !exists {
		t.m[newPath] = 0
	}
}

func (t *Tree) ProcessFile(size int) {
	parents := []string{}

	curr := t.path
	for curr != "/" {
		p := t.getParentPath(curr)
		parents = append(parents, p)
		curr = p
	}

	for _, p := range parents {
		t.m[p] += size
	}

	t.m[t.path] += size
}

func (t *Tree) ToString() string {
	return fmt.Sprintf("%#v", t.m)
}
