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
func (d *MyDijkstra) ShortestPath(from, to, minimize string, constraints raph.Constraints) ([]string, int) {
	// init dijkstra with distance 0 for first vertex
	d.Costs[from] = 0

	// run dijkstra until queue is empty
	for len(d.Q) > 0 {
		s1 := d.PickVertexFromQ()
		edges, weights := d.G.GetNeighborsWithEdgesAndWeights(s1, minimize, constraints)
		for s2, weight := range weights {
			edge := edges[s2]
			d.UpdateDistances(s1, s2, edge, weight)
		}
	}

	// gather path
	path := []string{to}
	current := to
	for d.PredsVertices[current] != "" {
		path = append(path, d.PredsEdges[current])
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

	noProps := map[string]string{}
	A := raph.NewVertex("A", noProps)
	B := raph.NewVertex("B", noProps)
	C := raph.NewEdge("C", "route", "A", "B", map[string]string{"cost": "1"})
	g.AddVertex(A)
	g.AddVertex(B)
	g.AddEdge(C)

	// init customized dijkstra
	d := NewMyDijkstra(*g)

	constraints := raph.NewConstraints("route")

	// call customized method
	path, cost := d.ShortestPath("A", "B", "cost", *constraints)
	fmt.Println(path, cost)
	// => [A C B] 1
}
