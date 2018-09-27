package raph

import (
	"encoding/json"
	"log"
)

// Query represents a shortest path query. Option is an optional vertex cost that should be included in the shortest path.
type Query struct {
	From       string      `json:"from"`
	To         string      `json:"to"`
	Constraint *Constraint `json:"constraint"`
	Minimize   []string    `json:"minimize"`
	Option     string      `json:"option"`
}

// NewQuery returns a query instance representing the specified JSON string.
func NewQuery(queryString string) *Query {
	query := Query{Constraint: &Constraint{Vertex: &Component{}, Edge: &Component{}}}
	err := json.Unmarshal([]byte(queryString), &query)
	if err != nil {
		log.Fatalln(err)
	}
	return &query
}

// Run executes and returns the query on the specified graph.
func (q Query) Run(graph Graph) map[string]interface{} {
	var path []map[string]interface{}
	var cost int
	dijkstra := NewDijkstra(graph)

	if q.Option == "" {
		path, cost = dijkstra.ShortestPath(q)
	} else {
		path, cost = dijkstra.ShortestPathOption(q)
	}

	if cost == MaxCost {
		cost = -1
	}

	return map[string]interface{}{"path": path, "cost": cost}
}
