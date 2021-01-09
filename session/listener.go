package session

type sessionListener func(session Session)

type Listener struct {
	Created     sessionListener
	Destroyed   sessionListener
	Invalidated sessionListener
	Refreshed   sessionListener
}
