package flygo

import (
	"path/filepath"
	"strings"
)

const (
	contentTypeText   = "text/plain;charset=utf-8"
	contentTypeHtml   = "text/html;charset=utf-8"
	contentTypeJS     = "text/javascript;charset=utf-8"
	contentTypeCSS    = "text/css;charset=utf-8"
	contentTypeJson   = "application/json;charset=utf-8"
	contentTypeBinary = "application/octet-stream"
	contentTypeImage  = "image/jpg"
	contentTypePng    = "image/png"
	contentTypeJpg    = "image/jpg"
	contentTypeGif    = "image/gif"
	contentTypeIco    = "image/x-icon"
)

//Define static handler
type StaticHandler func(c *Context, contentType, resourcePath string)

//Match favicon.ico
func (c *Context) matchFaviconStatic() (string, string) {
	//favicon.ico
	path := "/" + strings.Trim(c.ParsedRequestURI, "/")
	name := app.GetContextPath() + "/favicon.ico"
	realPath := ""
	contentType := ""

	if path == name {
		realPath = strings.Join([]string{app.GetWebRoot(), app.GetStaticPrefix(), name}, string(filepath.Separator))
		contentType = contentTypeIco
	}
	return contentType, realPath
}

//Match static res
func (c *Context) matchStatic() (string, string) {
	regex := "^" + app.GetContextPath() + app.GetStaticPattern() + "/.+$"
	if !c.matchPath(regex) {
		return "", ""
	}
	cpLen := len(app.GetContextPath())
	paLen := len(app.GetStaticPattern())
	subName := c.ParsedRequestURI[cpLen+paLen+1:]
	var contentType string
	var realPath string
	realPath = strings.Join([]string{app.GetWebRoot(), app.GetStaticPrefix(), subName}, string(filepath.Separator))
	suffix := subName[strings.LastIndexByte(subName, '.')+1:]
	contentType = app.statics[suffix]
	if contentType == "" {
		app.log.warn("static[%v] was not registered", suffix)
		return "", ""
	}
	return contentType, realPath
}

//Invoke static
func (c *Context) invokeStatic() bool {
	if !app.GetStaticEnable() {
		return false
	}
	if app.faviconIcon {
		contentType, staticPath := c.matchFaviconStatic()
		if staticPath != "" {
			app.faviconIconHandler(c, contentType, staticPath)
			return true
		}
	}
	contentType, staticPath := c.matchStatic()
	if staticPath != "" {
		app.defaultStaticHandler(c, contentType, staticPath)
		return true
	}
	return false
}

//Register a static res
func (a *App) RegisterStaticRes(fileExt, contentType string) *App {
	a.statics[fileExt] = contentType
	return a
}

//Print static
func (a *App) printStatic() {
	for ext := range a.statics {
		a.LogInfo("Static resource : %s", ext)
	}
}
