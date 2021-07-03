package flygo

import (
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/session"
	"github.com/billcoding/flygo/util"
	"regexp"
)

type defaultMWState struct {
	header           bool
	recovery         bool
	recoveryHandler  func(ctx *context.Context)
	recoveryCodeName string
	recoveryCodeVal  int
	recoveryMsgName  string
	notFound         bool
	notFoundHandler  func(ctx *context.Context)
	stdLogger        bool
	session          bool
	provider         session.Provider
	config           *session.Config
	listener         *session.Listener
	static           bool
	staticCache      bool
	staticRoot       string
}

// Use Middlewares
func (a *App) Use(middlewares ...middleware.Middleware) *App {
	for _, mw := range middlewares {
		if _, have := a.middlewareMap[mw.Name()]; have {
			a.Logger.Panicf("middleware: %s was registered", mw.Name())
		} else {
			a.middlewareMap[mw.Name()] = mw
		}
	}
	a.middlewares = append(a.middlewares, middlewares...)
	return a
}

// UseSession Use Session Middleware
func (a *App) UseSession(provider session.Provider, config *session.Config, listener *session.Listener) *App {
	a.defaultMWState.session = true
	a.defaultMWState.provider = provider
	a.defaultMWState.config = config
	a.defaultMWState.listener = listener
	return a
}

// UseHeader Use Header Middleware
func (a *App) UseHeader() *App {
	a.defaultMWState.header = true
	return a
}

// UseRecovery Use Recovery Middleware
func (a *App) UseRecovery() *App {
	a.defaultMWState.recovery = true
	return a
}

// RecoveryConfig Sets Recovery config
func (a *App) RecoveryConfig(codeName string, codeVal int, msgName string) *App {
	a.defaultMWState.recoveryCodeName = codeName
	a.defaultMWState.recoveryCodeVal = codeVal
	a.defaultMWState.recoveryMsgName = msgName
	return a
}

// RecoveryHandler Sets Recovery handler
func (a *App) RecoveryHandler(handlers ...func(ctx *context.Context)) *App {
	if len(handlers) > 0 {
		a.defaultMWState.recoveryHandler = handlers[0]
	}
	return a
}

// UseNotFound Use Not Found Middleware
func (a *App) UseNotFound() *App {
	a.defaultMWState.notFound = true
	return a
}

// NotFoundHandler Sets Not Found handler
func (a *App) NotFoundHandler(handlers ...func(ctx *context.Context)) *App {
	if len(handlers) > 0 {
		a.defaultMWState.notFoundHandler = handlers[0]
	}
	return a
}

// UseStdLogger Use Std Logger Middleware
func (a *App) UseStdLogger() *App {
	a.defaultMWState.stdLogger = true
	return a
}

// UseStatic Use Static Resources Middleware
func (a *App) UseStatic(cache bool, root string) *App {
	a.defaultMWState.static = true
	a.defaultMWState.staticCache = cache
	a.defaultMWState.staticRoot = root
	return a
}

func (a *App) useDefaultMWs() *App {

	if a.defaultMWState.session {
		a.middlewares[0] = middleware.Session(a.defaultMWState.provider, a.defaultMWState.config, a.defaultMWState.listener)
	}

	if a.defaultMWState.header {
		a.middlewares[1] = middleware.Header()
		a.middlewareMap[a.middlewares[1].Name()] = a.middlewares[1]
	}

	if a.defaultMWState.stdLogger {
		a.middlewares[2] = middleware.StdLogger()
	}

	if a.defaultMWState.recovery {
		if a.defaultMWState.recoveryMsgName != "" {
			a.middlewares[3] = middleware.RecoveryWithConfig(a.defaultMWState.recoveryCodeName, a.defaultMWState.recoveryCodeVal, a.defaultMWState.recoveryMsgName, a.defaultMWState.recoveryHandler)
		} else {
			a.middlewares[3] = middleware.Recovery(a.defaultMWState.recoveryHandler)
		}
	}

	if a.defaultMWState.notFound {
		a.middlewares[4] = middleware.NotFound(a.defaultMWState.notFoundHandler)
	}

	if a.defaultMWState.static {
		a.middlewares = append(a.middlewares, middleware.Static(a.defaultMWState.staticCache, a.defaultMWState.staticRoot))
	}

	a.setMiddlewareMap()

	return a
}

// Middlewares Filter Middlewares
func (a *App) Middlewares(ctx *context.Context, mwType *middleware.Type) []middleware.Middleware {
	mws := make([]middleware.Middleware, 0)
	if len(a.middlewares) > 0 {
		for _, mw := range a.middlewares {
			if mw == nil {
				continue
			}
			matched := false
			if mwType == mw.Type() {
				if mw.Method() == middleware.MethodAny || string(mw.Method()) == ctx.Request.Method {
					if mw.Pattern() == middleware.PatternNoRoute {
						matched = true
					} else if mw.Pattern() == middleware.PatternAny {
						matched = true
					} else if string(mw.Pattern()) == util.TrimLeftAndRight(ctx.Request.URL.Path) {
						matched = true
					} else {
						reEp := util.TrimPattern(string(mw.Pattern()))
						re := regexp.MustCompile(reEp)
						matched = re.MatchString(util.TrimLeftAndRight(ctx.Request.URL.Path))
					}
				}
			}
			if matched {
				mws = append(mws, mw)
				if mwType == middleware.TypeHandler {
					break
				}
			}
		}
	}
	a.setMiddlewareMap()
	return mws
}

func (a *App) Middleware(name string) middleware.Middleware {
	return a.middlewareMap[name]
}

func (a *App) setMiddlewareMap() {
	for _, mw := range a.middlewares {
		if mw != nil {
			a.middlewareMap[mw.Name()] = mw
		}
	}
}
