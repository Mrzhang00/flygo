package memory

import (
	se "github.com/billcoding/flygo/session"
	"time"
)

//Define memorySession struct
type session struct {
	id          string                 //The sessionId
	expiresTime time.Time              //The session expires time
	invalidated bool                   //invalidated?
	attributes  map[string]interface{} //The session attribute
}

//newSession
func newSession(id string, lifeTime time.Duration) se.Session {
	return &session{
		id:          id,
		expiresTime: time.Now().Add(lifeTime),
		attributes:  make(map[string]interface{}),
	}
}

//Id
func (s *session) Id() string {
	return s.id
}

//Renew
func (s *session) Renew(lifeTime time.Duration) {
	s.expiresTime = time.Now().Add(lifeTime)
}

//Invalidate
func (s *session) Invalidate() {
	s.invalidated = true
}

//Invalidated
func (s *session) Invalidated() bool {
	return s.invalidated
}

//Get all attributes
func (s *session) GetAll() map[string]interface{} {
	return s.attributes
}

//Get attribute
func (s *session) Get(name string) interface{} {
	val, have := s.attributes[name]
	if have {
		return val
	}
	return nil
}

//Set attribute
func (s *session) Set(name string, val interface{}) {
	s.attributes[name] = val
}

//Set attribute
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

//Del attribute
func (s *session) Del(name string) {
	delete(s.attributes, name)
}

//Clear attributes
func (s *session) Clear() {
	s.attributes = make(map[string]interface{})
}
