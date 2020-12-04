package flygo

import (
	"testing"
)

//Test Redirect
func TestRedirect(t *testing.T) {
	GetApp().Get("/", func(c *Context) {
		c.Redirect("/2")
	}).Get("/2", func(c *Context) {
		c.Text("2222")
	}).Run()
}
