package raph

import (
	"encoding/json"
	"log"
	"math"
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
	// origin and destination are equal
	if q.From == q.To {
		return map[string]interface{}{"path": []string{}, "cost": 0}
	}

	// origin or destination do not exist in the graph
	_, okFrom := graph.Vertices[q.From]
	_, okTo := graph.Vertices[q.To]
	if !okFrom || !okTo {
		return map[string]interface{}{"path": []string{}, "cost": -1}
	}

	var path []map[string]interface{}
	var cost float64
	dijkstra := NewDijkstra(graph)

	if q.Option == "" {
		path, cost = dijkstra.ShortestPath(q)
	} else {
		path, cost = dijkstra.ShortestPathOption(q)
	}

	if cost == math.Inf(0) {
		cost = -1
	}

	return map[string]interface{}{"path": path, "cost": cost}
}
