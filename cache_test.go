package flygo

import (
	"fmt"
	"testing"
)

//Test cache
func TestCache(t *testing.T) {

	app := NewApp()

	//Set cache
	app.SetCache("abc", "This is abc val")

	app.Get("/get", func(c *Context) {
		//Get abc cache val
		c.Text(fmt.Sprintf("%v", c.GetCache("abc")))
	})

	app.BeforeFilter("/**", func(c *FilterContext) {
		fmt.Println(c.GetCache("abc"))
	})

	app.Run()

}
