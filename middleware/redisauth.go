package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/log"
	"github.com/go-redis/redis"
)

type redisAuth struct {
	logger  log.Logger
	key     string
	msg     string
	code    int
	options *redis.Options
	client  *redis.Client
}

// RedisAuth return new redisAuth
func RedisAuth(options *redis.Options) *redisAuth {
	client := redis.NewClient(options)
	ping := client.Ping()
	r := &redisAuth{
		logger:  log.New("[RedisAuth]"),
		key:     "auth:authorization:",
		msg:     "Unauthorized",
		code:    777,
		options: options,
		client:  client,
	}
	if ping.Err() != nil {
		r.logger.Warn("%v", ping.Err())
	}
	return r
}

// Name implements
func (*redisAuth) Name() string {
	return "RedisAuth"
}

// Type implements
func (r *redisAuth) Type() *Type {
	return TypeBefore
}

// Method implements
func (r *redisAuth) Method() Method {
	return MethodAny
}

// Pattern implements
func (r *redisAuth) Pattern() Pattern {
	return PatternAny
}

// Handler implements
func (r *redisAuth) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		type jd struct {
			Msg  string `json:"msg"`
			Code int    `json:"code"`
		}
		getJd := func(r *redisAuth) *jd {
			return &jd{
				Msg:  r.msg,
				Code: r.code,
			}
		}
		client := redis.NewClient(r.options)
		defer func() {
			_ = client.Close()
		}()
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(getJd(r))
			return
		}
		getCmd := client.Get(r.key + authorization)
		authorizationData := getCmd.Val()
		if authorizationData == "" {
			c.JSON(getJd(r))
			return
		}
		setRedisAuthData(c, authorizationData)
	}
}

func setRedisAuthData(c *c.Context, data interface{}) {
	c.SetData("RedisAuth", data)
}

// GetRedisAuthData get data
func GetRedisAuthData(c *c.Context) interface{} {
	return c.GetData("RedisAuth")
}

// Msg set
func (r *redisAuth) Msg(msg string) *redisAuth {
	r.msg = msg
	return r
}

// Code set
func (r *redisAuth) Code(code int) *redisAuth {
	r.code = code
	return r
}
