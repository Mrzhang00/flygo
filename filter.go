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

//Define filter handler
type FilterHandler func(*FilterContext)

//Route filter
func filter(filterType, pattern string, filterHandler FilterHandler) {
	if pattern == "" {
		return
	}
	regex := fmt.Sprintf(`^%s%s$`, app.GetContextPath(), strings.ReplaceAll(trim(pattern), "*", "[/a-zA-Z0-9]+"))
	fr := filterRoute{
		pattern:       pattern,
		regex:         regex,
		filterHandler: filterHandler,
	}
	if filterType == "before" {
		app.beforeFilters[regex] = filterRouteChain{regex: regex, route: fr}
	} else {
		app.afterFilters[regex] = filterRouteChain{regex: regex, route: fr}
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
		filterChains = app.beforeFilters
		break
	case "after":
		filterChains = app.afterFilters
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
		Parameters:       c.Parameters,
		context:          c,
	})
}

//Route before filter
func (a *App) BeforeFilter(pattern string, filterHandler FilterHandler) *App {
	filter("before", pattern, filterHandler)
	return a
}

//Route after filter
func (a *App) AfterFilter(pattern string, filterHandler FilterHandler) *App {
	filter("after", pattern, filterHandler)
	return a
}

//Print filter
func (a *App) printFilter() {
	for _, filter := range a.beforeFilters {
		a.LogInfo("Before filter route [%s]", filter.route.pattern)
	}
	for _, filter := range a.afterFilters {
		a.LogInfo("After filter route [%s]", filter.route.pattern)
	}
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
