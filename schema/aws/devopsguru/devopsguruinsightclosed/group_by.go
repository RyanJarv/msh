package devopsguruinsightclosed

type GroupBy struct {
    Dimensions []string `json:"dimensions,omitempty"`
    Group string `json:"group,omitempty"`
}

func (g *GroupBy) SetDimensions(dimensions []string) {
    g.Dimensions = dimensions
}

func (g *GroupBy) SetGroup(group string) {
    g.Group = group
}
