package testdata

import "aiSearchAlgorithms/internal/graph"

type GraphName string
type GraphImplementation string

const (
	Graph07       GraphName = "Graph07"
	Graph07W      GraphName = "Graph07W"
	Graph07WDFS   GraphName = "Graph07WDFS"
	GraphRomania0 GraphName = "GraphRomania0"
)

const (
	matrixGraph         GraphImplementation = "Matrix Graph"
	weightedMatrixGraph GraphImplementation = "Weighted Matrix Graph"
)

func GetTestGraph(name GraphName) graph.Graph {
	var matrix [][]byte
	var weightedMatrix [][]int
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
		weightedMatrix = [][]int{
			{0, 9, 1, 0, 0, 0, 0},
			{9, 0, 0, 9, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 2},
			{0, 9, 0, 0, 9, 0, 0},
			{0, 0, 0, 9, 0, 9, 0},
			{0, 0, 0, 0, 9, 0, 9},
			{0, 0, 2, 0, 0, 9, 0},
		}
		implementation = weightedMatrixGraph
	case Graph07WDFS:
		weightedMatrix = [][]int{
			{0, 9, 1, 0, 0, 0, 0},
			{9, 0, 0, 9, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{0, 9, 0, 0, 9, 0, 0},
			{0, 0, 0, 9, 0, 9, 0},
			{0, 0, 0, 0, 9, 0, 9},
			{0, 0, 0, 0, 0, 9, 0},
		}
		implementation = weightedMatrixGraph
	case GraphRomania0:
		weightedMatrix = [][]int{
			{0, 80, 99, 0, 0},
			{80, 0, 0, 97, 0},
			{99, 0, 0, 0, 211},
			{0, 97, 0, 0, 101},
			{0, 0, 211, 101, 0},
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
