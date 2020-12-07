package flygo

import (
	"net/http"
	"strings"
	"sync"
)

//Define dispatcher struct
type dispatcher struct {
	mutex sync.Mutex
	app   *App
	c     *Context
}

//New dispatcher
func (a *App) newDispatcher() *dispatcher {
	var mutex sync.Mutex
	return &dispatcher{
		app:   a,
		mutex: mutex,
	}
}

//Create app c
func (d *dispatcher) initContext(writer http.ResponseWriter, request *http.Request) {
	d.c = &Context{
		app:              d.app,
		Request:          request,
		RequestURI:       request.RequestURI,
		ParsedRequestURI: request.RequestURI,
		RequestMethod:    strings.ToUpper(request.Method),
		RequestHeader:    request.Header,
		ResponseWriter:   writer,
		ParamMap:         make(map[string][]string, 0),
		Multipart:        make(map[string][]*MultipartFile, 0),
		Response: &Response{
			contentType: contentTypeText,
		},
		middlewareMap: d.app.initMiddlewareCtx(),
		funcMap:       make(map[string]interface{}),
	}

	if d.app.SessionConfig.Enable && d.app.SessionConfig.SessionProvider != nil {
		d.c.initSession()
	}

	header := writer.Header()

	d.c.ResponseHeader = &header
}

//ServeHTTP main entry point
func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	done := make(chan bool, 1)
	go d.invoke(w, r, done)
	<-done
}

func (d *dispatcher) writeDone() {
	//set default headers
	d.c.setDefaultHeaders()
	//set content type
	d.c.setContextType()
	//start write
	d.write()
}

func (d *dispatcher) invoke(w http.ResponseWriter, r *http.Request, done chan<- bool) {
	//init c
	d.initContext(w, r)

	//parse request uri
	d.c.parseRequestURI()

	//try parse form
	d.c.parseForm()

	//match static resource
	invokedStatic := d.c.invokeStatic()

	if invokedStatic {
		d.writeDone()
		done <- true
		return
	}

	//match and invoke before filter
	d.c.invokeBeforeFilter()

	//match and invoke before interceptor
	d.c.invokeBeforeInterceptor()

	//invoke handler
	d.c.invokeHandler()

	//match and invoke after interceptor
	d.c.invokeAfterInterceptor()

	//match and invoke after filter
	d.c.invokeAfterFilter()

	d.writeDone()
	done <- true
}

//write Response
func (d *dispatcher) write() {
	d.c.ResponseWriter.Write(d.c.Response.data)
}

//parse request uri
func (c *Context) parseRequestURI() {
	index := strings.IndexRune(c.RequestURI, '?')
	if index != -1 {
		//have parameter
		c.ParsedRequestURI = c.RequestURI[:index]
	}
}
