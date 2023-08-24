package graph

import "errors"

// Adjacency matrix alias.
type adjacencyMatrix = [][]byte

// Adjacency matrix graph implementation.
type MatrixGraph struct {
	matrix    adjacencyMatrix
	edges     WeightedEdges
	vertices  Vertices
	heuristic Heuristic
}

// NewMatrixGraph initializes a MatrixGraph.
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
		heuristic: *new(map[int]int),
	}, nil
}

// Neighbours returns the neighbour list of a given vertex of the graph.
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

// Vertices returns the vertex list of the graph.
func (g *MatrixGraph) Vertices() Vertices {
	return g.vertices
}

// Edges returns the weighted edges map of the graph.
func (g *MatrixGraph) Edges() WeightedEdges {
	return g.edges
}

// InitializeVertices sets to default the fields from all the grpah vertices.
func (g *MatrixGraph) InitializeVertices() {
	for _, v := range g.vertices {
		v.Time, v.Level = -1, -1
		v.Priority = -1
		v.Visited = false
		v.Parent = nil
	}
}

// Heuristic returns the heuristic value form a given vertex of the graph.
func (g *MatrixGraph) Heuristic(v *Vertex) int {
	return g.heuristic[v.Id]
}

// SetHeuristic associates a heuristic to the graph.
func (g *MatrixGraph) SetHeuristic(h Heuristic) {
	g.heuristic = h
}
