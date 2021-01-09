package main

import (
	"fmt"
	bind "github.com/billcoding/binding"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	mw "github.com/billcoding/flygo/middleware"
	se "github.com/billcoding/flygo/session"
	"github.com/billcoding/flygo/session/memory"
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

type model struct {
	Id1    uint8      `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id1不合法)"`
	Id2    uint16     `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id2不合法)"`
	Id3    uint32     `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id3不合法)"`
	Id4    uint       `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id4不合法)"`
	Id5    uint64     `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id5不合法)"`
	Float1 float32    `binding:"name(fl)" validate:"required(T) min(10) max(1000) message(Float1不合法)"`
	Float2 float64    `binding:"name(fl)" validate:"required(T) min(10) max(1000) message(Float2不合法)"`
	Name   string     `binding:"name(name)" validate:"required(T) minlength(2) maxlength(5) message(Name不合法)"`
	In     InnerModel `binding:"name(in)" validate:"required(T)"`
}

type InnerModel struct {
	Id   uint8  `binding:"name(id)" validate:"required(T) min(10) max(1000) message(in.Id1不合法)"`
	Name string `binding:"name(name)" validate:"required(T) minlength(2) maxlength(5) message(in.Name不合法)"`
}

func handler(c *Context) {
	m := model{}
	c.BindWithParamsAndValidate(&m, bind.Param, func() {
		c.JSON(&m)
	})
}

func handler2(c *Context) {
	m := model{}
	c.BindWithParamsAndValidate(&m, bind.Param, func() {
		c.JSON(&m)
	})
}

func main() {
	app := flygo.GetApp()
	app.UseRecovery()
	app.GET("/xx", handler)
	app.POST("/xx", handler2)
	app.POST("/file", func(c *Context) {
		c.ParseMultipart(0)
		file := c.MultipartFile("file")
		file.Copy(`/Users/local/Desktop/Workspaces/Goland/src/github.com/billcoding/flygo/main/asfdsfdsfds.txt`)
		c.Text("success")
	})
	app.UseSession(memory.Provider(), &se.Config{Timeout: 60}, nil)
	//	c.Text(fmt.Sprintf("%v", mw.GetSession(c).GetAll()))

	//	mw.GetSession(c).Set("xxx", "dsfdsfds")
	//	mw.GetSession(c).Set("xxx2", "dsfdsfds")
	//	mw.GetSession(c).Set("xxx222", "dsfdsfds")

	//	c.SetData("aaa", "zzzz")
	//	c.SetData("bbb", "zzzz")
	//	c.SetData("ccc", "zzzz")
	//	c.Template(`index`, nil)

	//	c.Text(c.Param("xxx"))

	//	c.Text(c.Param("yyy"))

	//	c.HTML(`<h1>	helloworld</h1>`)

	//	c.WriteCode(401)

	//	sess := mw.GetSession(c)
	//	sess.Set("name", "helloworld")
	//	c.Text([]byte("set done"))

	//	sess := mw.GetSession(c)
	//	c.Text([]byte("get " + sess.Get("name").(string)))

	//	Addr:     "139.196.40.100:6379",
	//	Password: "OQYG22dfd45gfgfgfB84V",
	//	DB:       15,

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
	//		Destroyed: func(s se.Session) {
	//			log.Println("Destroyed")
	//		},
	//	},
	//)

	//	&rds.Options{
	//		Addr:     "139.196.40.100:6379",
	//		Password: "OQYG22dfd45gfgfgfB84V",
	//		DB:       15,
	//	}),
	//	&se.Config{Timeout: time.Hour},
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
	//		Destroyed: func(s se.Session) {
	//			log.Println("Destroyed")
	//		},
	//	})

	app.UseNotFound()

	app.Run()
}
