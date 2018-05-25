package raph

type Constraint struct {
	Label string
	EdgeProps   map[string][]string
	VertexProps map[string][]string
	MinCosts   map[string]int
}

func NewConstraint(label string) *Constraint {
	return &Constraint{label, map[string][]string{}, map[string][]string{}, map[string]int{}}
}

func (c *Constraint) AddVertexConstraint(prop, value string) {
	c.VertexProps[prop] = append(c.VertexProps[prop], value)
}

func (c *Constraint) AddEdgeConstraint(prop, value string) {
	c.EdgeProps[prop] = append(c.EdgeProps[prop], value)
}

func (c *Constraint) SetMinCostConstraint(prop string, value int) {
	c.MinCosts[prop] = value
}

func (c Constraint) Copy() *Constraint {
	constraintCopy := NewConstraint(c.Label)
	for prop, values := range c.VertexProps {
		for _, value := range values {
			constraintCopy.AddVertexConstraint(prop, value)
		}
	}
	for prop, values := range c.EdgeProps {
		for _, value := range values {
			constraintCopy.AddEdgeConstraint(prop, value)
		}
	}
	return constraintCopy
}
