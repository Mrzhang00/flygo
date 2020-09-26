package flygo

import (
	"os"
	"strconv"
	"time"
)

const (
	flygoHost               = "FLYGO_HOST"                 //env host for web app
	flygoPort               = "FLYGO_PORT"                 //env port for web app
	flygoContextPath        = "FLYGO_CONTEXT_PATH"         //env contextPath for web app
	flygoSessionTimeout     = "FLYGO_SESSION_TIMEOUT"      //env session timeout for web app
	flygoWebRoot            = "FLYGO_WEB_ROOT"             //env webRoot for web app
	flygoStaticEnable       = "FLYGO_STATIC_ENABLE"        //env static enable for web app
	flygoStaticCache        = "FLYGO_STATIC_CACHE"         //env static cache for web app
	flygoViewEnable         = "FLYGO_VIEW_ENABLE"          //env view enable for web app
	flygoViewCache          = "FLYGO_VIEW_CACHE"           //env view cache for web app
	flygoTemplateEnable     = "FLYGO_TEMPLATE_ENABLE"      //env template for web app
	flygoTemplateDelimLeft  = "FLYGO_TEMPLATE_DELIM_LEFT"  //env template delim left for web app
	flygoTemplateDelimRight = "FLYGO_TEMPLATE_DELIM_RIGHT" //env template delim right for web app
	flygoSessionEnable      = "FLYGO_SESSION_ENABLE"       //env session enable for web app
)

//Parse env
func (a *App) parseEnv() {
	host := os.Getenv(flygoHost)
	if host != "" {
		a.address.host = host
	}
	port, err := strconv.Atoi(os.Getenv(flygoPort))
	if err == nil {
		a.address.port = port
	}
	sessionTimeout := os.Getenv(flygoSessionTimeout)
	if sessionTimeout != "" {
		timeout, err := time.ParseDuration(sessionTimeout)
		if err == nil {
			a.SessionConfig.Timeout = timeout
		}
	}
	webRoot := os.Getenv(flygoWebRoot)
	if webRoot != "" {
		a.SetWebRoot(webRoot)
	}
	contextPath := os.Getenv(flygoContextPath)
	if contextPath != "" {
		a.SetContextPath(contextPath)
	}
	staticCache := os.Getenv(flygoStaticCache)
	if staticCache != "" && (staticCache == "true" || staticCache == "false") {
		a.SetStaticCache(staticCache == "true")
	}
	staticEnable := os.Getenv(flygoStaticEnable)
	if staticEnable != "" && (staticEnable == "true" || staticEnable == "false") {
		a.SetStaticEnable(staticEnable == "true")
	}
	viewEnable := os.Getenv(flygoViewEnable)
	if viewEnable != "" && (viewEnable == "true" || viewEnable == "false") {
		a.SetViewEnable(viewEnable == "true")
	}
	viewCache := os.Getenv(flygoViewCache)
	if viewCache != "" && (viewCache == "true" || viewCache == "false") {
		a.SetViewCache(viewCache == "true")
	}
	templateEnable := os.Getenv(flygoTemplateEnable)
	if templateEnable != "" && (templateEnable == "true" || templateEnable == "false") {
		a.SetTemplateEnable(templateEnable == "true")
	}
	templateDelimLeft := os.Getenv(flygoTemplateDelimLeft)
	if templateDelimLeft != "" {
		a.SetTemplateDelimLeft(templateDelimLeft)
	}
	templateDelimRight := os.Getenv(flygoTemplateDelimRight)
	if templateDelimRight != "" {
		a.SetTemplateDelimRight(templateDelimRight)
	}
	sessionEnable := os.Getenv(flygoSessionEnable)
	if sessionEnable != "" && (sessionEnable == "true" || sessionEnable == "false") {
		a.SetSessionEnable(sessionEnable == "true")
	}
}
