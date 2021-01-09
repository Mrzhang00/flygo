package session

import (
	"fmt"
	"time"
)

var defaultTimeout = 60 * 24 // one Day

func GetTimeout(minute int) time.Duration {
	if minute <= 0 {
		minute = defaultTimeout
	}
	du, _ := time.ParseDuration(fmt.Sprintf("%dm", minute))
	return du
}

type Session interface {
	Id() string
	Renew(lifeTime time.Duration)
	Invalidate()
	Invalidated() bool
	Get(name string) interface{}
	GetAll() map[string]interface{}
	Set(name string, val interface{})
	SetAll(data map[string]interface{}, flush bool)
	Del(name string)
	Clear()
}
