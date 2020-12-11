package context

import (
	"fmt"
	"strconv"
	"strings"
)

//ParamMap
func (c *Context) ParamMap() map[string][]string {
	return c.paramMap
}

//Param
func (c *Context) Param(name string) string {
	return c.ParamWith(name, "")
}

//ParamInt
func (c *Context) ParamInt(name string) int {
	return c.ParamIntWith(name, 0)
}

//ParamIntWith
func (c *Context) ParamIntWith(name string, defaultVal int) int {
	val := c.ParamWith(name, fmt.Sprintf("%d", defaultVal))
	iv, err := strconv.Atoi(val)
	if err != nil {
		c.logger.Error("[Param]%v", err)
	}
	return iv
}

//ParamFloat
func (c *Context) ParamFloat(name string) float64 {
	return c.ParamFloatWith(name, 0)
}

//ParamFloatWith
func (c *Context) ParamFloatWith(name string, defaultVal float64) float64 {
	fval := c.ParamWith(name, fmt.Sprintf("%f", defaultVal))
	fv, err := strconv.ParseFloat(fval, 64)
	if err != nil {
		c.logger.Error("[Param]%v", err)
	}
	return fv
}

//Get request params
func (c *Context) Params(name string) []string {
	return c.ParamsWith(name, nil)
}

//Get request parameter with default
func (c *Context) ParamWith(name, defaultValue string) string {
	vals := c.paramMap[name]
	if vals != nil {
		if len(vals) == 0 || (len(vals) > 0 && vals[0] == "") {
			return defaultValue
		} else {
			return strings.TrimSpace(vals[0])
		}
	}
	return defaultValue
}

//ParamsWith
func (c *Context) ParamsWith(name string, defaultValue []string) []string {
	vals := c.paramMap[name]
	if vals != nil {
		return vals
	}
	return defaultValue
}

//transformParamMap
func (c *Context) transformParamMap(multiFunc func(vals []string) string) map[string]string {
	sm := make(map[string]string, 0)
	for k, v := range c.paramMap {
		sm[k] = multiFunc(v)
	}
	return sm
}

//SingleParamMap
func (c *Context) SingleParamMap() map[string]string {
	return c.transformParamMap(func(vals []string) string {
		if vals != nil && len(vals) > 0 {
			return vals[0]
		}
		return ""
	})
}

//JoinedParamMap
func (c *Context) JoinedParamMap(separator string) map[string]string {
	return c.transformParamMap(func(vals []string) string {
		if vals != nil && len(vals) > 0 {
			return strings.Join(vals, separator)
		}
		return ""
	})
}

//SetParamMap
func (c *Context) SetParamMap(paramMap map[string][]string) *Context {
	if paramMap != nil && len(paramMap) > 0 {
		for k, v := range paramMap {
			c.paramMap[k] = v
		}
	}
	return c
}

//RestId
func (c *Context) RestId() string {
	return c.Param("RESTFUL_ID")
}
