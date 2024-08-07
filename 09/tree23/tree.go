package tree23

import (
	"golang.org/x/exp/constraints"
)

type Tree[T constraints.Ordered] struct {
	root *node[T]
}

func New[T constraints.Ordered]() *Tree[T] {
	return &Tree[T]{root: nil}
}

func (t Tree[T]) Search(x T) bool {
	return t.root.Search(x) != nil
}

func (t *Tree[T]) Insert(x ...T) {
	for _, v := range x {
		t.root = t.root.Insert(v)
	}
}

func (t *Tree[T]) Remove(x T) {
	t.root = t.root.Remove(x)
}
