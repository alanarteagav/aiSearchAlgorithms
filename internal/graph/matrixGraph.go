package graph

import "errors"

type adjacencyMatrix = [][]byte

type MatrixGraph struct {
	matrix    adjacencyMatrix
	edges     WeightedEdges
	vertices  Vertices
	heuristic Heuristic
}

func NewMatrixGraph(matrix adjacencyMatrix) (*MatrixGraph, error) {
	vertices := map[int]*Vertex{}
	for i, v := range matrix {
		vertices[i] = NewVertex(i)
		for j, w := range v {
			if i < j && w != matrix[j][i] {
				return nil, errors.New("assymetric graph error")
			}
		}
	}
	return &MatrixGraph{
		matrix:    matrix,
		edges:     nil,
		vertices:  vertices,
		heuristic: *new(map[int]float64),
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

func (g *MatrixGraph) Edges() WeightedEdges {
	return g.edges
}

func (g *MatrixGraph) InitializeVertices() {
	for _, v := range g.vertices {
		v.Time, v.Level = -1, -1
		v.Priority = -1
		v.Visited = false
		v.Parent = nil
	}
}

func (g *MatrixGraph) SetHeuristic(h Heuristic) {
	g.heuristic = h
}

func (g *MatrixGraph) Heuristic(v *Vertex) float64 {
	return g.heuristic[v.Id]
}
