package raph

type Vertex struct {
	ID    string
	Props map[string][]string
	Costs map[string]int
}

func NewVertex(id string) *Vertex {
	return &Vertex{id, map[string][]string{}, map[string]int{}}
}

func (v *Vertex) AddProp(prop, value string) {
	v.Props[prop] = append(v.Props[prop], value)
}

func (v *Vertex) SetCost(cost string, value int) {
	v.Costs[cost] = value
}

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
