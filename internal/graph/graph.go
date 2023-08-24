// Package graph provides an interface and several implementations for graphs.
package graph

// Vertex represent a graph vertex.
type Vertex struct {
	Name     string
	Id       int
	Time     int
	Level    int
	Priority int
	Visited  bool
	Parent   *Vertex
	Value    int
}

// Returns a new initialized vertex, receives its id as an integer.
func NewVertex(id int) *Vertex {
	v := new(Vertex)
	v.Id = id
	v.Time, v.Level = -1, -1
	v.Priority = -1
	v.Visited = false
	v.Parent = nil
	return v
}

// Edge defines a graph edge through the ids of it two ends.
type Edge struct {
	U int
	V int
}

// Map with graph edges keys and weights values.
type WeightedEdges map[Edge]int

// Map with vertex ids keys and graph vertices values.
type Vertices map[int]*Vertex

// Map with vertex ids keys and heuristic values.
type Heuristic map[int]int

// Enum to define the different graph implementations.
type GraphImplementation string

const (
	Matrix         GraphImplementation = "Matrix Graph"
	WeightedMatrix GraphImplementation = "Weighted Matrix Graph"
)

// Graph interface. Defines the basic desired behaviour for a graph in order to
// be able to perform searches over it.
type Graph interface {
	// Neighbours returns the neighbour list of a given vertex of the graph.
	Neighbours(*Vertex) []*Vertex
	// Vertices returns the vertex list of the graph.
	Vertices() Vertices
	// Edges returns the weighted edges map of the graph.
	Edges() WeightedEdges
	// InitializeVertices sets to default the fields from all the grpah vertices.
	InitializeVertices()
	// Heuristic returns the heuristic value form a given vertex of the graph.
	Heuristic(*Vertex) int
	// SetHeuristic associates a heuristic to the graph.
	SetHeuristic(Heuristic)
}
