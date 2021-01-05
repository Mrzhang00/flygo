package flygo

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	. "github.com/billcoding/flygo/router"
	"net/http"
	"strings"
)

type Handler func(c *c.Context)

//Route Handler for all method
func (a *App) REQUEST(pattern string, handler Handler) *App {
	return a.Route("*", pattern, handler)
}

//Route Handler for GET method
func (a *App) GET(pattern string, handler Handler) *App {
	return a.Route(http.MethodGet, pattern, handler)
}

//Route Handler for POST method
func (a *App) POST(pattern string, handler Handler) *App {
	return a.Route(http.MethodPost, pattern, handler)
}

//Route Handler for PUT method
func (a *App) PUT(pattern string, handler Handler) *App {
	return a.Route(http.MethodPut, pattern, handler)
}

//Route Handler for DELETE method
func (a *App) DELETE(pattern string, handler Handler) *App {
	return a.Route(http.MethodDelete, pattern, handler)
}

//Route Handler for PATCH method
func (a *App) PATCH(pattern string, handler Handler) *App {
	return a.Route(http.MethodPatch, pattern, handler)
}

//Route Handler for HEAD method
func (a *App) HEAD(pattern string, handler Handler) *App {
	return a.Route(http.MethodHead, pattern, handler)
}

//Route Handler for OPTIONS method
func (a *App) OPTIONS(pattern string, handler Handler) *App {
	return a.Route(http.MethodOptions, pattern, handler)
}

//Route Handler for PATCH method
func (a *App) Route(method, pattern string, handler Handler) *App {
	a.routers[0].Route(method, pattern, handler)
	return a
}

//Add router
func (a *App) AddRouter(r ...*Router) *App {
	a.routers = append(a.routers, r...)
	return a
}

//Add router group
func (a *App) AddRouterGroup(g ...*Group) *App {
	a.groups = append(a.groups, g...)
	return a
}

//parseRouters
func (a *App) parseRouters() *App {

	simpleParseFunc := func(prefix string, simples []*Simple) {
		for _, simple := range simples {
			routeKey := fmt.Sprintf("%s:%s%s", simple.Method, prefix, simple.Pattern)
			a.parsedRouters.Simples[routeKey] = simple
		}
	}

	dynamicParseFunc := func(prefix string, dynamics []*Dynamic) {
		for _, dynamic := range dynamics {
			pattern := fmt.Sprintf("^%s%s$", prefix, strings.ReplaceAll(dynamic.Pattern, "*", `([\w-]+)`))
			dynamicsMap, have := a.parsedRouters.Dynamics[pattern]
			//FIX Params Pos
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
				a.parsedRouters.Dynamics[pattern] = map[string]*Dynamic{dynamic.Method: dynamic}
			}
		}
	}

	//start parse groups
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

	//start parse routers
	for _, r := range a.routers {
		simpleParseFunc("", r.Simples)
		dynamicParseFunc("", r.Dynamics)
	}
	return a
}
