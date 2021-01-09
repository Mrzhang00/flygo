package context

import "github.com/billcoding/calls"

// SetData Set data into context
func (c *Context) SetData(name string, value interface{}) *Context {
	c.dataMap[name] = value
	return c
}

// SetDataMap Set data map into context
func (c *Context) SetDataMap(dmap map[string]interface{}) *Context {
	calls.NNil(dmap, func() {
		for k, v := range dmap {
			c.SetData(k, v)
		}
	})
	return c
}

// GetData get data from context
func (c *Context) GetData(name string) interface{} {
	return c.dataMap[name]
}
