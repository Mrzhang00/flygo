package flygo

import (
	"os"
	"strconv"
	"time"
)

const (
	serverMaxHeaderSize     = "SERVER_MAX_HEADER_SIZE"
	serverReadTimeout       = "SERVER_READ_TIMEOUT"
	serverReadHeaderTimeout = "SERVER_READ_HEADER_TIMEOUT"
	serverWriteTimeout      = "SERVER_WRITE_TIMEOUT"
	serverIdleTimeout       = "SERVER_IDLE_TIMEOUT"
	flygoConfig             = "FLYGO_CONFIG"
	flygoDevDebug           = "FLYGO_DEV_DEBUG"
	flygoServerHost         = "FLYGO_SERVER_HOST"
	flygoServerPort         = "FLYGO_SERVER_PORT"
	flygoServerTLSEnable    = "FLYGO_SERVER_TLS_ENABLE"
	flygoServerTLSCertFile  = "FLYGO_SERVER_TLS_CERT_FILE"
	flygoServerTLSKeyFile   = "FLYGO_SERVER_TLS_KEY_FILE"
	flygoBannerEnable       = "FLYGO_BANNER_ENABLE"
	flygoBannerType         = "FLYGO_BANNER_TYPE"
	flygoBannerText         = "FLYGO_BANNER_TEXT"
	flygoBannerFile         = "FLYGO_BANNER_FILE"
	flygoTemplateEnable     = "FLYGO_TEMPLATE_ENABLE"
	flygoTemplateCache      = "FLYGO_TEMPLATE_CACHE"
	flygoTemplateRoot       = "FLYGO_TEMPLATE_ROOT"
	flygoTemplateSuffix     = "FLYGO_TEMPLATE_SUFFIX"
)

func (a *App) setServerMaxHeaderSize() {
	maxHeaderSize, err := intEnv(serverMaxHeaderSize)
	if err == nil {
		a.Config.Server.MaxHeaderSize = maxHeaderSize

		a.Logger.Debugf("[Env]Set [%v] = [%v]", serverMaxHeaderSize, maxHeaderSize)

	}
}

func (a *App) setServerReadTimeout() {
	readTimeout, err := durationEnv(serverReadTimeout)
	if err == nil {
		a.Config.Server.Timeout.Read = readTimeout
		a.Logger.Debugf("[Env]Set [%v] = [%v]", serverReadTimeout, readTimeout)
	}
}

func (a *App) setServerReadHeaderTimeout() {
	readHeaderTimeout, err := durationEnv(serverReadHeaderTimeout)
	if err == nil {
		a.Config.Server.Timeout.ReadHeader = readHeaderTimeout
		a.Logger.Debugf("[Env]Set [%v] = [%v]", serverReadHeaderTimeout, readHeaderTimeout)
	}
}

func (a *App) setServerWriteTimeout() {
	writeTimeout, err := durationEnv(serverWriteTimeout)
	if err == nil {
		a.Config.Server.Timeout.Write = writeTimeout
		a.Logger.Debugf("[Env]Set [%v] = [%v]", serverWriteTimeout, writeTimeout)
	}
}

func (a *App) setServerIdleTimeout() {
	idleTimeout, err := durationEnv(serverIdleTimeout)
	if err == nil {
		a.Config.Server.Timeout.Idle = idleTimeout

		a.Logger.Debugf("[Env]Set [%v] = [%v]", serverIdleTimeout, idleTimeout)

	}
}

func (a *App) setConfig() {
	config := stringEnv(flygoConfig)
	if config != "" {
		a.ConfigFile = config
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoConfig, config)

	}
}

func (a *App) setServerHost() {
	host := stringEnv(flygoServerHost)
	if host != "" {
		a.Config.Flygo.Server.Host = host
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoServerHost, host)
	}
}

func (a *App) setServerPort() {
	port, err := intEnv(flygoServerPort)
	if err == nil {
		a.Config.Flygo.Server.Port = port
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoServerPort, a.Config.Flygo.Server.Port)
	}
}

