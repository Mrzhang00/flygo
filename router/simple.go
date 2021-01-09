package router

import c "github.com/billcoding/flygo/context"

// Simple struct
type Simple struct {
	// Method route method
	Method string
	// Pattern route pattern
	Pattern string
	// Handler route handler
	Handler func(c *c.Context)
}
