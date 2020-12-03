package flygo

import (
	"fmt"
	"net/http"
	"time"
)

//Define Session interface
type Session interface {
	Id() string                                     //Get session Id
	SetExpiresTime(lifeTime time.Duration)          //Set expires time
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
	if app.SessionConfig.Enable && a.SessionConfig.SessionProvider != nil {
		a.Info(fmt.Sprintf("Use SessionProvider : %T", a.SessionConfig.SessionProvider))
	}
}

//Get sessions
func (a *App) Sessions() map[string]Session {
	return a.SessionConfig.SessionProvider.GetAll()
}

//Init CookieSession
func (c *Context) initSession() {
	sessionId := app.SessionConfig.SessionProvider.GetId(c.Request)
	have := app.SessionConfig.SessionProvider.Exists(sessionId)
	if have {
		c.SessionId = app.SessionConfig.SessionProvider.GetId(c.Request)
		c.Session = app.SessionConfig.SessionProvider.Get(c.SessionId)
		//Rrefresh session
		app.SessionConfig.SessionProvider.Refresh(c.Session, app.SessionConfig)
		//When Refreshed is set
		refreshed := app.SessionConfig.SessionListener.Refreshed
		if refreshed != nil {
			go func(session Session) {
				if session != nil {
					refreshed(c.Session)
				}
			}(c.Session)
		}
	} else {
		//Create new session
		session := app.SessionConfig.SessionProvider.New(app.SessionConfig)
		c.SessionId = session.Id()
		c.Session = session
		http.SetCookie(c.ResponseWriter, &http.Cookie{
			Name:  headerSessionId,
			Value: session.Id(),
			Path:  "/",
		})
		created := app.SessionConfig.SessionListener.Created
		if created != nil {
			go func(session Session) {
				if session != nil {
					created(session)
				}
			}(c.Session)
		}
	}
}
