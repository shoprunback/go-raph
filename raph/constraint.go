package raph

// Constraint is an instance used to filter out nodes. It inherits from Vertex structure because it behaves more or less like a Vertex against which we will compare vertices and edges. Its ID is not important.
type Constraint struct {
	Vertex *Component
	Edge   *Component
}

// NewConstraint returns a constraint with specified label.
func NewConstraint(vertexLabel, edgeLabel string) *Constraint {
	vertex := NewComponent(vertexLabel)
	edge := NewComponent(edgeLabel)
	return &Constraint{vertex, edge}
}

// Copy returns a copy of the constraint.
func (c Constraint) Copy() *Constraint {
	vertex := c.Vertex.Copy()
	edge := c.Edge.Copy()
	return &Constraint{vertex, edge}
}
