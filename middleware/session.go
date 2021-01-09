package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	se "github.com/billcoding/flygo/session"
	"net/http"
)

const sessionMWName = "Session"

type session struct {
	listener *se.Listener
	config   *se.Config
	provider se.Provider
}

// Type implements
func (s *session) Type() *Type {
	return TypeBefore
}

// Name implements
func (s *session) Name() string {
	return sessionMWName
}

// Method implements
func (s *session) Method() Method {
	return MethodAny
}

// Pattern implements
func (s *session) Pattern() Pattern {
	return PatternAny
}

func (s *session) setSession(c *c.Context, session se.Session) {
	c.MWData[s.Name()] = session
}

// GetSession get current session
func GetSession(c *c.Context) se.Session {
	sess, have := c.MWData[sessionMWName]
	if have {
		return sess.(se.Session)
	}
	return nil
}

// Handler implements
func (s *session) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		sessionId := s.provider.GetId(c.Request)
		have := false
		if sessionId != "" {
			have = s.provider.Exists(sessionId)
		}
		if have {
			session := s.provider.Get(sessionId)
			s.provider.Refresh(session, s.config, s.listener)
			c.SetData("session", session.GetAll())
		} else {
			session := s.provider.New(s.config, s.listener)
			s.setSession(c, session)
			c.Header().Add(headers.SetCookie, (&http.Cookie{
				Name:  s.provider.CookieName(),
				Value: session.Id(),
				Path:  "/",
			}).String())
			c.SetData("session", session.GetAll())
		}
		c.Chain()
	}
}

// Session return new session
func Session(provider se.Provider, config *se.Config, listener *se.Listener) *session {
	provider.Clean(config, listener)
	return &session{
		provider: provider,
		config:   config,
		listener: listener,
	}
}
