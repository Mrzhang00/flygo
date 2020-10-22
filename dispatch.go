package flygo

import (
	"net/http"
	"strings"
	"sync"
)

//Define dispatcher struct
type dispatcher struct {
	mutex sync.Mutex
	c     *Context
}

func newDispatcher() *dispatcher {
	var mutex sync.Mutex
	return &dispatcher{
		mutex: mutex,
	}
}

//Create app c
func (d *dispatcher) initContext(writer http.ResponseWriter, request *http.Request) {
	d.c = &Context{
		Request:          request,
		RequestURI:       request.RequestURI,
		ParsedRequestURI: request.RequestURI,
		RequestMethod:    strings.ToUpper(request.Method),
		RequestHeader:    request.Header,
		ResponseWriter:   writer,
		Parameters:       make(map[string][]string, 0),
		Multipart:        make(map[string][]*MultipartFile, 0),
		Response: &Response{
			contentType: contentTypeText,
		},
		middlewareCtx: app.initMiddlewareCtx(),
		funcMap:       make(map[string]interface{}),
	}

	if app.sessionEnable && app.sessionProvider != nil {
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
	go invoke(d, w, r, done)
	<-done
}

func invoke(d *dispatcher, w http.ResponseWriter, r *http.Request, done chan<- bool) {
	//init c
	d.initContext(w, r)

	defer func() {
		//set default headers
		d.c.setDefaultHeaders()
		//set content type
		d.c.setContextType()
		//start write
		d.write()
	}()

	//parse request uri
	d.c.parseRequestURI()

	//try parse form
	d.c.parseForm()

	//match static resource
	invokedStatic := d.c.invokeStatic()

	if invokedStatic {
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
