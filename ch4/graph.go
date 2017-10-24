package main

import (
	"fmt"
)

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("England", "London")
	addEdge("Ukraine", "Kyiv")
	addEdge("USA", "Orlean")
	fmt.Println(graph)
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdges(from, to string) bool {
	return graph[from][to]
}
