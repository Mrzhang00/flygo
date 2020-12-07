package flygo

import (
	"fmt"
	l "log"
	"os"
	"sync"
	"time"
)

//Define appGroup struct
type appGroup struct {
	sig          chan os.Signal
	ignoreMaster bool
	mu           sync.Mutex
	seq          int
	prefix       string
	listener     *listener
	appInfos     map[string]*AppInfo
}

type ListenerHandler func(ai *AppInfo)

//Default AppGroup
var AppGroup = &appGroup{
	sig:          make(chan os.Signal, 1),
	mu:           sync.Mutex{},
	seq:          0,
	ignoreMaster: false,
	prefix:       "FLYGO-APP-",
	listener:     &listener{},
	appInfos:     make(map[string]*AppInfo, 0),
}

//忽略Master
func (ag *appGroup) IgnoreMaster() {
	ag.ignoreMaster = true
}

//Get next AppId
func (ag *appGroup) nextAppId() string {
	ag.mu.Lock()
	defer ag.mu.Unlock()
	ag.seq++
	return fmt.Sprintf("%d", ag.seq)
}

//Return all appGroup apps
func (ag *appGroup) Apps() map[string]*AppInfo {
	return ag.appInfos
}

//Return appGroup apps zie
func (ag *appGroup) Count() int {
	return len(ag.appInfos)
}

//Return true if named id app exists
func (ag *appGroup) Have(id string) bool {
	_, have := ag.appInfos[id]
	return have
}

//Return identify app
func (ag *appGroup) add(app *App) *AppInfo {
	id := ag.nextAppId()
	ag.mu.Lock()
	defer ag.mu.Unlock()
	ai := newAppInfo(id, app)
	(*ag).appInfos[id] = ai
	if ag.listener.created != nil {
		ag.listener.created(ai)
	}
	return ai
}

//Return identify app
func (ag *appGroup) addWithId(id string, app *App) *AppInfo {
	if ag.Have(id) {
		l.Fatalf("[AppGroup]app[%s] was registered duplicate\n", id)
	}
	ag.mu.Lock()
	defer ag.mu.Unlock()
	ai := newAppInfo(id, app)
	(*ag).appInfos[id] = ai
	if ag.listener.created != nil {
		ag.listener.created(ai)
	}
	return ai
}

//Return identify appInfo
func (ag *appGroup) App(id string) *AppInfo {
	if !ag.Have(id) {
		l.Fatalf("[AppGroup]app[%s] was not registered\n", id)
	}
	return ag.appInfos[id]
}

//Run all registered app
func (ag *appGroup) Start() {
	bc := make(chan bool, 1)
	for _, ai := range ag.Apps() {
		if ag.ignoreMaster && ai.Id == "MASTER" {
			continue
		}
		go func() {
			time.AfterFunc(time.Millisecond*100, func() {
				bc <- true
				if ag.listener.started != nil {
					ag.listener.started(ai)
				}
			})
			ai.App.Config.Flygo.Banner.Enable = false
			ai.App.Run()
		}()
		<-bc
	}
	<-ag.sig
}

func (ag *appGroup) Stop() {
	if ag.listener.destoryed != nil {
		for _, ai := range ag.appInfos {
			if ag.listener.destoryed != nil {
				ag.listener.destoryed(ai)
			}
		}
	}
	ag.sig <- os.Interrupt
}
