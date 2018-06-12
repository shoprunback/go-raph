package raph

// Constraint is an instance used to filter out nodes. It inherits from Vertex structure because it behaves more or less like a Vertex against which we will compare vertices and edges. Its ID is not important.
type Constraint struct {
	Vertex *Component `json:"vertex"`
	Edge   *Component `json:"edge"`
	Label  string     `json:"label"`
}

// NewConstraint returns a constraint with specified label.
func NewConstraint(label string) *Constraint {
	vertex := NewComponent()
	edge := NewComponent()
	return &Constraint{vertex, edge, label}
}

// Copy returns a copy of the constraint.
func (c Constraint) Copy() *Constraint {
	vertex := c.Vertex.Copy()
	edge := c.Edge.Copy()
	return &Constraint{vertex, edge, c.Label}
}
