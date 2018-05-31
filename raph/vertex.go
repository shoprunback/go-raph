package raph

import (
	"encoding/json"
)

type Vertex struct {
	ID    string            `json:"id"`
	Label string            `json:"label"`
	Props map[string]string `json:"props"`
	Costs map[string]int    `json:"costs"`
}

func NewVertex(id, label string) *Vertex {
	return &Vertex{id, label,  map[string]string{}, map[string]int{}}
}

func (v *Vertex) SetProp(prop, value string) {
	v.Props[prop] = value
}

func (v *Vertex) SetCost(cost string, value int) {
	v.Costs[cost] = value
}

func (v Vertex) Satisfies(constraint Constraint) bool {
	props := constraint.Props
	for prop, constraintValues := range props {
		// retrieve property values of vertex
		value, ok := v.Props[prop]

		// considered satisfied if property does not exist
		if ok {
			if !Contains(constraintValues, value) {
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

func (v Vertex) ToJSON() map[string]interface{} {
	var data map[string]interface{}
	b, _ := json.Marshal(v)
	json.Unmarshal(b, &data)
	return data
}
