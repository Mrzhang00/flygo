package flygo

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	mw "github.com/billcoding/flygo/middleware"
	"net/http"
	"sync"
)

//Define dispatcher struct
type dispatcher struct {
	mu   *sync.Mutex //mu
	wmu  *sync.Mutex //wmu
	app  *App        //app bundle
	done chan bool   //done channel
}

//newDispatcher
func (a *App) newDispatcher() *dispatcher {
	return &dispatcher{
		app:  a,
		mu:   &sync.Mutex{},
		wmu:  &sync.Mutex{},
		done: make(chan bool, 1),
	}
}

//ServeHTTP main entry point
func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.dispatch(r, w)
}

//dispatch
func (d *dispatcher) dispatch(r *http.Request, w http.ResponseWriter) {
	//Init context
	ctx := c.New(r)

	//Add chains into context
	d.addChains(ctx,
		d.app.Middlewares(ctx, mw.TypeBefore),
		d.app.parsedRouters.Handlers(ctx),
		d.app.Middlewares(ctx, mw.TypeHandler),
		d.app.Middlewares(ctx, mw.TypeAfter))

	//Start chain
	ctx.Chain()

	//Finish done
	d.writeDone(ctx.Render(), w)
}

//addChains
func (d *dispatcher) addChains(c *c.Context,
	beforeMWs []mw.Middleware,
	handler func(c *c.Context),
	handlerMWs []mw.Middleware,
	afterMWs []mw.Middleware) {

	//First add before middlewares
	if len(beforeMWs) > 0 {
		for _, bmw := range beforeMWs {
			c.Add(bmw.Handler())
		}
	}

	if handler != nil {
		//And then add handler
		c.Add(handler)
	} else {
		//And then add handler middleware
		for _, hmw := range handlerMWs {
			c.Add(hmw.Handler())
			break
		}
	}

	//And then add after middlewares
	if len(afterMWs) > 0 {
		for _, amw := range afterMWs {
			c.Add(amw.Handler())
		}
	}
}

//writeDoned
func (d *dispatcher) writeDone(r *c.Render, w http.ResponseWriter) {
	d.wmu.Lock()
	defer d.wmu.Unlock()
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	for _, cookie := range r.Cookies {
		w.Header().Add(headers.SetCookie, cookie.String())
	}

	if r.Code != 0 {
		w.WriteHeader(r.Code)
	}

	if r.Buffer != nil {
		_, _ = w.Write(r.Buffer)
	}
}
