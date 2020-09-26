package flygo

import (
	"fmt"
	"testing"
)

//Test before filter
func TestBeforeFilter(t *testing.T) {

	app := NewApp()

	app.Get("/before/index", func(c *Context) {
		c.Text("index")
	})

	app.BeforeFilter("/before/*", func(c *FilterContext) {
		fmt.Println("before filter")
	})

	app.Run()

}

//Test after filter
func TestAfterFilter(t *testing.T) {

	app := NewApp()

	app.Get("/after/index/after", func(c *Context) {
		c.Text("index")
	})

	app.AfterFilter("/**", func(c *FilterContext) {
		fmt.Println("after filter")
	})

	app.Run()
}
