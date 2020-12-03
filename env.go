package flygo

import (
	"os"
	"strconv"
	"time"
)

const (
	flggoConfig = "FLYGO_CONFIG" //env config file for web app

	flygoDevDebug = "FLYGO_DEV_DEBUG" //env dev debug for web app

	flygoServerHost        = "FLYGO_SERVER_HOST"         //env server host for web app
	flygoServerPort        = "FLYGO_SERVER_PORT"         //env server port for web app
	flygoServerContextPath = "FLYGO_SERVER_CONTEXT_PATH" //env server contextPath for web app
	flygoServerWebRoot     = "FLYGO_SERVER_WEB_ROOT"     //env server webRoot for web app

	flygoBannerEnable = "FLYGO_BANNER_ENABLE" //env banner enable for web app
	flygoBannerType   = "FLYGO_BANNER_TYPE"   //env banner type for web app
	flygoBannerText   = "FLYGO_BANNER_TEXT"   //env banner text for web app
	flygoBannerFile   = "FLYGO_BANNER_ENABLE" //env banner file for web app

	flygoServerTlsEnable   = "FLYGO_SERVER_TLS_ENABLE"    //env server tls enable for web app
	flygoServerTlsCertFile = "FLYGO_SERVER_TLS_CERT_FILE" //env server tls cert file for web app
	flygoServerTlsKeyFile  = "FLYGO_SERVER_TLS_KEY_FILE"  //env server tls key file for web app

	flygoStaticEnable  = "FLYGO_STATIC_ENABLE"  //env static enable for web app
	flygoStaticPattern = "FLYGO_STATIC_PATTERN" //env static pattern for web app
	flygoStaticPrefix  = "FLYGO_STATIC_PREFIX"  //env static prefix for web app
	flygoStaticCache   = "FLYGO_STATIC_CACHE"   //env static cache for web app

	flygoStaticFaviconEnable = "FLYGO_STATIC_FAVICON_ENABLE" //env static favicon enable for web app

	flygoViewEnable = "FLYGO_VIEW_ENABLE" //env view enable for web app
	flygoViewPrefix = "FLYGO_VIEW_PREFIX" //env view prefix for web app
	flygoViewSuffix = "FLYGO_VIEW_SUFFIX" //env view suffix for web app
	flygoViewCache  = "FLYGO_VIEW_CACHE"  //env view cache for web app

	flygoTemplateEnable     = "FLYGO_TEMPLATE_ENABLE"      //env template for web app
	flygoTemplateDelimLeft  = "FLYGO_TEMPLATE_DELIM_LEFT"  //env template delim left for web app
	flygoTemplateDelimRight = "FLYGO_TEMPLATE_DELIM_RIGHT" //env template delim right for web app

	flygoSessionEnable  = "FLYGO_SESSION_ENABLE"  //env session enable for web app
	flygoSessionTimeout = "FLYGO_SESSION_TIMEOUT" //env session timeout for web app
)

//set config
func (a *App) _setConfig() {
	config := stringEnv(flggoConfig)
	if config != "" {
		a.ConfigFile = config
	}
}

//set dev debug
func (a *App) _setDevDebug() {
	debug := stringEnv(flygoDevDebug)
	if debug != "" {
		a.Config.Flygo.Dev.Debug = boolEnv(debug)
	}
}

//set server host
func (a *App) _setServerHost() {
	host := stringEnv(flygoServerHost)
	if host != "" {
		a.Config.Flygo.Server.Host = host
	}
}

//set server port
func (a *App) _setServerPort() {
	port, err := intEnv(flygoServerPort)
	if err == nil {
		a.Config.Flygo.Server.Port = port
	}
}

//set server webRoot
func (a *App) _setServerWebRoot() {
	webRoot := stringEnv(flygoServerWebRoot)
	if webRoot != "" {
		a.Config.Flygo.Server.WebRoot = webRoot
	}
}

//set server contextPath
func (a *App) _setServerContextPath() {
	contextPath := stringEnv(flygoServerContextPath)
	if contextPath != "" {
		a.Config.Flygo.Server.ContextPath = contextPath
	}
}

//set banner enable
func (a *App) _setBannerEnable() {
	bannerEnable := stringEnv(flygoBannerEnable)
	if bannerEnable != "" {
		a.Config.Flygo.Banner.Enable = boolEnv(flygoBannerEnable)
	}
}

//set banner type
func (a *App) _setBannerType() {
	bannerType := stringEnv(flygoBannerType)
	if bannerType != "" {
		a.Config.Flygo.Banner.Type = bannerType
	}
}

//set banner text
func (a *App) _setBannerText() {
	if stringEnv(flygoBannerText) != "" {
		a.Config.Flygo.Banner.Text = stringEnv(flygoBannerText)
	}
}

