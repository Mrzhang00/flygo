package flygo

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Define address struct
type address struct {
	host string
	port int
}

//Global app
var app *App

//Define app struct
type App struct {
	log                        log                                        //app log
	tls                        bool                                       //TLS mode
	tlsCertFile                string                                     //TLS cert file path
	tlsKeyFile                 string                                     //TLS key file path
	address                    address                                    //bind address
	faviconIcon                bool                                       //auto handler favicon icon
	statics                    map[string]string                          //static registry
	staticCaches               map[string][]byte                          //static res cache
	staticMimeCaches           map[string]string                          //static res mime cache
	viewCaches                 map[string]string                          //view cache
	configs                    map[string]interface{}                     //global, static, view configs
	patternRoutes              map[string]map[string]patternHandlerRoute  //pattern route handlers
	variableRoutes             map[string]map[string]variableHandlerRoute //variable route handlers
	beforeFilters              map[string]filterRouteChain                //before filters
	afterFilters               map[string]filterRouteChain                //after filters
	beforeInterceptors         map[string]interceptorRouteChain           //before interceptors
	afterInterceptors          map[string]interceptorRouteChain           //after interceptors
	faviconIconHandler         StaticHandler                              //default favicon ico handler
	defaultHandler             Handler                                    //default handler
	defaultStaticHandler       StaticHandler                              //default static handler
	preflightedHandler         Handler                                    //preflighted handler
	requestNotSupportedHandler Handler                                    //default not supported method handler
	banner                     bool                                       //Enable banner
	caches                     map[string]interface{}                     //Application LEVEL cache
	middlewares                map[string]int                             //Middleware names
	filterMiddlewares          map[string]int                             //Filter Middleware names
	interceptorMiddlewares     map[string]int                             //Interceptor Middleware names
	sessionEnable              bool                                       //Enable session
	sessionProvider            SessionProvider                            //Session Provider
	SessionConfig              *SessionConfig                             //Session Config
}

func init() {
	NewApp()
}

//Return app
func GetApp() *App {
	return app
}

//Return new App
func NewApp() *App {
	if app != nil {
		return app
	}

	app = &App{
		log:                    log{},
		tls:                    false,
		tlsCertFile:            "",
		tlsKeyFile:             "",
		address:                address{host: "", port: 80},
		statics:                make(map[string]string),
		staticCaches:           make(map[string][]byte),
		staticMimeCaches:       make(map[string]string),
		viewCaches:             make(map[string]string),
		configs:                make(map[string]interface{}),
		patternRoutes:          make(map[string]map[string]patternHandlerRoute),
		variableRoutes:         make(map[string]map[string]variableHandlerRoute),
		beforeFilters:          make(map[string]filterRouteChain),
		afterFilters:           make(map[string]filterRouteChain),
		beforeInterceptors:     make(map[string]interceptorRouteChain),
		afterInterceptors:      make(map[string]interceptorRouteChain),
		banner:                 true,
		caches:                 make(map[string]interface{}),
		middlewares:            make(map[string]int),
		filterMiddlewares:      make(map[string]int),
		interceptorMiddlewares: make(map[string]int),
		SessionConfig: &SessionConfig{
			Timeout: time.Hour * 24,
		},
	}
	app.init()
	return app
}

//Init
func (a *App) init() {
	a.initStatic()
	a.initDefaultHandler()
	a.initConfig()
}

//Init static
func (a *App) initStatic() {
	a.statics["css"] = contentTypeCSS
	a.statics["js"] = contentTypeJS
	a.statics["jpg"] = contentTypeJpg
	a.statics["png"] = contentTypePng
	a.statics["gif"] = contentTypeGif
	a.statics["ico"] = contentTypeIco
}

