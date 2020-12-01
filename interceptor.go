package flygo

import (
	"fmt"
	"strings"
)

//Define interceptor route struct
type interceptorRoute struct {
	pattern            string
	regex              string
	interceptorHandler InterceptorHandler
}

//Define interceptor route chain struct
type interceptorRouteChain struct {
	regex string
	route interceptorRoute
}

type iroute struct {
	t string
	interceptorRoute
}

//Define interceptor handler
type InterceptorHandler func(*Context)

//Route interceptor handler
func (a *App) interceptor(interceptorType, pattern string, interceptorHandler InterceptorHandler) {
	if pattern == "" {
		return
	}
	a.iroutes = append(a.iroutes, iroute{
		t: interceptorType,
		interceptorRoute: interceptorRoute{
			pattern:            pattern,
			interceptorHandler: interceptorHandler,
		},
	})
}

func (a *App) startInterceptor() {
	for _, iroute := range a.iroutes {
		pattern := iroute.pattern
		interceptorType := iroute.t
		interceptorHandler := iroute.interceptorHandler
		contexPath := app.Config.Flygo.Server.ContextPath
		regex := fmt.Sprintf(`^%s%s$`, contexPath, strings.ReplaceAll(trim(pattern), "*", "[/a-zA-Z0-9]+"))
		route := interceptorRoute{
			pattern:            pattern,
			regex:              regex,
			interceptorHandler: interceptorHandler,
		}
		if interceptorType == "before" {
			a.beforeInterceptors[regex] = interceptorRouteChain{regex: regex, route: route}
		} else {
			a.afterInterceptors[regex] = interceptorRouteChain{regex: regex, route: route}
		}
	}
}

//Match before interceptor
func (c *Context) matchBeforeInterceptor() []InterceptorHandler {
	return c.matchInterceptor("before")
}

//Match after interceptor
func (c *Context) matchAfterInterceptor() []InterceptorHandler {
	return c.matchInterceptor("after")
}

//Match base method
func (c *Context) matchInterceptor(interceptorType string) []InterceptorHandler {
	var interceptorChains map[string]interceptorRouteChain
	switch interceptorType {
	case "before":
		interceptorChains = app.beforeInterceptors
		break
	case "after":
		interceptorChains = app.afterInterceptors
		break
	}
	if interceptorChains == nil {
		return nil
	}
	interceptors := make([]InterceptorHandler, 0)
	for regex, chain := range interceptorChains {
		if c.matchPath(regex) {
			interceptors = append(interceptors, chain.route.interceptorHandler)
		}
	}
	return interceptors
}

//Invoke before interceptor
func (c *Context) invokeBeforeInterceptor() {
	interceptors := c.matchBeforeInterceptor()
	if interceptors != nil && len(interceptors) > 0 {
		for _, interceptor := range interceptors {
			c.invokeInterceptorHandler(interceptor)
		}
	}
}

//Invoke after interceptor
func (c *Context) invokeAfterInterceptor() {
	interceptors := c.matchAfterInterceptor()
	if interceptors != nil && len(interceptors) > 0 {
		for _, interceptor := range interceptors {
			c.invokeInterceptorHandler(interceptor)
		}
	}
}

//Invoke interceptor handler
func (c *Context) invokeInterceptorHandler(handler InterceptorHandler) {
	if handler == nil {
		return
	}
	handler(c)
}

//Route before interceptor
func (a *App) BeforeInterceptor(pattern string, interceptorHandler InterceptorHandler) *App {
	a.interceptor("before", pattern, interceptorHandler)
	return a
}

//Route after interceptor
func (a *App) AfterInterceptor(pattern string, interceptorHandler InterceptorHandler) *App {
	a.interceptor("after", pattern, interceptorHandler)
	return a
}

//Print interceptor
func (a *App) printInterceptor() {
	for _, chain := range a.beforeInterceptors {
		a.Info("Before interceptor route [%s]", chain.route.pattern)
	}
	for _, chain := range a.afterInterceptors {
		a.Info("After interceptor route [%s]", chain.route.pattern)
	}
}
