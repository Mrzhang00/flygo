package router

import (
	"fmt"
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"regexp"
	"strings"
)

// ParsedRouter struct
type ParsedRouter struct {
	// Simples routers
	Simples map[string]*Simple
	// Dynamics routers
	Dynamics map[string]map[string]*Dynamic
}

// Handler found simple or dynamic handler
func (pr *ParsedRouter) Handler(ctx *context.Context) func(ctx *context.Context) {
	simple := pr.simple(ctx)
	if simple != nil {
		return simple
	}
	return pr.dynamic(ctx)
}

func (pr *ParsedRouter) simple(ctx *context.Context) func(ctx *context.Context) {
	if len(pr.Simples) <= 0 {
		return nil
	}
	routeKey := fmt.Sprintf("%s:%s", ctx.Request.Method, util.TrimLeftAndRight(ctx.Request.URL.Path))
	simple, have := pr.Simples[routeKey]
	if have {
		return func(ctx *context.Context) {
			simple.Handler(ctx)
			ctx.Chain()
		}
	}
	return nil
}

func (pr *ParsedRouter) dynamic(ctx *context.Context) func(ctx *context.Context) {
	if len(pr.Dynamics) <= 0 {
		return nil
	}
	for pattern, mp := range pr.Dynamics {
		re := regexp.MustCompile(pattern)
		reqPath := util.TrimLeftAndRight(ctx.Request.URL.Path)
		matched := re.MatchString(reqPath)
		if matched {
			dy, routed := mp[ctx.Request.Method]
			if routed {
				return func(ctx *context.Context) {
					paramMap := make(map[string][]string, 0)
					paths := strings.Split(reqPath, "/")
					for i, paramVal := range paths {
						paramName, have := dy.Pos[i]
						if have {
							paramMap[paramName] = []string{paramVal}
						}
					}
					ctx.SetParamMap(paramMap)
					dy.Handler(ctx)
					ctx.Chain()
				}
			}
		}
	}
	return nil
}
