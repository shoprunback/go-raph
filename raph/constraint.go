package raph

// Constraint is an instance used to filter out nodes. It inheritates from Vertex structure because it behaves more or less like a Vertex against which we will compare vertices and edges. Its ID is not important.
type Constraint struct {
	Vertex
}

// NewConstraint returns a constraint with specified label.
func NewConstraint(label string) *Constraint {
	return &Constraint{*NewVertex("useless", label)}
}

// Copy returns a copy of the constraint.
func (c Constraint) Copy() *Constraint {
	constraint := NewConstraint(c.Label)
	for prop, values := range c.Props {
		for _, value := range values {
			constraint.AddProp(prop, value)
		}
	}
	for cost, value := range c.Costs {
		constraint.SetCost(cost, value)
	}
	return constraint
}
