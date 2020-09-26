package flygo

import (
	"fmt"
	"testing"
)

//Test parameter
func TestParameter(t *testing.T) {
	NewApp().Get("/", func(c *Context) {
		id := c.GetParameter("id")
		name := c.GetParameter("name")
		vals := c.GetParameters("vals")
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(vals)
	}).Run()

}
