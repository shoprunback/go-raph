package raph

type Constraint struct {
	Label       string
	Props       map[string][]string
	Costs       map[string]int
}

func NewConstraint(label string) *Constraint {
	return &Constraint{label, map[string][]string{}, map[string]int{}}
}

func (c *Constraint) AddProp(prop, value string) {
	c.Props[prop] = append(c.Props[prop], value)
}

func (c *Constraint) SetCost(prop string, value int) {
	c.Costs[prop] = value
}

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
