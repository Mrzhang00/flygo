package flygo

import (
	"os"
	"strconv"
	"time"
)

const (
	envServerMaxHeaderSize     = "SERVER_MAX_HEADER_SIZE"
	envServerReadTimeout       = "SERVER_READ_TIMEOUT"
	envServerReadHeaderTimeout = "SERVER_READ_HEADER_TIMEOUT"
	envServerWriteTimeout      = "SERVER_WRITE_TIMEOUT"
	envServerIdleTimeout       = "SERVER_IDLE_TIMEOUT"
	envConfigFile              = "CONFIG_FILE"
	envServerHost              = "SERVER_HOST"
	envServerPort              = "SERVER_PORT"
	envServerTLSEnable         = "SERVER_TLS_ENABLE"
	envServerTLSCertFile       = "SERVER_TLS_CERT_FILE"
	envServerTLSKeyFile        = "SERVER_TLS_KEY_FILE"
	envBannerEnable            = "BANNER_ENABLE"
	envBannerType              = "BANNER_TYPE"
	envBannerText              = "BANNER_TEXT"
	envBannerFile              = "BANNER_FILE"
	envTemplateEnable          = "TEMPLATE_ENABLE"
	envTemplateCache           = "TEMPLATE_CACHE"
	envTemplateRoot            = "TEMPLATE_ROOT"
	envTemplateSuffix          = "TEMPLATE_SUFFIX"
)

func (a *App) setServerMaxHeaderSize() {
	maxHeaderSize, err := intEnv(envServerMaxHeaderSize)
	if err == nil {
		a.Config.Server.MaxHeaderSize = maxHeaderSize
		a.Logger.Debugf("env: %v = %v", envServerMaxHeaderSize, maxHeaderSize)
	}
}

func (a *App) setServerReadTimeout() {
	readTimeout, err := durationEnv(envServerReadTimeout)
	if err == nil {
		a.Config.Server.ReadTimeout = readTimeout
		a.Logger.Debugf("env: %v = %v", envServerReadTimeout, readTimeout)
	}
}

func (a *App) setServerReadHeaderTimeout() {
	readHeaderTimeout, err := durationEnv(envServerReadHeaderTimeout)
	if err == nil {
		a.Config.Server.ReadHeaderTimeout = readHeaderTimeout
		a.Logger.Debugf("env: %v = %v", envServerReadHeaderTimeout, readHeaderTimeout)
	}
}

func (a *App) setServerWriteTimeout() {
	writeTimeout, err := durationEnv(envServerWriteTimeout)
	if err == nil {
		a.Config.Server.WriteTimeout = writeTimeout
		a.Logger.Debugf("env: %v = %v", envServerWriteTimeout, writeTimeout)
	}
}

func (a *App) setServerIdleTimeout() {
	idleTimeout, err := durationEnv(envServerIdleTimeout)
	if err == nil {
		a.Config.Server.IdleTimeout = idleTimeout
		a.Logger.Debugf("env: %v = %v", envServerIdleTimeout, idleTimeout)
	}
}

func (a *App) setConfigFile() {
	configFile := stringEnv(envConfigFile)
	if configFile != "" {
		a.ConfigFile = configFile
		a.Logger.Debugf("env: %v = %v", envConfigFile, configFile)
	}
}

func (a *App) setServerHost() {
	host := stringEnv(envServerHost)
	if host != "" {
		a.Config.Server.Host = host
		a.Logger.Debugf("env: %v = %v", envServerHost, host)
	}
}

func (a *App) setServerPort() {
	port, err := intEnv(envServerPort)
	if err == nil {
		a.Config.Server.Port = port
		a.Logger.Debugf("env: %v = %v", envServerPort, a.Config.Server.Port)
	}
}

func (a *App) setBannerEnable() {
	if stringEnv(envBannerEnable) != "" {
		a.Config.Banner.Enable = boolEnv(envBannerEnable)
		a.Logger.Debugf("env: %v = %v", envBannerEnable, a.Config.Banner.Enable)
	}
}

func (a *App) setBannerType() {
	bannerType := stringEnv(envBannerType)
	if bannerType != "" {
		a.Config.Banner.Type = bannerType
		a.Logger.Debugf("env: %v = %v", envBannerType, bannerType)
	}
}

func (a *App) setBannerText() {
	bannerText := stringEnv(envBannerText)
	if bannerText != "" {
		a.Config.Banner.Text = bannerText
		a.Logger.Debugf("env: %v = %v", envBannerText, bannerText)
	}
}

func (a *App) setBannerFile() {
	bannerFile := stringEnv(envBannerFile)
	if bannerFile != "" {
		a.Config.Banner.File = bannerFile
		a.Logger.Debugf("env: %v = %v", envBannerFile, bannerFile)
	}
}

func (a *App) setServerTLSEnable() {
	if stringEnv(envServerTLSEnable) != "" {
		a.Config.Server.TLS.Enable = boolEnv(envServerTLSEnable)
		a.Logger.Debugf("env: %v = %v", envServerTLSEnable, a.Config.Server.TLS.Enable)
	}
}

func (a *App) setServerTLSCertFile() {
	serverTLSCertFile := stringEnv(envServerTLSCertFile)
	if serverTLSCertFile != "" {
		a.Config.Server.TLS.CertFile = serverTLSCertFile
		a.Logger.Debugf("env: %v = %v", envServerTLSCertFile, envServerTLSCertFile)
	}
}

func (a *App) setServerTLSKeyFile() {
	serverTLSKeyFile := stringEnv(envServerTLSKeyFile)
	if serverTLSKeyFile != "" {
		a.Config.Server.TLS.KeyFile = serverTLSKeyFile
		a.Logger.Debugf("env: %v = %v", envServerTLSKeyFile, envServerTLSKeyFile)
	}
}

func (a *App) setTemplateEnable() {
	if stringEnv(envTemplateEnable) != "" {
		a.Config.Template.Enable = boolEnv(envTemplateEnable)
		a.Logger.Debugf("env: %v = %v", envTemplateEnable, a.Config.Template.Enable)
	}
}

func (a *App) setTemplateCache() {
	if stringEnv(envTemplateCache) != "" {
		a.Config.Template.Cache = boolEnv(envTemplateCache)
		a.Logger.Debugf("env: %v = %v", envTemplateCache, a.Config.Template.Cache)
	}
}

func (a *App) setTemplateRoot() {
	templateRoot := stringEnv(envTemplateRoot)
	if templateRoot != "" {
		a.Config.Template.Root = templateRoot
		a.Logger.Debugf("env: %v = %v", envTemplateRoot, templateRoot)
	}
}

func (a *App) setTemplateSuffix() {
	templateSuffix := stringEnv(envTemplateSuffix)
	if templateSuffix != "" {
		a.Config.Template.Suffix = templateSuffix
		a.Logger.Debugf("env: %v = %v", envTemplateSuffix, templateSuffix)
	}
}

func (a *App) parseEnv() {
	a.setConfigFile()
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
