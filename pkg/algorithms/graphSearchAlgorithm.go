// Package algorithms defines several AI search algorithms over graphs.
package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
)

// SearchAlgorithms defines an AI graph search algorithm as a function that
// receives a graph, a distinguished vertex, and a goal. Returns, if found,
// the vertex which value equals the goal.
type SearchAlgorithm func(g graph.Graph, v *graph.Vertex, goal int) *graph.Vertex

// Path receives a vertex and returs the list of its ancestors defined by a
// previously executed search algorithm.
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

// PathWeight returns the weight of the path returned by the Path method.
func PathWeight(v *graph.Vertex, edges graph.WeightedEdges) int {
	weight := 0
	path := Path(v)
	for i := 0; i < len(path)-1; i++ {
		edge := &graph.Edge{U: path[i].Id, V: path[i+1].Id}
		weight += edges[*edge]
	}
	return weight
}
