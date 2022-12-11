package main

type Block[T any] struct {
	Items []T
}

func (b *Block[T]) Pop() T {
	top := b.Items[len(b.Items)-1]
	b.Items = b.Items[:len(b.Items)-1]
	return top
}

func (b *Block[T]) Add(a T) {
	b.Items = append(b.Items, a)
}
