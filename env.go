package flygo

import (
	"os"
	"strconv"
	"time"
)

const (
	flggoConfig = "FLYGO_CONFIG" //env config file

	flygoDevDebug = "FLYGO_DEV_DEBUG" //env dev debug

	flygoServerHost        = "FLYGO_SERVER_HOST"         //env server host
	flygoServerPort        = "FLYGO_SERVER_PORT"         //env server port
	flygoServerContextPath = "FLYGO_SERVER_CONTEXT_PATH" //env server contextPath
	flygoServerWebRoot     = "FLYGO_SERVER_WEB_ROOT"     //env server webRoot

	flygoBannerEnable = "FLYGO_BANNER_ENABLE" //env banner enable
	flygoBannerType   = "FLYGO_BANNER_TYPE"   //env banner type
	flygoBannerText   = "FLYGO_BANNER_TEXT"   //env banner text
	flygoBannerFile   = "FLYGO_BANNER_ENABLE" //env banner file

	flygoServerTlsEnable   = "FLYGO_SERVER_TLS_ENABLE"    //env server tls enable
	flygoServerTlsCertFile = "FLYGO_SERVER_TLS_CERT_FILE" //env server tls cert file
	flygoServerTlsKeyFile  = "FLYGO_SERVER_TLS_KEY_FILE"  //env server tls key file

	flygoStaticEnable  = "FLYGO_STATIC_ENABLE"  //env static enable
	flygoStaticPattern = "FLYGO_STATIC_PATTERN" //env static pattern
	flygoStaticPrefix  = "FLYGO_STATIC_PREFIX"  //env static prefix
	flygoStaticCache   = "FLYGO_STATIC_CACHE"   //env static cache

	flygoStaticFaviconEnable = "FLYGO_STATIC_FAVICON_ENABLE" //env static favicon enable

	flygoViewEnable = "FLYGO_VIEW_ENABLE" //env view enable
	flygoViewPrefix = "FLYGO_VIEW_PREFIX" //env view prefix
	flygoViewSuffix = "FLYGO_VIEW_SUFFIX" //env view suffix
	flygoViewCache  = "FLYGO_VIEW_CACHE"  //env view cache

	flygoTemplateEnable      = "FLYGO_TEMPLATE_ENABLE"       //env template
	flygoTemplateDelimsLeft  = "FLYGO_TEMPLATE_DELIMS_LEFT"  //env template delims left
	flygoTemplateDelimsRight = "FLYGO_TEMPLATE_DELIMS_RIGHT" //env template delims right

	flygoSessionEnable  = "FLYGO_SESSION_ENABLE"  //env session enable
	flygoSessionTimeout = "FLYGO_SESSION_TIMEOUT" //env session timeout
)

//set config
func (a *App) setConfig() {
	config := stringEnv(flggoConfig)
	if config != "" {
		a.ConfigFile = config
	}
}

//set dev debug
func (a *App) setDevDebug() {
	debug := stringEnv(flygoDevDebug)
	if debug != "" {
		a.Config.Flygo.Dev.Debug = boolEnv(debug)
	}
}

//set server host
func (a *App) setServerHost() {
	host := stringEnv(flygoServerHost)
	if host != "" {
		a.Config.Flygo.Server.Host = host
	}
}

//set server port
func (a *App) setServerPort() {
	port, err := intEnv(flygoServerPort)
	if err == nil {
		a.Config.Flygo.Server.Port = port
	}
}

//set server webRoot
func (a *App) setServerWebRoot() {
	webRoot := stringEnv(flygoServerWebRoot)
	if webRoot != "" {
		a.Config.Flygo.Server.WebRoot = webRoot
	}
}

//set server contextPath
func (a *App) setServerContextPath() {
	contextPath := stringEnv(flygoServerContextPath)
	if contextPath != "" {
		a.Config.Flygo.Server.ContextPath = contextPath
	}
}

//set banner enable
func (a *App) setBannerEnable() {
	bannerEnable := stringEnv(flygoBannerEnable)
	if bannerEnable != "" {
		a.Config.Flygo.Banner.Enable = boolEnv(flygoBannerEnable)
	}
}

//set banner type
func (a *App) setBannerType() {
	bannerType := stringEnv(flygoBannerType)
	if bannerType != "" {
		a.Config.Flygo.Banner.Type = bannerType
	}
}

//set banner text
func (a *App) setBannerText() {
	if stringEnv(flygoBannerText) != "" {
		a.Config.Flygo.Banner.Text = stringEnv(flygoBannerText)
	}
}

//set banner file
func (a *App) setBannerFile() {
	bannerFile := stringEnv(flygoBannerFile)
	if bannerFile != "" {
		a.Config.Flygo.Banner.File = bannerFile
	}
}

//set server tls enable
func (a *App) setServerTlsEnable() {
	serverTlsEnable := stringEnv(flygoServerTlsEnable)
	if serverTlsEnable != "" {
		a.Config.Flygo.Server.Tls.Enable = boolEnv(flygoServerTlsEnable)
	}
}

//set server tls cert file
func (a *App) setServerTlsCertFile() {
	serverTlsCertFile := stringEnv(flygoServerTlsCertFile)
	if serverTlsCertFile != "" {
		a.Config.Flygo.Server.Tls.CertFile = serverTlsCertFile
	}
}

//set server tls key file
func (a *App) setServerTlsKeyFile() {
	serverTlsKeyFile := stringEnv(flygoServerTlsKeyFile)
	if serverTlsKeyFile != "" {
		a.Config.Flygo.Server.Tls.KeyFile = serverTlsKeyFile
	}
}

