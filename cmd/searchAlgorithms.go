package main

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/internal/graphbuilder"
	"aiSearchAlgorithms/pkg/algorithms"
	"fmt"
)

func main() {

	b := new(graphbuilder.GraphBuilder)
	b.WithGraphOrder(5)
	b.WithEdges(map[graph.Edge]int{
		{U: 0, V: 1}: 80,
		{U: 0, V: 2}: 99,
		{U: 1, V: 3}: 97,
		{U: 2, V: 4}: 211,
		{U: 3, V: 4}: 101,
	})
	g := b.Build(graph.WeightedMatrix)
	//g := testdata.GetTestGraph(testdata.GraphRomania0)

	vertices := g.Vertices()

	//setting the goal
	vertices[4].Value = 1

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
	answer := algorithms.UCS(g, vertices[0], 1)

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

	g.SetHeuristic(map[int]float64{0: 42.42})

	fmt.Println(g.Heuristic(vertices[0]))
}
