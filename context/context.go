package context

import (
	"github.com/billcoding/flygo/log"
	"net/http"
)

//Define Context struct
type Context struct {
	logger       log.Logger             //logger
	Request      *http.Request          //Request
	Response     http.ResponseWriter    //Response
	pos          int                    //Handler pos index
	handlers     []func(c *Context)     //All handlers chain
	Wrote        bool                   //Wrote
	MWData       map[string]interface{} //MWData
	paramMap     map[string][]string    //The ParamMap
	MultipartMap map[string][]*MultipartFile
}

//New
func New(r *http.Request, w http.ResponseWriter) *Context {
	c := &Context{
		logger:   log.New("[Context]"),
		Request:  r,
		Response: w,
		pos:      -1,
		handlers: make([]func(c *Context), 0),
		paramMap: make(map[string][]string, 0),
		MWData:   make(map[string]interface{}, 0),
	}
	c.onCreated()
	return c
}

//Add
func (c *Context) Add(handlers ...func(c *Context)) *Context {
	if len(handlers) <= 0 {
		return c
	}
	c.onPreparedAddOnce(handlers...)
	c.OnPreparedAdd(handlers...)
	c.handlers = append(c.handlers, handlers...)
	c.onAddedOnce(handlers...)
	c.onAdded(handlers...)
	return c
}

//Chain
func (c *Context) Chain() {
	if len(c.handlers) <= 0 {
		return
	}
	c.pos++
	if c.pos > len(c.handlers)-1 {
		//out of range
		c.onDestoryed()
		return
	}
	c.handlers[c.pos](c)
}

//Write
func (c *Context) Write(buffer []byte) {
	if !c.Wrote {
		c.Response.Write(buffer)
		c.Wrote = true
	}
}

//WriteHeader
func (c *Context) WriteHeader(code int) {
	c.Response.WriteHeader(code)
}

//Header
func (c *Context) Header() http.Header {
	return c.Response.Header()
}
