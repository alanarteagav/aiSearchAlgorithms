package knapsack

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/internal/graphbuilder"
	"aiSearchAlgorithms/internal/inputs"
	"aiSearchAlgorithms/pkg/algorithms"
	"fmt"
)

// Knapsack problem solver. Receives a path to the file containing an instance
// of the problem and a graph search function.. Returns an array of indexes of
// the solution items and an array of solution items.
func Knapsack(path string,
	algorithm algorithms.SearchAlgorithm) ([]int, []inputs.KnapsackItem) {

	knapsackIndexes := []int{}
	knapsackItems := []inputs.KnapsackItem{}

	bigPrize := 0

	items, capacity, pairs := inputs.GetKnapsackInfo(path)
	for _, p := range pairs {
		bigPrize += p.Value
	}

	builder := new(graphbuilder.GraphBuilder)
	var answer *graph.Vertex
	var root *graph.Vertex

	fmt.Printf("◉ Initial capacity: %d\n", capacity)

	for i := 0; len(pairs) > 0; i++ {
		edges := graph.WeightedEdges{}
		heuristic := graph.Heuristic{}

		for j, pair := range pairs {
			edge := graph.Edge{U: 0, V: j}
			weight := pair.Weight
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
				vertices[k].Value = 1
				if weight == 0 {
					fmt.Println("k", k)
				}
			}
		}

		answer = algorithm(g, g.Vertices()[0], 1)

		if answer == nil {
			break
		}
		root = answer
		knapsackItems = append(knapsackItems, pairs[answer.Id])
		knapsackIndexes = append(knapsackIndexes, answer.Id)

		newCapacity := edges[graph.Edge{U: 0, V: answer.Id}]

		if newCapacity <= 0 {
			break
		}

		if newCapacity < capacity {
			capacity = edges[graph.Edge{U: 0, V: answer.Id}]
		}

		fmt.Printf("‣ Added item, remaining capacity: %d\n", newCapacity)

		delete(pairs, answer.Id)
	}

	return knapsackIndexes, knapsackItems
}
