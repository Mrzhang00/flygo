package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"net/http"
)

//Define methodNotAllowed struct
type methodNotAllowed struct {
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
	return PatternAny
}

//Handler
func (m *methodNotAllowed) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		if !util.RequestSupport(c.Request.Method) {
			c.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			c.Chain()
		}
	}
}

//MethodNotAllowed
func MethodNotAllowed() Middleware {
	return &methodNotAllowed{}
}
