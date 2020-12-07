package flygo

//Define listener struct
type listener struct {
	created   ListenerHandler
	started   ListenerHandler
	destoryed ListenerHandler
}

//Register listener
func (ag *appGroup) Listener() *listener {
	return ag.listener
}

//Register created listener
func (l *listener) Created(handler ListenerHandler) {
	l.created = handler
}

//Register started listener
func (l *listener) Started(handler ListenerHandler) {
	l.started = handler
}

//Register destoryed listener
func (l *listener) Destoryed(handler ListenerHandler) {
	l.destoryed = handler
}
