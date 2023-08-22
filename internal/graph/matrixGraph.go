package graph

import "errors"

type adjacencyMatrix = [][]byte

type MatrixGraph struct {
	matrix   adjacencyMatrix
	edges    WeightedEdges
	vertices Vertices
}

func NewMatrixGraph(matrix adjacencyMatrix) (*MatrixGraph, error) {
	vertices := map[int]*Vertex{}
	for i, v := range matrix {
		vertices[i] = InitializeVertex(i)
		for j, w := range v {
			if i < j && w != matrix[j][i] {
				return nil, errors.New("assymetric graph error")
			}
		}
	}
	return &MatrixGraph{
		matrix:   matrix,
		edges:    nil,
		vertices: vertices,
	}, nil
}

func (g *MatrixGraph) Neighbours(node *Vertex) []*Vertex {
	neighbours := []*Vertex{}
	id := node.Id
	for i, adjacent := range g.matrix[id] {
		if adjacent != 0 {
			neighbours = append(neighbours, g.vertices[i])
		}
	}
	return neighbours
}

func (g *MatrixGraph) Vertices() Vertices {
	return g.vertices
}
