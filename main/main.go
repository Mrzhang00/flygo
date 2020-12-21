package main

import (
	"fmt"
	"github.com/billcoding/calls"
	"github.com/billcoding/flygo"
	"github.com/billcoding/flygo/binding"
	. "github.com/billcoding/flygo/context"
	mw "github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/validator"
	"reflect"
)

type HelloController struct {
}

func (h *HelloController) Prefix() string {
	return "/Hello"
}

func (h *HelloController) GET() func(c *Context) {
	return func(c *Context) {
		fmt.Println(c.RestId())
		c.Text("get =>")
	}
}

func (h *HelloController) GETS() func(c *Context) {
	return func(c *Context) {
		c.Text("gets")
	}
}

func (h *HelloController) POST() func(c *Context) {
	return func(c *Context) {
		c.Text("post")
	}
}

func (h *HelloController) PUT() func(c *Context) {
	return func(c *Context) {
		c.Text("put")
	}
}

func (h *HelloController) DELETE() func(c *Context) {
	return func(c *Context) {
		panic("helloworld")
		c.Text("delete")
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
	return mw.MethodGet
}

func (m *MyMW) Pattern() mw.Pattern {
	return "/index/*/index"
}

func (m *MyMW) Handler() func(c *Context) {
	return func(c *Context) {
		fmt.Println("MyMW........")
		c.Chain()
	}
}

func main() {
	type model struct {
		Id1  uint8  `binding:"name(id) default(100)" validate:"required(T) min(100) message(Id1不合法)"`
		Id2  uint16 `binding:"name(id) default(100)" validate:"required(T) min(100) message(Id2不合法)"`
		Id3  uint32 `binding:"name(id) default(100)" validate:"required(T) min(100) message(Id3不合法)"`
		Id4  uint   `binding:"name(id) default(100)" validate:"required(T) min(100) message(Id4不合法)"`
		Id5  uint64 `binding:"name(id) default(100)" validate:"required(T) min(100) message(Id5不合法)"`
		Name string `binding:"name(name) default(zhangsan)" validate:"required(T) maxlength(5) message(Name不合法)"`
	}
	aa := func(c *Context) {
		m := model{}
		fmt.Println(reflect.TypeOf(m).Name())
		binding.New(&m).Bind(c.Request)
		result := validator.New(&m).Validate()
		calls.False(result.Passed, func() {
			c.Text(result.Messages())
		})
		calls.True(result.Passed, func() {
			c.JSON(&m)
		})
	}
	app := flygo.GetApp()
	app.GET("/", aa)
	//app.GET("/set", func(c *Context) {
	//	sess := mw.GetSession(c)
	//	sess.Set("name", "helloworld")
	//	c.Text([]byte("set done"))
	//})
	//app.GET("/get", func(c *Context) {
	//	sess := mw.GetSession(c)
	//	c.Text([]byte("get " + sess.Get("name").(string)))
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
	//app.UseRecovery()
	//app.UseMethodNotAllowed()
	//app.UseNotFound()
	//app.UseStdLogger()
	//app.Config.Debug = true
	//app.ConfigFile = `/Users/local/Desktop/Workspaces/Goland/src/github.com/billcoding/flygo/main/flygo.yml`
	app.Run()
}
