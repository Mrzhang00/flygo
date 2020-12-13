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
	return mw.TypeBefore
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
	//app.GET("/set", func(c *Context) {
	//	sess := mw.GetSession(c)
	//	sess.Set("name", "helloworld")
	//	c.Write([]byte("set done"))
	//})
	//app.GET("/get", func(c *Context) {
	//	sess := mw.GetSession(c)
	//	c.Write([]byte("get " + sess.Get("name").(string)))
	//})
	//app.REST(&HelloController{})
	//app.Use(&MyMW{})
	//app.UseSession(redis.Provider(
	//	&Options{Password: "123"}),
	//	&se.Config{Timeout: time.Second * 20},
	//	&se.Listener{
	//		Created: func(s se.Session) {
	//			log.Println("Created")
	//		},
	//
	//		Refreshed: func(s se.Session) {
	//			log.Println("Refreshed")
	//		},
	//
	//		Invalidated: func(s se.Session) {
	//			log.Println("Invalidated")
	//		},
	//
	//		Destoryed: func(s se.Session) {
	//			log.Println("Destoryed")
	//		},
	//	})
	//app.UseRecovered()
	app.UseNotFound()
	app.Config.Debug = true
	//app.ConfigFile = `/Users/local/Desktop/Workspaces/Goland/src/github.com/billcoding/flygo/main/flygo.yml`
	app.Run()
}
