package graph

type BaseGraph interface {
	AddEdge(u, v string)
	RemoveEdge(u, v string)
	HasEdge(u, v string) bool
	GetNeighbors(node string) []string
	GetNodes() []string
	BFS(start string) []string
	DFS(start string) []string
}

type Edge struct {
	From   string
	To     string
	Weight int
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
