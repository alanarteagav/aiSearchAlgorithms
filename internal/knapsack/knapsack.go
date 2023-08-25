package knapsack

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/internal/graphbuilder"
	"aiSearchAlgorithms/internal/inputs"
	"aiSearchAlgorithms/pkg/algorithms"
	"fmt"
)

func Knapsack(path string,
	algorithm algorithms.SearchAlgorithm) []inputs.KnapsackItem {

	knapsackItems := []inputs.KnapsackItem{}
	bigPrize := 0

	items, capacity, pairs := inputs.GetKnapsackInfo(path)
	for _, p := range pairs {
		bigPrize += p.Value
	}

	builder := new(graphbuilder.GraphBuilder)
	var answer *graph.Vertex
	var root *graph.Vertex
	answerPath := []*graph.Vertex{}

	for i := 0; len(pairs) > 0; i++ {
		edges := graph.WeightedEdges{}
		heuristic := graph.Heuristic{}

		for j, pair := range pairs {
			edge := graph.Edge{U: 0, V: j}
			weight := capacity - pair.Weight
			heuristic[j] = bigPrize - pair.Value
			edges[edge] = weight
		}

		builder.WithGraphOrder(items + 1)
		builder.WithEdges(edges)
		builder.WithHeuristic(heuristic)
		g := builder.Build(graph.WeightedMatrix)
		vertices := g.Vertices()
		vertices[0].Parent = root

		for k := 1; k < len(vertices); k++ {
			weight, ok := edges[graph.Edge{U: 0, V: k}]
			if ok && weight >= 0 {
				vertices[k].Value = -1
			}
		}

		answer = algorithm(g, g.Vertices()[0], -1)
		if answer == nil {
			break
		}
		root = answer
		answerPath = append(answerPath, answer)

		newCapacity := edges[graph.Edge{U: 0, V: answer.Id}]

		if newCapacity < 0 {
			break
		}

		if newCapacity < capacity {
			capacity = edges[graph.Edge{U: 0, V: answer.Id}]
		}

		delete(pairs, answer.Id)
	}
	fmt.Println("path")
	for _, v := range answerPath {
		fmt.Println(v)
	}

	return knapsackItems
}
