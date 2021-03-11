package context

// Listener interface
type Listener interface {
	// Created Listener
	Created(c *Context)
	// Before Listener
	Before(c *Context)
	// After Listener
	After(c *Context)
	// Destroyed Listener
	Destroyed(c *Context)
}

var listeners = make([]Listener, 0)

// AddListeners add listeners
func AddListeners(ls ...Listener) {
	listeners = append(listeners, ls...)
}

func (ctx *Context) onCreated() {
	for _, listener := range listeners {
		listener.Created(ctx)
	}
}

func (ctx *Context) onBefore() {
	for _, listener := range listeners {
		listener.Before(ctx)
	}
}

func (ctx *Context) onAfter() {
	for _, listener := range listeners {
		listener.After(ctx)
	}
}

func (ctx *Context) onDestroyed() {
	for _, listener := range listeners {
		listener.Destroyed(ctx)
	}
}
