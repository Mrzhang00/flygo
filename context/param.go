package context

import (
	"fmt"
	"strconv"
	"strings"
)

// ParamMap return param map
func (c *Context) ParamMap() map[string][]string {
	return c.paramMap
}

// Param return named param
func (c *Context) Param(name string) string {
	return c.ParamWith(name, "")
}

// ParamInt return named param of int
func (c *Context) ParamInt(name string) int {
	return c.ParamIntWith(name, 0)
}

// ParamIntWith return named param of int
func (c *Context) ParamIntWith(name string, defaultVal int) int {
	val := c.ParamWith(name, fmt.Sprintf("%d", defaultVal))
	iv, err := strconv.Atoi(val)
	if err != nil {
		c.logger.Error("[Param]%v", err)
	}
	return iv
}

// ParamFloat return named param of float
func (c *Context) ParamFloat(name string) float64 {
	return c.ParamFloatWith(name, 0)
}

// ParamFloatWith return named param of float
func (c *Context) ParamFloatWith(name string, defaultVal float64) float64 {
	fval := c.ParamWith(name, fmt.Sprintf("%f", defaultVal))
	fv, err := strconv.ParseFloat(fval, 64)
	if err != nil {
		c.logger.Error("[Param]%v", err)
	}
	return fv
}

// Params return named param values
func (c *Context) Params(name string) []string {
	return c.ParamsWith(name, nil)
}

// ParamWith return named param with default value
func (c *Context) ParamWith(name, defaultValue string) string {
	vals := c.paramMap[name]
	if vals != nil {
		if len(vals) != 0 && !(len(vals) > 0 && vals[0] == "") {
			return strings.TrimSpace(vals[0])
		} else {
			return defaultValue
		}
	}
	return defaultValue
}

// ParamsWith return named param with default value
func (c *Context) ParamsWith(name string, defaultValue []string) []string {
	vals := c.paramMap[name]
	if vals != nil {
		return vals
	}
	return defaultValue
}

func (c *Context) transformParamMap(multiFunc func(vals []string) string) map[string]string {
	sm := make(map[string]string, 0)
	for k, v := range c.paramMap {
		sm[k] = multiFunc(v)
	}
	return sm
}

// SingleParamMap return single value param map
func (c *Context) SingleParamMap() map[string]string {
	return c.transformParamMap(func(vals []string) string {
		if vals != nil && len(vals) > 0 {
			return vals[0]
		}
		return ""
	})
}

// JoinedParamMap return joined single value param map
func (c *Context) JoinedParamMap(separator string) map[string]string {
	return c.transformParamMap(func(vals []string) string {
		if vals != nil && len(vals) > 0 {
			return strings.Join(vals, separator)
		}
		return ""
	})
}

// SetParamMap set param map
func (c *Context) SetParamMap(paramMap map[string][]string) *Context {
	if paramMap != nil && len(paramMap) > 0 {
		for k, v := range paramMap {
			c.paramMap[k] = v
		}
	}
	return c
}

// RestId return RESTful ID
func (c *Context) RestId() string {
	return c.Param("RESTFUL_ID")
}
