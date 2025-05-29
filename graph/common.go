package graph

type Graph interface {
	AddEdge(u, v string)
	RemoveEdge(u, v string)
	HasEdge(u, v string) bool
	AddWeightedEdge(u, v string, weight int)

	GetNeighbors(node string) []string
	GetNodes() []string
	IsConnected() bool
	ConnectedComponents() [][]string

	BFS(start string) []string
	DFS(start string) []string

	ShortestPathUnweighted(start, end string) []string
	Dijkstra(start string) map[string]int
	FloydWarshall() map[string]map[string]int

	TopologicalSort() ([]string, bool)
	HasCycle() bool

	MinimumSpanningTree() []Edge
	ToString() string
}

type Edge struct {
	From   string
	To     string
	Weight int
}
