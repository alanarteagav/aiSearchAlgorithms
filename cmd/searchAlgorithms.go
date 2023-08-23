package main

import (
	"aiSearchAlgorithms/pkg/algorithms"
	"aiSearchAlgorithms/testdata"
	"fmt"
)

func main() {

	g := testdata.GetTestGraph(testdata.Graph07WDFS)
	vertices := g.Vertices()

	//setting the goal
	vertices[5].Value = 1

	/*
		// =================== DFS ================
		fmt.Println("\nDFS execution:")
		answer := algorithms.DFS(g, vertices[0], 1)

		fmt.Println("\nVertices after DFS:")
		for _, v := range vertices {
			fmt.Println(v)
		}

		fmt.Println("\nPath from root:")

		path := algorithms.Path(answer)
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Println(path[i])
		}

		fmt.Println("\nTotal path weight:")
		fmt.Println(algorithms.PathWeight(answer, g.Edges()))

		// =================== DFS ================
		fmt.Println("\nBFS execution:")
		answer = algorithms.BFS(g, vertices[0], 1)

		vertices = g.Vertices()

		fmt.Println("\nVertices after BFS:")
		for _, v := range vertices {
			fmt.Println(v)
		}

		fmt.Println("\nPath from root:")

		path = algorithms.Path(answer)
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Println(path[i])
		}

		fmt.Println("\nTotal path weight:")
		fmt.Println(algorithms.PathWeight(answer, g.Edges()))

	*/

	// =================== IDS ================
	fmt.Println("\nIDS execution:")
	answer := algorithms.IDS(g, vertices[0], 1)

	vertices = g.Vertices()

	fmt.Println("\nVertices after IDS:")
	for _, v := range vertices {
		fmt.Println(v)
	}

	fmt.Println("\nPath from root:")

	path := algorithms.Path(answer)
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println(path[i])
	}

	fmt.Println("\nTotal path weight:")
	fmt.Println(algorithms.PathWeight(answer, g.Edges()))
}
