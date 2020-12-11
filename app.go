package flygo

import (
	. "github.com/billcoding/flygo/log"
	. "github.com/billcoding/flygo/middleware"
	. "github.com/billcoding/flygo/rest"
	. "github.com/billcoding/flygo/router"
	"net/http"
	"os"
	"strconv"
)

//Define app struct
type App struct {
	ConfigFile     string          //Config file
	Config         *Config         //Yml config
	Logger         Logger          //App logger
	controllers    []Controller    //Rest controllers
	groups         []*Group        //groups
	routers        []*Router       //routers
	parsedRouters  *ParsedRouter   //parsed routers
	middlewares    []Middleware    //middlewares
	defaultMWState *defaultMWState //defaultMWState
}

var defaultApp = NewApp()

//GetApp
func GetApp() *App {
	return defaultApp
}

//NewApp
func NewApp() *App {
	return &App{
		ConfigFile:  "flygo.yml",
		Config:      defaultConfig(),
		Logger:      New("[FLYGO]"),
		controllers: make([]Controller, 0),
		groups:      make([]*Group, 0),
		routers:     []*Router{NewRouter()},
		parsedRouters: &ParsedRouter{
			Simples:  make(map[string]*Simple),
			Dynamics: make(map[string]map[string]*Dynamic),
		},
		middlewares:    make([]Middleware, 0),
		defaultMWState: &defaultMWState{},
	}
}

//Run
func (a *App) Run() {
	a.parseYml()
	a.parseEnv()
	a.printBanner()
	a.routeRestControllers()
	a.parseRouters()
	a.useDefaultMWs()
	a.serve()
}

//parseAddr
func (a *App) parseAddr() {
	host := a.Config.Host
	port := a.Config.Port
	if host == "" || host == "*" {
		host = "0.0.0.0"
	}
	minPort := 0
	maxPort := 65536
	if port < minPort || port > maxPort {
		a.Logger.Error("[parseAddr]The port `%v` is invalid.[valid : %v - %v]", port, minPort, maxPort)
		os.Exit(0)
	}
}

//serve
func (a *App) serve() {
	host := a.Config.Host
	port := a.Config.Port
	tlsEnable := a.Config.TLS.Enable
	addr := host + ":" + strconv.Itoa(port)
	a.Logger.Info("[serve]Bind on %s", addr)
	a.Logger.Info("[serve]Server started")
	var err error
	if tlsEnable {
		//tls support
		certFile := a.Config.TLS.CertFile
		keyFile := a.Config.TLS.KeyFile
		err = http.ListenAndServeTLS(addr, certFile, keyFile, a.newDispatcher())
	} else {
		//http
		err = http.ListenAndServe(addr, a.newDispatcher())
	}
	if err != nil {
		a.Logger.Error("[serve]%v", err.Error())
	}
}
