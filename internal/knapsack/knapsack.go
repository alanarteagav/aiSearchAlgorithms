package knapsack

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/internal/graphbuilder"
	"aiSearchAlgorithms/internal/inputs"
	"aiSearchAlgorithms/pkg/algorithms"
	"fmt"
)

func Knapsack(path string, algorithm algorithms.SearchAlgorithm) {
	var bigPrize int

	items, capacity, pairs := inputs.GetKnapsackInfo(path)
	for _, p := range pairs {
		bigPrize += p.Value
	}

	builder := new(graphbuilder.GraphBuilder)
	var answer *graph.Vertex
	var parent *graph.Vertex
	answerPath := []*graph.Vertex{}

	for i := 0; i < items; i++ {
		edges := graph.WeightedEdges{}
		heuristic := graph.Heuristic{}

		edgeYes := graph.Edge{U: 0, V: 1}
		weightYes := capacity - pairs[i+1].Weight
		heuristic[1] = bigPrize - pairs[i+1].Value

		edgeNo := graph.Edge{U: 0, V: 2}
		weightNo := capacity
		heuristic[2] = bigPrize

		edges[edgeYes] = weightYes
		edges[edgeNo] = weightNo
		builder.WithGraphOrder(3)
		builder.WithEdges(edges)
		builder.WithHeuristic(heuristic)
		g := builder.Build(graph.WeightedMatrix)
		vertices := g.Vertices()
		vertices[0].Parent = parent
		vertices[1].Value = -1
		vertices[2].Value = -1

		answer = algorithm(g, g.Vertices()[0], -1)
		parent = answer
		answerPath = append(answerPath, answer)

		if capacity > edges[graph.Edge{U: 0, V: answer.Id}] {
			capacity = edges[graph.Edge{U: 0, V: answer.Id}]
			fmt.Println("new capacity:", capacity)
		}

		fmt.Println("ANSWER")
		fmt.Println(answer)
		fmt.Println("V2==========================")
		for _, v := range g.Vertices() {
			fmt.Println(v)
		}
		//fmt.Println(algorithms.Path(answer))
		fmt.Println("path")
		for _, v := range answerPath {
			fmt.Println(v)
		}
	}
}
