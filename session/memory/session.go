package memory

import (
	se "github.com/billcoding/flygo/session"
	"time"
)

type session struct {
	id          string
	expiresTime time.Time
	invalidated bool
	attributes  map[string]interface{}
}

func newSession(id string, minute int) se.Session {
	return &session{
		id:          id,
		expiresTime: time.Now().Add(se.GetTimeout(minute)),
		attributes:  make(map[string]interface{}),
	}
}

// Id return session id
func (s *session) Id() string {
	return s.id
}

// Renew session
func (s *session) Renew(lifeTime time.Duration) {
	s.expiresTime = time.Now().Add(lifeTime)
}

// Invalidate session
func (s *session) Invalidate() {
	s.invalidated = true
}

// Invalidated session
func (s *session) Invalidated() bool {
	return s.invalidated
}

// GetAll session's vals
func (s *session) GetAll() map[string]interface{} {
	return s.attributes
}

// Get session named val
func (s *session) Get(name string) interface{} {
	val, have := s.attributes[name]
	if have {
		return val
	}
	return nil
}

// Set val into session
func (s *session) Set(name string, val interface{}) {
	s.attributes[name] = val
}

// SetAll vals into session
func (s *session) SetAll(data map[string]interface{}, flush bool) {
	if data == nil {
		return
	}

	if flush {
		s.attributes = data
		return
	}

	for k, v := range data {
		s.Set(k, v)
	}
}

// Del named val from session
func (s *session) Del(name string) {
	delete(s.attributes, name)
}

// Clear session's vals
func (s *session) Clear() {
	s.attributes = make(map[string]interface{})
}
