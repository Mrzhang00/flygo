package router

import (
	"errors"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"net/http"
	"strings"
)

// Router struct
type Router struct {
	Simples  []*Simple
	Dynamics []*Dynamic
}

// NewRouter return new router
func NewRouter() *Router {
	return &Router{
		Simples:  make([]*Simple, 0),
		Dynamics: make([]*Dynamic, 0),
	}
}

// REQUEST Route all Methods
func (r *Router) REQUEST(pattern string, handler func(c *c.Context)) *Router {
	return r.Route("*", pattern, handler)
}

// GET Route Get Method
func (r *Router) GET(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodGet, pattern, handler)
}

// POST Route POST Method
func (r *Router) POST(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodPost, pattern, handler)
}

// PUT Route PUT Method
func (r *Router) PUT(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodPut, pattern, handler)
}

// DELETE Route DELETE Method
func (r *Router) DELETE(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodDelete, pattern, handler)
}

// PATCH Route PATCH Method
func (r *Router) PATCH(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodPatch, pattern, handler)
}

// HEAD Route HEAD Method
func (r *Router) HEAD(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodHead, pattern, handler)
}

// OPTIONS Route OPTIONS Method
func (r *Router) OPTIONS(pattern string, handler func(c *c.Context)) *Router {
	return r.Route(http.MethodOptions, pattern, handler)
}

// Route Route DIY Method
func (r *Router) Route(method, pattern string, handler func(c *c.Context)) *Router {

	if !util.RouteSupport(method) {
		panic(errors.New("method not supported : " + method))
	}

	pattern = util.TrimSpecialChars(pattern)

	if r.isSimpleRoute(pattern) {
		r.simpleRoute(method, pattern, handler)
	} else {
		r.dynamicRoute(method, pattern, handler)
	}

	return r
}

func _simpleRoute(r *Router, method, pattern string, handler func(c *c.Context)) {
	r.Simples = append(r.Simples, &Simple{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	})
}

func (r *Router) simpleRoute(method, pattern string, handler func(c *c.Context)) *Router {

	methods := []string{method}
	if method == "*" {

		methods = util.AllMethods()
	}
	for _, m := range methods {
		_simpleRoute(r, m, pattern, handler)
	}
	return r
}

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

func (r *Router) dynamicRoute(method, pattern string, handler func(c *c.Context)) *Router {

	patterns := strings.Split(pattern, "/")
	pos := make(map[int]string, 0)
	newPatterns := make([]string, len(patterns))
	for i, p := range patterns {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {

			paramName := p[1 : len(p)-1]
			newPatterns[i] = "*"

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

func (r *Router) isSimpleRoute(pattern string) bool {
	ps := strings.Split(pattern, "/")
	for _, p := range ps {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			return false
		}
	}
	return true
}
