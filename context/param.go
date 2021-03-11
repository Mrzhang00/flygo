package context

import (
	"fmt"
	"strconv"
	"strings"
)

// ParamMap return param map
func (ctx *Context) ParamMap() map[string][]string {
	return ctx.paramMap
}

// Param return named param
func (ctx *Context) Param(name string) string {
	return ctx.ParamWith(name, "")
}

// ParamInt return named param of int
func (ctx *Context) ParamInt(name string) int64 {
	val := ctx.Param(name)
	iv, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}
	return iv
}

// ParamIntWith return named param of int
func (ctx *Context) ParamIntWith(name string, defaultVal int) int {
	val := ctx.ParamWith(name, fmt.Sprintf("%d", defaultVal))
	iv, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return iv
}

// ParamFloat return named param of float
func (ctx *Context) ParamFloat(name string) float64 {
	val := ctx.Param(name)
	iv, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}
	return iv
}

// ParamFloatWith return named param of float
func (ctx *Context) ParamFloatWith(name string, defaultVal float64) float64 {
	fval := ctx.ParamWith(name, fmt.Sprintf("%f", defaultVal))
	fv, err := strconv.ParseFloat(fval, 64)
	if err != nil {
		panic(err)
	}
	return fv
}

// Params return named param values
func (ctx *Context) Params(name string) []string {
	return ctx.ParamsWith(name, nil)
}

// ParamWith return named param with default value
func (ctx *Context) ParamWith(name, defaultValue string) string {
	vals := ctx.paramMap[name]
	if vals != nil && len(vals) > 0 && strings.TrimSpace(vals[0]) != "" {
		return strings.TrimSpace(vals[0])
	}
	return defaultValue
}

// ParamsWith return named param with default value
func (ctx *Context) ParamsWith(name string, defaultValue []string) []string {
	vals := ctx.paramMap[name]
	if vals != nil {
		return vals
	}
	return defaultValue
}

func (ctx *Context) transformParamMap(multiFunc func(vals []string) string) map[string]string {
	sm := make(map[string]string, 0)
	for k, v := range ctx.paramMap {
		sm[k] = multiFunc(v)
	}
	return sm
}

// SingleParamMap return single value param map
func (ctx *Context) SingleParamMap() map[string]string {
	return ctx.transformParamMap(func(vals []string) string {
		if vals != nil && len(vals) > 0 {
			return vals[0]
		}
		return ""
	})
}

// JoinedParamMap return joined single value param map
func (ctx *Context) JoinedParamMap(separator string) map[string]string {
	return ctx.transformParamMap(func(vals []string) string {
		if vals != nil && len(vals) > 0 {
			return strings.Join(vals, separator)
		}
		return ""
	})
}

// SetParamMap set param map
func (ctx *Context) SetParamMap(paramMap map[string][]string) *Context {
	if paramMap != nil && len(paramMap) > 0 {
		for k, v := range paramMap {
			ctx.paramMap[k] = v
		}
	}
	return ctx
}

// RESTID return RESTful ID
func (ctx *Context) RESTID() string {
	return ctx.Param("RESTFUL_ID")
}

// RESTIntID return int RESTful ID
func (ctx *Context) RESTIntID() int64 {
	intID, err := strconv.ParseInt(ctx.RESTID(), 10, 64)
	if err != nil {
		panic(err)
	}
	return intID
}

// RESTStringID return RESTful ID
func (ctx *Context) RESTStringID() string {
	return ctx.Param("RESTFUL_ID")
}
