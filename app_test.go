package flygo

import (
	"testing"
)

var app2 = NewApp()

func TestHelloworld(t *testing.T) {
	app2.BeforeInterceptor("/**", func(c *Context) {
		c.JSON(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{Code: 1,
			Msg: "hellowofffrld"})
	})
	app2.Get("/", func(c *Context) {
		c.Text("index")
	}).Run()
}
