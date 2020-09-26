package flygo

import "testing"

//Test all route
func TestAllRoute(t *testing.T) {
	NewApp().Req("/req", func(c *Context) {
		c.Text("req")
	}).Run()
}

//Test get route
func TestGetRoute(t *testing.T) {
	NewApp().Get("/get", func(c *Context) {
		c.Text("get")
	}).Run()
}

//Test post route
func TestPostRoute(t *testing.T) {
	NewApp().Post("/post", func(c *Context) {
		c.Text("post")
	}).Run()
}

//Test delete route
func TestDeleteRoute(t *testing.T) {
	NewApp().Delete("/delete", func(c *Context) {
		c.Text("delete")
	}).Run()
}

//Test put route
func TestPutRoute(t *testing.T) {
	NewApp().Put("/put", func(c *Context) {
		c.Text("put")
	}).Run()
}

//Test patch route
func TestPatchRoute(t *testing.T) {
	NewApp().Patch("/patch", func(c *Context) {
		c.Text("patch")
	}).Run()
}

func TestFix(t *testing.T) {
	h := func(c *Context) {}
	app := NewApp().SetContextPath("/go")
	app.Get("/shop/account", h)
	app.Post("/shop/account", h)
	app.Delete("/shop/account", h)
	app.Put("/shop/account", h)
	app.Run()
}
