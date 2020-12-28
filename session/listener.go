package session

//Define sessionListener type
type sessionListener func(session Session)

//Define Listener struct
type Listener struct {
	Created     sessionListener //Session Created
	Destoryed   sessionListener //Session Destoryed
	Invalidated sessionListener //Session Invalidated
	Refreshed   sessionListener //Session Refreshed
}
