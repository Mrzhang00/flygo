package main

import (
	"github.com/billcoding/flygo"
	"github.com/billcoding/flygo/context"
)

type MyController struct {
}

func (ctl *MyController) Prefix() string {
	return "//////gfdgdgfdg/////////////////"
}

func (ctl *MyController) GET() func(c *context.Context) {
	return func(ctx *context.Context) {

	}
}

func (ctl *MyController) GETS() func(c *context.Context) {
	return func(c *context.Context) {

	}
}

func (ctl *MyController) POST() func(c *context.Context) {
	return func(c *context.Context) {

	}
}

func (ctl *MyController) PUT() func(c *context.Context) {
	return func(c *context.Context) {

	}
}

func (ctl *MyController) DELETE() func(c *context.Context) {
	return func(c *context.Context) {

	}
}

func main() {
	flygo.GetApp().GET("/aaaaaaaaaa", func(ctx *context.Context) {
		ctx.JS(`alert(00000000000000000)`)
	}).REST().Run()
}
