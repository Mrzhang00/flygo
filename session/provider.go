package session

import "net/http"

type Provider interface {
	CookieName() string
	Exists(id string) bool
	GetId(r *http.Request) string
	Del(id string)
	Get(id string) Session
	GetAll() map[string]Session
	Clear()
	New(config *Config, listener *Listener) Session
	Refresh(session Session, config *Config, listener *Listener)
	Clean(config *Config, listener *Listener)
}
