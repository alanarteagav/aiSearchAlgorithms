package vpriorityqueue

import (
	"aiSearchAlgorithms/internal/graph"
	"container/heap"
)

type VPriorityQueue []*graph.Vertex

func (pq VPriorityQueue) Len() int {
	return len(pq)
}

func (pq VPriorityQueue) Less(i int, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq VPriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *VPriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*graph.Vertex))
}

func (pq *VPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq *VPriorityQueue) AddVertex(v *graph.Vertex) {
	heap.Push(pq, v)
}

func (pq *VPriorityQueue) PopVertex() *graph.Vertex {
	v := heap.Pop(pq).(*graph.Vertex)
	return v
}