//Init config
func (a *App) initConfig() {
	a.SetWebRoot("./")
	a.SetContextPath("")

	a.SetStaticEnable(false)
	a.SetStaticCache(true)
	a.SetStaticPattern("/static")
	a.SetStaticPrefix("static")

	a.SetViewEnable(false)
	a.SetViewCache(true)
	a.SetViewPrefix("templates")
	a.SetViewSuffix("html")

	a.SetTemplateEnable(false)
	a.SetTemplateFuncs(nil)
	a.SetTemplateDelimLeft("{{")
	a.SetTemplateDelimRight("}}")

	a.SetSessionEnable(false)

	a.SetValidateErrCode(1)
}

//Init default handler
func (a *App) initDefaultHandler() {
	a.defaultHandler = func(c *Context) {
		c.ResponseWriter.WriteHeader(404)
	}
	a.defaultStaticHandler = func(c *Context, contentType, resourcePath string) {
		data := a.staticCaches[resourcePath]
		if data != nil {
			c.BinaryWith(data, a.staticMimeCaches[resourcePath])
			return
		}
		buffer, err := ioutil.ReadFile(resourcePath)
		if err != nil {
			c.ResponseWriter.WriteHeader(404)
			return
		}
		if a.GetStaticCache() {
			a.staticCaches[resourcePath] = buffer
			a.staticMimeCaches[resourcePath] = contentType
		}
		c.BinaryWith(buffer, contentType)
	}
	a.requestNotSupportedHandler = func(c *Context) {
		c.ResponseWriter.WriteHeader(405)
	}
	a.preflightedHandler = func(c *Context) {
		c.ResponseHeader.Set("Allow", "GET,POST,DELETE,PUT,PATCH,HEAD,OPTIONS")
	}
}

//Enable TLS mode
func (a *App) TLS(certFile, keyFile string) {
	a.tls = true
	a.tlsCertFile = certFile
	a.tlsKeyFile = keyFile
}

//Enable default favicon ico handler
func (a *App) FaviconIco() *App {
	a.faviconIcon = true
	a.faviconIconHandler = func(c *Context, contentType, resourcePath string) {
		a.defaultStaticHandler(c, contentTypeIco, resourcePath)
	}
	return a
}

//Run as the server
func (a *App) RunAs(host string, port int) {
	a.address.host = host
	a.address.port = port
	a.Run()
}

//Enable banner print
func (a *App) Banner(banner bool) *App {
	a.banner = banner
	return a
}

//Run the server
func (a *App) Run() {
	//print banner
	a.printBanner()

	//parse env
	a.parseEnv()

	//parse bind address
	a.parseAddr()

	//print static
	a.printStatic()

	//print middleware
	a.printMiddleware()

	//print route
	a.printRoute()

	//print filter
	a.printFilter()

	//print interceptor
	a.printInterceptor()

	//print config
	a.printConfig()

	//print session provider
	a.printSessionProvider()

	//when sessionProvider is nil, turnoff sessionEnable
	if a.sessionProvider == nil {
		a.sessionEnable = false
	}

	//start server
	a.serve()
}

//Parse bind address
func (a *App) parseAddr() {
	host := a.address.host
	port := a.address.port
	if host == "" || host == "*" {
		host = "0.0.0.0"
	}
	minPort := 0
	maxPort := 65536
	if port < minPort || port > maxPort {
		a.log.fatal("The port `%v` is invalid.[valid : %v - %v]", port, minPort, maxPort)
		os.Exit(0)
	}
	a.address.host = host
	a.address.port = port
}

//Start serve bind
func (a *App) serve() {
	defer func() {
		if re := recover(); re != nil {
			a.log.fatal("%v", re)
		}
	}()
	addr := a.address.host + ":" + strconv.Itoa(a.address.port)
	a.LogInfo("Bind on %s", addr)
	a.LogInfo("Server started")
	var err error
	if a.tls {
		//tls support
		err = http.ListenAndServeTLS(addr, a.tlsCertFile, a.tlsKeyFile, newDispatcher())
	} else {
		//http
		err = http.ListenAndServe(addr, newDispatcher())
	}
	if err != nil {
		a.log.fatal(err.Error())
	}
}
