package flygo

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	mw "github.com/billcoding/flygo/middleware"
	se "github.com/billcoding/flygo/session"
	"regexp"
	"strings"
)

type defaultMWState struct {
	header           bool
	methodNotAllowed bool
	recovery         bool
	notFound         bool
	stdLogger        bool
	session          bool
	provider         se.Provider
	config           *se.Config
	listener         *se.Listener
	static           bool
	staticCache      bool
	staticRoot       string
}

// Use Middlewares
func (a *App) Use(middlewares ...mw.Middleware) *App {
	a.middlewares = append(a.middlewares, middlewares...)
	return a
}

// UseSession Use Session Middleware
func (a *App) UseSession(provider se.Provider, config *se.Config, listener *se.Listener) *App {
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

// UseRecovery Use Recovery Middleware
func (a *App) UseRecovery() *App {
	a.defaultMWState.recovery = true
	return a
}

// UseNotFound Use Not Found Middleware
func (a *App) UseNotFound() *App {
	a.defaultMWState.notFound = true
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
		a.middlewares[0] = mw.Session(a.defaultMWState.provider, a.defaultMWState.config, a.defaultMWState.listener)
	}

	if a.defaultMWState.header {
		a.middlewares[1] = mw.Header()
	}

	if a.defaultMWState.methodNotAllowed {
		a.middlewares[2] = mw.MethodNotAllowed()
	}

	if a.defaultMWState.stdLogger {
		a.middlewares[3] = mw.StdLogger()
	}

	if a.defaultMWState.recovery {
		a.middlewares[4] = mw.Recovery()
	}

	if a.defaultMWState.notFound {
		a.middlewares[5] = mw.NotFound()
	}

	if a.defaultMWState.static {
		a.middlewares = append(a.middlewares, mw.Static(a.defaultMWState.staticCache, a.defaultMWState.staticRoot))
	}

	return a
}

// Middlewares Filter Middlewares
func (a *App) Middlewares(c *c.Context, mtype *mw.Type) []mw.Middleware {
	mws := make([]mw.Middleware, 0)
	if len(a.middlewares) > 0 {
		for _, middleware := range a.middlewares {
			if middleware == nil {
				continue
			}
			matched := false
			if mtype == middleware.Type() {
				if middleware.Method() == mw.MethodAny || string(middleware.Method()) == c.Request.Method {
					if middleware.Pattern() == mw.PatternNoRoute {
						matched = true
					} else if middleware.Pattern() == mw.PatternAny {
						matched = true
					} else if string(middleware.Pattern()) == c.Request.URL.Path {
						matched = true
					} else {
						reEp := trimPattern(string(middleware.Pattern()))
						re := regexp.MustCompile(reEp)
						matched = re.MatchString(c.Request.URL.Path)
					}
				}
			}
			if matched {
				mws = append(mws, middleware)
				if mtype == mw.TypeHandler {
					break
				}
			}
		}
	}
	return mws
}

func trimPattern(pattern string) string {
	re := regexp.MustCompile(`[^/\w-._*]`)
	np := re.ReplaceAllString(pattern, "")
	np = strings.ReplaceAll(np, "**", "*")
	np = strings.ReplaceAll(np, "*", `[\w-._/]+`)
	return fmt.Sprintf(`^%s$`, np)
}
