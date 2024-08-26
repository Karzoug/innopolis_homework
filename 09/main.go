package main

import (
	"09/bfs"
)

func main() {
	graph := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 0, 0, 0},
	}
	bfs.Search(graph, 1, func(index int) {
		println(index)
	})
}
