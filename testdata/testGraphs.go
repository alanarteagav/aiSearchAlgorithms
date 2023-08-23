package testdata

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/internal/graphbuilder"
)

type GraphName string

const (
	Graph07       GraphName = "Graph07"
	Graph07W      GraphName = "Graph07W"
	Graph07WDFS   GraphName = "Graph07WDFS"
	GraphRomania0 GraphName = "GraphRomania0"
	GraphRomania  GraphName = "GraphRomania"
)

func GetTestGraph(name GraphName) graph.Graph {
	var matrix [][]byte
	var weightedMatrix [][]int
	var implementation graph.GraphImplementation
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
		implementation = graph.Matrix
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
		implementation = graph.WeightedMatrix
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
		implementation = graph.WeightedMatrix
	case GraphRomania0:
		weightedMatrix = [][]int{
			{0, 80, 99, 0, 0},
			{80, 0, 0, 97, 0},
			{99, 0, 0, 0, 211},
			{0, 97, 0, 0, 101},
			{0, 0, 211, 101, 0},
		}
		implementation = graph.WeightedMatrix
	case GraphRomania:
		b := new(graphbuilder.GraphBuilder)
		b.WithGraphOrder(20)
		b.WithEdges(map[graph.Edge]int{
			{U: 0, V: 1}:   118,
			{U: 0, V: 2}:   75,
			{U: 2, V: 3}:   71,
			{U: 0, V: 5}:   140,
			{U: 1, V: 4}:   111,
			{U: 3, V: 5}:   151,
			{U: 4, V: 6}:   70,
			{U: 5, V: 7}:   80,
			{U: 5, V: 8}:   99,
			{U: 6, V: 9}:   75,
			{U: 9, V: 10}:  120,
			{U: 7, V: 10}:  146,
			{U: 7, V: 11}:  97,
			{U: 8, V: 12}:  211,
			{U: 11, V: 12}: 101,
			{U: 12, V: 13}: 90,
			{U: 12, V: 14}: 85,
			{U: 14, V: 15}: 98,
			{U: 15, V: 17}: 86,
			{U: 14, V: 16}: 142,
			{U: 16, V: 18}: 92,
			{U: 18, V: 19}: 87,
		})
		g := b.Build(graph.WeightedMatrix)
		return g
	}

	switch implementation {
	case graph.Matrix:
		g, _ = graph.NewMatrixGraph(matrix)
	case graph.WeightedMatrix:
		g, _ = graph.NewWeightedMatrixGraph(weightedMatrix)
	}
	return g
}
