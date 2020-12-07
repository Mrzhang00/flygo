package flygo

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Define app struct
type App struct {
	Id                     string                  //app id
	Name                   string                  //app name
	ConfigFile             string                  //conf file
	Config                 *YmlConfig              //yml config
	Logger                 *log                    //App logger
	staticCaches           staticCache             //static res cache
	staticMimeCaches       staticMimeCache         //static res mime cache
	viewCaches             viewCache               //view cache
	routes                 []handlerRouteCache     //list routes
	patternRoutes          patternRoute            //pattern route handlers
	variableRoutes         variableRoute           //variable route handlers
	filterRouteCaches      []filterRouteCache      //list filterRouteCaches
	beforeFilters          aroundFilter            //before filters
	afterFilters           aroundFilter            //after filters
	interceptorRouteCaches []interceptorRouteCache //list interceptorRouteCaches
	beforeInterceptors     aroundInterceptor       //before interceptors
	afterInterceptors      aroundInterceptor       //after interceptors
	FaviconIconHandler     StaticHandler           //default favicon ico handler
	StaticHandler          StaticHandler           //default static handler
	PreflightedHandler     Handler                 //preflighted handler

	NotFoundHandler         Handler //not found handler
	MethodNotAllowedHandler Handler //method not allowed  handler

	Caches                 appCache   //Application LEVEL cache
	middlewares            middleware //Middleware names
	filterMiddlewares      middleware //Filter Middleware names
	interceptorMiddlewares middleware //Interceptor Middleware names

	TemplateFuncs template.FuncMap //Template funcs
	SessionConfig *SessionConfig   //Session config
}

var defaultApp *App

func init() {
	id := "MASTER"
	defaultApp = NewAppWithId(id)
	if AppGroup.Listener().created != nil {
		AppGroup.Listener().created(newAppInfo(id, defaultApp))
	}
}

//Get default app
func GetApp() *App {
	return defaultApp
}

//New app with named seq index
func NewApp() *App {
	return NewAppWithId("")
}

//New app with named id
func NewAppWithId(id string) *App {
	if id == "" {
		id = AppGroup.nextAppId()
	}
	app := createApp(id)
	AppGroup.addWithId(id, app)
	return app
}

func createApp(id string) *App {
	if id == "" {
		panic("[App]app is empty")
	}
	return &App{
		Id:                      id,
		Name:                    AppGroup.prefix + id,
		ConfigFile:              fmt.Sprintf("flygo-%s.yml", id),
		Config:                  defaultYmlConfig(),
		staticCaches:            make(map[string][]byte),
		staticMimeCaches:        make(map[string]string),
		viewCaches:              make(map[string]string),
		routes:                  make([]handlerRouteCache, 0),
		patternRoutes:           make(map[string]map[string]patternHandlerRoute),
		variableRoutes:          make(map[string]map[string]variableHandlerRoute),
		filterRouteCaches:       make([]filterRouteCache, 0),
		beforeFilters:           make(map[string]filterRouteChain),
		afterFilters:            make(map[string]filterRouteChain),
		interceptorRouteCaches:  make([]interceptorRouteCache, 0),
		beforeInterceptors:      make(map[string]interceptorRouteChain),
		afterInterceptors:       make(map[string]interceptorRouteChain),
		FaviconIconHandler:      faviconIconHandler,
		StaticHandler:           staticHandler,
		PreflightedHandler:      preflightedHandler,
		NotFoundHandler:         notFoundHandler,
		MethodNotAllowedHandler: methodNotAllowedHandler,
		Caches:                  make(map[string]interface{}),
		middlewares:             make(map[string]int),
		filterMiddlewares:       make(map[string]int),
		interceptorMiddlewares:  make(map[string]int),
		TemplateFuncs:           make(map[string]interface{}),
		SessionConfig: &SessionConfig{
			SessionListener: &SessionListener{},
			Timeout:         time.Hour * 24 * 30, //1 month
		},
	}
}

//Init
func (a *App) inita() {
	a.checkConfig()
	a.setLoggers()
}

//Run the server
func (a *App) Run() {
	a.parseYml()

	//parse env
	a.parseEnv()

	//start route
	a.startRoute()

	//start filterRouteCache
	a.startFilter()

	//start interceptorRouteCache
	a.startInterceptor()

	//print banner
	a.printBanner()

	//parse bind address
	a.parseAddr()

	//print config
	a.DebugTrace(a.printConfig)

	//print middleware
	a.DebugTrace(a.printMiddleware)

	//print route
	a.DebugTrace(a.printRoute)

	//print filter
	a.DebugTrace(a.printFilter)

	//print interceptor
	a.DebugTrace(a.printInterceptor)

	//print session provider
	a.DebugTrace(a.printSessionProvider)

	//start server
	a.serve()
}

//Parse bind address
func (a *App) parseAddr() {
	host := a.Config.Flygo.Server.Host
	port := a.Config.Flygo.Server.Port
	if host == "" || host == "*" {
		host = "0.0.0.0"
	}
	minPort := 0
	maxPort := 65536
	if port < minPort || port > maxPort {
		a.Logger.Error("The port `%v` is invalid.[valid : %v - %v]", port, minPort, maxPort)
		os.Exit(0)
	}
}

//Start serve bind
func (a *App) serve() {
	defer func() {
		if re := recover(); re != nil {
			a.Logger.Error("%v", re)
		}
	}()
	host := a.Config.Flygo.Server.Host
	port := a.Config.Flygo.Server.Port
	tlsEnable := a.Config.Flygo.Server.Tls.Enable
	addr := host + ":" + strconv.Itoa(port)
	a.Logger.Info("Bind on %s", addr)
	a.Logger.Info("Server started")
	var err error
	if tlsEnable {
		//tls support
		certFile := a.Config.Flygo.Server.Tls.CertFile
		keyFile := a.Config.Flygo.Server.Tls.KeyFile
		err = http.ListenAndServeTLS(addr, certFile, keyFile, a.newDispatcher())
	} else {
		//http
		err = http.ListenAndServe(addr, a.newDispatcher())
	}
	if err != nil {
		a.Logger.Error(err.Error())
	}
}
