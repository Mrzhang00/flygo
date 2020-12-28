package middleware

import (
	"github.com/billcoding/calls"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/log"
	"github.com/go-redis/redis"
)

//Define redisAuth struct
type redisAuth struct {
	logger  log.Logger
	key     string         //Redis key
	msg     string         //The msg
	code    int            //The code
	options *redis.Options //Redis options
	client  *redis.Client  //Redis client
}

//RedisAuth
func RedisAuth(options *redis.Options) Middleware {
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
	calls.NNil(ping.Err(), func() {
		r.logger.Warn("%v", ping.Err())
	})
	return r
}

//Name
func (*redisAuth) Name() string {
	return "RedisAuth"
}

//Type
func (r *redisAuth) Type() *Type {
	return TypeBefore
}

//Method
func (r *redisAuth) Method() Method {
	return MethodAny
}

//Pattern
func (r *redisAuth) Pattern() Pattern {
	return PatternAny
}

//Handler
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
		defer client.Close()
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

func GetRedisAuthData(c *c.Context) interface{} {
	return c.GetData("RedisAuth")
}

//Set return msg when unauthorized
func (r *redisAuth) Msg(msg string) *redisAuth {
	r.msg = msg
	return r
}

//Set return code when unauthorized
func (r *redisAuth) Code(code int) *redisAuth {
	r.code = code
	return r
}
