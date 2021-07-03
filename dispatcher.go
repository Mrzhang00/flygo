package flygo

import (
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/middleware"
	"net/http"
	"sync"
)

type dispatcher struct {
	mu  *sync.Mutex
	app *App
}

func (a *App) newDispatcher() *dispatcher {
	return &dispatcher{
		app: a,
		mu:  &sync.Mutex{},
	}
}

func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.dispatch(r, w)
}

func (d *dispatcher) dispatch(r *http.Request, w http.ResponseWriter) {
	ctx := context.New(d.app.Logger, r, d.app.Config.Template)
	d.addChains(ctx,
		d.app.parsedRouters.Handler(ctx),
		d.app.Middlewares(ctx, middleware.TypeBefore),
		d.app.Middlewares(ctx, middleware.TypeHandler),
		d.app.Middlewares(ctx, middleware.TypeAfter))
	ctx.Chain()
	d.writeDone(ctx.Rendered(), w)
}

func (d *dispatcher) addChains(ctx *context.Context, handler func(ctx *context.Context), beforeMWs, handlerMWs, afterMWs []middleware.Middleware) {
	if handler != nil || len(handlerMWs) > 0 {
		for _, bmw := range beforeMWs {
			// Add all route before MW
			ctx.Add(bmw.Handler())
		}
	} else {
		// Add No route before MW
		for _, bmw := range beforeMWs {
			if bmw.Pattern() == middleware.PatternNoRoute {
				ctx.Add(bmw.Handler())
			}
		}
	}

	if handler != nil {
		ctx.Add(handler)
		ctx.Route()
	} else {
		for _, hmw := range handlerMWs {
			ctx.Add(hmw.Handler())
			ctx.Route()
			break
		}
	}

	if handler != nil || len(handlerMWs) > 0 {
		// Add all route after MW
		for _, amw := range afterMWs {
			ctx.Add(amw.Handler())
		}
	} else {
		// Add No route after MW
		for _, amw := range afterMWs {
			if amw.Pattern() == middleware.PatternNoRoute {
				ctx.Add(amw.Handler())
			}
		}
	}
}

func (d *dispatcher) writeDone(r *context.Render, w http.ResponseWriter) {
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	if r.ContentType != "" {
		w.Header().Set(headers.MIME, r.ContentType)
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
