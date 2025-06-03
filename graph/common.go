package graph

import "math"

type BaseGraph interface {
	RemoveEdge(u, v string)
	HasEdge(u, v string) bool
	GetNeighbors(node string) []string
	GetNodes() []string
	BFS(start string) []string
	DFS(start string) []string
	IsConnected() bool
	ConnectedComponents() [][]string
	HasCycle() bool
	ToString() string
}

type WeightedEdge struct {
	To     string
	Weight int
}

type DijkstraResult struct {
	Path []string
	Cost int
}

func dfs(node string, visited map[string]bool, g BaseGraph) []string {
	if visited[node] {
		return nil
	}
	var result []string
	visited[node] = true
	result = append(result, node)
	for _, neighbor := range g.GetNeighbors(node) {
		if !visited[neighbor] {
			result = append(result, dfs(neighbor, visited, g)...)
		}
	}
	return result
}

func hasCycleUtil(node string, visited map[string]bool, parent string, g BaseGraph) bool {
	visited[node] = true
	for _, neighbor := range g.GetNeighbors(node) {
		if !visited[neighbor] {
			if hasCycleUtil(neighbor, visited, node, g) {
				return true
			}
		} else if neighbor != parent {
			return true
		}
	}
	return false
}

func minimumDistance(distances map[string]int, visited map[string]bool, g BaseGraph) string {
	min := math.MaxInt
	minNode := ""
	for _, node := range g.GetNodes() {
		if visited[node] == false && distances[node] <= min {
			min = distances[node]
			minNode = node
		}
	}
	return minNode
}
