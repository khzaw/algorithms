package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(x int, y int) bool {
		return records[x].ID < records[y].ID
	})

	t := make([]*Node, len(records))

	for i, r := range records {
		if r.ID != i || r.Parent > i || i > 0 && r.Parent == i {
			return nil, errors.New("invalid tree")
		}

		t[i] = &Node{i, nil}

		if i > 0 {
			t[r.Parent].Children = append(t[r.Parent].Children, t[i])
		}

	}
	return t[0], nil
}
