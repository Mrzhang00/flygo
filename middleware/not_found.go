package middleware

import (
	c "github.com/billcoding/flygo/context"
	"net/http"
)

type notFound struct {
	handler func(c *c.Context)
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
func (n *notFound) Handler() func(c *c.Context) {
	return n.handler
}

// NotFound return new notFound
func NotFound(handlers ...func(c *c.Context)) Middleware {
	if len(handlers) > 0 {
		return &notFound{handlers[0]}
	}
	return &notFound{notFoundHandler}
}

var notFoundHandler = func(c *c.Context) {
	if !c.Render().Rended() {
		c.WriteCode(http.StatusNotFound)
		c.Write([]byte("404 Not Found"))
	} else {
		c.Chain()
	}
}
