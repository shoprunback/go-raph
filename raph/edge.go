package raph

import (
    "encoding/json"
)

// Edge represents an edge instance. It inherits from Vertex structure. Froms & Tos fields are removed from JSON Marshaling.
type Edge struct {
	Vertex
	Label string          `json:"label"`
	Froms map[string]bool `json:"-"` // list of vertices from which the edge is reachable
	Tos   map[string]bool `json:"-"` // list of vertices that the edge can reach
}

// NewEdge returns a new edge.
func NewEdge(id, label, from, to string) *Edge {
	froms := map[string]bool{from: true}
	tos := map[string]bool{to: true}
	return NewMultiEdge(id, label, froms, tos)
}

// NewMultiEdge returns a new multiedge.
func NewMultiEdge(id, label string, froms, tos map[string]bool) *Edge {
	return &Edge{*NewVertex(id), label, froms, tos}
}

// ToJSON formats the edge to JSON.
func (e Edge) ToJSON() map[string]interface{} {
	var data map[string]interface{}
	b, _ := json.Marshal(e)
	json.Unmarshal(b, &data)
	return data
}
