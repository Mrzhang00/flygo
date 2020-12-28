package router

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	"regexp"
	"strings"
)

//Define ParsedRouter struct
type ParsedRouter struct {
	//K<routeKey> V<Simple>
	Simples map[string]*Simple
	//K<regex> V< K<method> V<Dynamic> >
	Dynamics map[string]map[string]*Dynamic
}

//Handlers
func (pr *ParsedRouter) Handlers(ctx *c.Context) func(c *c.Context) {
	simple := pr.simple(ctx)
	if simple != nil {
		return simple
	}
	return pr.dynamic(ctx)
}

//simple
func (pr *ParsedRouter) simple(ctx *c.Context) func(c *c.Context) {
	if len(pr.Simples) <= 0 {
		return nil
	}
	routeKey := fmt.Sprintf("%s:%s", ctx.Request.Method, ctx.Request.URL.Path)
	simple, have := pr.Simples[routeKey]
	if have {
		return func(c *c.Context) {
			simple.Handler(c)
			ctx.Chain()
		}
	}
	return nil
}

//dynamic
func (pr *ParsedRouter) dynamic(ctx *c.Context) func(c *c.Context) {
	if len(pr.Dynamics) <= 0 {
		return nil
	}
	for pattern, mp := range pr.Dynamics {
		re := regexp.MustCompile(pattern)
		matched := re.MatchString(ctx.Request.URL.Path)
		if matched {
			dy, routed := mp[ctx.Request.Method]
			if routed {
				return func(c *c.Context) {

					//Add params into context
					paramMap := make(map[string][]string, 0)
					paths := strings.Split(c.Request.URL.Path, "/")
					for i, paramVal := range paths {
						paramName, have := dy.Pos[i]
						if have {
							//found match variable
							paramMap[paramName] = []string{paramVal}
						}
					}

					//Set paramMap
					c.SetParamMap(paramMap)

					//Add handler
					dy.Handler(c)

					//Next Chain
					c.Chain()
				}
			}
		}
	}
	return nil
}
