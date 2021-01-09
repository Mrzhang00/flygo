package middleware

import (
	"fmt"
	"github.com/billcoding/calls"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/log"
	"github.com/go-redis/redis"
	"math/rand"
	"strings"
	"time"
)

type redisToken struct {
	logger  log.Logger
	key     string
	appKey  string
	expires time.Duration
	length  int
	msg     string
	code    int
	options *redis.Options
	client  *redis.Client
}

// RedisToken return new redisToken
func RedisToken(options *redis.Options) *redisToken {
	client := redis.NewClient(options)
	ping := client.Ping()
	rt := &redisToken{
		logger:  log.New("[RedisToken]"),
		key:     "auth:authorization",
		appKey:  "auth:apps",
		expires: 24 * time.Hour,
		length:  64,
		msg:     "Error credentials",
		code:    400,
		options: options,
		client:  client,
	}
	calls.NNil(ping.Err(), func() {
		rt.logger.Warn("%v", ping.Err())
	})
	return rt
}

// Name implements
func (*redisToken) Name() string {
	return "RedisToken"
}

// Type implements
func (r *redisToken) Type() *Type {
	return TypeHandler
}

// Method implements
func (r *redisToken) Method() Method {
	return MethodGet
}

// Pattern implements
func (r *redisToken) Pattern() Pattern {
	return "/authorization/authorize"
}

// Handler implements
func (r *redisToken) Handler() func(c *c.Context) {
	type jd struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	type jdt struct {
		Msg   string `json:"msg"`
		Code  int    `json:"code"`
		Token string `json:"token"`
	}
	getJson := func(msg string, code int) *jd {
		return &jd{
			Msg:  msg,
			Code: code,
		}
	}
	getJsonToken := func(token string) *jdt {
		return &jdt{
			Msg:   "success",
			Code:  0,
			Token: token,
		}
	}
	type model struct {
		AppKey    string `binding:"name(appKey)" validate:"required(T) minlength(1) message(appKey is empty)"`
		AppSecret string `binding:"name(appSecret)" validate:"required(T) minlength(1) message(appSecret is empty)"`
	}
	return func(c *c.Context) {
		m := model{}
		c.BindAndValidate(&m, func() {
			hGet := r.client.HGet(r.appKey, m.AppKey)
			if err := hGet.Err(); err != nil {
				c.JSON(getJson(r.msg, r.code))
				return
			}
			if m.AppSecret != hGet.Val() {
				c.JSON(getJson(r.msg, r.code))
				return
			}
			token := newToken(r.length)
			r.client.Set(r.key+":"+token, getAppJson(m.AppKey, m.AppSecret), r.expires)
			c.JSON(getJsonToken(token))
		})
	}
}

// Key set
func (r *redisToken) Key(key string) *redisToken {
	r.key = key
	return r
}

// AppKey set
func (r *redisToken) AppKey(appKey string) *redisToken {
	r.appKey = appKey
	return r
}

// Expires set
func (r *redisToken) Expires(expires time.Duration) *redisToken {
	r.expires = expires
	return r
}

// Length set
func (r *redisToken) Length(length int) *redisToken {
	r.length = length
	return r
}

// Msg set
func (r *redisToken) Msg(msg string) *redisToken {
	r.msg = msg
	return r
}

// Code set
func (r *redisToken) Code(code int) *redisToken {
	r.code = code
	return r
}

// Options set
func (r *redisToken) Options(options *redis.Options) *redisToken {
	r.options = options
	return r
}

func getAppJson(appKey, appSecret string) string {
	return fmt.Sprintf("{\"%s\":\"%s\",\"%s\":\"%s\"}", "appKey", appKey, "appSecret", appSecret)
}

func newToken(length int) string {
	symbols := []byte(`0123456789abcdefghijklmnopqrstuvwxyz`)
	token := ""
	for i := 0; i < length; i++ {
		cc := fmt.Sprintf("%c", symbols[rand.Intn(len(symbols))])
		if rand.Intn(2) == 1 {
			token += strings.ToUpper(cc)
		} else {
			token += cc
		}
	}
	return token
}
