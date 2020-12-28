package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"net/http"
)

//Define methodNotAllowed struct
type methodNotAllowed struct {
	handler func(c *c.Context)
}

//Type
func (m *methodNotAllowed) Type() *Type {
	return TypeBefore
}

//Name
func (m *methodNotAllowed) Name() string {
	return "MethodNotAllowed"
}

//Method
func (m *methodNotAllowed) Method() Method {
	return MethodAny
}

//Pattern
func (m *methodNotAllowed) Pattern() Pattern {
	return PatternNoRoute
}

//Handler
func (m *methodNotAllowed) Handler() func(c *c.Context) {
	return m.handler
}

//MethodNotAllowed
func MethodNotAllowed(handlers ...func(c *c.Context)) Middleware {
	if len(handlers) > 0 {
		return &methodNotAllowed{handlers[0]}
	}
	return &methodNotAllowed{methodNotAllowedHandler}
}

//methodNotAllowedHandler
var methodNotAllowedHandler = func(c *c.Context) {
	if !util.RequestSupport(c.Request.Method) {
		c.WriteCode(http.StatusMethodNotAllowed)
	} else {
		c.Chain()
	}
}
