package raph

type Edge struct {
	ID    string
	Props map[string][]string
	Costs map[string]int
	Label string
	Froms map[string]bool
	Tos   map[string]bool
}

func NewEdge(id, label, from, to string) *Edge {
	froms := map[string]bool{from: true}
	tos := map[string]bool{to: true}
	return &Edge{id, map[string][]string{}, map[string]int{}, label, froms, tos}
}

func NewMultiEdge(id, label string, froms, tos map[string]bool) *Edge {
	return &Edge{id, map[string][]string{}, map[string]int{}, label, froms, tos}
}

func (e *Edge) AddProp(prop, value string) {
	e.Props[prop] = append(e.Props[prop], value)
}

func (e *Edge) SetCost(cost string, value int) {
	e.Costs[cost] = value
}

func (e *Edge) Satisfies(constraint Constraint) bool {
	propsToSatisfy := constraint.EdgeProps
	for propToSatisfy, satisfyingValues := range propsToSatisfy {
		// retrieve property values of vertex
		values, ok := e.Props[propToSatisfy]

		// considered satisfied if property does not exist
		if ok {
			if !ContainsOne(values, satisfyingValues) {
				return false
			}
		}
	}

	costsToSatisfy := constraint.MinCosts
	for costToSatisfy, minCost := range costsToSatisfy {
		value, ok := e.Costs[costToSatisfy]

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
