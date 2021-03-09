package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"net/http"
)

type methodNotAllowed struct {
	handler func(c *c.Context)
}

// Type implements
func (m *methodNotAllowed) Type() *Type {
	return TypeBefore
}

// Name implements
func (m *methodNotAllowed) Name() string {
	return "MethodNotAllowed"
}

// Method implements
func (m *methodNotAllowed) Method() Method {
	return MethodAny
}

// Pattern implements
func (m *methodNotAllowed) Pattern() Pattern {
	return PatternNoRoute
}

// Handler implements
func (m *methodNotAllowed) Handler() func(c *c.Context) {
	return m.handler
}

// MethodNotAllowed return new methodNotAllowed
func MethodNotAllowed(handlers ...func(c *c.Context)) Middleware {
	if len(handlers) > 0 && handlers[0] != nil {
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
