package graph

type UnweightedGraphInterface interface {
	BaseGraph
	AddEdge(u, v string)
	ShortestPathUnweighted(start, end string) []string
}

type UnweightedGraph struct {
	adj map[string][]string
}

func NewUnweightedGraph() *UnweightedGraph {
	return &UnweightedGraph{adj: make(map[string][]string)}
}

func (g *UnweightedGraph) AddEdge(u, v string) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *UnweightedGraph) HasEdge(u, v string) bool {
	neighbors := g.GetNeighbors(u)
	for _, neighbor := range neighbors {
		if neighbor == v {
			return true
		}
	}
	return false
}

func (g *UnweightedGraph) GetNeighbors(node string) []string {
	if neighbors, ok := g.adj[node]; ok {
		return neighbors
	}
	return nil
}

func (g *UnweightedGraph) RemoveEdge(u, v string) {
	uNeighbors := g.GetNeighbors(u)
	for i, neighbor := range uNeighbors {
		if neighbor == v {
			g.adj[u] = append(uNeighbors[:i], uNeighbors[i+1:]...)
			break
		}
	}
	vNeighbors := g.GetNeighbors(v)
	for i, neighbor := range vNeighbors {
		if neighbor == u {
			g.adj[v] = append(vNeighbors[:i], vNeighbors[i+1:]...)
			break
		}
	}
}

func (g *UnweightedGraph) GetNodes() []string {
	nodes := make([]string, 0, len(g.adj))
	for node := range g.adj {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g *UnweightedGraph) BFS(start string) []string {
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
		for _, neighbor := range g.GetNeighbors(node) {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

func (g *UnweightedGraph) DFS(start string) []string {
	visited := make(map[string]bool)

	return dfs(start, visited, g)
}

func (g *UnweightedGraph) IsConnected() bool {
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

func (g *UnweightedGraph) ShortestPathUnweighted(start, end string) []string {
	if start == end {
		return []string{start}
	}

	visited := make(map[string]bool)
	prev := make(map[string]string)
	queue := []string{start}
	visited[start] = true
	found := false

	for len(queue) > 0 && !found {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.GetNeighbors(node) {
			if !visited[neighbor] {
				visited[neighbor] = true
				prev[neighbor] = node
				if neighbor == end {
					found = true
					break
				}
				queue = append(queue, neighbor)
			}
		}
	}

	if !found {
		return nil
	}

	path := []string{}
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
		if at == start {
			break
		}
	}

	if len(path) > 0 && path[0] == start {
		return path
	}

	return nil
}

func (g *UnweightedGraph) ConnectedComponents() [][]string {
	visited := make(map[string]bool)
	var components [][]string
	for _, node := range g.GetNodes() {
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
				for _, neighbor := range g.GetNeighbors(current) {
					if !visited[neighbor] {
						queue = append(queue, neighbor)
					}
				}
			}
			components = append(components, component)
		}
	}
	return components
}

func (g *UnweightedGraph) HasCycle() bool {
	visited := make(map[string]bool)
	for _, node := range g.GetNodes() {
		if !visited[node] {
			if hasCycleUtil(node, visited, "", g) {
				return true
			}
		}
	}
	return false
}

func (g *UnweightedGraph) ToString() string {
	result := ""
	for node, neighbors := range g.adj {
		result += node + ": "
		for _, neighbor := range neighbors {
			result += neighbor + " "
		}
		result += "\n"
	}
	return result
}
