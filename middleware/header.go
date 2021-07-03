package middleware

import (
	"fmt"
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"runtime"
)

type header struct {
	headers map[string]string
}

// Type implements
func (h *header) Type() *Type {
	return TypeBefore
}

// Name implements
func (h *header) Name() string {
	return "Header"
}

// Method implements
func (h *header) Method() Method {
	return MethodAny
}

// Pattern implements
func (h *header) Pattern() Pattern {
	return PatternNoRoute
}

// Handler implements
func (h *header) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		for k, v := range h.headers {
			ctx.Header().Set(k, v)
		}
		ctx.Chain()
	}
}

// Header return header
func Header() Middleware {
	return &header{map[string]string{
		headers.Server:        fmt.Sprintf("flygo/%s", runtime.Version()),
		headers.BackendServer: fmt.Sprintf("flygo/%s", runtime.Version()),
	}}
}
