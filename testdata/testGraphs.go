package testdata

import "aiSearchAlgorithms/internal/graph"

type GraphName string
type GraphImplementation string

const (
	Graph07  GraphName = "Graph07"
	Graph07W GraphName = "Graph07W"
)

const (
	matrixGraph         GraphImplementation = "Matrix Graph"
	weightedMatrixGraph GraphImplementation = "Weighted Matrix Graph"
)

func GetTestGraph(name GraphName) graph.Graph {
	var matrix [][]byte
	var weightedMatrix [][]float64
	var implementation GraphImplementation
	var g graph.Graph
	switch name {
	case Graph07:
		matrix = [][]byte{
			{0, 1, 1, 0, 0, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0, 1, 0},
			{0, 0, 0, 0, 1, 0, 1},
			{0, 0, 1, 0, 0, 1, 0},
		}
		implementation = matrixGraph
	case Graph07W:
		weightedMatrix = [][]float64{
			{0, 9, 1, 0, 0, 0, 0},
			{9, 0, 0, 9, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 2},
			{0, 9, 0, 0, 9, 0, 0},
			{0, 0, 0, 9, 0, 9, 0},
			{0, 0, 0, 0, 9, 0, 9},
			{0, 0, 2, 0, 0, 9, 0},
		}
		implementation = weightedMatrixGraph
	}

	switch implementation {
	case matrixGraph:
		g, _ = graph.NewMatrixGraph(matrix)
	case weightedMatrixGraph:
		g, _ = graph.NewWeightedMatrixGraph(weightedMatrix)
	}
	return g
}
