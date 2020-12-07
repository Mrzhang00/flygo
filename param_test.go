package flygo

import (
	"fmt"
	"testing"
)

//Test param
func TestParam(t *testing.T) {
	NewApp().Get("/", func(c *Context) {
		id := c.Param("id")
		name := c.Param("name")
		vals := c.Params("vals")
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(vals)
	}).Run()
}