package flygo

import "testing"

//Test config c path
func TestConfigContextPath(t *testing.T) {

	app := NewApp()

	//Config c path
	app.SetContextPath("/flygo")

	app.Get("/", func(c *Context) {
		c.Text("index")
	})

	app.Run()

}

//Test config web root
func TestConfigWebRoot(t *testing.T) {

	app := NewApp()

	//Config app webroot
	app.SetWebRoot("d:\\webroot")

	app.SetStaticEnable(true)

	app.Run()

}

//Test config static
func TestConfigStatic(t *testing.T) {

	app := NewApp()

	app.SetWebRoot("D:\\webroot")

	//Config enable static
	app.SetStaticEnable(true)

	//Config enable static cache
	app.SetStaticCache(false)

	//Config static pattern, default `/static`
	app.SetStaticPattern("/staticres")

	//Config static prefix, default `static`
	app.SetStaticPrefix("staticres")

	app.Run()

}

//Test config view
func TestConfigView(t *testing.T) {

	app := NewApp()

	app.SetWebRoot("D:\\webroot")

	//Config enable view cache
	app.SetViewCache(false)

	//Config view prefix,default `templates`
	app.SetViewPrefix("tpls")

	//Config static pattern, default `html`
	app.SetViewSuffix("tpl")

	app.Get("/", func(c *Context) {
		c.View("index")
	})

	app.Run()

}
