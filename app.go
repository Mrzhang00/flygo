package flygo

import (
	cfg "github.com/billcoding/flygo/config"
	l "github.com/billcoding/flygo/log"
	mw "github.com/billcoding/flygo/middleware"
	rt "github.com/billcoding/flygo/rest"
	rr "github.com/billcoding/flygo/router"
	"log"
	"net/http"
	"os"
	"strconv"
)

// App struct
type App struct {
	ConfigFile     string
	Config         *cfg.Config
	Logger         l.Logger
	controllers    []rt.Controller
	groups         []*rr.Group
	routers        []*rr.Router
	parsedRouters  *rr.ParsedRouter
	middlewares    []mw.Middleware
	defaultMWState *defaultMWState
}

var defaultApp = NewApp()

// GetApp return default App
func GetApp() *App {
	return defaultApp
}

// NewApp return new App
func NewApp() *App {
	return &App{
		ConfigFile:  "flygo.yml",
		Config:      cfg.Default(),
		Logger:      l.New("[FLYGO]"),
		controllers: make([]rt.Controller, 0),
		groups:      make([]*rr.Group, 0),
		routers:     []*rr.Router{rr.NewRouter()},
		parsedRouters: &rr.ParsedRouter{
			Simples:  make(map[string]*rr.Simple),
			Dynamics: make(map[string]map[string]*rr.Dynamic),
		},
		middlewares: make([]mw.Middleware, 6, 6),
		defaultMWState: &defaultMWState{
			header: true,
		},
	}
}

// Run App
func (a *App) Run() {
	a.parseYml()
	a.parseEnv()
	a.printBanner()
	a.routeRestControllers()
	a.parseRouters()
	a.useDefaultMWs()
	a.parseConfig()
	a.debugTrace()
	a.serve()
}

func (a *App) parseAddr() {
	port := a.Config.Flygo.Server.Port
	minPort := 0
	maxPort := 65536
	if port < minPort || port > maxPort {
		a.Logger.Error("[parseAddr]The port `%v` is invalid.[valid : %v - %v]", port, minPort, maxPort)
		os.Exit(0)
	}
}

func (a *App) serve() {
	host := a.Config.Flygo.Server.Host
	port := a.Config.Flygo.Server.Port
	tlsEnable := a.Config.Flygo.Server.TLS.Enable
	addr := host + ":" + strconv.Itoa(port)
	a.Logger.Info("[Serve]Bind on %s", addr)
	a.Logger.Info("[Serve]Server started")
	var err error
	server := &http.Server{
		Addr:              addr,
		Handler:           a.newDispatcher(),
		MaxHeaderBytes:    a.Config.Server.MaxHeaderSize,
		ReadTimeout:       a.Config.Server.Timeout.Read,
		ReadHeaderTimeout: a.Config.Server.Timeout.ReadHeader,
		WriteTimeout:      a.Config.Server.Timeout.Write,
		IdleTimeout:       a.Config.Server.Timeout.Idle,
		ErrorLog:          log.New(os.Stderr, "[http]", log.LstdFlags),
	}
	if tlsEnable {

		certFile := a.Config.Flygo.Server.TLS.CertFile
		keyFile := a.Config.Flygo.Server.TLS.KeyFile
		err = server.ListenAndServeTLS(certFile, keyFile)
	} else {

		err = server.ListenAndServe()
	}
	if err != nil {
		a.Logger.Error("[Serve]%v", err.Error())
	}
}
