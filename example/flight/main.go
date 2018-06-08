package main

import (
	"fmt"

	"github.com/shoprunback/go-raph/raph"
)

func main() {
	// create graph
	g := raph.NewGraph()

	// create vertices
	A := raph.NewVertex("Paris", "city")
	B := raph.NewVertex("Amsterdam", "city")
	C := raph.NewVertex("Beijing", "city")

	// create edges
	D := raph.NewEdge("P->B", "flight", "Paris", "Beijing")
	D.AddProp("maxLuggageSize", "S")
	D.SetCost("price", 500)
	D.SetCost("time", 11)
	E := raph.NewEdge("P->A", "flight", "Paris", "Amsterdam")
	E.AddProp("maxLuggageSize", "L")
	E.SetCost("price", 100)
	E.SetCost("time", 5)
	F := raph.NewEdge("A->B", "flight", "Amsterdam", "Beijing")
	F.AddProp("maxLuggageSize", "L")
	F.SetCost("price", 300)
	F.SetCost("time", 10)

	// populate graph
	g.AddVertex(A)
	g.AddVertex(B)
	g.AddVertex(C)
	g.AddEdge(D)
	g.AddEdge(E)
	g.AddEdge(F)

	// init dijkstra
	d := raph.NewDijkstra(*g)

	var constraint *raph.Constraint
	var path []string
	var cost int

	// find shortest path between Paris and Beijing, minimizing time
	constraint = raph.NewConstraint("city", "flight")
	path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "time")
	fmt.Println(path, cost)
	// => [Paris P->B Beijing] 11

	// find shortest path between Paris and Beijing, minimizing price
	constraint = raph.NewConstraint("city", "flight")
	path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "price")
	fmt.Println(path, cost)
	// => [Paris P->A Amsterdam A->B Beijing] 400

	// find shortest path between Paris and Beijing accepting M or L luggages, minimizing time
	constraint = raph.NewConstraint("city", "flight")
	constraint.Edge.AddProp("maxLuggageSize", "M", "L")
	path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "time")
	fmt.Println(path, cost)
	// => [Paris P->A Amsterdam A->B Beijing] 15

	// find shortest path between Paris and Beijing, avoiding flights shorter than 10 hours, minimizing price
	constraint = raph.NewConstraint("city", "flight")
	constraint.Edge.SetCost("time", 10)
	path, cost = d.ShortestPath("Paris", "Beijing", *constraint, "price")
	fmt.Println(path, cost)
	// => [Paris P->B Beijing] 500
}
