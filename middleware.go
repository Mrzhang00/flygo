package flygo

import (
	"strings"
)

//Define Middleware interface
type Middleware interface {
	//Set Middleware name
	Name() string
	//Set Middleware method
	Method() string
	//Set Middleware pattern
	Pattern() string
	//Process Middleware
	Process() Handler
	//Set Middleware fields
	Fields() []*Field
}

//Define MiddlewareConfig struct
type MiddlewareConfig struct {
	Name    string
	Method  string
	Pattern string
	Fields  []*Field
}

//Define Filter Middleware interface
type FilterMiddleware interface {
	//Set FilterMiddleware name
	Name() string
	//Set FilterMiddleware type
	Type() string
	//Set FilterMiddleware pattern
	Pattern() string
	//Process FilterMiddleware
	Process() FilterHandler
}

//Define FilterMiddlewareConfig struct
type FilterMiddlewareConfig struct {
	Name    string
	Type    string
	Pattern string
}

//Define Interceptor Middleware interface
type InterceptorMiddleware interface {
	//Set InterceptorMiddleware name
	Name() string
	//Set InterceptorMiddleware type
	Type() string
	//Set InterceptorMiddleware pattern
	Pattern() string
	//Process InterceptorMiddleware
	Process() InterceptorHandler
}

//Define InterceptorMiddlewareConfig struct
type InterceptorMiddlewareConfig struct {
	Name    string
	Type    string
	Pattern string
}

//Using middlewares
func (a *App) Use(middlewares ...Middleware) *App {
	for _, middleware := range middlewares {
		a.middlewares[middleware.Name()] = 0
		a.route(strings.ToUpper(middleware.Name()), middleware.Pattern(), middleware.Process(), middleware.Fields()...)
	}
	return a
}

//Using middlewares with config
func (a *App) UseWithConfig(mwConfig MiddlewareConfig, middleware Middleware) *App {
	if mwConfig.Name == "" {
		mwConfig.Name = middleware.Name()
	}
	if mwConfig.Method == "" {
		mwConfig.Method = middleware.Method()
	}
	if mwConfig.Pattern == "" {
		mwConfig.Pattern = middleware.Pattern()
	}
	if mwConfig.Fields == nil {
		mwConfig.Fields = middleware.Fields()
	}
	a.middlewares[mwConfig.Name] = 0
	a.route(strings.ToUpper(mwConfig.Method), mwConfig.Pattern, middleware.Process(), mwConfig.Fields...)
	return a
}

//Using filter middlewares
func (a *App) UseFilter(filterMiddlewares ...FilterMiddleware) *App {
	for _, middleware := range filterMiddlewares {
		a.filterMiddlewares[middleware.Name()] = 0
		if strings.ToUpper(middleware.Type()) == "BEFORE" {
			a.BeforeFilter(middleware.Pattern(), middleware.Process())
		} else if strings.ToUpper(middleware.Type()) == "AFTER" {
			a.AfterFilter(middleware.Pattern(), middleware.Process())
		}
	}
	return a
}

//Using filter middlewares with config
func (a *App) UseFilterWithConfig(fmwConfig FilterMiddlewareConfig, filterMiddleware FilterMiddleware) *App {
	if fmwConfig.Name == "" {
		fmwConfig.Name = filterMiddleware.Name()
	}
	if fmwConfig.Type == "" {
		fmwConfig.Type = filterMiddleware.Type()
	}
	if fmwConfig.Pattern == "" {
		fmwConfig.Pattern = filterMiddleware.Pattern()
	}
	a.filterMiddlewares[fmwConfig.Name] = 0
	if strings.ToUpper(fmwConfig.Type) == "BEFORE" {
		a.BeforeFilter(fmwConfig.Pattern, filterMiddleware.Process())
	} else if strings.ToUpper(fmwConfig.Type) == "AFTER" {
		a.AfterFilter(fmwConfig.Pattern, filterMiddleware.Process())
	}
	return a
}

//Using interceptor middlewares
func (a *App) UseInterceptor(interceptorMiddlewares ...InterceptorMiddleware) *App {
	for _, middleware := range interceptorMiddlewares {
		a.interceptorMiddlewares[middleware.Name()] = 0
		if strings.ToUpper(middleware.Type()) == "BEFORE" {
			a.BeforeInterceptor(middleware.Pattern(), middleware.Process())
		} else if strings.ToUpper(middleware.Type()) == "AFTER" {
			a.AfterInterceptor(middleware.Pattern(), middleware.Process())
		}
	}
	return a
}

//Using interceptor middlewares with config
func (a *App) UseInterceptorWithConfig(imwConfig InterceptorMiddlewareConfig, interceptorMiddleware InterceptorMiddleware) *App {
	if imwConfig.Name == "" {
		imwConfig.Name = interceptorMiddleware.Name()
	}
	if imwConfig.Type == "" {
		imwConfig.Type = interceptorMiddleware.Type()
	}
	if imwConfig.Pattern == "" {
		imwConfig.Pattern = interceptorMiddleware.Pattern()
	}
	a.filterMiddlewares[imwConfig.Name] = 0
	if strings.ToUpper(imwConfig.Type) == "BEFORE" {
		a.BeforeInterceptor(imwConfig.Pattern, interceptorMiddleware.Process())
	} else if strings.ToUpper(imwConfig.Type) == "AFTER" {
		a.AfterInterceptor(imwConfig.Pattern, interceptorMiddleware.Process())
	}
	return a
}

//Print middlewares
func (a *App) printMiddleware() {
	for name := range a.middlewares {
		a.LogInfo("Middleware detected [%s]", name)
	}
	for name := range a.filterMiddlewares {
		a.LogInfo("FilterMiddleware detected [%s]", name)
	}
	for name := range a.interceptorMiddlewares {
		a.LogInfo("InterceptorMiddleware detected [%s]", name)
	}
}

//Init middleware c
func (a *App) initMiddlewareCtx() map[string]map[string]interface{} {
	middlewareCtx := make(map[string]map[string]interface{}, 0)
	for name := range a.middlewares {
		middlewareCtx[name] = make(map[string]interface{}, 0)
	}
	for name := range a.filterMiddlewares {
		middlewareCtx[name] = make(map[string]interface{}, 0)
	}
	for name := range a.interceptorMiddlewares {
		middlewareCtx[name] = make(map[string]interface{}, 0)
	}
	return middlewareCtx
}
