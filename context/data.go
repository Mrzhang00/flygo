package context

import "github.com/billcoding/calls"

func (c *Context) SetData(name string, value interface{}) *Context {
	c.dataMap[name] = value
	return c
}

func (c *Context) SetDataMap(dmap map[string]interface{}) *Context {
	calls.NNil(dmap, func() {
		for k, v := range dmap {
			c.SetData(k, v)
		}
	})
	return c
}

func (c *Context) GetData(name string) interface{} {
	return c.dataMap[name]
}
