package flygo

import (
	"io/ioutil"
	"net/http"
)

type staticCache map[string][]byte
type staticMimeCache map[string]string
type viewCache map[string]string
type patternRoute map[string]map[string]patternHandlerRoute
type variableRoute map[string]map[string]variableHandlerRoute
type aroundFilter map[string]filterRouteChain
type aroundInterceptor map[string]interceptorRouteChain
type appCache map[string]interface{}
type middleware map[string]int

var faviconIconHandler = func(c *Context, contentType, resourcePath string) {
	app.StaticHandler(c, contentTypeIco, resourcePath)
}

var notFoundHandler = func(c *Context) {
	c.ResponseWriter.WriteHeader(http.StatusNotFound)
}

var methodNotAllowedHandler = func(c *Context) {
	c.ResponseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

var preflightedHandler = func(c *Context) {
	c.ResponseHeader.Set("Allow", "GET,POST,DELETE,PUT,PATCH,HEAD,OPTIONS")
}

var staticHandler = func(c *Context, contentType, resourcePath string) {
	data := app.staticCaches[resourcePath]
	if data != nil {
		c.Render(data, app.staticMimeCaches[resourcePath])
		return
	}
	buffer, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		c.ResponseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	if app.Config.Flygo.Static.Cache {
		app.staticCaches[resourcePath] = buffer
		app.staticMimeCaches[resourcePath] = contentType
	}
	c.Render(buffer, contentType)
}
