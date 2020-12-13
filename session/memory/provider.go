package memory

import (
	"crypto/md5"
	"fmt"
	se "github.com/billcoding/flygo/session"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//Define provider struct
type provider struct {
	sessions map[string]se.Session
	mu       sync.RWMutex
}

//Provider
func Provider() se.Provider {
	p := &provider{
		sessions: make(map[string]se.Session),
		mu:       sync.RWMutex{},
	}
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

//Exists
func (p *provider) Exists(id string) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	_, have := p.sessions[id]
	return have
}

//Get
func (p *provider) Get(id string) se.Session {
	p.mu.RLock()
	defer p.mu.RUnlock()
	session, _ := p.sessions[id]
	return session
}

//Del
func (p *provider) Del(id string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.sessions, id)
}

//GetAll
func (p *provider) GetAll() map[string]se.Session {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.sessions
}

//Clear
func (p *provider) Clear() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.sessions = make(map[string]se.Session)
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
	p.mu.Lock()
	defer p.mu.Unlock()
	sessionId := newSID()
	session := newSession(sessionId, config.Timeout)
	p.sessions[sessionId] = session
	go func() {
		if listener != nil && listener.Created != nil {
			listener.Created(session)
		}
	}()
	return session
}

//Refresh
func (p *provider) Refresh(session se.Session, config *se.Config, listener *se.Listener) {
	session.Renew(config.Timeout)
	go func() {
		if listener != nil && listener.Refreshed != nil {
			listener.Refreshed(session)
		}
	}()
}

//Clean
func (p *provider) Clean(config *se.Config, listener *se.Listener) {
	blocked := make(chan bool, 1)
	go p.invalidatedSession(listener)
	go p.destoryedSession(listener)
	<-blocked
}

//invalidatedSession
func (p *provider) invalidatedSession(listener *se.Listener) {
	for {
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
		}
		time.Sleep(time.Microsecond * 10)
	}
}

//destoryedSession
func (p *provider) destoryedSession(listener *se.Listener) {
	for {
		for sessionId, sess := range p.GetAll() {
			cs := sess.(*session)
			if cs.Invalidated() {
				p.Del(sessionId)
				go func() {
					if listener != nil && listener.Destoryed != nil {
						listener.Destoryed(sess)
					}
				}()
			}
		}
	}
	time.Sleep(time.Microsecond * 10)
}
