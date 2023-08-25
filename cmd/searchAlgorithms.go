package main

import (
	"aiSearchAlgorithms/internal/graph"
	"aiSearchAlgorithms/internal/inputs"
	"aiSearchAlgorithms/internal/knapsack"
	"aiSearchAlgorithms/pkg/algorithms"
	"aiSearchAlgorithms/testdata"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("[r]omania (press r) or [k]napsack (press k), then hit enter:")
	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	option = strings.TrimSpace(option)
	switch option {
	case "r":
		fmt.Println("===================")
		fmt.Println("Select algorithm:")
		fmt.Println("[d]fs (depth first search) [press d & hit enter]\n" +
			"[b]fs (breadth first search) [press b & hit enter]\n" +
			"[i]ds (iterative deepening depth-first search) [press i & hit enter]\n" +
			"[u]cs (uniform cost search) [press u & hit enter]\n" +
			"[g]bs (greedy best-first search ) [press g & hit enter]\n" +
			"[a]star (a star) [press a & hit enter]")
		algorithm, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		algorithm = strings.TrimSpace(algorithm)
		fmt.Println("______________________")
		fmt.Println("Select source vertex [0-19]:")
		source, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		i, err := strconv.Atoi(strings.TrimSpace(source))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("______________________")
		fmt.Println("Select goal vertex [0-19]:")
		goal, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		j, err := strconv.Atoi(strings.TrimSpace(goal))
		if err != nil {
			log.Fatal(err)
		}
		var answer *graph.Vertex
		g := testdata.GetTestGraph(testdata.GraphRomania)
		u := g.Vertices()[i]
		g.Vertices()[j].Value = 1

		fmt.Println("======================")
		switch algorithm {
		case "b":
			fmt.Println("[BFS algorithm]:")
			answer = algorithms.BFS(g, u, 1)
		case "d":
			fmt.Println("[DFS algorithm]:")
			answer = algorithms.DFS(g, u, 1)
		case "i":
			fmt.Println("[IDS algorithm]:")
			answer = algorithms.IDS(g, u, 1)
		case "u":
			fmt.Println("[UCS algorithm]:")
			answer = algorithms.UCS(g, u, 1)
		case "g":
			fmt.Println("[GBS algorithm]:")
			answer = algorithms.GBS(g, u, 1)
		case "a":
			fmt.Println("[A Star algorithm]:")
			answer = algorithms.AStar(g, u, 1)
		default:
			fmt.Println("no algorithm chosen. exiting")
			return
		}
		fmt.Println("ANSWER:")
		fmt.Printf("• Name: %s, ID: %d. \n", answer.Name, answer.Id)
		fmt.Println("......................")
		fmt.Println("path from root:")
		path := algorithms.Path(answer)
		for i := len(path) - 1; i >= 0; i-- {
			pathVertex := path[i]
			if pathVertex.Parent != nil {
				fmt.Printf("↓ /%d/\n",
					g.Edges()[graph.Edge{U: pathVertex.Parent.Id, V: path[i].Id}])
			}
			fmt.Printf("• Name: %s, ID: %d. \n", pathVertex.Name, pathVertex.Id)
		}
		fmt.Println("......................")
		fmt.Println("Total path weight:")
		fmt.Println(algorithms.PathWeight(answer, g.Edges()))
		fmt.Println("......................")
	case "k":
		fmt.Println("===================")
		fmt.Println("Select KS problem instance [press number inside brackets & hit enter]:")
		fmt.Println("[0] KS 3\n" +
			"[1] KS 5\n" +
			"[2] KS 15\n" +
			"[3] KS 19 (★)\n" +
			"[4] KS 10000 (★)\n" +
			"[5] KS 100001")
		instance, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		ks, err := strconv.Atoi(strings.TrimSpace(instance))
		if err != nil {
			log.Fatal(err)
		}

		var indexes []int
		var items []inputs.KnapsackItem
		fmt.Println("======================")
		fmt.Println("Running with A* algorithm:")
		switch ks {
		case 0:
			fmt.Println("[KS 3]")
			indexes, items = knapsack.Knapsack("testdata/ks_3_0", algorithms.AStar)
		case 1:
			fmt.Println("[KS 5]")
			indexes, items = knapsack.Knapsack("testdata/ks_5_0", algorithms.AStar)
		case 2:
			fmt.Println("[KS 15]")
			indexes, items = knapsack.Knapsack("testdata/ks_15_0", algorithms.AStar)
		case 3:
			fmt.Println("[KS 19]")
			indexes, items = knapsack.Knapsack("testdata/ks_19_0", algorithms.AStar)
		case 4:
			fmt.Println("[KS 10000]")
			indexes, items = knapsack.Knapsack("testdata/ks_10000_0", algorithms.AStar)
		case 5:
			fmt.Println("[KS 10001]")
			indexes, items = knapsack.Knapsack("testdata/ks_10001_0", algorithms.AStar)
		default:
			fmt.Println("no instance chosen. exiting")
			return
		}
		fmt.Println("______________________")
		fmt.Println("SOLUTION:")
		fmt.Println("KS items indexes:", indexes)
		fmt.Println("......................")
		fmt.Println("KS items values and weights:")
		for _, item := range items {
			fmt.Printf("• 🅅 %d, 🅆 %d. \n", item.Value, item.Weight)
		}
		fmt.Println("......................")
		fmt.Println("KS total value and weight:")
		totalValue := 0
		totalWeight := 0

		for _, item := range items {
			totalValue += item.Value
			totalWeight += item.Weight
		}
		fmt.Printf("• 🅅 %d, 🅆 %d. \n", totalValue, totalWeight)
	default:
		fmt.Println("no option chosen. exiting")
	}
}
