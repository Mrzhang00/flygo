package flygo

import "testing"

//Test handler
func TestHandler(t *testing.T) {

	app := NewApp()

	//Config default 404 not found
	app.DefaultHandler(func(c *Context) {
		c.Text("404")
	})

	//Config request not supported handler
	app.RequestNotSupportedHandler(func(c *Context) {
		c.Text("No impls")
	})

	app.Post("/post", func(c *Context) {
		c.Text("post")
	})

	app.Run()

}
