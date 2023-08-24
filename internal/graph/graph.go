package graph

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

func NewVertex(id int) *Vertex {
	v := new(Vertex)
	v.Id = id
	v.Time, v.Level = -1, -1
	v.Priority = -1
	v.Visited = false
	v.Parent = nil
	return v
}

type Edge struct {
	U int
	V int
}

type WeightedEdges map[Edge]int
type Vertices map[int]*Vertex
type Heuristic map[int]int

type Graph interface {
	Neighbours(*Vertex) []*Vertex
	Vertices() Vertices
	Edges() WeightedEdges
	InitializeVertices()
	Heuristic(*Vertex) int
	SetHeuristic(Heuristic)
}

type GraphImplementation string

const (
	Matrix         GraphImplementation = "Matrix Graph"
	WeightedMatrix GraphImplementation = "Weighted Matrix Graph"
)