//set banner file
func (a *App) _setBannerFile() {
	bannerFile := stringEnv(flygoBannerFile)
	if bannerFile != "" {
		a.Config.Flygo.Banner.File = bannerFile
	}
}

//set server tls enable
func (a *App) _setServerTlsEnable() {
	serverTlsEnable := stringEnv(flygoServerTlsEnable)
	if serverTlsEnable != "" {
		a.Config.Flygo.Server.Tls.Enable = boolEnv(flygoServerTlsEnable)
	}
}

//set server tls cert file
func (a *App) _setServerTlsCertFile() {
	serverTlsCertFile := stringEnv(flygoServerTlsCertFile)
	if serverTlsCertFile != "" {
		a.Config.Flygo.Server.Tls.CertFile = serverTlsCertFile
	}
}

//set server tls key file
func (a *App) _setServerTlsKeyFile() {
	serverTlsKeyFile := stringEnv(flygoServerTlsKeyFile)
	if serverTlsKeyFile != "" {
		a.Config.Flygo.Server.Tls.KeyFile = serverTlsKeyFile
	}
}

//set static enable
func (a *App) _setStaticEnable() {
	staticEnable := stringEnv(flygoStaticEnable)
	if staticEnable != "" {
		a.Config.Flygo.Static.Enable = boolEnv(flygoStaticEnable)
	}
}

//set static cache
func (a *App) _setStaticCache() {
	staticCache := stringEnv(flygoStaticCache)
	if staticCache != "" {
		a.Config.Flygo.Static.Enable = boolEnv(flygoStaticCache)
	}
}

//set static favicon enable
func (a *App) _setStaticFaviconEnable() {
	staticFaviconEnable := stringEnv(flygoStaticFaviconEnable)
	if staticFaviconEnable != "" {
		a.Config.Flygo.Static.Favicon.Enable = boolEnv(flygoStaticFaviconEnable)
	}
}

//set static pattern
func (a *App) _setStaticPattern() {
	staticPattern := stringEnv(flygoStaticPattern)
	if staticPattern != "" {
		a.Config.Flygo.Static.Pattern = staticPattern
	}
}

//set static prefix
func (a *App) _setStaticPrefix() {
	staticPrefix := stringEnv(flygoStaticPrefix)
	if staticPrefix != "" {
		a.Config.Flygo.Static.Prefix = staticPrefix
	}
}

//Parse env
func (a *App) parseEnv() {
	//set config
	a._setConfig()

	if a.ConfigFile != "" {
		//parse yml config
		a.parseYml()
	}

	//set dev debug
	a._setDevDebug()

	//set server host
	a._setServerHost()

	//set server port
	a._setServerPort()

	//set server webRoot
	a._setServerWebRoot()

	//set server contextPath
	a._setServerContextPath()

	//set banner enable
	a._setBannerEnable()

	//set banner type
	a._setBannerType()

	//set banner type
	a._setBannerText()

	//set banner file
	a._setBannerFile()

	//set server tls enable
	a._setServerTlsEnable()

	//set server tls cert file
	a._setServerTlsCertFile()

	//set server tls key file
	a._setServerTlsKeyFile()

	//set static enable
	a._setStaticEnable()

	//set static cache
	a._setStaticCache()

	//set static favicon
	a._setStaticFaviconEnable()

	//set static pattern
	a._setStaticPattern()

	//set static prefix
	a._setStaticPrefix()

	if stringEnv(flygoStaticPrefix) != "" {
		a.Config.Flygo.Static.Prefix = stringEnv(flygoStaticPrefix)
	}

	a.Config.Flygo.View.Enable = boolEnv(flygoViewEnable)
	a.Config.Flygo.View.Cache = boolEnv(flygoViewCache)
	if stringEnv(flygoViewPrefix) != "" {
		a.Config.Flygo.View.Prefix = stringEnv(flygoStaticPrefix)
	}
	if stringEnv(flygoViewSuffix) != "" {
		a.Config.Flygo.View.Suffix = stringEnv(flygoViewSuffix)
	}

	a.Config.Flygo.Template.Enable = boolEnv(flygoTemplateEnable)
	if stringEnv(flygoTemplateDelimLeft) != "" {
		a.Config.Flygo.Template.Delims.Left = stringEnv(flygoTemplateDelimLeft)
	}
	if stringEnv(flygoTemplateDelimRight) != "" {
		a.Config.Flygo.Template.Delims.Right = stringEnv(flygoTemplateDelimRight)
	}

	a.Config.Flygo.Session.Enable = boolEnv(flygoSessionEnable)
	sessionTimeout, sessionTimeoutErr := durationEnv(flygoSessionTimeout)
	if sessionTimeoutErr == nil {
		a.SessionConfig.Timeout = sessionTimeout
	}
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
