package context

// Listener interface
type Listener interface {
	// Created Listener
	Created(c *Context)
	// Before Listener
	Before(c *Context, handler func(c *Context))
	// After Listener
	After(c *Context, handler func(c *Context))
	// Destroyed Listener
	Destroyed(c *Context)
}

var listeners = make([]Listener, 0)

// AddListeners add listeners
func AddListeners(ls ...Listener) {
	listeners = append(listeners, ls...)
}

func (c *Context) onCreated() {
	for _, listener := range listeners {
		listener.Created(c)
	}
}

func (c *Context) onBefore(handler func(c *Context)) {
	for _, listener := range listeners {
		listener.Before(c, handler)
	}
}

func (c *Context) onAfter(handler func(c *Context)) {
	for _, listener := range listeners {
		listener.After(c, handler)
	}
}

func (c *Context) onDestroyed() {
	for _, listener := range listeners {
		listener.Destroyed(c)
	}
}
