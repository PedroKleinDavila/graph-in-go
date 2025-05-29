package main

import (
	"fmt"

	"github.com/PedroKleinDavila/graph-in-go/graph"
)

func main() {
	testUnweightedGraph()
}

func testUnweightedGraph() {
	g := graph.NewUnweightedGraph()

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("E", "F")

	fmt.Println("Grafo:")
	fmt.Println(g.ToString())

	fmt.Println("HasEdge A-B:", g.HasEdge("A", "B"))
	fmt.Println("HasEdge A-D:", g.HasEdge("A", "D"))

	fmt.Println("Vizinhos de A:", g.GetNeighbors("A"))
	fmt.Println("Vizinhos de E:", g.GetNeighbors("E"))

	fmt.Println("BFS a partir de A:", g.BFS("A"))
	fmt.Println("DFS a partir de A:", g.DFS("A"))

	fmt.Println("É conectado?", g.IsConnected())

	fmt.Println("Caminho mais curto entre A e D:", g.ShortestPathUnweighted("A", "D"))
	fmt.Println("Caminho mais curto entre A e F:", g.ShortestPathUnweighted("A", "F"))

	fmt.Println("Componentes conectados:", g.ConnectedComponents())

	fmt.Println("Tem ciclo?", g.HasCycle())

	g.RemoveEdge("A", "B")
	fmt.Println("Após remover A-B:")
	fmt.Println(g.ToString())
	fmt.Println("HasEdge A-B:", g.HasEdge("A", "B"))
}
