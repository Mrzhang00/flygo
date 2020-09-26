package flygo

import "testing"

func TestView(t *testing.T) {
	app := NewApp()
	//app.SetWebRoot(`E:\Workspaces\Goland\src\gitee.com\billcoding\flygo\main`)
	app.SetViewEnable(true)
	app.SetViewCache(false)
	app.Get("/v", func(c *Context) {
		c.Session.Set("name", "zhansgan")
		c.Session.Set("age", "123")
		c.View("index")
	})
	app.SetCache("globalName", "sdfdsfdsfdsf")
	app.Run()
}
