package flygo

import (
	"github.com/billcoding/calls"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	mw "github.com/billcoding/flygo/middleware"
	"net/http"
	"sync"
)

type dispatcher struct {
	mu   *sync.Mutex
	app  *App
	done chan bool
}

func (a *App) newDispatcher() *dispatcher {
	return &dispatcher{
		app:  a,
		mu:   &sync.Mutex{},
		done: make(chan bool, 1),
	}
}

func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.dispatch(r, w)
}

func (d *dispatcher) dispatch(r *http.Request, w http.ResponseWriter) {

	ctx := c.New(r, d.app.Config.Flygo.Template)

	d.addChains(ctx,
		d.app.Middlewares(ctx, mw.TypeBefore),
		d.app.parsedRouters.Handlers(ctx),
		d.app.Middlewares(ctx, mw.TypeHandler),
		d.app.Middlewares(ctx, mw.TypeAfter))

	ctx.Chain()

	d.writeDone(ctx.Render(), w)
}

func (d *dispatcher) addChains(c *c.Context,
	beforeMWs []mw.Middleware,
	handler func(c *c.Context),
	handlerMWs []mw.Middleware,
	afterMWs []mw.Middleware) {

	if len(beforeMWs) > 0 {
		for _, bmw := range beforeMWs {
			c.Add(bmw.Handler())
		}
	}

	if handler != nil {

		c.Add(handler)
	} else {

		for _, hmw := range handlerMWs {
			c.Add(hmw.Handler())
			break
		}
	}

	if len(afterMWs) > 0 {
		for _, amw := range afterMWs {
			c.Add(amw.Handler())
		}
	}
}

func (d *dispatcher) writeDone(r *c.Render, w http.ResponseWriter) {
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	calls.NEmpty(r.ContentType, func() {
		w.Header().Set(headers.MIME, r.ContentType)
	})

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
