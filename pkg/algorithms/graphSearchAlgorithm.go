package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
)

type SearchAlgorithm func(g *graph.Graph, v *graph.Vertex, goal int)

func Path(v *graph.Vertex) []*graph.Vertex {
	path := []*graph.Vertex{}
	path = append(path, v)
	parent := v.Parent
	for parent != nil {
		path = append(path, parent)
		parent = parent.Parent
	}
	return path
}

func PathWeight(v *graph.Vertex, edges graph.WeightedEdges) float64 {
	weight := 0.0
	path := Path(v)
	for i := 0; i < len(path)-1; i++ {
		edge := &graph.Edge{U: path[i].Id, V: path[i+1].Id}
		weight += edges[*edge]
	}
	return weight
}
