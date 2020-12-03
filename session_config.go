package flygo

import "time"

type sessionListener func(session Session)

//call when the session be created
type SessionCreatedListener sessionListener

//call when the session be destoryed
type SessionDestoryedListener sessionListener

//call when the session be refreshed
type SessionRefreshedListener sessionListener

//Define SessionListener struct
type SessionListener struct {
	Created   sessionListener //SessionCreated
	Destoryed sessionListener //SessionDestoryed
	Refreshed sessionListener //SessionRefreshed
}

//Define SessionConfig struct
type SessionConfig struct {
	Enable          bool
	SessionProvider SessionProvider
	*SessionListener
	Timeout time.Duration
}
