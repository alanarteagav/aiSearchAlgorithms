package algorithms

import (
	"aiSearchAlgorithms/internal/graph"
)

type SearchAlgorithm func(g *graph.Graph, v *graph.Vertex, goal int)
