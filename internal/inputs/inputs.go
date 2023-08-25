// Package inputs allows to read lines from some simple files.
package inputs

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read reads lines from a file.
func Read(path string) []string {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// KnapsackItem stores the value-weight pair of a KS problem item.
type KnapsackItem struct {
	Value  int
	Weight int
}

// GetKnapsackInfo gets the data from a file with a KS problem instance.
func GetKnapsackInfo(path string) (int, int, map[int]KnapsackItem) {
	knapsackItems := map[int]KnapsackItem{}

	lines := Read(path)
	header := strings.Fields(strings.TrimSpace(lines[0]))
	items, err := strconv.Atoi(header[0])
	if err != nil {
		log.Fatal(err)
	}
	capacity, err := strconv.Atoi(header[1])
	if err != nil {
		log.Fatal(err)
	}
	lines[0] = "0 0"
	for i, line := range lines {
		pair := strings.Fields(strings.TrimSpace(line))
		value, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatal(err)
		}
		weight, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatal(err)
		}
		knapsackItems[i] = KnapsackItem{Value: value, Weight: weight}
	}
	return items, capacity, knapsackItems
}
