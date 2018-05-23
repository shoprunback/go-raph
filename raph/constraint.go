package raph

type Constraints struct {
	EdgeLabel string
	VertexProps map[string][]string
	EdgeProps   map[string][]string
}

func NewConstraints(label string) *Constraints {
	return &Constraints{label, map[string][]string{}, map[string][]string{}}
}

func (c *Constraints) AddVertexConstraint(prop, value string) {
	c.VertexProps[prop] = append(c.VertexProps[prop], value)
}

func (c *Constraints) AddEdgeConstraint(prop, value string) {
	c.EdgeProps[prop] = append(c.EdgeProps[prop], value)
}

func (c Constraints) Copy() *Constraints {
	constraintsCopy := NewConstraints(c.EdgeLabel)
	for prop, values := range c.VertexProps {
		for _, value := range values {
			constraintsCopy.AddVertexConstraint(prop, value)
		}
	}
	for prop, values := range c.EdgeProps {
		for _, value := range values {
			constraintsCopy.AddEdgeConstraint(prop, value)
		}
	}
	return constraintsCopy
}
