package flygo

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	mw "github.com/billcoding/flygo/middleware"
	se "github.com/billcoding/flygo/session"
	"regexp"
	"strings"
)

//Define defaultMWState struct
type defaultMWState struct {
	session          bool
	header           bool
	methodNotAllowed bool
	recovered        bool
	notFound         bool
	stdLogger        bool
	provider         se.Provider
	config           *se.Config
	listener         *se.Listener
}

//Use
func (a *App) Use(middlewares ...mw.Middleware) *App {
	a.middlewares = append(a.middlewares, middlewares...)
	return a
}

//UseSession
func (a *App) UseSession(provider se.Provider, config *se.Config, listener *se.Listener) *App {
	a.defaultMWState.session = true
	a.defaultMWState.provider = provider
	a.defaultMWState.config = config
	a.defaultMWState.listener = listener
	return a
}

//UseHeader
func (a *App) UseHeader() *App {
	a.defaultMWState.header = true
	return a
}

//UseMethodAllowed
func (a *App) UseMethodAllowed() *App {
	a.defaultMWState.methodNotAllowed = true
	return a
}

//UseRecovered
func (a *App) UseRecovered() *App {
	a.defaultMWState.recovered = true
	return a
}

//UseNotFound
func (a *App) UseNotFound() *App {
	a.defaultMWState.notFound = true
	return a
}

//UseStdLogger
func (a *App) UseStdLogger() *App {
	a.defaultMWState.stdLogger = true
	return a
}

//Use default
func (a *App) useDefaultMWs() *App {

	//use session middleware
	if a.defaultMWState.session {
		a.middlewares = append(a.middlewares, mw.Session(
			a.defaultMWState.provider,
			a.defaultMWState.config,
			a.defaultMWState.listener),
		)
	}

	//use header middleware
	if a.defaultMWState.header {
		a.middlewares = append(a.middlewares, mw.Header())
	}

	//use method not allowed middleware
	if a.defaultMWState.methodNotAllowed {
		a.middlewares = append(a.middlewares, mw.MethodNotAllowed())
	}

	//use recovered middleware
	if a.defaultMWState.recovered {
		a.middlewares = append(a.middlewares, mw.Recovered())
	}

	//use std logger middleware
	if a.defaultMWState.stdLogger {
		a.middlewares = append(a.middlewares, mw.StdLogger())
	}

	//use not found middleware
	if a.defaultMWState.notFound {
		a.middlewares = append(a.middlewares, mw.NotFound())
	}

	return a
}

//Middlewares
func (a *App) Middlewares(c *c.Context, mtype *mw.Type) []mw.Middleware {
	mws := make([]mw.Middleware, 0)
	if len(a.middlewares) > 0 {
		for _, middleware := range a.middlewares {
			matched := false
			if mtype == middleware.Type() {
				if middleware.Method() == mw.MethodAny || string(middleware.Method()) == c.Request.Method {
					if middleware.Pattern() == mw.PatternAny {
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

//trimPattern
func trimPattern(pattern string) string {
	re := regexp.MustCompile(`[^/\w-._*]`)
	np := re.ReplaceAllString(pattern, "")
	np = strings.ReplaceAll(np, "**", "*")
	np = strings.ReplaceAll(np, "*", `[\w-._/]+`)
	return fmt.Sprintf(`^%s$`, np)
}
