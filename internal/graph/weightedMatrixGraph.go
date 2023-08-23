package graph

import "errors"

type weightedAdjacencyMatrix = [][]int

type WeightedMatrixGraph struct {
	matrix   weightedAdjacencyMatrix
	edges    WeightedEdges
	vertices Vertices
}

func NewWeightedMatrixGraph(matrix weightedAdjacencyMatrix) (*WeightedMatrixGraph,
	error) {
	vertices := map[int]*Vertex{}
	edges := map[Edge]int{}
	for i, v := range matrix {
		vertices[i] = NewVertex(i)
		for j, w := range v {
			if i < j && w != matrix[j][i] {
				return nil, errors.New("assymetric graph error")
			}
			if matrix[i][j] != 0 {
				edge := &Edge{U: i, V: j}
				edges[*edge] = matrix[i][j]
			}
		}
	}
	return &WeightedMatrixGraph{
		matrix:   matrix,
		edges:    edges,
		vertices: vertices,
	}, nil
}

func (g *WeightedMatrixGraph) Neighbours(node *Vertex) []*Vertex {
	neighbours := []*Vertex{}
	id := node.Id
	for i, adjacent := range g.matrix[id] {
		if adjacent != 0 {
			neighbours = append(neighbours, g.vertices[i])
		}
	}
	return neighbours
}

func (g *WeightedMatrixGraph) Vertices() Vertices {
	return g.vertices
}

func (g *WeightedMatrixGraph) Edges() WeightedEdges {
	return g.edges
}

func (g *WeightedMatrixGraph) InitializeVertices() {
	for i, _ := range g.matrix {
		g.vertices[i] = NewVertex(i)
	}
}
