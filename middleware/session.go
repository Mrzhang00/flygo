package middleware

import (
	"github.com/billcoding/flygo/context"
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

func (s *session) setSession(ctx *context.Context, session se.Session) {
	ctx.MWData[s.Name()] = session
}

// GetSession get current session
func GetSession(ctx *context.Context) se.Session {
	sess, have := ctx.MWData[sessionMWName]
	if have {
		return sess.(se.Session)
	}
	return nil
}

// Handler implements
func (s *session) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		sessionId := s.provider.GetId(ctx.Request)
		have := false
		if sessionId != "" {
			have = s.provider.Exists(sessionId)
		}
		if have {
			session := s.provider.Get(sessionId)
			s.provider.Refresh(session, s.config, s.listener)
			ctx.SetData("session", session.GetAll())
		} else {
			session := s.provider.New(s.config, s.listener)
			s.setSession(ctx, session)
			ctx.Header().Add(headers.SetCookie, (&http.Cookie{
				Name:  s.provider.CookieName(),
				Value: session.Id(),
				Path:  "/",
			}).String())
			ctx.SetData("session", session.GetAll())
		}
		ctx.Chain()
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
