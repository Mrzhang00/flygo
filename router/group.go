package router

//Define Group struct
type Group struct {
	prefix  string //router prefix
	routers []*Router
}

//NewRouterGroup
func NewGroup() *Group {
	return NewGroupWithPrefix("")
}

//NewGroupWithPrefix
func NewGroupWithPrefix(prefix string) *Group {
	return &Group{
		prefix:  prefix,
		routers: make([]*Router, 0),
	}
}

//Add
func (g *Group) Add(r ...*Router) *Group {
	g.routers = append(g.routers, r...)
	return g
}

//Prefix
func (g *Group) Prefix() string {
	return g.prefix
}

//Routers
func (g *Group) Routers() []*Router {
	return g.routers
}
