package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/pkg/algorithms/vpriorityqueue"
)

// Iterative Deepening depth-first Search (IDS) implementation.
func IDS(g graph.Graph, r *graph.Vertex, goal int) *graph.Vertex {
	for l := 1; l < len(g.Vertices()); l++ {
		pq := new(vpriorityqueue.VPriorityQueue)
		g.InitializeVertices()

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
				if !y.Visited && x.Level+1 < l {
					y.Time = i
					y.Level = x.Level + 1
					y.Priority = n - y.Level // 1 / y.Level
					y.Visited = true
					y.Parent = x
					pq.AddVertex(y)
				}
			}
		}
	}
	return nil
}
