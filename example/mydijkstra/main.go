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
func (d *MyDijkstra) ShortestPath(from, to, minimize string, constraint raph.Constraint) ([]string, int) {
	// init dijkstra with distance 0 for first vertex
	d.Costs[from] = 0

	// run dijkstra until queue is empty
	for len(d.Q) > 0 {
		s1 := d.PickVertexFromQ()
		neighbors := d.G.GetNeighborsWithCosts(s1, minimize, constraint)
		for s2, cost := range neighbors {
			d.UpdateDistances(s1, s2, cost)
		}
	}

	// gather path
	path := []string{to}
	current := to
	for d.PredsVertices[current] != "" {
		path = append(path, d.PredsVertices[current])
		current = d.PredsVertices[current]
	}

	// arrange return variables
	raph.Reverse(path)
	cost := d.GetCost(to)
	if cost == raph.MaxCost {
		cost = -1
	}

	// reset dijkstra for further use
	d.Reset()

	if from != to && len(path) == 1 {
		return []string{}, cost
	} else {
		return path, cost
	}
}

func main() {
	// create & populate graph
	g := raph.NewGraph()
	A := raph.NewVertex("A")
	B := raph.NewVertex("B")
	C := raph.NewEdge("C", "route", "A", "B")
	C.SetCost("cost", 1)
	g.AddVertex(A)
	g.AddVertex(B)
	g.AddEdge(C)

	// init customized dijkstra
	d := NewMyDijkstra(*g)

	constraint := raph.NewConstraint("route")

	// call customized method
	path, cost := d.ShortestPath("A", "B", "cost", *constraint)
	fmt.Println(path, cost)
	// => [A B] 1
}
