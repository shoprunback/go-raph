package raph

import (
	"encoding/json"
)

type Edge struct {
	Vertex
	Label string          `json:"label"`
	Froms map[string]bool `json:"-"`
	Tos   map[string]bool `json:"-"`
}

func NewEdge(id, label, from, to string) *Edge {
	froms := map[string]bool{from: true}
	tos := map[string]bool{to: true}
	return NewMultiEdge(id, label, froms, tos)
}

func NewMultiEdge(id, label string, froms, tos map[string]bool) *Edge {
	return &Edge{*NewVertex(id), label, froms, tos}
}

func (e Edge) ToJSON() map[string]interface{} {
	var data map[string]interface{}
	b, _ := json.Marshal(e)
	json.Unmarshal(b, &data)

	froms := []string{}
	tos := []string{}
	for id, ok := range e.Froms {
		if ok {
			froms = append(froms, id)
		}
	}
	for id, ok := range e.Tos {
		if ok {
			tos = append(tos, id)
		}
	}
	data["froms"] = froms
	data["tos"] = tos
    
	return data
}
