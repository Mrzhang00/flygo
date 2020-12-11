package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	se "github.com/billcoding/flygo/session"
	"net/http"
)

//Define session struct
type session struct {
	listener *se.Listener
	config   *se.Config
	provider se.Provider
}

//Type
func (s *session) Type() *Type {
	return TypeBefore
}

const sessionMWName = "Session"

//Name
func (s *session) Name() string {
	return sessionMWName
}

//Method
func (s *session) Method() Method {
	return MethodAny
}

//Pattern
func (s *session) Pattern() Pattern {
	return PatternAny
}

//setSession
func (s *session) setSession(c *c.Context, session se.Session) {
	c.MWData[s.Name()] = session
}

//GetSession
func GetSession(c *c.Context) se.Session {
	session, have := c.MWData[sessionMWName]
	if have {
		return session.(se.Session)
	}
	return nil
}

//Handler
func (s *session) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		sessionId := s.provider.GetId(c.Request)
		have := false
		if sessionId != "" {
			have = s.provider.Exists(sessionId)
		}
		if have {
			//Get session
			session := s.provider.Get(sessionId)
			//Refresh session
			s.provider.Refresh(session, s.config, s.listener)
		} else {
			//Create new session
			session := s.provider.New(s.config, s.listener)
			//Set session
			s.setSession(c, session)
			//Add set cookie header
			c.Header().Add(headers.SetCookie, (&http.Cookie{
				Name:  s.provider.CookieName(),
				Value: session.Id(),
				Path:  "/",
			}).String())
		}
		c.Chain()
	}
}

//Session
func Session(provider se.Provider, config *se.Config, listener *se.Listener) Middleware {
	//Start clean goroutine
	go provider.Clean(config, listener)
	return &session{
		provider: provider,
		config:   config,
		listener: listener,
	}
}
