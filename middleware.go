package flygo

import (
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/session"
	"github.com/billcoding/flygo/util"
	"regexp"
)

type defaultMWState struct {
	header                  bool
	methodNotAllowed        bool
	methodNotAllowedHandler func(ctx *context.Context)
	recovery                bool
	recoveryHandler         func(ctx *context.Context)
	recoveryCodeName        string
	recoveryCodeVal         int
	recoveryMsgName         string
	notFound                bool
	notFoundHandler         func(ctx *context.Context)
	stdLogger               bool
	session                 bool
	provider                session.Provider
	config                  *session.Config
	listener                *session.Listener
	static                  bool
	staticHandler           func(ctx *context.Context)
	staticCache             bool
	staticRoot              string
}

// Use Middlewares
func (a *App) Use(middlewares ...middleware.Middleware) *App {
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

// UseMethodNotAllowed Use Method Not Allowed Middleware
func (a *App) UseMethodNotAllowed() *App {
	a.defaultMWState.methodNotAllowed = true
	return a
}

// MethodNotAllowedHandler Sets MethodNotAllowed handler
func (a *App) MethodNotAllowedHandler(handlers ...func(ctx *context.Context)) *App {
	if len(handlers) > 0 {
		a.defaultMWState.methodNotAllowedHandler = handlers[0]
	}
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

// StaticHandler Sets Static Resources handler
func (a *App) StaticHandler(handlers ...func(ctx *context.Context)) *App {
	if len(handlers) > 0 {
		a.defaultMWState.staticHandler = handlers[0]
	}
	return a
}

func (a *App) useDefaultMWs() *App {

	if a.defaultMWState.session {
		a.middlewares[0] = middleware.Session(a.defaultMWState.provider, a.defaultMWState.config, a.defaultMWState.listener)
	}

	if a.defaultMWState.header {
		a.middlewares[1] = middleware.Header()
	}

	if a.defaultMWState.methodNotAllowed {
		a.middlewares[2] = middleware.MethodNotAllowed(a.defaultMWState.methodNotAllowedHandler)
	}

	if a.defaultMWState.stdLogger {
		a.middlewares[3] = middleware.StdLogger()
	}

	if a.defaultMWState.recovery {
		if a.defaultMWState.recoveryMsgName != "" {
			a.middlewares[4] = middleware.RecoveryWithConfig(a.defaultMWState.recoveryCodeName, a.defaultMWState.recoveryCodeVal, a.defaultMWState.recoveryMsgName, a.defaultMWState.recoveryHandler)
		} else {
			a.middlewares[4] = middleware.Recovery(a.defaultMWState.recoveryHandler)
		}
	}

	if a.defaultMWState.notFound {
		a.middlewares[5] = middleware.NotFound(a.defaultMWState.notFoundHandler)
	}

	if a.defaultMWState.static {
		a.middlewares = append(a.middlewares, middleware.Static(a.defaultMWState.staticCache, a.defaultMWState.staticRoot, a.defaultMWState.staticHandler))
	}

	return a
}

// Middlewares Filter Middlewares
func (a *App) Middlewares(ctx *context.Context, mtype *middleware.Type) []middleware.Middleware {
	mws := make([]middleware.Middleware, 0)
	if len(a.middlewares) > 0 {
		for _, mw := range a.middlewares {
			if mw == nil {
				continue
			}
			matched := false
			if mtype == mw.Type() {
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
				if mtype == middleware.TypeHandler {
					break
				}
			}
		}
	}
	return mws
}
