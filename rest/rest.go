package rest

import "github.com/billcoding/flygo/context"

// Controller interface
type Controller interface {
	// Prefix route pattern prefix
	Prefix() string
	// GET route
	GET() func(ctx *context.Context)
	// GETS route
	GETS() func(ctx *context.Context)
	// POST route
	POST() func(ctx *context.Context)
	// PUT route
	PUT() func(ctx *context.Context)
	// DELETE route
	DELETE() func(ctx *context.Context)
}
