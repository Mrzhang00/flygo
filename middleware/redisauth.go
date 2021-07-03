package middleware

import (
	"fmt"
	"github.com/billcoding/flygo/context"
	"github.com/go-redis/redis"
	"os"
	"regexp"
)

type redisAuth struct {
	key             string
	msg             string
	code            int
	options         *redis.Options
	client          *redis.Client
	excludePaths    []string
	excludePatterns []string
}

// RedisAuth return new redisAuth
func RedisAuth(options *redis.Options) *redisAuth {
	client := redis.NewClient(options)
	ping := client.Ping()
	r := &redisAuth{
		key:     "auth:authorization:",
		msg:     "Unauthorized",
		code:    777,
		options: options,
		client:  client,
	}
	if ping.Err() != nil {
		fmt.Fprintln(os.Stderr, "redis:", ping.Err())
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
func (r *redisAuth) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		if r.excludePaths != nil && len(r.excludePaths) > 0 {
			exMap := make(map[string]struct{})
			for _, e := range r.excludePaths {
				exMap[e] = struct{}{}
			}
			p := ctx.Request.URL.Path
			for k := range exMap {
				if p == k {
					ctx.Chain()
					return
				}
			}
		}
		if r.excludePatterns != nil && len(r.excludePatterns) > 0 {
			exMap := make(map[string]struct{})
			for _, e := range r.excludePatterns {
				exMap[e] = struct{}{}
			}
			p := ctx.Request.URL.Path
			for k := range exMap {
				re, err := regexp.Compile(k)
				if err != nil {
					panic(err)
				}
				if re.MatchString(p) {
					ctx.Chain()
					return
				}
			}
		}
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
		authorization := ctx.Request.Header.Get("Authorization")
		if authorization == "" {
			ctx.JSON(getJd(r))
			return
		}
		getCmd := client.Get(r.key + authorization)
		authorizationData := getCmd.Val()
		if authorizationData == "" {
			ctx.JSON(getJd(r))
			return
		}
		ctx.SetData("RedisAuth", authorizationData)
		ctx.Chain()
	}
}

// RedisAuthData get data
func RedisAuthData(ctx *context.Context) interface{} {
	return ctx.GetData("RedisAuth")
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

// ExcludePaths set
func (r *redisAuth) ExcludePaths(excludePath ...string) *redisAuth {
	r.excludePaths = excludePath
	return r
}

// ExcludePatterns set
func (r *redisAuth) ExcludePatterns(excludePattern ...string) *redisAuth {
	r.excludePatterns = excludePattern
	return r
}
