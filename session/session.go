package session

import (
	"time"
)

//Define Session interface
type Session interface {
	Id() string                                     //Get session Id
	Renew(lifeTime time.Duration)                   //Renew session
	Invalidate()                                    //Invalidate session
	Invalidated() bool                              //Invalidated
	Get(name string) interface{}                    //Get data
	GetAll() map[string]interface{}                 //Get all data
	Set(name string, val interface{})               //Set data
	SetAll(data map[string]interface{}, flush bool) //Set all data
	Del(name string)                                //Del by name
	Clear()                                         //Clear all data
}
