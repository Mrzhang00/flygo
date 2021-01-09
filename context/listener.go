package context

// Listener interface
type Listener interface {
	// Created Listener
	Created(c *Context)
	// PreparedAdd Listener
	PreparedAdd(c *Context, handlers ...func(c *Context))
	// Added Listener
	Added(c *Context, handlers ...func(c *Context))
	// PreparedAddOnce Listener
	PreparedAddOnce(c *Context, handlers ...func(c *Context))
	// AddedOnce Listener
	AddedOnce(c *Context, handlers ...func(c *Context))
	// Destroyed Listener
	Destroyed(c *Context)
}

var listeners = make([]Listener, 0)

// AddListeners add listeners
func AddListeners(ls ...Listener) {
	listeners = append(listeners, ls...)
}

func (c *Context) onCreated() {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.Created != nil {
				listener.Created(c)
			}
		}
	}
}

func (c *Context) onDestroyed() {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.Destroyed != nil {
				listener.Destroyed(c)
			}
		}
	}
}

func (c *Context) onPreparedAdd(handlers ...func(c *Context)) {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.PreparedAdd != nil {
				listener.PreparedAdd(c, handlers...)
			}
		}
	}
}

func (c *Context) onAdded(handlers ...func(c *Context)) {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.Added != nil {
				listener.Added(c, handlers...)
			}
		}
	}
}

func (c *Context) onPreparedAddOnce(handlers ...func(c *Context)) {
	if len(listeners) > 0 {
		if len(c.handlers) <= 0 {
			for _, listener := range listeners {
				if listener.PreparedAddOnce != nil {
					listener.PreparedAddOnce(c, handlers...)
				}
			}
		}
	}
}

func (c *Context) onAddedOnce(handlers ...func(c *Context)) {
	if len(listeners) > 0 {
		if len(c.handlers) <= 0 {
			for _, listener := range listeners {
				if listener.AddedOnce != nil {
					listener.AddedOnce(c, handlers...)
				}
			}
		}
	}
}
