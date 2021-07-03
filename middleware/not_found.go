package middleware

import (
	"github.com/billcoding/flygo/context"
	"net/http"
)

type notFound struct {
	handler func(ctx *context.Context)
}

// Name implements
func (n *notFound) Name() string {
	return "NotFound"
}

// Type implements
func (n *notFound) Type() *Type {
	return TypeAfter
}

// Method implements
func (n *notFound) Method() Method {
	return MethodAny
}

// Pattern implements
func (n *notFound) Pattern() Pattern {
	return PatternNoRoute
}

// Handler implements
func (n *notFound) Handler() func(ctx *context.Context) {
	return n.handler
}

// NotFound return new notFound
func NotFound(handlers ...func(ctx *context.Context)) Middleware {
	if len(handlers) > 0 && handlers[0] != nil {
		return &notFound{handlers[0]}
	}
	return &notFound{notFoundHandler}
}

var notFoundHandler = func(ctx *context.Context) {
	if ctx.Routed() {
		ctx.Chain()
		return
	}
	ctx.WriteCode(http.StatusNotFound)
	ctx.Write([]byte("404 Not Found"))
}
