package main

import (
	"fmt"

	"github.com/shoprunback/go-raph/raph"
)

type MyDijkstra struct {
	raph.Dijkstra
}

func NewMyDijkstra(g raph.Graph) *MyDijkstra {
	dijkstra := raph.NewDijkstra(g)
	return &MyDijkstra{*dijkstra}
}

// override raph.Dijkstra.ShortestPath()
func (d *MyDijkstra) ShortestPath(from, to string, constraint raph.Constraint, minimize string) ([]string, int) {
	d.Reset()

	// init dijkstra with distance 0 for first vertex
	d.Costs[from] = 0

	// run dijkstra until queue is empty
	for len(d.Q) > 0 {
		s1 := d.PickVertexFromQ()
		neighbors, edges := d.G.GetNeighborsWithCostsAndEdges(s1, constraint, minimize)
		for s2, cost := range neighbors {
			edge := edges[s2]
			d.UpdateDistances(s1, s2, edge, cost)
		}
	}

	// arrange return variables
	path := raph.GetPath(from, to, d.PredsV, d.PredsE)
	cost := d.GetCost(to)
	if from == to {
		cost = 0
	} else if len(path) == 0 || cost == raph.MaxCost {
		cost = -1
	}

	return path, cost
}

func main() {
	// create & populate graph
	g := raph.NewGraph()
	A := raph.NewVertex("A", "place")
	B := raph.NewVertex("B", "place")
	C := raph.NewEdge("C", "route", "A", "B")
	C.SetCost("cost", 1)
	g.AddVertex(A)
	g.AddVertex(B)
	g.AddEdge(C)

	// init customized dijkstra
	d := NewMyDijkstra(*g)

	constraint := raph.NewConstraint("route")

	// call customized method
	path, cost := d.ShortestPath("A", "B", *constraint, "cost")
	fmt.Println(path, cost)
	// => [A C B] 1
}
