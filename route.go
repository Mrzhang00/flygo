package flygo

import (
	"fmt"
	"regexp"
	"strings"
)

//Route Handler for all method
func (a *App) Req(pattern string, handler Handler, fields ...*Field) *App {
	return a.route(methodAll, pattern, handler, fields...)
}

//Route Handler for GET method
func (a *App) Get(pattern string, handler Handler, fields ...*Field) *App {
	return a.route(methodGet, pattern, handler, fields...)
}

//Route Handler for POST method
func (a *App) Post(pattern string, handler Handler, fields ...*Field) *App {
	return a.route(methodPost, pattern, handler, fields...)
}

//Route Handler for PUT method
func (a *App) Put(pattern string, handler Handler, fields ...*Field) *App {
	return a.route(methodPut, pattern, handler, fields...)
}

//Route Handler for DELETE method
func (a *App) Delete(pattern string, handler Handler, fields ...*Field) *App {
	return a.route(methodDelete, pattern, handler, fields...)
}

//Route Handler for PATCH method
func (a *App) Patch(pattern string, handler Handler, fields ...*Field) *App {
	return a.route(methodPatch, pattern, handler, fields...)
}

//Route handler
func (a *App) startRoute() *App {
	for _, route := range a.routes {
		method := route.method
		pattern := route.pattern
		handler := route.handler
		fields := route.fields

		pattern = trim(pattern)

		//pattern route : /index/path1 or /index/*.go
		dynamicRoute := isVariableRoute(pattern)

		contextPath := a.Config.Flygo.Server.ContextPath
		if !dynamicRoute {
			regex := fmt.Sprintf(`^%s%s$`, contextPath, strings.ReplaceAll(pattern, "*", "[a-zA-Z0-9]+"))
			//pattern
			phr := patternHandlerRoute{
				regex:   regex,
				pattern: pattern,
				method:  method,
				handler: handler,
				fields:  fields,
			}
			routes, have := a.patternRoutes[regex]
			if have {
				//have same route
				routes[method] = phr
				a.patternRoutes[regex] = routes
			} else {
				a.patternRoutes[regex] = map[string]patternHandlerRoute{method: phr}
			}
		} else {
			//variable route : /index/{id}/{name}
			reg := regexp.MustCompile("{[^{]+}")
			matches := findMatches(pattern, "{[^{]+}")
			if matches == nil || len(matches) <= 0 {
				return a
			}
			params := make([]string, 0)
			for _, mat := range matches {
				p := strings.TrimRight(strings.TrimLeft(mat, "{"), "}")
				params = append(params, p)
			}
			regex := reg.ReplaceAllString(pattern, "([a-zA-Z0-9]+)")
			regex = "^" + contextPath + regex + "$"
			vhr := variableHandlerRoute{
				regex:   regex,
				pattern: pattern,
				method:  method,
				params:  params,
				handler: handler,
				fields:  fields,
			}
			routes, have := a.variableRoutes[regex]
			if have {
				//have same route
				routes[method] = vhr
				a.variableRoutes[regex] = routes
			} else {
				a.variableRoutes[regex] = map[string]variableHandlerRoute{method: vhr}
			}
		}
	}
	return a
}

func (a *App) route(method, pattern string, handler Handler, fields ...*Field) *App {
	a.routes = append(a.routes, handlerRouteCache{
		pattern: pattern,
		method:  method,
		handler: &handler,
		fields:  fields,
	})
	return a
}

//Print routes
func (a *App) printRoute() {
	for _, routes := range a.patternRoutes {
		for method, route := range routes {
			a.Logger.Info("Route route [%v:%v] with %v fields", method, route.pattern, len(route.fields))
		}
	}
	for _, routes := range a.variableRoutes {
		for method, route := range routes {
			a.Logger.Info("Variables route [%v:%v] with %v fields", method, route.pattern, len(route.fields))
		}
	}
}
