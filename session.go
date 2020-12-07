package flygo

import (
	"fmt"
	"net/http"
	"time"
)

//Define Session interface
type Session interface {
	Id() string                                     //Get session Id
	Renew(lifeTime time.Duration)                   //Renew session
	Invalidate()                                    //Invalidate session
	Invalidated() bool                              //Invalidated?
	Get(name string) interface{}                    //Get data
	GetAll() map[string]interface{}                 //Get all data
	Set(name string, val interface{})               //Set data
	SetAll(data map[string]interface{}, flush bool) //Set all data
	Del(name string)                                //Del by name
	Clear()                                         //Clear all data
}

//Define SessionProvider interface
type SessionProvider interface {
	Exists(id string) bool                          //Get session is exists
	GetId(r *http.Request) string                   //Get SessionId from client
	Del(id string)                                  //Del Session from provider
	Get(id string) Session                          //Get Session from provider
	GetAll() map[string]Session                     //Get All Session from provider
	New(config *SessionConfig) Session              //Get new Session from provider
	Clear()                                         //Clear all sessions
	Refresh(session Session, config *SessionConfig) //Refresh handler from provider
}

//Print SessionProvider
func (a *App) printSessionProvider() {
	if a.SessionConfig.Enable && a.SessionConfig.SessionProvider != nil {
		a.Logger.Info(fmt.Sprintf("Use SessionProvider : %T", a.SessionConfig.SessionProvider))
	}
}

//Get sessions
func (a *App) Sessions() map[string]Session {
	return a.SessionConfig.SessionProvider.GetAll()
}

//Init CookieSession
func (c *Context) initSession() {
	sessionId := c.app.SessionConfig.SessionProvider.GetId(c.Request)
	have := c.app.SessionConfig.SessionProvider.Exists(sessionId)
	if have {
		c.SessionId = c.app.SessionConfig.SessionProvider.GetId(c.Request)
		c.Session = c.app.SessionConfig.SessionProvider.Get(c.SessionId)
		//Rrefresh session
		c.app.SessionConfig.SessionProvider.Refresh(c.Session, c.app.SessionConfig)
		//When Refreshed is set
		refreshed := c.app.SessionConfig.SessionListener.Refreshed
		if refreshed != nil {
			go func(session Session) {
				if session != nil {
					refreshed(c.Session)
				}
			}(c.Session)
		}
	} else {
		//Create new session
		session := c.app.SessionConfig.SessionProvider.New(c.app.SessionConfig)
		c.SessionId = session.Id()
		c.Session = session
		http.SetCookie(c.ResponseWriter, &http.Cookie{
			Name:  headerSessionId,
			Value: session.Id(),
			Path:  "/",
		})
		created := c.app.SessionConfig.SessionListener.Created
		if created != nil {
			go func(session Session) {
				if session != nil {
					created(session)
				}
			}(c.Session)
		}
	}
}
