package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/pkg/algorithms/vpriorityqueue"
	"fmt"
)

func Dfs(g graph.Graph, v *graph.Vertex, goal int) {
	pq := new(vpriorityqueue.VPriorityQueue)
	for _, v := range g.Vertices() {
		v.Priority = 10 - v.Id
		pq.AddVertex(v)
	}

	for len(*pq) != 0 {
		u := pq.PopVertex()
		fmt.Println(u)
	}
}
