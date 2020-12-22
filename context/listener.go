package context

//Define Listener interface
type Listener interface {
	Created(c *Context)                                       //Created Listener
	PreparedAdd(c *Context, handlers ...func(c *Context))     //PreparedAdd Listener
	Added(c *Context, handlers ...func(c *Context))           //Added Listener
	PreparedAddOnce(c *Context, handlers ...func(c *Context)) //PreparedAddOnce Listener
	AddedOnce(c *Context, handlers ...func(c *Context))       //AddedOnce Listener
	Destoryed(c *Context)                                     //Destoryed Listener
}

//listeners
var listeners = make([]Listener, 0)

//AddListener
func AddListeners(ls ...Listener) {
	listeners = append(listeners, ls...)
}

//onCreated
func (c *Context) onCreated() {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.Created != nil {
				listener.Created(c)
			}
		}
	}
}

//onDestoryed
func (c *Context) onDestoryed() {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.Destoryed != nil {
				listener.Destoryed(c)
			}
		}
	}
}

//OnPreparedAdd
func (c *Context) onPreparedAdd(handlers ...func(c *Context)) {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.PreparedAdd != nil {
				listener.PreparedAdd(c, handlers...)
			}
		}
	}
}

//onAdded
func (c *Context) onAdded(handlers ...func(c *Context)) {
	if len(listeners) > 0 {
		for _, listener := range listeners {
			if listener.Added != nil {
				listener.Added(c, handlers...)
			}
		}
	}
}

//onPreparedAddOnce
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

//onAddedOnce
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
