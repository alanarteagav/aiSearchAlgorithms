package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/pkg/algorithms/vpriorityqueue"
)

// Uniform Cost Search implementation
func UCS(g graph.Graph, r *graph.Vertex, goal int) *graph.Vertex {
	pq := new(vpriorityqueue.VPriorityQueue)
	g.InitializeVertices()

	r.Visited = true
	r.Priority = 0
	pq.AddVertex(r)

	for len(*pq) != 0 {
		x := pq.PopVertex()
		if x.Value == goal {
			return x
		}
		for _, y := range g.Neighbours(x) {
			edge := &graph.Edge{U: x.Id, V: y.Id}
			if !y.Visited {
				y.Priority = g.Edges()[*edge] + x.Priority
				y.Visited = true
				y.Parent = x
				pq.AddVertex(y)
			} else if y.Priority > g.Edges()[*edge]+x.Priority {
				y.Priority = g.Edges()[*edge] + x.Priority
				y.Parent = x
			}
		}
	}
	return nil
}
