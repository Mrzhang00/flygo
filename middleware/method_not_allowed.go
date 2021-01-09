package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"net/http"
)

type methodNotAllowed struct {
	handler func(c *c.Context)
}

func (m *methodNotAllowed) Type() *Type {
	return TypeBefore
}

func (m *methodNotAllowed) Name() string {
	return "MethodNotAllowed"
}

func (m *methodNotAllowed) Method() Method {
	return MethodAny
}

func (m *methodNotAllowed) Pattern() Pattern {
	return PatternNoRoute
}

func (m *methodNotAllowed) Handler() func(c *c.Context) {
	return m.handler
}

func MethodNotAllowed(handlers ...func(c *c.Context)) Middleware {
	if len(handlers) > 0 {
		return &methodNotAllowed{handlers[0]}
	}
	return &methodNotAllowed{methodNotAllowedHandler}
}

var methodNotAllowedHandler = func(c *c.Context) {
	if !util.RequestSupport(c.Request.Method) {
		c.WriteCode(http.StatusMethodNotAllowed)
	} else {
		c.Chain()
	}
}
