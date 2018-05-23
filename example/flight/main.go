package main

import (
	"fmt"

	"github.com/shoprunback/go-raph/raph"
)

func main() {
	// create graph
	g := raph.NewGraph()

	// create vertigithub.com/shoprunback/ces
	noProps := map[string]string{}
	A := raph.NewVertex("Paris", noProps)
	B := raph.NewVertex("Amsterdam", noProps)
	C := raph.NewVertex("Beijing", noProps)

	// create edges
	D := raph.NewEdge("P->B", "flight", "Paris", "Beijing", map[string]string{"price": "500", "time": "11", "maxLuggageSize": "S"})
	E := raph.NewEdge("P->A", "flight", "Paris", "Amsterdam", map[string]string{"price": "100", "time": "5", "maxLuggageSize": "L"})
	F := raph.NewEdge("A->B", "flight", "Amsterdam", "Beijing", map[string]string{"price": "300", "time": "10", "maxLuggageSize": "L"})

	// populate graph
	g.AddVertex(A)
	g.AddVertex(B)
	g.AddVertex(C)
	g.AddEdge(D)
	g.AddEdge(E)
	g.AddEdge(F)

	// init dijkstra
	d := raph.NewDijkstra(*g)

	var constraints *raph.Constraints
	var path []string
	var cost int

	// find shortest path between Paris and Beijing, minimizing time
	constraints = raph.NewConstraints("flight")
	path, cost = d.ShortestPath("Paris", "Beijing", "time", *constraints)
	fmt.Println(path, cost)
	// => [Paris P->B Beijing] 11

	// find shortest path between Paris and Beijing, minimizing price
	constraints = raph.NewConstraints("flight")
	path, cost = d.ShortestPath("Paris", "Beijing", "price", *constraints)
	fmt.Println(path, cost)
	// => [Paris P->A Amsterdam A->B Beijing] 400

	// find shortest path between Paris and Beijing accepting M or L luggages, minimizing time
	constraints = raph.NewConstraints("flight")
	constraints.AddEdgeConstraint("maxLuggageSize", "M")
	constraints.AddEdgeConstraint("maxLuggageSize", "L")
	path, cost = d.ShortestPath("Paris", "Beijing", "time", *constraints)
	fmt.Println(path, cost)
	// => [Paris P->A Amsterdam A->B Beijing] 15

	// find shortest path between Paris and Beijing, avoiding flights shorter than 10 hours, minimizing price
	constraints = raph.NewConstraints("flight")
	constraints.AddEdgeConstraint("time", "10")
	path, cost = d.ShortestPath("Paris", "Beijing", "price", *constraints)
	fmt.Println(path, cost)
	// => [Paris P->B Beijing] 500
}
