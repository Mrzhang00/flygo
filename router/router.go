package router

import (
	"errors"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"net/http"
	"strings"
)

//Define Router struct
type Router struct {
	Simples  []*Simple
	Dynamics []*Dynamic
}

//NewRouter
func NewRouter() *Router {
	return &Router{
		Simples:  make([]*Simple, 0),
		Dynamics: make([]*Dynamic, 0),
	}
}

//Route Handler for all method
func (r *Router) REQUEST(pattern string, handler func(c *c.Context)) *Router {
	return r.Route("*", pattern, handler)
}

//Route Handler for GET method
func (r *Router) GET(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodGet, pattern, handler)
}

//Route Handler for POST method
func (r *Router) POST(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodPost, pattern, handler)
}

//Route Handler for PUT method
func (r *Router) PUT(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodPut, pattern, handler)
}

//Route Handler for DELETE method
func (r *Router) DELETE(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodDelete, pattern, handler)
}

//Route Handler for PATCH method
func (r *Router) PATCH(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodPatch, pattern, handler)
}

//Route Handler for HEAD method
func (r *Router) HEAD(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodHead, pattern, handler)
}

//Route Handler for OPTIONS method
func (r *Router) OPTIONS(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodOptions, pattern, handler)
}

//Route Handler for PATCH method
func (r *Router) Route(method, pattern string, handler func(c *c.Context)) *Router {
	//first check method supported?
	if !util.RouteSupport(method) {
		panic(errors.New("method not supported : " + method))
	}

	//check route type
	//1.simple only static route : /to/path
	//2.dynamic contains some variables in route pattern : /to/path/{name}
	pattern = util.TrimSpecialChars(pattern)

	//check route is simple?
	if r.isSimpleRoute(pattern) {
		r.simpleRoute(method, pattern, handler)
	} else {
		r.dynamicRoute(method, pattern, handler)
	}

	return r
}

//_simpleRoute
func _simpleRoute(r *Router, method, pattern string, handler func(c *c.Context)) {
	r.Simples = append(r.Simples, &Simple{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	})
}

//simpleRoute
func (r *Router) simpleRoute(method, pattern string, handler func(c *c.Context)) *Router {
	//map[string]map[string]*simple
	methods := []string{method}
	if method == "*" {
		//Route all
		methods = util.AllMethods()
	}
	for _, m := range methods {
		_simpleRoute(r, m, pattern, handler)
	}
	return r
}

//_dynamicRoute
func _dynamicRoute(r *Router, method, pattern string, pos map[int]string, handler func(c *c.Context)) {
	r.Dynamics = append(r.Dynamics, &Dynamic{
		Pos: pos,
		Simple: &Simple{
			Method:  method,
			Pattern: pattern,
			Handler: handler,
		},
	})
}

//dynamicRoute
func (r *Router) dynamicRoute(method, pattern string, handler func(c *c.Context)) *Router {
	//map[string]map[string]*dynamic
	patterns := strings.Split(pattern, "/")
	pos := make(map[int]string, 0)
	newPatterns := make([]string, len(patterns))
	for i, p := range patterns {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			//exclude start '{' & end '}'
			paramName := p[1 : len(p)-1]
			newPatterns[i] = "*"
			//Add param pos
			pos[i] = paramName
		} else {
			newPatterns[i] = p
		}
	}
	newPattern := strings.Join(newPatterns, "/")
	methods := []string{method}
	if method == "*" {
		methods = util.AllMethods()
	}
	for _, m := range methods {
		_dynamicRoute(r, m, newPattern, pos, handler)
	}
	return r
}

//isSimpleRoute
func (r *Router) isSimpleRoute(pattern string) bool {
	ps := strings.Split(pattern, "/")
	for _, p := range ps {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			return false
		}
	}
	return true
}
