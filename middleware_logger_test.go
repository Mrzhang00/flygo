package flygo

import (
	"testing"
)

//Test logger
func TestLogger(t *testing.T) {
	app := NewApp()

	app.UseFilter(Logger)

	app.Get("/", func(c *Context) {
		c.Text("index")
	})

	app.Run()
}
