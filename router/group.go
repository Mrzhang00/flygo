package router

type Group struct {
	prefix  string
	routers []*Router
}

func NewGroup() *Group {
	return NewGroupWithPrefix("")
}

func NewGroupWithPrefix(prefix string) *Group {
	return &Group{
		prefix:  prefix,
		routers: make([]*Router, 0),
	}
}

func (g *Group) Add(r ...*Router) *Group {
	g.routers = append(g.routers, r...)
	return g
}

func (g *Group) Prefix() string {
	return g.prefix
}

func (g *Group) Routers() []*Router {
	return g.routers
}
