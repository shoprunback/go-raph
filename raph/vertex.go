package raph

import (
	"encoding/json"
)

// Vertex represents a vertex instance.
type Vertex struct {
	ID string `json:"id"`
	Label string          `json:"label"`
	Component
}

// NewVertex returns a new vertex.
func NewVertex(id, label string) *Vertex {
	return &Vertex{id, label, *NewComponent()}
}

// ToJSON formats the vertex to JSON.
func (v Vertex) ToJSON() map[string]interface{} {
	var data map[string]interface{}
	b, _ := json.Marshal(v)
	json.Unmarshal(b, &data)
	return data
}
