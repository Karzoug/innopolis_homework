package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	graph := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 0, 0, 0},
	}

	res := make([]int, 0, 4)
	Search(graph, 1, func(index int) {
		res = append(res, index)
	})
	assert.Equal(t, []int{1, 3, 4, 2}, res)

	res = make([]int, 0)
	Search(graph, 0, func(index int) {
		res = append(res, index)
	})
	assert.Equal(t, []int{0}, res)
}