//set static enable
func (a *App) setStaticEnable() {
	staticEnable := stringEnv(flygoStaticEnable)
	if staticEnable != "" {
		a.Config.Flygo.Static.Enable = boolEnv(flygoStaticEnable)
	}
}

//set static cache
func (a *App) setStaticCache() {
	staticCache := stringEnv(flygoStaticCache)
	if staticCache != "" {
		a.Config.Flygo.Static.Enable = boolEnv(flygoStaticCache)
	}
}

//set static favicon enable
func (a *App) setStaticFaviconEnable() {
	staticFaviconEnable := stringEnv(flygoStaticFaviconEnable)
	if staticFaviconEnable != "" {
		a.Config.Flygo.Static.Favicon.Enable = boolEnv(flygoStaticFaviconEnable)
	}
}

//set static pattern
func (a *App) setStaticPattern() {
	staticPattern := stringEnv(flygoStaticPattern)
	if staticPattern != "" {
		a.Config.Flygo.Static.Pattern = staticPattern
	}
}

//set static prefix
func (a *App) setStaticPrefix() {
	staticPrefix := stringEnv(flygoStaticPrefix)
	if staticPrefix != "" {
		a.Config.Flygo.Static.Prefix = staticPrefix
	}
}

//set view enable
func (a *App) setViewEnable() {
	viewEnable := stringEnv(flygoViewEnable)
	if viewEnable != "" {
		a.Config.Flygo.View.Enable = boolEnv(flygoViewEnable)
	}
}

//set view cache
func (a *App) setViewCache() {
	viewCache := stringEnv(flygoViewCache)
	if viewCache != "" {
		a.Config.Flygo.View.Cache = boolEnv(flygoViewCache)
	}
}

//set view prefix
func (a *App) setViewPrefix() {
	viewPrefix := stringEnv(flygoViewPrefix)
	if viewPrefix != "" {
		a.Config.Flygo.View.Prefix = viewPrefix
	}
}

//set view suffix
func (a *App) setViewSuffix() {
	viewSuffix := stringEnv(flygoViewSuffix)
	if viewSuffix != "" {
		a.Config.Flygo.View.Suffix = viewSuffix
	}
}

//set template enable
func (a *App) setTemplateEnable() {
	templateEnable := stringEnv(flygoTemplateEnable)
	if templateEnable != "" {
		a.Config.Flygo.Template.Enable = boolEnv(flygoTemplateEnable)
	}
}

//set template delims left
func (a *App) setTemplateDelimsLeft() {
	templateDelimsLeft := stringEnv(flygoTemplateDelimsLeft)
	if templateDelimsLeft != "" {
		a.Config.Flygo.Template.Delims.Left = templateDelimsLeft
	}
}

//set template delims right
func (a *App) setTemplateDelimsRight() {
	templateDelimsRight := stringEnv(flygoTemplateDelimsRight)
	if templateDelimsRight != "" {
		a.Config.Flygo.Template.Delims.Left = flygoTemplateDelimsRight
	}
}

//set session enable
func (a *App) setSessionEnable() {
	sessionEnable := stringEnv(flygoSessionEnable)
	if sessionEnable != "" {
		a.SessionConfig.Enable = boolEnv(flygoSessionEnable)
	}
}

//set session timeout
func (a *App) setSessionTimeout() {
	sessionTimeout := stringEnv(flygoSessionTimeout)
	if sessionTimeout != "" {
		sessionTimeout, sessionTimeoutErr := durationEnv(flygoSessionTimeout)
		if sessionTimeoutErr == nil {
			a.SessionConfig.Timeout = sessionTimeout
		}
	}
}

//Parse env
func (a *App) parseEnv() {
	//set config
	a.setConfig()

	if a.ConfigFile != "" {
		//parse yml config
		a.parseYml()
	}

	//set dev debug
	a.setDevDebug()

	//set server host
	a.setServerHost()

	//set server port
	a.setServerPort()

	//set server webRoot
	a.setServerWebRoot()

	//set server contextPath
	a.setServerContextPath()

	//set banner enable
	a.setBannerEnable()

	//set banner type
	a.setBannerType()

	//set banner type
	a.setBannerText()

	//set banner file
	a.setBannerFile()

	//set server tls enable
	a.setServerTlsEnable()

	//set server tls cert file
	a.setServerTlsCertFile()

	//set server tls key file
	a.setServerTlsKeyFile()

	//set static enable
	a.setStaticEnable()

	//set static cache
	a.setStaticCache()

	//set static favicon
	a.setStaticFaviconEnable()

	//set static pattern
	a.setStaticPattern()

	//set static prefix
	a.setStaticPrefix()

	//set view enable
	a.setViewEnable()

	//set view cache
	a.setViewCache()

	//set view prefix
	a.setViewPrefix()

	//set view suffix
	a.setViewSuffix()

	//set template enable
	a.setTemplateEnable()

	//set template delims left
	a.setTemplateDelimsLeft()

	//set template delims right
	a.setTemplateDelimsRight()

	//set session enable
	a.setSessionEnable()

	//set session timeout
	a.setSessionTimeout()
}

func stringEnv(key string) string {
	return os.Getenv(key)
}

func boolEnv(key string) bool {
	be := os.Getenv(key)
	if be == "" {
		return false
	}
	return be == "ON" || be == "on" || be == "1" || be == "true" || be == "TRUE"
}

func intEnv(key string) (int, error) {
	be := os.Getenv(key)
	return strconv.Atoi(be)
}

func durationEnv(key string) (time.Duration, error) {
	be := os.Getenv(key)
	return time.ParseDuration(be)
}
