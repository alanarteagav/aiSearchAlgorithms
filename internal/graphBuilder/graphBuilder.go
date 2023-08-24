// graphbuilder provides a struct and methods for intuitive graph generation.
package graphbuilder

import "aiSearchAlgorithms/internal/graph"

// GraphBuilder is a graph generator struct. Stores the desired graph order and
// its desired edges.
type GraphBuilder struct {
	graphOrder int
	edges      graph.WeightedEdges
}

// WithGraphOrder defines the desired graph order.
func (b *GraphBuilder) WithGraphOrder(n int) {
	b.graphOrder = n
}

// WithEdges defines the desired graph edges through an specified weighed edges map.
func (b *GraphBuilder) WithEdges(w graph.WeightedEdges) {
	b.edges = w
}

// Build returns a graph built with the desired implementation.
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
