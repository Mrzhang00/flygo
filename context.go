package flygo

import (
	"encoding/json"
	"html/template"
	"net/http"
	"reflect"
)

//Define Context struct
type Context struct {
	Request          *http.Request                     //request
	RequestURI       string                            //request uri
	ParsedRequestURI string                            //parsed request uri
	RequestMethod    string                            //request method
	RequestHeader    http.Header                       //request header
	ResponseWriter   http.ResponseWriter               //Response writer
	ResponseHeader   *http.Header                      //Response header
	Multipart        map[string][]*MultipartFile       //multipart map
	Parameters       map[string][]string               //map[string][]string
	MultipartParsed  bool                              //multipart parsed ?
	Response         *Response                         //Response
	middlewareCtx    map[string]map[string]interface{} //Middleware c
	Cookies          map[string]*http.Cookie           //cookies
	SessionId        string                            //session id
	Session          Session                           //session
	funcMap          template.FuncMap                  //Template funcMap
}

//Set content type
func (c *Context) SetHeader(name, value string) *Context {
	c.ResponseWriter.Header().Set(name, value)
	return c
}

//Add Response header
func (c *Context) AddHeader(name, value string) *Context {
	c.ResponseWriter.Header().Add(name, value)
	return c
}

//Set view funcMap
func (c *Context) SetViewFuncMap(funcMap template.FuncMap) *Context {
	if app.GetTemplateEnable() {
		c.funcMap = funcMap
	}
	return c
}

//Add view funcMap
func (c *Context) AddViewFuncMap(name string, fn interface{}) *Context {
	if fn == nil {
		return c
	}
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		return c
	}
	if app.GetTemplateEnable() {
		c.funcMap[name] = fn
	}
	return c
}

//Response view to client
func (c *Context) View(name string) {
	c.ViewWithData(name, nil)
}

//Response view to client with data
func (c *Context) ViewWithData(name string, data map[string]interface{}) {
	if !app.GetViewEnable() {
		return
	}
	viewData, err := parseView(name)
	if err != nil {
		app.log.fatal(err.Error())
		return
	}

	if app.GetTemplateEnable() {
		tt := template.New("template")
		if app.GetTemplateFuncs() != nil {
			for k, fc := range app.GetTemplateFuncs() {
				c.funcMap[k] = fc
			}
		}
		if len(c.funcMap) > 0 {
			tt.Funcs(c.funcMap)
		}
		if app.GetTemplateDelimLeft() != "" && app.GetTemplateDelimRight() != "" {
			tt.Delims(app.GetTemplateDelimLeft(), app.GetTemplateDelimRight())
		}
		templ := template.Must(tt.Parse(viewData))
		if err != nil {
			app.log.fatal(err.Error())
			return
		}
		if data == nil {
			data = make(map[string]interface{})
		}
		if app.GetSessionEnable() && app.sessionProvider != nil {
			data["session"] = c.Session.GetAll()
		}
		data["application"] = app.GetAllCaches()
		c.setDefaultHeaders()
		err = templ.Execute(c.ResponseWriter, data)
		if err != nil {
			app.log.fatal(err.Error())
			return
		}
	} else {
		c.render([]byte(viewData), contentTypeHtml)
	}
}

//Response text to client
func (c *Context) Text(text string) {
	c.render([]byte(text), contentTypeText)
}

//Response json to client
func (c *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		app.LogFatal(err.Error())
		panic(err.Error())
		return
	}
	c.render(jsonData, contentTypeJson)
}

//Response image to client
func (c *Context) Image(buffer []byte) {
	c.render(buffer, contentTypeImage)
}

//Response binary to client
func (c *Context) Binary(buffer []byte) {
	c.render(buffer, contentTypeBinary)
}

//Response binary to client with ContentType
func (c *Context) BinaryWith(buffer []byte, contentType string) {
	c.render(buffer, contentType)
}

//Base Response
func (c *Context) render(buffer []byte, contentType string) {
	if !c.Response.done {
		c.Response.data = buffer
		c.Response.contentType = contentType
		c.Response.done = true
	}
}

//Set default headers
func (c *Context) setDefaultHeaders() {
	//Server mark
	c.ResponseHeader.Set(headerServer, "golang/flygo")
}

//Set c type
func (c *Context) setContextType() {
	c.ResponseHeader.Set(headerContentType, c.Response.contentType)
}

//Get middleware c
func (c *Context) GetMiddlewareCtx(name string) map[string]interface{} {
	return c.middlewareCtx[name]
}

//Clear middleware c
func (c *Context) ClearMiddlewareCtx(name string) {
	delete(c.middlewareCtx, name)
}

//Remove middleware data
func (c *Context) RemoveMiddlewareData(name, key string) {
	ctx := c.GetMiddlewareCtx(name)
	if ctx != nil {
		delete(ctx, key)
	}
}

//Set middleware data
func (c *Context) SetMiddlewareData(name, key string, val interface{}) {
	v := c.GetMiddlewareCtx(name)
	if v == nil {
		return
	}
	v = map[string]interface{}{
		key: val,
	}
	c.middlewareCtx[name] = v
}

//Get middleware data
func (c *Context) GetMiddlewareData(name, key string) interface{} {
	v := c.GetMiddlewareCtx(name)
	if v == nil {
		return nil
	}
	return v[key]
}
