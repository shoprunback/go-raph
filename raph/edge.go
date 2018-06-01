package raph

import (
	"encoding/json"
)

// Edge represents an edge instance. It inheritates from Vertex structure. Froms & Tos fields are removed from JSON Marshaling to be added manually as slices (cf ToJSON).
type Edge struct {
	Vertex
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
	return &Edge{*NewVertex(id, label), froms, tos}
}

// ToJSON formats the vertex to JSON. Froms and Tos fields are converted to slice.
func (e Edge) ToJSON() map[string]interface{} {
	var data map[string]interface{}
	b, _ := json.Marshal(e)
	json.Unmarshal(b, &data)

	// convert froms/tos from map to slice
	froms := []string{}
	tos := []string{}
	for id, present := range e.Froms {
		if present {
			froms = append(froms, id)
		}
	}
	for id, present := range e.Tos {
		if present {
			tos = append(tos, id)
		}
	}
	data["froms"] = froms
	data["tos"] = tos

	return data
}
