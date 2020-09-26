package flygo

//Set app cache
func (c *FilterContext) SetCache(key string, data interface{}) {
	c.context.SetCache(key, data)
}

//Set app cache
func (c *Context) SetCache(key string, data interface{}) {
	app.SetCache(key, data)
}

//Get app cache
func (c *FilterContext) GetCache(key string) interface{} {
	return c.context.GetCache(key)
}

//Get app cache
func (c *Context) GetCache(key string) interface{} {
	return app.GetCache(key)
}

//Remove app cache
func (c *FilterContext) RemoveCache(key string) {
	c.context.RemoveCache(key)
}

//Remove app cache
func (c *Context) RemoveCache(key string) {
	app.RemoveCache(key)
}

//Clear app cache
func (c *FilterContext) ClearCaches() {
	c.context.ClearCaches()
}

//Clear app cache
func (c *Context) ClearCaches() {
	app.ClearCaches()
}

//Get data from app storage
func (a *App) GetCache(key string) interface{} {
	return a.caches[key]
}

//Set data into app storage
func (a *App) SetCache(key string, data interface{}) {
	a.caches[key] = data
}

//remove data from app storage
func (a *App) RemoveCache(key string) {
	delete(a.caches, key)
}

//clear caches
func (a *App) ClearCaches() {
	a.caches = map[string]interface{}{}
}

//get all caches
func (a *App) GetAllCaches() map[string]interface{} {
	return a.caches
}
