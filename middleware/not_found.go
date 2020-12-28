package middleware

import (
	c "github.com/billcoding/flygo/context"
	"net/http"
)

//Define notFound struct
type notFound struct {
	handler func(c *c.Context)
}

//Name
func (n *notFound) Name() string {
	return "NotFound"
}

//Method
func (n *notFound) Type() *Type {
	return TypeAfter
}

//Method
func (n *notFound) Method() Method {
	return MethodAny
}

//Pattern
func (n *notFound) Pattern() Pattern {
	return PatternNoRoute
}

//Handler
func (n *notFound) Handler() func(c *c.Context) {
	return n.handler
}

//NotFound
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
