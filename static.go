package flygo

import (
	filepath "path/filepath"
	"strings"
)

const (
	contentTypeText   = "text/plain;charset=utf-8"
	contentTypeHtml   = "text/html;charset=utf-8"
	contentTypeJS     = "text/javascript;charset=utf-8"
	contentTypeCSS    = "text/css;charset=utf-8"
	contentTypeJson   = "application/json;charset=utf-8"
	contentTypeXml    = "text/xml;charset=utf-8"
	contentTypeImage  = "image/jpg"
	contentTypePng    = "image/png"
	contentTypeBmp    = "image/bmp"
	contentTypeJpg    = "image/jpg"
	contentTypeGif    = "image/gif"
	contentTypeIco    = "image/x-icon"
	contentTypeBinary = "application/octet-stream"
)

//Define static handler
type StaticHandler func(c *Context, contentType, resourcePath string)

//Match favicon.ico
func (c *Context) matchFaviconStatic() (string, string) {
	//favicon.ico
	webRoot := app.Config.Flygo.Server.WebRoot
	contextPath := app.Config.Flygo.Server.ContextPath
	staticPrefix := app.Config.Flygo.Static.Prefix
	path := "/" + strings.Trim(c.ParsedRequestURI, "/")
	name := contextPath + "/favicon.ico"
	realPath := ""
	contentType := ""
	if path == name {
		realPath = filepath.Join(webRoot, staticPrefix, name)
		contentType = contentTypeIco
	}
	return contentType, realPath
}

//Match static res
func (c *Context) matchStatic() (string, string) {
	webRoot := app.Config.Flygo.Server.WebRoot
	contextPath := app.Config.Flygo.Server.ContextPath
	staticPattern := app.Config.Flygo.Static.Pattern
	staticPrefix := app.Config.Flygo.Static.Prefix
	regex := "^" + strings.Join([]string{contextPath, staticPattern}, "/") + "/.+$"
	if !c.matchPath(regex) {
		return "", ""
	}
	cpLen := len(contextPath)
	paLen := len(staticPattern) + 1
	subName := c.ParsedRequestURI[cpLen+paLen+1:]
	var contentType string
	var realPath string
	realPath = filepath.Join(webRoot, staticPrefix, subName)
	suffix := subName[strings.LastIndexByte(subName, '.')+1:]
	contentType = app.Config.Flygo.Static.Mimes[suffix]
	if contentType == "" {
		app.Warn("static[%v] was not registered", suffix)
		return "", ""
	}
	return contentType, realPath
}

//Invoke static
func (c *Context) invokeStatic() bool {
	if !app.Config.Flygo.Static.Enable {
		return false
	}
	if app.Config.Flygo.Static.Favicon.Enable {
		contentType, staticPath := c.matchFaviconStatic()
		if staticPath != "" {
			app.FaviconIconHandler(c, contentType, staticPath)
			return true
		}
	}
	contentType, staticPath := c.matchStatic()
	if staticPath != "" {
		app.StaticHandler(c, contentType, staticPath)
		return true
	}
	return false
}
