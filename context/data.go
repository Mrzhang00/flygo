package context

// SetData Set data into context
func (ctx *Context) SetData(name string, value interface{}) *Context {
	ctx.dataMap[name] = value
	return ctx
}

// SetDataMap Set data map into context
func (ctx *Context) SetDataMap(dmap map[string]interface{}) *Context {
	if dmap != nil {
		for k, v := range dmap {
			ctx.SetData(k, v)
		}
	}
	return ctx
}

// GetData get data from context
func (ctx *Context) GetData(name string) interface{} {
	return ctx.dataMap[name]
}
