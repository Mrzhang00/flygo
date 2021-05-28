package flygo

import (
	"github.com/billcoding/flygo/config"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/rest"
	"github.com/billcoding/flygo/router"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

// App struct
type App struct {
	ConfigFile     string
	Config         *config.Config
	Logger         *logrus.Logger
	controllers    []rest.Controller
	groups         []*router.Group
	routers        []*router.Router
	parsedRouters  *router.ParsedRouter
	middlewares    []middleware.Middleware
	middlewareMap  map[string]middleware.Middleware
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
		Config:      config.Default(),
		Logger:      logrus.StandardLogger(),
		controllers: make([]rest.Controller, 0),
		groups:      make([]*router.Group, 0),
		routers:     []*router.Router{router.NewRouter()},
		parsedRouters: &router.ParsedRouter{
			Simples:  make(map[string]*router.Simple),
			Dynamics: make(map[string]map[string]*router.Dynamic),
		},
		middlewares:   make([]middleware.Middleware, 6),
		middlewareMap: make(map[string]middleware.Middleware, 0),
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
		a.Logger.Errorf("[parseAddr]The port `%v` is invalid.[valid : %v - %v]", port, minPort, maxPort)
		os.Exit(0)
	}
}

func (a *App) serve() {
	host := a.Config.Flygo.Server.Host
	port := a.Config.Flygo.Server.Port
	tlsEnable := a.Config.Flygo.Server.TLS.Enable
	addr := host + ":" + strconv.Itoa(port)
	a.Logger.Infof("[Serve]Bind on %s", addr)
	a.Logger.Infof("[Serve]Server started")
	var err error
	server := &http.Server{
		Addr:              addr,
		Handler:           a.newDispatcher(),
		MaxHeaderBytes:    a.Config.Server.MaxHeaderSize,
		ReadTimeout:       a.Config.Server.Timeout.Read,
		ReadHeaderTimeout: a.Config.Server.Timeout.ReadHeader,
		WriteTimeout:      a.Config.Server.Timeout.Write,
		IdleTimeout:       a.Config.Server.Timeout.Idle,
	}
	if tlsEnable {
		certFile := a.Config.Flygo.Server.TLS.CertFile
		keyFile := a.Config.Flygo.Server.TLS.KeyFile
		err = server.ListenAndServeTLS(certFile, keyFile)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		a.Logger.Errorf("[Serve]%v", err.Error())
	}
}
