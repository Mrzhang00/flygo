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
	if app.Config.Flygo.Template.Enable {
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
	if app.Config.Flygo.Template.Enable {
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
	if !app.Config.Flygo.View.Enable {
		return
	}
	viewData, err := parseView(name)
	if err != nil {
		app.Error(err.Error())
		return
	}

	if app.Config.Flygo.Template.Enable {
		tt := template.New("template")
		if app.TemplateFuncs != nil {
			for k, fc := range app.TemplateFuncs {
				c.funcMap[k] = fc
			}
		}
		if len(c.funcMap) > 0 {
			tt.Funcs(c.funcMap)
		}
		left := app.Config.Flygo.Template.Delims.Left
		right := app.Config.Flygo.Template.Delims.Right
		if left != "" && right != "" {
			tt.Delims(left, right)
		}
		templ := template.Must(tt.Parse(viewData))
		if err != nil {
			app.Error(err.Error())
			return
		}
		if data == nil {
			data = make(map[string]interface{})
		}
		if app.Config.Flygo.Session.Enable && app.SessionProvider != nil {
			data["session"] = c.Session.GetAll()
		}
		data["application"] = app.Caches
		c.setDefaultHeaders()
		err = templ.Execute(c.ResponseWriter, data)
		if err != nil {
			app.Error(err.Error())
			return
		}
	} else {
		c.Render([]byte(viewData), contentTypeHtml)
	}
}

//Response text to client
func (c *Context) Text(text string) {
	c.Render([]byte(text), contentTypeText)
}

//Response json to client
func (c *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		app.Error(err.Error())
		return
	}
	c.Render(jsonData, contentTypeJson)
}

//Response image to client
func (c *Context) Image(buffer []byte) {
	c.Render(buffer, contentTypeImage)
}

//Response jpg image to client
func (c *Context) JPG(buffer []byte) {
	c.Render(buffer, contentTypeJpg)
}

//Response jpeg image to client
func (c *Context) JPEG(buffer []byte) {
	c.Render(buffer, contentTypeJpg)
}

//Response png image to client
func (c *Context) PNG(buffer []byte) {
	c.Render(buffer, contentTypePng)
}

//Response gif image to client
func (c *Context) GIF(buffer []byte) {
	c.Render(buffer, contentTypeGif)
}

//Response css to client
func (c *Context) CSS(buffer []byte) {
	c.Render(buffer, contentTypeCSS)
}

//Response css to client
func (c *Context) JS(buffer []byte) {
	c.Render(buffer, contentTypeJS)
}

//Base Response
func (c *Context) Render(buffer []byte, contentType string) {
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
