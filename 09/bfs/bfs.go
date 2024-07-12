package bfs

import (
	"container/list"
)

func Search(graph [][]int, node int, fn func(index int)) {
	isVisited := make(map[int]struct{})
	isVisited[node] = struct{}{}

	var bfsQueue list.List

	bfsQueue.PushBack(node)

	for e := bfsQueue.Front(); e != nil; e = e.Next() {
		curr := e.Value.(int)
		fn(curr)

		for nodes := 0; nodes < len(graph[curr]); nodes++ {
			if graph[curr][nodes] != 0 {
				if _, ok := isVisited[nodes]; !ok {
					bfsQueue.PushBack(nodes)
					isVisited[nodes] = struct{}{}
				}
			}
		}
	}
}
