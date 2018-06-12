package raph

type Component struct {
	Props map[string][]string `json:"props"`
	Costs map[string]int      `json:"costs"`
}

func NewComponent() *Component {
	return &Component{map[string][]string{}, map[string]int{}}
}

// AddProp adds a list of property values to the vertex.
func (c *Component) AddProp(prop string, values ...string) {
	c.Props[prop] = append(c.Props[prop], values...)
}

// SetCost sets a cost value for the vertex.
func (c *Component) SetCost(cost string, value int) {
	c.Costs[cost] = value
}

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
		// retrieve property values of component
		values, ok := c.Props[prop]

		// considered satisfied if property does not exist
		if ok {
			if !ContainsOne(values, satisfyingValues) {
				return false
			}
		}
	}

	// check costs
	for cost, threshold := range component.Costs {
		value, ok := c.Costs[cost]
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
