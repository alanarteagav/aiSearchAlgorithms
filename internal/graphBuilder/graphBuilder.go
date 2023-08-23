package graphbuilder

import "aiSearchAlgorithms/internal/graph"

type GraphBuilder struct {
	graphOrder int
	edges      graph.WeightedEdges
}

func (b *GraphBuilder) WithGraphOrder(n int) {
	b.graphOrder = n
}

func (b *GraphBuilder) WithEdges(w graph.WeightedEdges) {
	b.edges = w
}

func (b *GraphBuilder) Build(implementation graph.GraphImplementation) graph.Graph {
	switch implementation {
	case graph.WeightedMatrix:
		matrix := make([][]int, b.graphOrder)
		for i := range matrix {
			matrix[i] = make([]int, b.graphOrder)
		}
		for e, w := range b.edges {
			matrix[e.U][e.V] = w
			matrix[e.V][e.U] = w
		}
		g, _ := graph.NewWeightedMatrixGraph(matrix)
		return g
	}
	return nil
}
