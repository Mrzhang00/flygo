package flygo

import "time"

type sessionListener func(session Session)

//call when the session be created
type SessionCreatedListener sessionListener

//call when the session be destoryed
type SessionDestoryedListener sessionListener

//call when the session be refreshed
type SessionRefreshedListener sessionListener

//Define SessionConfig struct
type SessionConfig struct {
	Timeout           time.Duration   //session idle timeout
	CreatedListener   sessionListener //SessionCreatedListener
	DestoryedListener sessionListener //SessionDestoryedListener
	RefreshedListener sessionListener //SessionRefreshedListener
}
