package context

import (
	"github.com/billcoding/calls"
	"github.com/billcoding/flygo/config"
	"github.com/billcoding/flygo/log"
	"html/template"
	"net/http"
)

type Context struct {
	logger       log.Logger
	Request      *http.Request
	render       *Render
	pos          int
	handlers     []func(c *Context)
	MWData       map[string]interface{}
	paramMap     map[string][]string
	MultipartMap map[string][]*MultipartFile

	dataMap        map[string]interface{}
	funcMap        template.FuncMap
	templateConfig *config.YmlConfigTemplate
}

func New(r *http.Request, templateConfig *config.YmlConfigTemplate) *Context {
	funcMap := make(map[string]interface{}, 0)
	calls.NNil(templateConfig.FuncMap, func() {
		for k, v := range templateConfig.FuncMap {
			funcMap[k] = v
		}
	})
	c := &Context{
		logger:         log.New("[Context]"),
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
	c.onCreated()
	return c
}

func (c *Context) Add(handlers ...func(c *Context)) *Context {
	if len(handlers) <= 0 {
		return c
	}
	c.onPreparedAddOnce(handlers...)
	c.onPreparedAdd(handlers...)
	c.handlers = append(c.handlers, handlers...)
	c.onAddedOnce(handlers...)
	c.onAdded(handlers...)
	return c
}

func (c *Context) Chain() {
	if len(c.handlers) <= 0 {
		return
	}
	c.pos++
	if c.pos > len(c.handlers)-1 {

		c.onDestroyed()
		return
	}
	c.handlers[c.pos](c)
}

func (c *Context) Write(buffer []byte) {
	c.render.Buffer = buffer
}

func (c *Context) WriteCode(code int) {
	c.render.Code = code
}

func (c *Context) Header() http.Header {
	return c.render.Header
}

func (c *Context) AddCookie(cookies ...*http.Cookie) *Context {
	c.render.Cookies = append(c.render.Cookies, cookies...)
	return c
}
