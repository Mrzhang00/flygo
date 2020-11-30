package flygo

import "testing"

func TestView(t *testing.T) {
	app := NewApp()
	app.Get("/v", func(c *Context) {
		c.Session.Set("name", "zhansgan")
		c.Session.Set("age", "123")
		c.View("index")
	})
	app.Caches["globalName"] = "sdfdsfdsfdsf"
	app.Run()
}
