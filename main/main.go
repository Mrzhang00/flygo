package main

import (
	"fmt"
	bind "github.com/billcoding/binding"
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
	In     InnerModel `binding:"name(in)" validate:"required(T) message(In不合法)"`
}

type InnerModel struct {
	Id   uint8  `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id1不合法)"`
	Name string `binding:"name(name)" validate:"required(T) minlength(2) maxlength(5) message(Name不合法)"`
}

func handler(c *Context) {
	m := model{}
	c.BindWithParamsAndValidate(&m, bind.Body, func() {
		c.JSON(&m)
	})
}

func main() {
	app := flygo.GetApp()
	app.UseStatic(false, `./abc`)
	app.GET("/", func(c *Context) {
		c.Text(fmt.Sprintf("%v", mw.GetSession(c).GetAll()))
	})
	app.GET("/1", func(c *Context) {
		mw.GetSession(c).Set("xxx", "dsfdsfds")
		mw.GetSession(c).Set("xxx2", "dsfdsfds")
		mw.GetSession(c).Set("xxx222", "dsfdsfds")
	})
	app.GET("/2", func(c *Context) {
		c.SetData("aaa", "zzzz")
		c.SetData("bbb", "zzzz")
		c.SetData("ccc", "zzzz")
		c.Template(`index`, nil)
	})
	//app.Use(mw.Cors())
	//app.HEAD("/*", func(c *Context) {
	//	c.WriteCode(401)
	//})
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
	//app.Use(mw.RedisToken(&rds.Options{
	//	Addr:     "139.196.40.100:6379",
	//	Password: "OQYG22dfd45gfgfgfB84V",
	//	DB:       15,
	//}))
	//app.UseSession(memory.Provider(), &se.Config{Timeout: time.Minute},
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
	//	},
	//)
	//app.UseSession(redis.Provider(
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
	//		Destoryed: func(s se.Session) {
	//			log.Println("Destoryed")
	//		},
	//	})
	//app.UseRecovery()
	//app.UseMethodNotAllowed()
	app.UseNotFound()
	//app.UseStdLogger()
	//app.Config.Debug = true
	//app.ConfigFile = `/Users/local/Desktop/Workspaces/Goland/src/github.com/billcoding/flygo/main/flygo.yml`
	app.Run()
}
