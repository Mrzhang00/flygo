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

func (s *session) Id() string {
	return s.id
}

func (s *session) Renew(lifeTime time.Duration) {
	s.client.Expire(s.key, lifeTime)
}

func (s *session) Invalidated() bool {
	return s.invalidated
}

func (s *session) Invalidate() {
	s.invalidated = true
}

func (s *session) Get(name string) interface{} {
	getCmd := s.client.HGet(s.key, name)
	val := ""
	err := getCmd.Scan(&val)
	if err != nil {
		s.logger.Error("[Get]%v", err)
	}
	return val
}

func (s *session) GetAll() map[string]interface{} {
	getAllCmd := s.client.HGetAll(s.key)
	vals := getAllCmd.Val()
	nvals := make(map[string]interface{}, 0)
	for k, v := range vals {
		nvals[k] = v
	}
	return nvals
}

func (s *session) Set(name string, val interface{}) {
	s.supportedHandle(name, func() {
		s.client.HSet(s.key, name, val)
	})
}

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

func (s *session) Del(name string) {
	s.supportedHandle(name, func() {
		s.client.HDel(s.key, name)
	})
}

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
