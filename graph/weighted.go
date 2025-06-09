package graph

import (
	"fmt"
	"math"
)

type WeightedGraphInterface interface {
	BaseGraph
	AddEdge(u, v string, weight int)
	GetWeightedNeighbors(node string) []WeightedEdge
	Dijkstra(start, end string) map[string][]WeightedEdge
	FloydWarshall() map[string]map[string]WeightedEdge
	TopologicalSort() ([]string, bool)
	MinimumSpanningTree() []WeightedEdge
}

type WeightedGraph struct {
	adj map[string][]WeightedEdge
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{adj: make(map[string][]WeightedEdge)}
}

func (g *WeightedGraph) AddEdge(u, v string, weight int) {
	if weight < 0 {
		return
	}
	g.adj[u] = append(g.adj[u], WeightedEdge{To: v, Weight: weight})
	g.adj[v] = append(g.adj[v], WeightedEdge{To: u, Weight: weight})
}

func (g *WeightedGraph) HasEdge(u, v string) bool {
	neighbors := g.GetWeightedNeighbors(u)
	for _, edge := range neighbors {
		if edge.To == v {
			return true
		}
	}
	return false
}

func (g *WeightedGraph) GetNeighbors(node string) []string {
	if edges, ok := g.adj[node]; ok {
		neighbors := make([]string, len(edges))
		for i, edge := range edges {
			neighbors[i] = edge.To
		}
		return neighbors
	}
	return nil
}

func (g *WeightedGraph) GetWeightedNeighbors(node string) []WeightedEdge {
	if edges, ok := g.adj[node]; ok {
		return edges
	}
	return nil
}

func (g *WeightedGraph) RemoveEdge(u, v string) {
	uNeighbors := g.GetWeightedNeighbors(u)
	for i, edge := range uNeighbors {
		if edge.To == v {
			g.adj[u] = append(uNeighbors[:i], uNeighbors[i+1:]...)
			break
		}
	}
	vNeighbors := g.GetWeightedNeighbors(v)
	for i, edge := range vNeighbors {
		if edge.To == u {
			g.adj[v] = append(vNeighbors[:i], vNeighbors[i+1:]...)
			break
		}
	}
}

func (g *WeightedGraph) GetNodes() []string {
	nodeSet := make(map[string]struct{})
	for u, neighbors := range g.adj {
		nodeSet[u] = struct{}{}
		for _, edge := range neighbors {
			nodeSet[edge.To] = struct{}{}
		}
	}
	nodes := make([]string, 0, len(nodeSet))
	for node := range nodeSet {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g *WeightedGraph) BFS(start string) []string {
	visited := make(map[string]bool)
	queue := []string{start}
	var result []string
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		visited[node] = true
		result = append(result, node)
		for _, edge := range g.GetWeightedNeighbors(node) {
			if !visited[edge.To] {
				queue = append(queue, edge.To)
			}
		}
	}
	return result
}

func (g *WeightedGraph) DFS(start string) []string {
	visited := make(map[string]bool)
	result := dfs(start, visited, g)
	return result
}

func (g *WeightedGraph) IsConnected() bool {
	nodes := g.GetNodes()
	if len(nodes) == 0 {
		return true
	}
	visited := make(map[string]bool)
	dfs(nodes[0], visited, g)
	for _, node := range nodes {
		if !visited[node] {
			return false
		}
	}
	return true
}

func (g *WeightedGraph) ConnectedComponents() [][]string {
	nodes := g.GetNodes()
	visited := make(map[string]bool)
	var components [][]string

	for _, node := range nodes {
		if !visited[node] {
			component := []string{}
			queue := []string{node}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				if visited[current] {
					continue
				}
				visited[current] = true
				component = append(component, current)
				for _, edge := range g.GetWeightedNeighbors(current) {
					if !visited[edge.To] {
						queue = append(queue, edge.To)
					}
				}
			}
			components = append(components, component)
		}
	}
	return components
}

func (g *WeightedGraph) HasCycle() bool {
	nodes := g.GetNodes()
	visited := make(map[string]bool)

	for _, node := range nodes {
		if !visited[node] {
			if hasCycleUtil(node, visited, "", g) {
				return true
			}
		}
	}
	return false
}

func (g *WeightedGraph) ToString() string {
	var result string
	for u, neighbors := range g.adj {
		result += u + ": "
		for _, edge := range neighbors {
			result += edge.To + "(" + fmt.Sprint(edge.Weight) + ") "
		}
		result += "\n"
	}
	return result
}

func (g *WeightedGraph) Dijkstra(start, end string) DijkstraResult {
	distances := make(map[string]int)
	visited := make(map[string]bool)
	parent := make(map[string]string)
	nodes := g.GetNodes()
	for _, node := range nodes {
		distances[node] = math.MaxInt
		visited[node] = false
		parent[node] = ""
	}
	distances[start] = 0
	for range nodes {
		u := minimumDistance(distances, visited, g)
		if u == "" {
			break
		}
		visited[u] = true
		for _, neighbor := range g.GetWeightedNeighbors(u) {
			v := neighbor.To
			weight := neighbor.Weight
			if !visited[v] && distances[u] != math.MaxInt && distances[u]+weight < distances[v] {
				distances[v] = distances[u] + weight
				parent[v] = u
			}
		}
	}
	path := []string{}
	current := end
	totalCost := 0
	for parent[current] != "" {
		path = append([]string{current}, path...)
		prev := parent[current]

		for _, edge := range g.GetWeightedNeighbors(prev) {
			if edge.To == current {
				totalCost += edge.Weight
				break
			}
		}
		current = prev
	}
	path = append([]string{start}, path...)

	return DijkstraResult{
		Path: path,
		Cost: totalCost,
	}
}

func (g *WeightedGraph) FloydWarshall() map[string]map[string]WeightedEdge {
	return nil
}

func (g *WeightedGraph) MinimumSpanningTree() []WeightedEdge {
	return nil
}