func (a *App) setBannerEnable() {
	if stringEnv(flygoBannerEnable) != "" {
		a.Config.Flygo.Banner.Enable = boolEnv(flygoBannerEnable)
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoBannerEnable, a.Config.Flygo.Banner.Enable)
	}
}

func (a *App) setBannerType() {
	bannerType := stringEnv(flygoBannerType)
	if bannerType != "" {
		a.Config.Flygo.Banner.Type = bannerType
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoBannerType, bannerType)
	}
}

func (a *App) setBannerText() {
	bannerText := stringEnv(flygoBannerText)
	if bannerText != "" {
		a.Config.Flygo.Banner.Text = bannerText
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoBannerText, bannerText)
	}
}

func (a *App) setBannerFile() {
	bannerFile := stringEnv(flygoBannerFile)
	if bannerFile != "" {
		a.Config.Flygo.Banner.File = bannerFile
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoBannerFile, bannerFile)
	}
}

func (a *App) setServerTLSEnable() {
	if stringEnv(flygoServerTLSEnable) != "" {
		a.Config.Flygo.Server.TLS.Enable = boolEnv(flygoServerTLSEnable)
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoServerTLSEnable, a.Config.Flygo.Server.TLS.Enable)
	}
}

func (a *App) setServerTLSCertFile() {
	serverTLSCertFile := stringEnv(flygoServerTLSCertFile)
	if serverTLSCertFile != "" {
		a.Config.Flygo.Server.TLS.CertFile = serverTLSCertFile
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoServerTLSCertFile, serverTLSCertFile)
	}
}

func (a *App) setServerTLSKeyFile() {
	serverTLSKeyFile := stringEnv(flygoServerTLSKeyFile)
	if serverTLSKeyFile != "" {
		a.Config.Flygo.Server.TLS.KeyFile = serverTLSKeyFile
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoServerTLSKeyFile, serverTLSKeyFile)
	}
}

func (a *App) setTemplateEnable() {
	if stringEnv(flygoTemplateEnable) != "" {
		a.Config.Flygo.Template.Enable = boolEnv(flygoTemplateEnable)
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoTemplateEnable, a.Config.Flygo.Template.Enable)
	}
}

func (a *App) setTemplateCache() {
	if stringEnv(flygoTemplateCache) != "" {
		a.Config.Flygo.Template.Cache = boolEnv(flygoTemplateCache)
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoTemplateCache, a.Config.Flygo.Template.Cache)
	}
}

func (a *App) setTemplateRoot() {
	templateRoot := stringEnv(flygoTemplateRoot)
	if templateRoot != "" {
		a.Config.Flygo.Template.Root = templateRoot
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoTemplateRoot, templateRoot)
	}
}

func (a *App) setTemplateSuffix() {
	templateSuffix := stringEnv(flygoTemplateSuffix)
	if templateSuffix != "" {
		a.Config.Flygo.Template.Suffix = templateSuffix
		a.Logger.Debugf("[Env]Set [%v] = [%v]", flygoTemplateSuffix, templateSuffix)
	}
}

func (a *App) parseEnv() {
	a.setConfig()
	if a.ConfigFile != "" {
		a.parseYml()
	}
	a.setServerMaxHeaderSize()
	a.setServerReadTimeout()
	a.setServerReadHeaderTimeout()
	a.setServerWriteTimeout()
	a.setServerIdleTimeout()
	a.setServerHost()
	a.setServerPort()
	a.setServerTLSEnable()
	a.setServerTLSCertFile()
	a.setServerTLSKeyFile()
	a.setBannerEnable()
	a.setBannerType()
	a.setBannerText()
	a.setBannerFile()
	a.setTemplateEnable()
	a.setTemplateCache()
	a.setTemplateRoot()
	a.setTemplateSuffix()
}

func stringEnv(key string) string {
	return os.Getenv(key)
}

func boolEnv(key string) bool {
	be := os.Getenv(key)
	b, err := strconv.ParseBool(be)
	return err == nil && b
}

func intEnv(key string) (int, error) {
	ie := os.Getenv(key)
	return strconv.Atoi(ie)
}

func durationEnv(key string) (time.Duration, error) {
	de := os.Getenv(key)
	return time.ParseDuration(de)
}
