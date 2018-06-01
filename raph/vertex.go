package raph

import (
	"encoding/json"
)

// Vertex represents a vertex instance.
type Vertex struct {
	ID    string              `json:"id"`
	Label string              `json:"label"`
	Props map[string][]string `json:"props"`
	Costs map[string]int      `json:"costs"`
}

// NewVertex returns a new vertex.
func NewVertex(id, label string) *Vertex {
	return &Vertex{id, label, map[string][]string{}, map[string]int{}}
}

// AddProp adds a list of property values to the vertex.
func (v *Vertex) AddProp(prop, values ...string) {
	v.Props[prop] = append(v.Props[prop], values...)
}

// SetCost sets a cost value for the vertex.
func (v *Vertex) SetCost(cost string, value int) {
	v.Costs[cost] = value
}

// Satisfies returns whether or not the vertex satisfies the props and costs (threshold) of the constraint.
func (v Vertex) Satisfies(constraint Constraint) bool {
	props := constraint.Props
	for prop, constraintValues := range props {
		// retrieve property values of vertex
		values, ok := v.Props[prop]

		// considered satisfied if property does not exist
		if ok {
			if !ContainsOne(values, constraintValues) {
				return false
			}
		}
	}

	costs := constraint.Costs
	for cost, threshold := range costs {
		value, ok := v.Costs[cost]
		if ok {
			// cost should be greater or equal than threshold
			if !(value >= threshold) {
				return false
			}
		}
	}

	// all constraints are satisfied
	return true
}

// ToJSON formats the vertex to JSON.
func (v Vertex) ToJSON() map[string]interface{} {
	var data map[string]interface{}
	b, _ := json.Marshal(v)
	json.Unmarshal(b, &data)
	return data
}
