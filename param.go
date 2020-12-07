package flygo

import (
	"strings"
)

//Get request parameter
func (c *Context) Param(name string) string {
	return c.ParamWith(name, "")
}

//Get request parameters
func (c *Context) Params(name string) []string {
	return c.ParamsWith(name, nil)
}

//Get request parameter with default
func (c *Context) ParamWith(name, defaultValue string) string {
	vals := c.ParamMap[name]
	if vals != nil {
		if len(vals) == 0 || (len(vals) > 0 && vals[0] == "") {
			return defaultValue
		} else {
			return strings.TrimSpace(vals[0])
		}
	}
	return defaultValue
}

//Get request parameters with default
func (c *Context) ParamsWith(name string, defaultValue []string) []string {
	vals := c.ParamMap[name]
	if vals != nil {
		return vals
	}
	return defaultValue
}