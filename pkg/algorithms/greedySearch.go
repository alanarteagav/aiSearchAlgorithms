package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/pkg/algorithms/vpriorityqueue"
)

// Greedy Best-first Search (GBS) implementation.
func GBS(g graph.Graph, r *graph.Vertex, goal int) *graph.Vertex {
	pq := new(vpriorityqueue.VPriorityQueue)
	g.InitializeVertices()

	r.Visited = true
	r.Priority = g.Heuristic(r)
	pq.AddVertex(r)

	for len(*pq) != 0 {
		x := pq.PopVertex()
		if x.Value == goal {
			return x
		}
		for _, y := range g.Neighbours(x) {
			if !y.Visited {
				y.Priority = g.Heuristic(y)
				y.Visited = true
				y.Parent = x
				pq.AddVertex(y)
			}
		}
	}
	return nil
}
