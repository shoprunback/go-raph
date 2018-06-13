package main

import (
	"fmt"

	"github.com/shoprunback/go-raph/raph"
)

func main() {
	// create graph
	g := raph.NewGraph()

	// create vertices
	A := raph.NewVertex("Paris")
	B := raph.NewVertex("Amsterdam")
	B.SetCost("chewam", 10)
	C := raph.NewVertex("Beijing")

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

	var query *raph.Query
	var res map[string]interface{}
	var ids []string

	// find shortest path between Paris and Beijing, minimizing time
	query = raph.NewQuery(`
		{
			"from": "Paris",
			"to": "Beijing",
			"constraint": {
				"label": "flight"
			},
			"minimize": ["time"],
			"option": "chewam"
		}
	`)
	res = query.Run(*g)
	ids = ExtractIDS(res["path"].([]map[string]interface{}))
	fmt.Println(ids, res["cost"])
	// => [Paris P->B Beijing] 11

	// find shortest path between Paris and Beijing, minimizing price
	query = raph.NewQuery(`
		{
			"from": "Paris",
			"to": "Beijing",
			"constraint": {
				"label": "flight"
			},
			"minimize": ["price"]
		}
	`)
	res = query.Run(*g)
	ids = ExtractIDS(res["path"].([]map[string]interface{}))
	fmt.Println(ids, res["cost"])
	// => [Paris P->A Amsterdam A->B Beijing] 400

	// find shortest path between Paris and Beijing accepting M or L luggages, minimizing time
	query = raph.NewQuery(`
		{
			"from": "Paris",
			"to": "Beijing",
			"constraint": {
				"edge": {
					"props": {
						"maxLuggageSize": ["M", "L"]
					}
				},
				"label": "flight"
			},
			"minimize": ["time"]
		}
	`)
	res = query.Run(*g)
	ids = ExtractIDS(res["path"].([]map[string]interface{}))
	fmt.Println(ids, res["cost"])
	// => [Paris P->A Amsterdam A->B Beijing] 15

	// find shortest path between Paris and Beijing, avoiding flights shorter than 10 hours, minimizing price
	query = raph.NewQuery(`
		{
			"from": "Paris",
			"to": "Beijing",
			"constraint": {
				"edge": {
					"costs": {
						"time": 10
					}
				},
				"label": "flight"
			},
			"minimize": ["price"]
		}
	`)
	res = query.Run(*g)
	ids = ExtractIDS(res["path"].([]map[string]interface{}))
	fmt.Println(ids, res["cost"])
	// => [Paris P->B Beijing] 500
}

func ExtractIDS(path []map[string]interface{}) []string {
	ids := []string{}
	for _, v := range path {
		ids = append(ids, v["id"].(string))
	}
	return ids
}
