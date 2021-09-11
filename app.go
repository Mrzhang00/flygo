package flygo

import (
	"github.com/billcoding/flygo/config"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/rest"
	"github.com/billcoding/flygo/router"
	"github.com/sirupsen/logrus"
	"net/http"
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
		ConfigFile:  "app.yml",
		Config:      config.Default(),
		Logger:      logrus.New(),
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
	a.debugTrace()
	a.serve()
}

func (a *App) serve() {
	host := a.Config.Server.Host
	port := a.Config.Server.Port
	tlsEnable := a.Config.Server.TLS.Enable
	addr := host + ":" + strconv.Itoa(port)
	a.Logger.Infof("serve: Bind on %s", addr)
	a.Logger.Infof("serve: Server started")
	var err error
	server := &http.Server{
		Addr:              addr,
		Handler:           a.newDispatcher(),
		MaxHeaderBytes:    a.Config.Server.MaxHeaderSize,
		ReadTimeout:       a.Config.Server.ReadTimeout,
		ReadHeaderTimeout: a.Config.Server.ReadHeaderTimeout,
		WriteTimeout:      a.Config.Server.WriteTimeout,
		IdleTimeout:       a.Config.Server.IdleTimeout,
	}
	if tlsEnable {
		certFile := a.Config.Server.TLS.CertFile
		keyFile := a.Config.Server.TLS.KeyFile
		err = server.ListenAndServeTLS(certFile, keyFile)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		a.Logger.Errorf("serve: %v", err.Error())
	}
}
