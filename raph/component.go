package raph

import (
	"strings"
)

// Component represents an instance that can have properties and costs.
type Component struct {
	Props map[string][]string `json:"props"`
	Costs map[string]float64  `json:"costs"`
}

// NewEdge returns a new component.
func NewComponent() *Component {
	return &Component{map[string][]string{}, map[string]float64{}}
}

// AddProp adds a property values to the component.
func (c *Component) AddProp(prop string, values ...string) {
	c.Props[prop] = append(c.Props[prop], values...)
}

// SetCost sets a cost value for the component.
func (c *Component) SetCost(cost string, value float64) {
	c.Costs[cost] = value
}

// Copy returns a copy of the component.
func (c Component) Copy() *Component {
	component := NewComponent()
	for prop, values := range c.Props {
		component.AddProp(prop, values...)
	}
	for cost, value := range c.Costs {
		component.SetCost(cost, value)
	}
	return component
}

// Satisfies returns whether or not the component satisfies the props and costs (threshold) of the specified component.
func (c Component) Satisfies(component Component) bool {
	// check props
	for prop, satisfyingValues := range component.Props {
		// negation is the case where a property should not contain any of satisfying values
		negation := strings.HasPrefix(prop, "~")
		if negation {
			prop = prop[1:]
		}

		// retrieve property values of component
		values, ok := c.Props[prop]

		// considered satisfied if property does not exist
		if ok {
			// if !negation, intersection of values and satisfying values should not be empty
			// if negation, intersection of values and satisfying values should be empty
			if (!negation && !ContainsOne(values, satisfyingValues)) || (negation && ContainsOne(values, satisfyingValues)) {
				return false
			}
		}
	}

	// check costs
	for cost, threshold := range component.Costs {
		value, ok := c.Costs[cost]

		// considered satisfied if cost does not exist
		if ok {
			// cost should be greater or equal than specified threshold
			if !(value >= threshold) {
				return false
			}
		}
	}

	// all constraints are satisfied
	return true
}
