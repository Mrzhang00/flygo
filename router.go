package flygo

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	rr "github.com/billcoding/flygo/router"
	"net/http"
	"strings"
)

// REQUEST Route all Methods
func (a *App) REQUEST(pattern string, handler func(c *c.Context)) *App {
	return a.Route("*", pattern, handler)
}

// GET Route Get Method
func (a *App) GET(pattern string, handler func(c *c.Context)) *App {
	return a.Route(http.MethodGet, pattern, handler)
}

// POST Route POST Method
func (a *App) POST(pattern string, handler func(c *c.Context)) *App {
	return a.Route(http.MethodPost, pattern, handler)
}

// PUT Route PUT Method
func (a *App) PUT(pattern string, handler func(c *c.Context)) *App {
	return a.Route(http.MethodPut, pattern, handler)
}

// DELETE Route DELETE Method
func (a *App) DELETE(pattern string, handler func(c *c.Context)) *App {
	return a.Route(http.MethodDelete, pattern, handler)
}

// PATCH Route PATCH Method
func (a *App) PATCH(pattern string, handler func(c *c.Context)) *App {
	return a.Route(http.MethodPatch, pattern, handler)
}

// HEAD Route HEAD Method
func (a *App) HEAD(pattern string, handler func(c *c.Context)) *App {
	return a.Route(http.MethodHead, pattern, handler)
}

// OPTIONS Route OPTIONS Method
func (a *App) OPTIONS(pattern string, handler func(c *c.Context)) *App {
	return a.Route(http.MethodOptions, pattern, handler)
}

// Route Route DIY Method
func (a *App) Route(method, pattern string, handler func(c *c.Context)) *App {
	a.routers[0].Route(method, pattern, handler)
	return a
}

// AddRouter Add Routers
func (a *App) AddRouter(r ...*rr.Router) *App {
	a.routers = append(a.routers, r...)
	return a
}

// AddRouterGroup Add Group Routers
func (a *App) AddRouterGroup(g ...*rr.Group) *App {
	a.groups = append(a.groups, g...)
	return a
}

func (a *App) parseRouters() *App {

	simpleParseFunc := func(prefix string, simples []*rr.Simple) {
		for _, simple := range simples {
			routeKey := fmt.Sprintf("%s:%s%s", simple.Method, prefix, simple.Pattern)
			a.parsedRouters.Simples[routeKey] = simple
		}
	}

	dynamicParseFunc := func(prefix string, dynamics []*rr.Dynamic) {
		for _, dynamic := range dynamics {
			pattern := fmt.Sprintf("^%s%s$", prefix, strings.ReplaceAll(dynamic.Pattern, "*", `([\w-]+)`))
			dynamicsMap, have := a.parsedRouters.Dynamics[pattern]

			if prefix != "" {
				pslen := len(strings.Split(prefix, "/"))
				if pslen > 1 {
					for k, v := range dynamic.Pos {
						delete(dynamic.Pos, k)
						dynamic.Pos[k+pslen-1] = v
					}
				}
			}
			if have {
				dynamicsMap[dynamic.Method] = dynamic
			} else {
				a.parsedRouters.Dynamics[pattern] = map[string]*rr.Dynamic{dynamic.Method: dynamic}
			}
		}
	}

	for _, g := range a.groups {
		for _, gr := range g.Routers() {
			if g.Prefix() == "" {
				a.routers = append(a.routers, gr)
			} else {
				simpleParseFunc(g.Prefix(), gr.Simples)
				dynamicParseFunc(g.Prefix(), gr.Dynamics)
			}
		}
	}

	for _, r := range a.routers {
		simpleParseFunc("", r.Simples)
		dynamicParseFunc("", r.Dynamics)
	}
	return a
}
