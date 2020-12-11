package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	mw "github.com/billcoding/flygo/middleware"
)

type HelloController struct {
}

func (h *HelloController) Prefix() string {
	return "/Hello"
}

func (h *HelloController) GET() func(c *Context) {
	return func(c *Context) {
		fmt.Println(c.RestId())
		c.Write([]byte("get =>"))
	}
}

func (h *HelloController) GETS() func(c *Context) {
	return func(c *Context) {
		c.Write([]byte("gets"))
	}
}

func (h *HelloController) POST() func(c *Context) {
	return func(c *Context) {
		c.Write([]byte("post"))
	}
}

func (h *HelloController) PUT() func(c *Context) {
	return func(c *Context) {
		c.Write([]byte("put"))
	}
}

func (h *HelloController) DELETE() func(c *Context) {
	return func(c *Context) {
		c.Write([]byte("delete"))
	}
}

type MyMW struct {
}

func (m *MyMW) Type() *mw.Type {
	return mw.TypeAfter
}

func (m *MyMW) Name() string {
	return "MyMW"
}

func (m *MyMW) Method() mw.Method {
	return mw.MethodAny
}

func (m *MyMW) Pattern() mw.Pattern {
	return mw.Pattern("/index/*/index")
}

func (m *MyMW) Handler() func(c *Context) {
	return func(c *Context) {
		fmt.Println("MyMW........")
		c.Chain()
	}
}

func main() {
	app := flygo.GetApp()
	app.GET("/index", func(c *Context) {
		c.Write([]byte("index"))
	})
	app.REST(&HelloController{})
	app.Use(&MyMW{})
	app.Config.Debug = true
	app.ConfigFile = `/Users/local/Desktop/Workspaces/Goland/src/github.com/billcoding/flygo/main/flygo.yml`
	app.Run()
}
