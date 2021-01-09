package memory

import (
	"crypto/md5"
	"fmt"
	se "github.com/billcoding/flygo/session"
	"io"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type provider struct {
	sessions *sync.Map
}

// Provider return session provider
func Provider() se.Provider {
	p := &provider{
		sessions: &sync.Map{},
	}
	return p
}

// CookieName return cookie name
func (p *provider) CookieName() string {
	return "GSESSIONID"
}

// GetId get session id
func (p *provider) GetId(r *http.Request) string {
	cookie, err := r.Cookie(p.CookieName())
	if err == nil && cookie != nil {
		return cookie.Value
	}
	return ""
}

// Exists session
func (p *provider) Exists(id string) bool {
	_, have := p.sessions.Load(id)
	return have
}

// Get session
func (p *provider) Get(id string) se.Session {
	value, have := p.sessions.Load(id)
	if !have {
		return nil
	}
	return value.(se.Session)
}

// Del session
func (p *provider) Del(id string) {
	p.sessions.Delete(id)
}

// GetAll sessions
func (p *provider) GetAll() map[string]se.Session {
	m := make(map[string]se.Session, 0)
	p.sessions.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(se.Session)
		return true
	})
	return m
}

// Clear session's vals
func (p *provider) Clear() {
	copyx := p.sessions
	p.sessions = &sync.Map{}
	go func(copy *sync.Map) {
		keys := make([]string, 0)
		copy.Range(func(key, value interface{}) bool {
			keys = append(keys, key.(string))
			return true
		})
		for _, key := range keys {
			copy.Delete(key)
		}
		runtime.GC()
	}(copyx)
}

func tmd5(text string) string {
	hashMd5 := md5.New()
	_, _ = io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func newSID() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	return strings.ToUpper(tmd5(tmd5(strconv.FormatInt(nano, 10)) + tmd5(strconv.FormatInt(rndNum, 10))))
}

// New return new session
func (p *provider) New(config *se.Config, listener *se.Listener) se.Session {
	sessionId := newSID()
	sess := newSession(sessionId, config.Timeout)
	p.sessions.Store(sessionId, sess)
	go func() {
		if listener != nil && listener.Created != nil {
			listener.Created(sess)
		}
	}()
	return sess
}

// Refresh session
func (p *provider) Refresh(session se.Session, config *se.Config, listener *se.Listener) {
	session.Renew(se.GetTimeout(config.Timeout))
	go func() {
		if listener != nil && listener.Refreshed != nil {
			listener.Refreshed(session)
		}
	}()
}

// Clean session
func (p *provider) Clean(_ *se.Config, listener *se.Listener) {
	go func() {
		for {
			p.cleanSession(listener)
			time.Sleep(time.Minute)
		}
	}()
}

func (p *provider) cleanSession(listener *se.Listener) {
	if len(p.GetAll()) <= 0 {
		return
	}
	for _, sess := range p.GetAll() {
		nu := time.Now().Unix()
		cu := sess.(*session).expiresTime.Unix()
		invalidate := false
		if cu == 0 {
			invalidate = true
		} else if cu <= nu {
			invalidate = true
		}
		if invalidate {
			sess.Invalidate()
			go func() {
				if listener != nil && listener.Invalidated != nil {
					listener.Invalidated(sess)
				}
			}()
		}
		if sess.Invalidated() {
			p.Del(sess.Id())
			go func() {
				if listener != nil && listener.Destroyed != nil {
					listener.Destroyed(sess)
				}
			}()
		}
	}
}
