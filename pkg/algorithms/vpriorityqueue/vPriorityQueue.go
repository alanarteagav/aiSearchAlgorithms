// Package vpriorityqueue defines a min-heap priority queue for graph vertices.
package vpriorityqueue

import (
	"aiSearchAlgorithms/internal/graph"
	"container/heap"
)

// VPriorityQueue min-heap priority queue for graph vertices.
type VPriorityQueue []*graph.Vertex

// Len returns the length of the priority queue.
func (pq VPriorityQueue) Len() int {
	return len(pq)
}

// Less defines the comparison method of the priority queue.
func (pq VPriorityQueue) Less(i int, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap defines the swap method of the priority queue.
func (pq VPriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push defines the push method of the priority queue.
func (pq *VPriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*graph.Vertex))
}

// Pop defines the pop method of the priority queue.
func (pq *VPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// AddVertex adds a vertex to the priority queue.
func (pq *VPriorityQueue) AddVertex(v *graph.Vertex) {
	heap.Push(pq, v)
}

// PopVertex removes the vertex at the top of the priority queue and returns it.
func (pq *VPriorityQueue) PopVertex() *graph.Vertex {
	v := heap.Pop(pq).(*graph.Vertex)
	return v
}
