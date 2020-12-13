package flygo

import (
	c "github.com/billcoding/flygo/context"
	mw "github.com/billcoding/flygo/middleware"
	"net/http"
	"sync"
)

//Define dispatcher struct
type dispatcher struct {
	pool *sync.Pool
	mu   *sync.RWMutex //mu
	app  *App          //app bundle
	done chan bool     //done channel
}

//newDispatcher
func (a *App) newDispatcher() *dispatcher {
	return &dispatcher{
		pool: &sync.Pool{
			New: func() interface{} {
				return c.New(nil, nil)
			},
		},
		app:  a,
		mu:   &sync.RWMutex{},
		done: make(chan bool, 1),
	}
}

//ServeHTTP main entry point
func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.dispatch(r, w)
	//d.waitDone()
}

//dispatch
func (d *dispatcher) dispatch(r *http.Request, w http.ResponseWriter) {
	func() {
		//Finish done
		//defer d.doned()

		//Init context
		ctx := d.pool.Get().(*c.Context)
		ctx.Request = r
		ctx.Response = w

		//Add chains into context
		d.addChains(ctx,
			d.app.Middlewares(ctx, mw.TypeBefore),
			d.app.parsedRouters.Handlers(ctx),
			d.app.Middlewares(ctx, mw.TypeHandler),
			d.app.Middlewares(ctx, mw.TypeAfter))

		//Start chain
		ctx.Chain()
	}()
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
		}
	}

	//And then add after middlewares
	if len(afterMWs) > 0 {
		for _, amw := range afterMWs {
			c.Add(amw.Handler())
		}
	}
}

//waitDone
func (d *dispatcher) waitDone() {
	<-d.done
}

//makeDone
func (d *dispatcher) doned() {
	d.done <- true
}
