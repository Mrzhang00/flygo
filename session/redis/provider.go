package redis

import (
	"crypto/md5"
	"fmt"
	"github.com/billcoding/flygo/log"
	se "github.com/billcoding/flygo/session"
	"github.com/go-redis/redis"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const defaultPrefixKey = "session:"

//Define provider struct
type provider struct {
	logger    log.Logger
	mu        sync.Mutex
	keyPrefix string
	options   *redis.Options
	client    *redis.Client
	sessions  map[string]se.Session
}

//Get new SessionProvider
func Provider(options *redis.Options) se.Provider {
	return NewWithPrefixKey(options, defaultPrefixKey)
}

//Get new SessionProvider with prefix key
func NewWithPrefixKey(options *redis.Options, prefixKey string) se.Provider {
	client := redis.NewClient(options)
	p := &provider{
		logger:    log.New("[Provider]"),
		mu:        sync.Mutex{},
		keyPrefix: prefixKey,
		options:   options,
		client:    client,
		sessions:  make(map[string]se.Session),
	}
	//First sync session from redis
	p.syncSession()
	return p
}

//CookieName
func (p *provider) CookieName() string {
	return "GSESSIONID"
}

//GetId
func (p *provider) GetId(r *http.Request) string {
	cookie, err := r.Cookie(p.CookieName())
	if err == nil && cookie != nil {
		return cookie.Value
	}
	return ""
}

//getRedisKey
func (p *provider) getRedisKey(id string) string {
	return fmt.Sprintf("%s%s", p.keyPrefix, id)
}

//Exists
func (p *provider) Exists(id string) bool {
	session, have := p.sessions[id]
	return have && !session.Invalidated()
}

//Get
func (p *provider) Get(id string) se.Session {
	session, have := p.sessions[id]
	if have && !session.Invalidated() {
		return session
	}
	return nil
}

//Del
func (p *provider) Del(id string) {
	p.client.Del(p.getRedisKey(id))
}

//GetAll
func (p *provider) GetAll() map[string]se.Session {
	keys := p.client.Keys(p.keyPrefix)
	result, err := keys.Result()
	if err != nil {
		fmt.Errorf("%s", err)
		return nil
	}
	sessionMap := make(map[string]se.Session)
	for _, key := range result {
		keys := strings.Split(key, ":")
		if len(keys) > 1 {
			sessionId := keys[1]
			sessionMap[sessionId] = p.Get(sessionId)
		}
	}
	return sessionMap
}

//Clear
func (p *provider) Clear() {
	keys := p.client.Keys(p.keyPrefix)
	result, err := keys.Result()
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
	for _, key := range result {
		p.client.Del(key)
	}
}

//New
func (p *provider) New(config *se.Config, listener *se.Listener) se.Session {
	tmd5 := func(text string) string {
		hashMd5 := md5.New()
		io.WriteString(hashMd5, text)
		return fmt.Sprintf("%x", hashMd5.Sum(nil))
	}
	newSID := func() string {
		nano := time.Now().UnixNano()
		rand.Seed(nano)
		rndNum := rand.Int63()
		return strings.ToUpper(tmd5(tmd5(strconv.FormatInt(nano, 10)) + tmd5(strconv.FormatInt(rndNum, 10))))
	}
	sessionId := newSID()
	session := newSession(p.client, sessionId, p.getRedisKey(sessionId))
	boolCmd := p.client.HSet(p.getRedisKey(sessionId), sessionIdName, sessionId)
	p.client.Expire(p.getRedisKey(sessionId), config.Timeout)
	_, err := boolCmd.Result()
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}
	p.sessions[sessionId] = session
	return session
}

//Refresh
func (p *provider) Refresh(session se.Session, config *se.Config, listener *se.Listener) {
	expireCmd := p.client.Expire(p.getRedisKey(session.Id()), config.Timeout)
	if expireCmd.Err() != nil {
		p.logger.Error("[Refresh]%v", expireCmd.Err())
	} else {
		go func() {
			if listener != nil && listener.Refreshed != nil {
				listener.Refreshed(session)
			}
		}()
	}
}

//Clean
func (p *provider) Clean(config *se.Config, listener *se.Listener) {
	blocked := make(chan bool, 1)
	go p.invalidatedSession(listener)
	go p.cleanSession(listener)
	<-blocked
}

//syncSession
func (p *provider) syncSession() {
	wait := make(chan bool, 1)
	go func() {
		keys := p.client.Keys(p.keyPrefix + "*")
		sessionMap := make(map[string]se.Session, 0)
		for _, key := range keys.Val() {
			getAll := p.client.HGetAll(key)
			vals := getAll.Val()
			sessionId := vals[sessionIdName]
			rs := newSession(p.client, sessionId, key)
			sessionMap[sessionId] = rs
		}
		p.sessions = sessionMap
	}()
	wait <- true
}

//invalidatedSession
func (p *provider) invalidatedSession(listener *se.Listener) {
	for {
		p.mu.Lock()
		for sessionId, sess := range p.sessions {
			key := p.getRedisKey(sessionId)
			existsCmd := p.client.Exists(key)
			if existsCmd.Val() <= 0 {
				sess.Invalidate()
				go func() {
					if listener != nil && listener.Invalidated != nil {
						listener.Invalidated(sess)
					}
				}()
			}
		}
		p.mu.Unlock()
		time.Sleep(time.Second)
	}
}

//cleanSession
func (p *provider) cleanSession(listener *se.Listener) {
	for {
		p.mu.Lock()
		for sessionId, session := range p.sessions {
			if session.Invalidated() {
				delete(p.sessions, sessionId)
				p.Del(sessionId)
				go func() {
					if listener != nil && listener.Destoryed != nil {
						listener.Destoryed(session)
					}
				}()
			}
		}
		p.mu.Unlock()
		time.Sleep(time.Second)
	}
}
