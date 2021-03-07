package session

import (
	"fmt"
	"time"
)

// Session interface
type Session interface {
	// Id session id
	Id() string
	// Renew session
	Renew(lifeTime time.Duration)
	// Invalidate session
	Invalidate()
	// Invalidated session
	Invalidated() bool
	// Get named val
	Get(name string) interface{}
	// GetAll session's values
	GetAll() map[string]interface{}
	// Set named val
	Set(name string, val interface{})
	// SetAll values
	SetAll(data map[string]interface{}, flush bool)
	// Del named val
	Del(name string)
	// Clear session's val
	Clear()
}

var defaultTimeout = 60 * 24 // one Day

// GetTimeout session timeout
func GetTimeout(minute int) time.Duration {
	if minute <= 0 {
		minute = defaultTimeout
	}
	du, _ := time.ParseDuration(fmt.Sprintf("%dm", minute))
	return du
}
