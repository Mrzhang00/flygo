package flygo

import (
	. "github.com/billcoding/flygo/log"
	. "github.com/billcoding/flygo/middleware"
	. "github.com/billcoding/flygo/rest"
	. "github.com/billcoding/flygo/router"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Define app struct
type App struct {
	ServerConfig   *serverConfig   //Server Config
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

//Define ServerConfig struct
type serverConfig struct {
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

var defaultApp = NewApp()

//GetApp
func GetApp() *App {
	return defaultApp
}

//NewApp
func NewApp() *App {
	return &App{
		ServerConfig: &serverConfig{
			ReadTimeout:       time.Hour,
			ReadHeaderTimeout: time.Hour,
			WriteTimeout:      time.Hour,
			IdleTimeout:       0,
			MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		},
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
		middlewares: make([]Middleware, 6, 6),
		defaultMWState: &defaultMWState{
			header: true,
		},
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
	host := a.Config.Server.Host
	port := a.Config.Server.Port
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
	host := a.Config.Server.Host
	port := a.Config.Server.Port
	tlsEnable := a.Config.Server.TLS.Enable
	addr := host + ":" + strconv.Itoa(port)
	a.Logger.Info("[serve]Bind on %s", addr)
	a.Logger.Info("[serve]Server started")
	var err error
	server := &http.Server{
		Addr:              addr,
		Handler:           a.newDispatcher(),
		ReadTimeout:       a.ServerConfig.ReadTimeout,
		ReadHeaderTimeout: a.ServerConfig.ReadHeaderTimeout,
		WriteTimeout:      a.ServerConfig.WriteTimeout,
		IdleTimeout:       a.ServerConfig.IdleTimeout,
		MaxHeaderBytes:    a.ServerConfig.MaxHeaderBytes,
		ErrorLog:          log.New(os.Stderr, "[http]", log.LstdFlags),
	}
	if tlsEnable {
		//tls support
		certFile := a.Config.Server.TLS.CertFile
		keyFile := a.Config.Server.TLS.KeyFile
		err = server.ListenAndServeTLS(certFile, keyFile)
	} else {
		//http
		err = server.ListenAndServe()
	}
	if err != nil {
		a.Logger.Error("[serve]%v", err.Error())
	}
}
