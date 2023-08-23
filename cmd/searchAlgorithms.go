package main

import (
	"aiSearchAlgorithms/pkg/algorithms"
	"aiSearchAlgorithms/testdata"
	"fmt"
)

func main() {
	g := testdata.GetTestGraph(testdata.Graph07W)

	vertices := g.Vertices()

	fmt.Println("\nVertices before DFS:")
	for _, v := range vertices {
		fmt.Println(v)
	}

	fmt.Println("\nDFS execution:")
	algorithms.DFS(g, vertices[0], 0)

	fmt.Println("\nVertices after DFS:")
	for _, v := range vertices {
		fmt.Println(v)
	}

	fmt.Println("\nPath from root:")

	path := algorithms.Path(vertices[6])
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println(path[i])
	}

	fmt.Println("\nTotal path weight:")
	fmt.Println(algorithms.PathWeight(vertices[6], g.Edges()))

	fmt.Println("\nBFS execution:")
	algorithms.BFS(g, vertices[0], 0)

	vertices = g.Vertices()

	fmt.Println("\nVertices after BFS:")
	for _, v := range vertices {
		fmt.Println(v)
	}

	fmt.Println("\nPath from root:")

	path = algorithms.Path(vertices[6])
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println(path[i])
	}

	fmt.Println("\nTotal path weight:")
	fmt.Println(algorithms.PathWeight(vertices[6], g.Edges()))
}
