package graph

type DirectedGraphInterface interface {
	BaseGraph
	AddEdge(u, v string)
	ShortestPathUnweighted(start, end string) []string
	TopologicalSort() ([]string, bool)
}

type DirectedGraph struct {
	adj map[string][]string
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{adj: make(map[string][]string)}
}

func (g *DirectedGraph) AddEdge(u, v string) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *DirectedGraph) HasEdge(u, v string) bool {
	neighbors := g.GetNeighbors(u)
	for _, neighbor := range neighbors {
		if neighbor == v {
			return true
		}
	}
	return false
}

func (g *DirectedGraph) GetNeighbors(node string) []string {
	if neighbors, ok := g.adj[node]; ok {
		return neighbors
	}
	return nil
}

func (g *DirectedGraph) RemoveEdge(u, v string) {
	uNeighbors := g.GetNeighbors(u)
	for i, neighbor := range uNeighbors {
		if neighbor == v {
			g.adj[u] = append(uNeighbors[:i], uNeighbors[i+1:]...)
			break
		}
	}
}

func (g *DirectedGraph) GetNodes() []string {
	nodeSet := make(map[string]struct{})
	for u, neighbors := range g.adj {
		nodeSet[u] = struct{}{}
		for _, v := range neighbors {
			nodeSet[v] = struct{}{}
		}
	}
	nodes := make([]string, 0, len(nodeSet))
	for node := range nodeSet {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g *DirectedGraph) BFS(start string) []string {
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

func (g *DirectedGraph) DFS(start string) []string {
	visited := make(map[string]bool)

	return dfs(start, visited, g)
}

func (g *DirectedGraph) IsConnected() bool {
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

func (g *DirectedGraph) ShortestPathUnweighted(start, end string) []string {
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

func (g *DirectedGraph) ConnectedComponents() [][]string {
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

func (g *DirectedGraph) HasCycle() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for _, node := range g.GetNodes() {
		if g.hasCycleUtil(node, visited, recStack) {
			return true
		}
	}
	return false
}

func (g *DirectedGraph) hasCycleUtil(node string, visited map[string]bool, recStack map[string]bool) bool {
	if recStack[node] {
		return true
	}
	if visited[node] {
		return false
	}

	visited[node] = true
	recStack[node] = true

	for _, neighbor := range g.GetNeighbors(node) {
		if g.hasCycleUtil(neighbor, visited, recStack) {
			return true
		}
	}

	recStack[node] = false
	return false
}

func (g *DirectedGraph) ToString() string {
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

func (g *DirectedGraph) TopologicalSort() ([]string, bool) {
	if g.HasCycle() {
		return nil, false
	}

	visited := make(map[string]bool)
	stack := []string{}

	var dfsTopo func(string)
	dfsTopo = func(node string) {
		visited[node] = true
		for _, neighbor := range g.GetNeighbors(node) {
			if !visited[neighbor] {
				dfsTopo(neighbor)
			}
		}
		stack = append([]string{node}, stack...)
	}

	for _, node := range g.GetNodes() {
		if !visited[node] {
			dfsTopo(node)
		}
	}

	return stack, true
}
