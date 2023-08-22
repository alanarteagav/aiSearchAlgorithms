package main

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/pkg/algorithms"
	"fmt"
)

func main() {
	fmt.Println("hello")

	matrix := [][]byte{
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 0, 1, 0},
	}

	var g graph.Graph
	g, _ = graph.NewMatrixGraph(matrix)

	vertices := g.Vertices()
	fmt.Println(vertices)

	/*
		for _, v := range vertices {
			fmt.Println(v)
		}
	*/

	for _, n := range g.Neighbours(vertices[1]) {
		fmt.Println(n)
	}

	fmt.Println("===")

	algorithms.Dfs(g, vertices[1], 0)

}
