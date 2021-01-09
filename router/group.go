package router

// Group struct
type Group struct {
	prefix  string
	routers []*Router
}

// NewGroup return new group
func NewGroup() *Group {
	return NewGroupWithPrefix("")
}

// NewGroupWithPrefix return new group
func NewGroupWithPrefix(prefix string) *Group {
	return &Group{
		prefix:  prefix,
		routers: make([]*Router, 0),
	}
}

// Add routers
func (g *Group) Add(r ...*Router) *Group {
	g.routers = append(g.routers, r...)
	return g
}

// Prefix return group's prefix
func (g *Group) Prefix() string {
	return g.prefix
}

// Routers return group's routers
func (g *Group) Routers() []*Router {
	return g.routers
}
