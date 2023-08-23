package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/pkg/algorithms/vpriorityqueue"
)

func DFS(g graph.Graph, r *graph.Vertex, goal int) *graph.Vertex {
	pq := new(vpriorityqueue.VPriorityQueue)
	n := len(g.Vertices())
	i := 0
	r.Visited = true
	r.Parent = nil
	r.Time = i
	r.Level = 0
	r.Priority = n
	pq.AddVertex(r)

	for len(*pq) != 0 {
		x := pq.PopVertex()
		if x.Value == goal {
			return x
		}
		i += 1
		for _, y := range g.Neighbours(x) {
			if !y.Visited {
				y.Time = i
				y.Level = x.Level + 1
				y.Priority = 1 / y.Level
				y.Visited = true
				y.Parent = x
				pq.AddVertex(y)
			}
		}
	}
	return nil
}
