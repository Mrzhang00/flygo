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

//Define provider struct
type provider struct {
	sessions *sync.Map //map[string]se.Session
}

//Provider
func Provider() se.Provider {
	p := &provider{
		sessions: &sync.Map{},
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
	_, have := p.sessions.Load(id)
	return have
}

//Get
func (p *provider) Get(id string) se.Session {
	value, have := p.sessions.Load(id)
	if !have {
		return nil
	}
	return value.(se.Session)
}

//Del
func (p *provider) Del(id string) {
	p.sessions.Delete(id)
}

//GetAll
func (p *provider) GetAll() map[string]se.Session {
	m := make(map[string]se.Session, 0)
	p.sessions.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(se.Session)
		return true
	})
	return m
}

//Clear
func (p *provider) Clear() {
	copy := p.sessions
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
	}(copy)
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
	sess := newSession(sessionId, config.Timeout)
	p.sessions.Store(sessionId, sess)
	go func() {
		if listener != nil && listener.Created != nil {
			listener.Created(sess)
		}
	}()
	return sess
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
	//go func() {
	//	for {
	//		p.cleanSession(listener)
	//		time.Sleep(time.Second)
	//	}
	//}()
}

//cleanSession
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
				if listener != nil && listener.Destoryed != nil {
					listener.Destoryed(sess)
				}
			}()
		}
	}
}
