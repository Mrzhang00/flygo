package flygo

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Global app
var app *App

//Define app struct
type App struct {
	ConfigFile         string                //conf file
	Config             *YmlConfig            //yml config
	outLogger          *log.Logger           //app out log
	errLogger          *log.Logger           //app err log
	staticCaches       staticCache           //static res cache
	staticMimeCaches   staticMimeCache       //static res mime cache
	viewCaches         viewCache             //view cache
	routes             []patternHandlerRoute //list routes
	patternRoutes      patternRoute          //pattern route handlers
	variableRoutes     variableRoute         //variable route handlers
	froutes            []froute              //list froutes
	beforeFilters      aroundFilter          //before filters
	afterFilters       aroundFilter          //after filters
	iroutes            []iroute              //list iroutes
	beforeInterceptors aroundInterceptor     //before interceptors
	afterInterceptors  aroundInterceptor     //after interceptors
	FaviconIconHandler StaticHandler         //default favicon ico handler
	StaticHandler      StaticHandler         //default static handler
	PreflightedHandler Handler               //preflighted handler

	NotFoundHandler         Handler //not found handler
	MethodNotAllowedHandler Handler //method not allowed  handler

	Caches                 appCache   //Application LEVEL cache
	middlewares            middleware //Middleware names
	filterMiddlewares      middleware //Filter Middleware names
	interceptorMiddlewares middleware //Interceptor Middleware names

	TemplateFuncs   template.FuncMap //Template funcs
	SessionProvider SessionProvider  //Session Provider
	SessionConfig   *SessionConfig   //Session Config
}

func GetApp() *App {
	return app
}

func init() {
	NewApp()
}

func defaultApp() *App {
	return &App{
		ConfigFile:              "flygo.yml",
		Config:                  defaultYmlConfig(),
		staticCaches:            make(map[string][]byte),
		staticMimeCaches:        make(map[string]string),
		viewCaches:              make(map[string]string),
		routes:                  make([]patternHandlerRoute, 0),
		patternRoutes:           make(map[string]map[string]patternHandlerRoute),
		variableRoutes:          make(map[string]map[string]variableHandlerRoute),
		froutes:                 make([]froute, 0),
		beforeFilters:           make(map[string]filterRouteChain),
		afterFilters:            make(map[string]filterRouteChain),
		iroutes:                 make([]iroute, 0),
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
			Timeout: time.Hour * 24,
		},
	}
}

//Return new App
func NewApp() *App {
	if app != nil {
		return app
	}
	app = defaultApp()
	return app
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

	//start froute
	a.startFilter()

	//start iroute
	a.startInterceptor()

	//print banner
	a.printBanner()

	//parse bind address
	a.parseAddr()

	if a.Config.Flygo.Dev.Debug {

		//print middleware
		a.printMiddleware()

		//print route
		a.printRoute()

		//print filter
		a.printFilter()

		//print interceptor
		a.printInterceptor()

		//print session provider
		a.printSessionProvider()
	}

	//when sessionProvider is nil, turnoff sessionEnable
	if a.SessionProvider == nil {
		a.Config.Flygo.Session.Enable = false
	}

	//start server
	a.serve()
}

//Parse bind address
func (a *App) parseAddr() {
	host := app.Config.Flygo.Server.Host
	port := app.Config.Flygo.Server.Port
	if host == "" || host == "*" {
		host = "0.0.0.0"
	}
	minPort := 0
	maxPort := 65536
	if port < minPort || port > maxPort {
		a.Error("The port `%v` is invalid.[valid : %v - %v]", port, minPort, maxPort)
		os.Exit(0)
	}
}

//Start serve bind
func (a *App) serve() {
	defer func() {
		if re := recover(); re != nil {
			a.Error("%v", re)
		}
	}()
	log.SetPrefix("[FLYGO]")
	host := app.Config.Flygo.Server.Host
	port := app.Config.Flygo.Server.Port
	tlsEnable := app.Config.Flygo.Server.Tls.Enable
	addr := host + ":" + strconv.Itoa(port)
	log.Printf("Bind on %s\n", addr)
	log.Println("Server started")
	var err error
	if tlsEnable {
		//tls support
		certFile := app.Config.Flygo.Server.Tls.CertFile
		keyFile := app.Config.Flygo.Server.Tls.KeyFile
		err = http.ListenAndServeTLS(addr, certFile, keyFile, newDispatcher())
	} else {
		//http
		err = http.ListenAndServe(addr, newDispatcher())
	}
	if err != nil {
		a.Error(err.Error())
	}
}
