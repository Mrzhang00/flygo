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

func (s *session) Id() string {
	return s.id
}

func (s *session) Renew(lifeTime time.Duration) {
	s.expiresTime = time.Now().Add(lifeTime)
}

func (s *session) Invalidate() {
	s.invalidated = true
}

func (s *session) Invalidated() bool {
	return s.invalidated
}

func (s *session) GetAll() map[string]interface{} {
	return s.attributes
}

func (s *session) Get(name string) interface{} {
	val, have := s.attributes[name]
	if have {
		return val
	}
	return nil
}

func (s *session) Set(name string, val interface{}) {
	s.attributes[name] = val
}

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

func (s *session) Del(name string) {
	delete(s.attributes, name)
}

func (s *session) Clear() {
	s.attributes = make(map[string]interface{})
}
