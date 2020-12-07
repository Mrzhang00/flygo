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

type multipartFileMap map[string][]*MultipartFile
type paramMap map[string][]string
type headerMap map[string][]string
type middlewareMap map[string]map[string]interface{}
type cookieMap map[string]*http.Cookie

var faviconIconHandler = func(c *Context, contentType, resourcePath string) {
	c.app.StaticHandler(c, contentTypeIco, resourcePath)
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
	data := c.app.staticCaches[resourcePath]
	if data != nil {
		c.Render(data, c.app.staticMimeCaches[resourcePath])
		return
	}
	buffer, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		c.ResponseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	if c.app.Config.Flygo.Static.Cache {
		c.app.staticCaches[resourcePath] = buffer
		c.app.staticMimeCaches[resourcePath] = contentType
	}
	c.Render(buffer, contentType)
}
