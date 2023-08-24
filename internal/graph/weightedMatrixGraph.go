package graph

import "errors"

// Weighted adjacency matrix alias.
type weightedAdjacencyMatrix = [][]int

// Weighted adjacency matrix graph implementation.
type WeightedMatrixGraph struct {
	matrix    weightedAdjacencyMatrix
	edges     WeightedEdges
	vertices  Vertices
	heuristic Heuristic
}

// NewWeightedMatrixGraph initializes a WeightedMatrixGraph.
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
		matrix:    matrix,
		edges:     edges,
		vertices:  vertices,
		heuristic: *new(map[int]int),
	}, nil
}

// Neighbours returns the neighbour list of a given vertex of the graph.
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

// Vertices returns the vertex list of the graph.
func (g *WeightedMatrixGraph) Vertices() Vertices {
	return g.vertices
}

// Edges returns the weighted edges map of the graph.
func (g *WeightedMatrixGraph) Edges() WeightedEdges {
	return g.edges
}

// InitializeVertices sets to default the fields from all the grpah vertices.
func (g *WeightedMatrixGraph) InitializeVertices() {
	for _, v := range g.vertices {
		v.Time, v.Level = -1, -1
		v.Priority = -1
		v.Visited = false
		v.Parent = nil
	}
}

// Heuristic returns the heuristic value form a given vertex of the graph.
func (g *WeightedMatrixGraph) Heuristic(v *Vertex) int {
	return g.heuristic[v.Id]
}

// SetHeuristic associates a heuristic to the graph.
func (g *WeightedMatrixGraph) SetHeuristic(h Heuristic) {
	g.heuristic = h
}
