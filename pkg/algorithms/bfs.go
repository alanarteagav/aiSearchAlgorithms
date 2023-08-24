package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/pkg/algorithms/vpriorityqueue"
)

// Breadth First Search (BFS) implementation.
func BFS(g graph.Graph, r *graph.Vertex, goal int) *graph.Vertex {
	pq := new(vpriorityqueue.VPriorityQueue)
	g.InitializeVertices()

	i := 1
	r.Visited = true
	r.Parent = nil
	r.Time = i
	r.Level = 0
	r.Priority = -1
	pq.AddVertex(r)

	for len(*pq) != 0 {
		x := pq.PopVertex()
		if x.Value == goal {
			return x
		}
		for _, y := range g.Neighbours(x) {
			if !y.Visited {
				i += 1
				y.Visited = true
				y.Parent = x
				y.Time = i
				y.Level = x.Level + 1
				y.Priority = i
				pq.AddVertex(y)
			}
		}
	}
	return nil
}
