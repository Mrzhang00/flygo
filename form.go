package flygo

//Try parse form, put to c
func (c *Context) parseForm() {
	err := c.Request.ParseForm()
	if err != nil {
		return
	}
	for name, values := range c.Request.Form {
		c.ParamMap[name] = values
	}
}
