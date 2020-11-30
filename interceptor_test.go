package flygo

import "testing"

//Test before interceptor
func TestBeforeInteceptor(t *testing.T) {
	app := NewApp()

	app.BeforeInterceptor("/before/*", func(c *Context) {
		if c.Param("id") == "" {
			c.Text("before interceptor")
		}
	})

	app.Get("/before/helloworld", func(c *Context) {
		c.Text("helloworld")
	})

	app.Run()
}

//Test after interceptor
func TestAfterInteceptor(t *testing.T) {
	app := NewApp()

	app.AfterInterceptor("/**", func(c *Context) {
		c.Response.SetDone(false)
		c.Text("after interceptor")
	})

	app.Get("/after/helloworld/index", func(c *Context) {
		c.Text("helloworld")
	})

	app.Run()
}
