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

//Get app session provider
func (a *App) SessionProvider(sessionProvider SessionProvider) *App {
	a.sessionProvider = sessionProvider
	return a
}

//Print SessionProvider
func (a *App) printSessionProvider() {
	if a.sessionEnable && a.sessionProvider != nil {
		a.LogInfo(fmt.Sprintf("Use SessionProvider : %T", a.sessionProvider))
	}
}

//Get sessions
func (a *App) Sessions() map[string]Session {
	return a.sessionProvider.GetAll()
}

//Enable sessions
func (a *App) SessionEnable(sessionEnable bool) *App {
	a.sessionEnable = sessionEnable
	return a
}

//Init CookieSession
func (c *Context) initSession() {
	sessionId := app.sessionProvider.GetId(c.Request)
	have := app.sessionProvider.Exists(sessionId)
	if have {
		c.SessionId = app.sessionProvider.GetId(c.Request)
		c.Session = app.sessionProvider.Get(c.SessionId)
		//Rrefresh session
		app.sessionProvider.Refresh(c.Session, app.SessionConfig)
		//When RefreshedListener is set
		refreshedListener := app.SessionConfig.RefreshedListener
		if refreshedListener != nil {
			go func(session Session) {
				if session != nil {
					refreshedListener(c.Session)
				}
			}(c.Session)
		}
	} else {
		//Create new session
		session := app.sessionProvider.New(app.SessionConfig)
		c.SessionId = session.Id()
		c.Session = session
		http.SetCookie(c.ResponseWriter, &http.Cookie{Name: headerSessionId, Value: session.Id()})
		createdListener := app.SessionConfig.CreatedListener
		if createdListener != nil {
			go func(session Session) {
				if session != nil {
					createdListener(session)
				}
			}(c.Session)
		}
	}
}
