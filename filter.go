package flygo

import (
	"fmt"
	"net/http"
	"strings"
)

//Define FilterContext struct
type FilterContext struct {
	RequestURI       string
	ParsedRequestURI string
	RequestMethod    string
	RequestHeader    http.Header
	ResponseHeader   *http.Header
	Parameters       map[string][]string
	context          *Context
}

//Define filter route
type filterRoute struct {
	pattern       string
	regex         string
	filterHandler FilterHandler
}

//Define filter route chain
type filterRouteChain struct {
	regex string
	route filterRoute
}

type filterRouteCache struct {
	t string
	filterRoute
}

//Define filter handler
type FilterHandler func(*FilterContext)

//Route filter
func (a *App) filter(filterType, pattern string, filterHandler FilterHandler) {
	if pattern == "" {
		return
	}
	a.filterRouteCaches = append(a.filterRouteCaches, filterRouteCache{
		t: filterType,
		filterRoute: filterRoute{
			pattern:       pattern,
			filterHandler: filterHandler,
		},
	})
}

func (a *App) startFilter() {
	for _, froute := range a.filterRouteCaches {
		filterType := froute.t
		pattern := froute.pattern
		filterHandler := froute.filterHandler
		contextPath := a.Config.Flygo.Server.ContextPath
		regex := fmt.Sprintf(`^%s%s$`, contextPath, strings.ReplaceAll(trim(pattern), "*", "[/a-zA-Z0-9]+"))
		fr := filterRoute{
			pattern:       pattern,
			regex:         regex,
			filterHandler: filterHandler,
		}
		if filterType == "before" {
			a.beforeFilters[regex] = filterRouteChain{regex: regex, route: fr}
		} else {
			a.afterFilters[regex] = filterRouteChain{regex: regex, route: fr}
		}
	}
}

//Match before filter
func (c *Context) matchBeforeFilter() []FilterHandler {
	return c.matchFilter("before")
}

//Match after filter
func (c *Context) matchAfterFilter() []FilterHandler {
	return c.matchFilter("after")
}

//Match base method
func (c *Context) matchFilter(filterType string) []FilterHandler {
	var filterChains map[string]filterRouteChain
	switch filterType {
	case "before":
		filterChains = c.app.beforeFilters
		break
	case "after":
		filterChains = c.app.afterFilters
		break
	}
	if filterChains == nil {
		return nil
	}
	filters := make([]FilterHandler, 0)
	for regex, chain := range filterChains {
		if c.matchPath(regex) {
			filters = append(filters, chain.route.filterHandler)
		}
	}
	return filters
}

//Invoke before filters
func (c *Context) invokeBeforeFilter() {
	beforeFilters := c.matchBeforeFilter()
	if beforeFilters != nil && len(beforeFilters) > 0 {
		for _, filter := range beforeFilters {
			c.invokeFilterHandler(filter)
		}
	}
}

//Invoke after filter
func (c *Context) invokeAfterFilter() {
	afterFilters := c.matchAfterFilter()
	if afterFilters != nil && len(afterFilters) > 0 {
		for _, filter := range afterFilters {
			c.invokeFilterHandler(filter)
		}
	}
}

//Invoke filter handler
func (c *Context) invokeFilterHandler(handler FilterHandler) {
	if handler == nil {
		return
	}
	handler(&FilterContext{
		RequestURI:       c.RequestURI,
		ParsedRequestURI: c.ParsedRequestURI,
		RequestMethod:    c.RequestMethod,
		RequestHeader:    c.RequestHeader,
		ResponseHeader:   c.ResponseHeader,
		Parameters:       c.ParamMap,
		context:          c,
	})
}

//Route before filter
func (a *App) BeforeFilter(pattern string, filterHandler FilterHandler) *App {
	a.filter("before", pattern, filterHandler)
	return a
}

//Route after filter
func (a *App) AfterFilter(pattern string, filterHandler FilterHandler) *App {
	a.filter("after", pattern, filterHandler)
	return a
}

//Print filter
func (a *App) printFilter() {
	for _, filter := range a.beforeFilters {
		a.Logger.Info("Before filter route [%s]", filter.route.pattern)
	}
	for _, filter := range a.afterFilters {
		a.Logger.Info("After filter route [%s]", filter.route.pattern)
	}
}

//Get bundle app
func (c *FilterContext) App() *App {
	return c.context.App()
}

//Get middleware ctx
func (c *FilterContext) GetMiddlewareCtx(name string) map[string]interface{} {
	return c.context.GetMiddlewareCtx(name)
}

//Clear middleware c
func (c *FilterContext) ClearMiddlewareCtx(name string) {
	c.context.ClearMiddlewareCtx(name)
}

//Remove middleware data
func (c *FilterContext) RemoveMiddlewareData(name, key string) {
	c.context.RemoveMiddlewareData(name, key)
}

//Set middleware data
func (c *FilterContext) SetMiddlewareData(name, key string, val interface{}) {
	c.context.SetMiddlewareData(name, key, val)
}

//Get middleware data
func (c *FilterContext) GetMiddlewareData(name, key string) interface{} {
	return c.context.GetMiddlewareData(name, key)
}
