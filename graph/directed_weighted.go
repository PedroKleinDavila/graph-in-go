package graph

type WeightedDirectedGraphInterface interface {
	BaseGraph
	AddEdge(u, v string, weight int)
	Dijkstra(start, end string) map[string][]WeightedEdge
	FloydWarshall() map[string]map[string]WeightedEdge
	TopologicalSort() ([]string, bool)
}

type WeightedDirectedGraph struct {
	adj map[string][]WeightedEdge
}

func NewWeightedDirectedGraph() *WeightedDirectedGraph {
	return &WeightedDirectedGraph{adj: make(map[string][]WeightedEdge)}
}
