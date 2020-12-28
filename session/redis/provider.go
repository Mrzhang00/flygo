package redis

import (
	"crypto/md5"
	"fmt"
	"github.com/billcoding/calls"
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
	keyPrefix string
	options   *redis.Options
	client    *redis.Client
	sessions  *sync.Map
}

//Get new SessionProvider
func Provider(options *redis.Options) se.Provider {
	return NewWithPrefixKey(options, defaultPrefixKey)
}

//Get new SessionProvider with prefix key
func NewWithPrefixKey(options *redis.Options, prefixKey string) se.Provider {
	client := redis.NewClient(options)
	ping := client.Ping()
	p := &provider{
		logger:    log.New("[Provider]"),
		keyPrefix: prefixKey,
		options:   options,
		client:    client,
		sessions:  &sync.Map{},
	}
	calls.NNil(ping.Err(), func() {
		p.logger.Warn("%v", ping.Err())
	})
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
	get := p.Get(id)
	return get != nil && !get.Invalidated()
}

//Get
func (p *provider) Get(id string) se.Session {
	sess, have := p.sessions.Load(id)
	if !have {
		return nil
	}
	return sess.(se.Session)
}

//Del
func (p *provider) Del(id string) {
	delCmd := p.client.Del(p.getRedisKey(id))
	if delCmd.Err() != nil {
		p.logger.Error("[Del]%s", delCmd.Err())
	}
}

//GetAll
func (p *provider) GetAll() map[string]se.Session {
	keysCmd := p.client.Keys(p.keyPrefix)
	if keysCmd.Err() != nil {
		p.logger.Error("[GetAll]%s", keysCmd.Err())
		return nil
	}
	keys, _ := keysCmd.Result()
	sessionMap := make(map[string]se.Session)
	for _, key := range keys {
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
	keysCmd := p.client.Keys(p.keyPrefix)
	if keysCmd.Err() != nil {
		p.logger.Error("[Clear]%s", keysCmd.Err())
		return
	}
	keys, _ := keysCmd.Result()
	for _, key := range keys {
		delCmd := p.client.Del(key)
		if delCmd.Err() != nil {
			p.logger.Error("[Clear]%s", delCmd.Err())
		}
	}
}

//tmd5
func tmd5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

//newSID
func newSID() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	return strings.ToUpper(tmd5(tmd5(strconv.FormatInt(nano, 10)) + tmd5(strconv.FormatInt(rndNum, 10))))
}

//New
func (p *provider) New(config *se.Config, listener *se.Listener) se.Session {
	sessionId := newSID()
	session := newSession(p.client, sessionId, p.getRedisKey(sessionId))
	hsetCmd := p.client.HSet(p.getRedisKey(sessionId), sessionIdName, sessionId)
	if hsetCmd.Err() != nil {
		p.logger.Error("[New]%v", hsetCmd.Err())
		return nil
	}
	expireCmd := p.client.Expire(p.getRedisKey(sessionId), config.Timeout)
	if expireCmd.Err() != nil {
		p.logger.Error("[New]%v", expireCmd.Err())
		return nil
	}
	p.sessions.Store(sessionId, session)
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
	go p.cleanSession(listener)
}

//syncSession
func (p *provider) syncSession() {
	wait := make(chan bool, 1)
	go func() {
		keysCmd := p.client.Keys(p.keyPrefix + "*")
		if keysCmd.Err() != nil {
			p.logger.Error("[syncSession]%v", keysCmd.Err())
			return
		}
		keys := keysCmd.Val()
		sessionMap := make(map[string]se.Session, 0)
		for _, key := range keys {
			hgetAllCmd := p.client.HGetAll(key)
			if hgetAllCmd.Err() != nil {
				p.logger.Error("[syncSession]%v", hgetAllCmd.Err())
				continue
			}
			vals := hgetAllCmd.Val()
			sessionId := vals[sessionIdName]
			rs := newSession(p.client, sessionId, key)
			sessionMap[sessionId] = rs
			p.sessions.Store(sessionId, newSession(p.client, sessionId, key))
		}
	}()
	wait <- true
}

//cleanSession
func (p *provider) cleanSession(listener *se.Listener) {
	for {
		p.sessions.Range(func(k, v interface{}) bool {
			sessionId := k.(string)
			sess := v.(se.Session)
			key := p.getRedisKey(sessionId)
			existsCmd := p.client.Exists(key)
			if existsCmd.Err() != nil {
				p.logger.Error("[invalidatedSession]%v", existsCmd.Err())
			} else {
				if existsCmd.Val() <= 0 {
					sess.Invalidate()
					go func() {
						if listener != nil && listener.Invalidated != nil {
							listener.Invalidated(sess)
						}
					}()
				}
			}
			if sess.Invalidated() {
				p.sessions.Delete(sessionId)
				go func() {
					if listener != nil && listener.Destoryed != nil {
						listener.Destoryed(sess)
					}
				}()
			}
			return true
		})
		time.Sleep(time.Second)
	}
}
