package flygo

import (
	"strings"
)

const (
	methodAll     = "*"
	methodGet     = "GET"
	methodPost    = "POST"
	methodPut     = "PUT"
	methodDelete  = "DELETE"
	methodPatch   = "PATCH"
	methodOptions = "OPTIONS"
	methodHead    = "HEAD"
)

type handlerRouteCache patternHandlerRoute

//Define pattern handler route struct
type patternHandlerRoute struct {
	regex   string
	pattern string
	method  string
	handler *Handler
	fields  []*Field
}

//Define variable route struct
type variableHandlerRoute struct {
	regex   string
	pattern string
	method  string
	handler *Handler
	params  []string
	fields  []*Field
}

//Define handler
type Handler func(*Context)

//match handler
func (c *Context) matchHandler() ([]*Field, Handler) {
	var handler Handler
	fields, handler := c.matchPatternHandler()
	if handler != nil {
		return fields, handler
	}

	fields, handler = c.matchVariableHandler()
	if handler != nil {
		return fields, handler
	}

	return nil, nil
}

//match pattern handler
func (c *Context) matchPatternHandler() ([]*Field, Handler) {
	var phrs []patternHandlerRoute
	for regex, routes := range app.patternRoutes {
		if c.matchPath(regex) {
			for _, route := range routes {
				phrs = append(phrs, route)
			}
		}
	}

	if phrs == nil || len(phrs) <= 0 {
		return nil, nil
	}

	var phr *patternHandlerRoute
	for _, r := range phrs {

		if r.method == methodAll ||
			c.RequestMethod == methodOptions ||
			c.RequestMethod == methodHead ||
			r.method == c.RequestMethod {
			phr = &r
			break
		}
	}

	if phr == nil {
		var ms []string
		for _, r := range phrs {
			ms = append(ms, r.method)
		}
		app.Info("Request is not supported {config: [%s], request: [%s]}", strings.Join(ms, ","), c.RequestMethod)
		return nil, app.NotFoundHandler
	}

	if c.RequestMethod == methodOptions || c.RequestMethod == methodHead {
		return nil, app.PreflightedHandler
	}

	return phr.fields, *phr.handler
}

//match variable handler
func (c *Context) matchVariableHandler() ([]*Field, Handler) {
	//  path : /index/123/name
	//  pattern : /index/{id}/name
	//  regex : /index/.+/name
	var vhrs []variableHandlerRoute
	for regex, routes := range app.variableRoutes {
		for _, route := range routes {
			if route.pattern == c.ParsedRequestURI || c.matchPath(regex) {
				vhrs = append(vhrs, route)
			}
		}
	}

	if vhrs == nil || len(vhrs) <= 0 {
		return nil, nil
	}

	var vhr *variableHandlerRoute
	for _, r := range vhrs {
		if r.method == methodAll ||
			c.RequestMethod == methodOptions ||
			c.RequestMethod == methodHead ||
			r.method == c.RequestMethod {
			vhr = &r
			break
		}
	}

	if vhr == nil {
		return nil, nil
	}

	if vhr == nil {
		var ms []string
		for _, r := range vhrs {
			ms = append(ms, r.method)
		}
		app.Warn("Request is not supported {config: [%s], request: [%s]}", strings.Join(ms, ","), c.RequestMethod)
		return nil, app.MethodNotAllowedHandler
	}

	if c.RequestMethod == methodOptions || c.RequestMethod == methodHead {
		return nil, app.PreflightedHandler
	}

	//Setting params
	matches := c.contextMatches(vhr.regex)
	if nil != matches && len(matches)-1 == len(vhr.params) {
		for i, p := range matches {
			pn := vhr.params[i]
			c.ParamMap[pn] = []string{p}
		}
	}

	return vhr.fields, *vhr.handler
}

//Invoke handler
func (c *Context) invokeHandler() {
	//If the response is done, direct return
	if c.Response.done {
		return
	}
	//match handler
	fields, handler := c.matchHandler()
	if handler == nil {
		app.NotFoundHandler(c)
		return
	}

	if fields != nil {
		err := presetAndValidate(fields, c)
		if err != nil {
			c.Text(err.Error())
			c.Response.SetDone(true)
			return
		}
	}

	handler(c)
}
