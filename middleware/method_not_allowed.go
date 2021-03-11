package middleware

import (
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/util"
	"net/http"
)

type methodNotAllowed struct {
	handler func(ctx *context.Context)
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
	return PatternAny
}

// Handler implements
func (m *methodNotAllowed) Handler() func(ctx *context.Context) {
	return m.handler
}

// MethodNotAllowed return new methodNotAllowed
func MethodNotAllowed(handlers ...func(ctx *context.Context)) Middleware {
	if len(handlers) > 0 && handlers[0] != nil {
		return &methodNotAllowed{handlers[0]}
	}
	return &methodNotAllowed{methodNotAllowedHandler}
}

var methodNotAllowedHandler = func(ctx *context.Context) {
	if !util.RequestSupport(ctx.Request.Method) {
		ctx.WriteCode(http.StatusMethodNotAllowed)
	} else {
		ctx.Chain()
	}
}
