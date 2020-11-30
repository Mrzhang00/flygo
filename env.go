package flygo

import (
	"os"
	"strconv"
	"time"
)

const (
	flggoConfig = "FLYGO_CONFIG" //env config file for web app
	flygoHost   = "FLYGO_HOST"   //env host for web app
	flygoPort   = "FLYGO_PORT"   //env port for web app

	flygoContextPath = "FLYGO_CONTEXT_PATH" //env contextPath for web app
	flygoWebRoot     = "FLYGO_WEB_ROOT"     //env webRoot for web app

	flygoBannerEnable = "FLYGO_BANNER_ENABLE" //env banner enable for web app
	flygoBannerType   = "FLYGO_BANNER_TYPE"   //env banner type for web app
	flygoBannerText   = "FLYGO_BANNER_TEXT"   //env banner text for web app
	flygoBannerFile   = "FLYGO_BANNER_ENABLE" //env banner file for web app

	flygoTlsEnable   = "FLYGO_TLS_ENABLE"    //env tls enable for web app
	flygoTlsCertFile = "FLYGO_TLS_CERT_FILE" //env tls cert file for web app
	flygoTlsKeyFile  = "FLYGO_TLS_KEY_FILE"  //env tls key file for web app

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

//Parse env
func (a *App) parseEnv() {
	a.ConfigFile = stringEnv(flggoConfig)

	if a.ConfigFile != "" {
		a.parseYml()
	}

	if stringEnv(flygoHost) != "" {
		a.Config.Flygo.Server.Host = stringEnv(flygoHost)
	}

	port, err := intEnv(flygoPort)
	if err == nil {
		a.Config.Flygo.Server.Port = port
	}

	if stringEnv(flygoWebRoot) != "" {
		a.Config.Flygo.Server.WebRoot = stringEnv(flygoWebRoot)
	}

	if stringEnv(flygoContextPath) != "" {
		a.Config.Flygo.Server.ContextPath = stringEnv(flygoContextPath)
	}

	a.Config.Flygo.Banner.Enable = boolEnv(flygoBannerEnable)
	if stringEnv(flygoBannerType) != "" {
		a.Config.Flygo.Banner.Type = stringEnv(flygoBannerType)
	}
	if stringEnv(flygoBannerText) != "" {
		a.Config.Flygo.Banner.Text = stringEnv(flygoBannerText)
	}
	if stringEnv(flygoBannerFile) != "" {
		a.Config.Flygo.Banner.File = stringEnv(flygoBannerFile)
	}

	a.Config.Flygo.Server.Tls.Enable = boolEnv(flygoTlsEnable)
	if stringEnv(flygoTlsCertFile) != "" {
		a.Config.Flygo.Server.Tls.CertFile = stringEnv(flygoTlsCertFile)
	}
	if stringEnv(flygoTlsKeyFile) != "" {
		a.Config.Flygo.Server.Tls.KeyFile = stringEnv(flygoTlsKeyFile)
	}

	a.Config.Flygo.Static.Enable = boolEnv(flygoStaticEnable)
	a.Config.Flygo.Static.Cache = boolEnv(flygoStaticCache)
	a.Config.Flygo.Static.Favicon.Enable = boolEnv(flygoStaticFaviconEnable)
	if stringEnv(flygoStaticPattern) != "" {
		a.Config.Flygo.Static.Pattern = stringEnv(flygoStaticPattern)
	}
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
