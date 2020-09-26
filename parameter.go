package flygo

import (
	"strings"
)

//Get request parameter
func (c *Context) GetParameter(name string) string {
	return c.GetParameterWithDefault(name, "")
}

//Get request parameters
func (c *Context) GetParameters(name string) []string {
	return c.GetParametersWithDefault(name, nil)
}

//Get request parameter with default
func (c *Context) GetParameterWithDefault(name, defaultValue string) string {
	vals := c.Parameters[name]
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
func (c *Context) GetParametersWithDefault(name string, defaultValue []string) []string {
	vals := c.Parameters[name]
	if vals != nil {
		return vals
	}
	return defaultValue
}
