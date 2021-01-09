package redis

import (
	"github.com/billcoding/flygo/log"
	se "github.com/billcoding/flygo/session"
	"github.com/go-redis/redis"
	"time"
)

type session struct {
	logger      log.Logger
	id          string
	key         string
	invalidated bool
	client      *redis.Client
}

func newSession(client *redis.Client, id, key string) se.Session {
	return &session{
		logger:      log.New("[Session]"),
		id:          id,
		key:         key,
		invalidated: false,
		client:      client,
	}
}

const sessionIdName = "sessionId"

// Id return session id
func (s *session) Id() string {
	return s.id
}

// Renew session
func (s *session) Renew(lifeTime time.Duration) {
	s.client.Expire(s.key, lifeTime)
}

// Invalidated session
func (s *session) Invalidated() bool {
	return s.invalidated
}

// Invalidate session
func (s *session) Invalidate() {
	s.invalidated = true
}

// Get session named val
func (s *session) Get(name string) interface{} {
	getCmd := s.client.HGet(s.key, name)
	val := ""
	err := getCmd.Scan(&val)
	if err != nil {
		s.logger.Error("[Get]%v", err)
	}
	return val
}

// GetAll session's vals
func (s *session) GetAll() map[string]interface{} {
	getAllCmd := s.client.HGetAll(s.key)
	vals := getAllCmd.Val()
	nvals := make(map[string]interface{}, 0)
	for k, v := range vals {
		nvals[k] = v
	}
	return nvals
}

// Set named val into session
func (s *session) Set(name string, val interface{}) {
	s.supportedHandle(name, func() {
		s.client.HSet(s.key, name, val)
	})
}

// SetAll vals into session
func (s *session) SetAll(data map[string]interface{}, flush bool) {
	if flush {
		s.Clear()
	}
	_, have := data[sessionIdName]
	if have {
		delete(data, sessionIdName)
	}
	s.client.HMSet(s.key, data)
}

// Del named val from session
func (s *session) Del(name string) {
	s.supportedHandle(name, func() {
		s.client.HDel(s.key, name)
	})
}

// Clear session's vals
func (s *session) Clear() {
	all := s.GetAll()
	ks := make([]string, 0)
	for k := range all {
		if k != sessionIdName {
			ks = append(ks, k)
		}
	}
	s.client.HDel(s.key, ks...)
}

func (s *session) supportedHandle(name string, fn func()) {
	if name != sessionIdName {
		fn()
	}
}
