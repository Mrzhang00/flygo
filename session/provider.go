package session

import "net/http"

// Provider interface
type Provider interface {
	// CookieName session cookie name
	CookieName() string
	// Exists session
	Exists(id string) bool
	// GetId session id
	GetId(r *http.Request) string
	// Del session
	Del(id string)
	// Get session id
	Get(id string) Session
	// GetAll session
	GetAll() map[string]Session
	// Clear sessions
	Clear()
	// New return session
	New(config *Config, listener *Listener) Session
	// Refresh session
	Refresh(session Session, config *Config, listener *Listener)
	// Clean session
	Clean(config *Config, listener *Listener)
}
