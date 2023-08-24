package main

import (
	"aiSearchAlgorithms/pkg/algorithms"
	"aiSearchAlgorithms/testdata"
	"fmt"
)

func main() {

	g := testdata.GetTestGraph(testdata.GraphRomania)

	vertices := g.Vertices()

	//setting the goal
	vertices[12].Value = 1

	fmt.Println("\nUCS execution:")
	answer := algorithms.GBS(g, vertices[0], 1)

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

	g.SetHeuristic(map[int]int{0: 42})

	//fmt.Println(g.Heuristic(vertices[0]))
}
