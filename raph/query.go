package raph

import (
	"encoding/json"
	"log"
)

type Query struct {
	From       string      `json:"from"`
	To         string      `json:"to"`
	Constraint *Constraint `json:"constraint"`
	Minimize   []string    `json:"minimize"`
	Option     string      `json:"option"`
}

func NewQuery(queryString string) *Query {
	query := Query{Constraint: &Constraint{Vertex: &Component{}, Edge: &Component{}}}
	err := json.Unmarshal([]byte(queryString), &query)
	if err != nil {
		log.Fatalln(err)
	}
	return &query
}

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

    return map[string]interface{}{ "path": path, "cost": cost }
}
