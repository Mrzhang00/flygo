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
	return ctx.ParamDefault(name, "")
}

// HasParam return true if named param in param map
func (ctx *Context) HasParam(name string) bool {
	return ctx.ParamMap()[name] != nil
}

// Int return named param of int
func (ctx *Context) Int(name string) (int64, error) {
	val := ctx.Param(name)
	return strconv.ParseInt(val, 10, 64)
}

// IntDefault return named param of int with default `defaultVal`
func (ctx *Context) IntDefault(name string, defaultVal int) int {
	val := ctx.ParamDefault(name, fmt.Sprintf("%d", defaultVal))
	iv, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return iv
}

// Float return named param of float
func (ctx *Context) Float(name string) float64 {
	val := ctx.Param(name)
	iv, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}
	return iv
}

// FloatDefault return named param of float
func (ctx *Context) FloatDefault(name string, defaultVal float64) float64 {
	val := ctx.ParamDefault(name, fmt.Sprintf("%f", defaultVal))
	fv, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}
	return fv
}

// Params return named param values
func (ctx *Context) Params(name string) []string {
	return ctx.ParamsDefault(name, nil)
}

// ParamDefault return named param with default value
func (ctx *Context) ParamDefault(name, defaultVal string) string {
	params := ctx.paramMap[name]
	//FIX: when param is empty
	if params == nil || len(params) <= 0 || params[0] == "" {
		return defaultVal
	}
	return params[0]
}

// ParamsDefault return named param with default value
func (ctx *Context) ParamsDefault(name string, defaultVal []string) []string {
	params := ctx.paramMap[name]
	if params != nil {
		return params
	}
	return defaultVal
}

func (ctx *Context) transformParamMap(multiFunc func(name string, params []string) string) map[string]string {
	sm := make(map[string]string, 0)
	for k, v := range ctx.paramMap {
		sm[k] = multiFunc(k, v)
	}
	return sm
}

// SingleParamMap return single value param map
func (ctx *Context) SingleParamMap() map[string]string {
	return ctx.transformParamMap(func(name string, params []string) string {
		if params != nil && len(params) > 0 {
			return params[0]
		}
		return ""
	})
}

// JoinedParamMap return joined single value param map
func (ctx *Context) JoinedParamMap(separator string) map[string]string {
	return ctx.transformParamMap(func(name string, params []string) string {
		if params != nil && len(params) > 0 {
			return strings.Join(params, separator)
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

// Key return REST-ful key
func (ctx *Context) Key() string {
	return ctx.Param("RESTFUL_KEY")
}

// IntKey return int REST-ful key
func (ctx *Context) IntKey() int64 {
	intID, err := strconv.ParseInt(ctx.Key(), 10, 64)
	if err != nil {
		panic(err)
	}
	return intID
}
