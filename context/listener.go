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

func (c *Context) onCreated() {
	for _, listener := range listeners {
		listener.Created(c)
	}
}

func (c *Context) onBefore() {
	for _, listener := range listeners {
		listener.Before(c)
	}
}

func (c *Context) onAfter() {
	for _, listener := range listeners {
		listener.After(c)
	}
}

func (c *Context) onDestroyed() {
	for _, listener := range listeners {
		listener.Destroyed(c)
	}
}
