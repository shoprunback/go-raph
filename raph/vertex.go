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

func (v *Vertex) Satisfies(constraint Constraint) bool {
	propsToSatisfy := constraint.EdgeProps
	for propToSatisfy, satisfyingValues := range propsToSatisfy {
		// retrieve property values of vertex
		values, ok := v.Props[propToSatisfy]

		// considered satisfied if property does not exist
		if ok {
			if !ContainsOne(values, satisfyingValues) {
				return false
			}
		}
	}

	costsToSatisfy := constraint.MinCosts
	for costToSatisfy, minCost := range costsToSatisfy {
		value, ok := v.Costs[costToSatisfy]

		if ok {
			// if value exists, value should be greater or equal than constraint
			if !(value >= minCost) {
				return false
			}
		}
	}

	// all constraints are satisfied
	return true
}
