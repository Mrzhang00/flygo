package router

import "github.com/billcoding/flygo/context"

// Simple struct
type Simple struct {
	// Method route method
	Method string
	// Pattern route pattern
	Pattern string
	// Handler route handler
	Handler func(ctx *context.Context)
}
