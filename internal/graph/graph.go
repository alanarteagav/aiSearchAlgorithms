package graph

type Vertex struct {
	Id       int
	TimeA    int
	TimeB    int
	Level    int
	Priority int
	Visited  bool
	Parent   *Vertex
}

func InitializeVertex(id int) *Vertex {
	v := new(Vertex)
	v.Id = id
	v.TimeA, v.TimeB, v.Level = -1, -1, -1
	v.Priority = -1
	v.Visited = false
	v.Parent = nil
	return v
}

type Edge struct {
	u int
	v int
}

type WeightedEdges map[Edge]float64
type Vertices map[int]*Vertex

type Graph interface {
	Neighbours(*Vertex) []*Vertex
	Vertices() Vertices
}
