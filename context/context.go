package context

import (
	"github.com/billcoding/flygo/config"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

// Context struct
type Context struct {
	Logger         *logrus.Logger
	Request        *http.Request
	render         *Render
	pos            int
	handlers       []func(c *Context)
	MWData         map[string]interface{}
	paramMap       map[string][]string
	MultipartMap   map[string][]*MultipartFile
	dataMap        map[string]interface{}
	funcMap        template.FuncMap
	templateConfig *config.YmlConfigTemplate
	handlerRouted  bool
}

// New context
func New(logger *logrus.Logger, r *http.Request, templateConfig *config.YmlConfigTemplate) *Context {
	funcMap := make(map[string]interface{}, 0)
	if templateConfig.FuncMap != nil {
		for k, v := range templateConfig.FuncMap {
			funcMap[k] = v
		}
	}
	ctx := &Context{
		Logger:         logger,
		Request:        r,
		render:         RenderBuilder().DefaultBuild(),
		pos:            -1,
		handlers:       make([]func(c *Context), 0),
		MWData:         make(map[string]interface{}, 0),
		paramMap:       make(map[string][]string, 0),
		MultipartMap:   make(map[string][]*MultipartFile, 0),
		dataMap:        make(map[string]interface{}, 0),
		funcMap:        funcMap,
		templateConfig: templateConfig,
	}
	ctx.onCreated()

	// try parse form
	_ = ctx.Request.ParseForm()

	if ctx.Request.Form != nil {
		for k := range ctx.Request.Form {
			ctx.paramMap[k] = ctx.Request.Form[k]
		}
	}

	return ctx
}

// Add context handler
func (ctx *Context) Add(handlers ...func(ctx *Context)) *Context {
	if len(handlers) <= 0 {
		return ctx
	}
	ctx.handlers = append(ctx.handlers, handlers...)
	return ctx
}

// Chain execute context handler
func (ctx *Context) Chain() {
	if len(ctx.handlers) <= 0 {
		return
	}
	ctx.pos++
	if ctx.pos == 0 {
		ctx.onBefore()
	}
	if ctx.pos > len(ctx.handlers)-1 {
		ctx.onAfter()
		ctx.onDestroyed()
		return
	}
	ctx.handlers[ctx.pos](ctx)
}

// Write buffer
func (ctx *Context) Write(buffer []byte) {
	ctx.render.Buffer = buffer
}

// WriteCode code
func (ctx *Context) WriteCode(code int) {
	ctx.render.Code = code
}

// Header return header
func (ctx *Context) Header() http.Header {
	return ctx.render.Header
}

// AddCookie add cookie
func (ctx *Context) AddCookie(cookies ...*http.Cookie) *Context {
	ctx.render.Cookies = append(ctx.render.Cookies, cookies...)
	return ctx
}
