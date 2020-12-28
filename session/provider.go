package session

import "net/http"

//Define Provider interface
type Provider interface {
	CookieName() string                                          //Set cookie Name
	Exists(id string) bool                                       //Get session is exists
	GetId(r *http.Request) string                                //Get Session Id
	Del(id string)                                               //Del Session
	Get(id string) Session                                       //Get Session
	GetAll() map[string]Session                                  //Get all Session
	Clear()                                                      //Clear all Sessions
	New(config *Config, listener *Listener) Session              //Get new Session
	Refresh(session Session, config *Config, listener *Listener) //Refresh Session
	Clean(config *Config, listener *Listener)                    //Clean Session
}
