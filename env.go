package flygo

import (
	"github.com/billcoding/calls"
	"os"
	"strconv"
)

const (
	flygoConfig = "FLYGO_CONFIG" //env config file

	flygoDevDebug = "FLYGO_DEV_DEBUG" //env dev debug

	flygoServerHost = "FLYGO_SERVER_HOST" //env server host
	flygoServerPort = "FLYGO_SERVER_PORT" //env server port

	flygoBannerEnable = "FLYGO_BANNER_ENABLE" //env banner enable
	flygoBannerType   = "FLYGO_BANNER_TYPE"   //env banner type
	flygoBannerText   = "FLYGO_BANNER_TEXT"   //env banner text
	flygoBannerFile   = "FLYGO_BANNER_FILE"   //env banner file

	flygoServerTlsEnable   = "FLYGO_SERVER_TLS_ENABLE"    //env server tls enable
	flygoServerTlsCertFile = "FLYGO_SERVER_TLS_CERT_FILE" //env server tls cert file
	flygoServerTlsKeyFile  = "FLYGO_SERVER_TLS_KEY_FILE"  //env server tls key file

	flygoTemplateEnable = "FLYGO_TEMPLATE_ENABLE" //env template enable
	flygoTemplateCache  = "FLYGO_TEMPLATE_CACHE"  //env template cache
	flygoTemplateRoot   = "FLYGO_TEMPLATE_ROOT"   //env template root
	flygoTemplateSuffix = "FLYGO_TEMPLATE_SUFFIX" //env template suffix
)

//set config
func (a *App) setConfig() {
	config := stringEnv(flygoConfig)
	calls.NEmpty(config, func() {
		a.ConfigFile = config
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoConfig, config)
		})
	})
}

//set dev debug
func (a *App) setDevDebug() {
	calls.NEmpty(stringEnv(flygoDevDebug), func() {
		a.Config.Dev.Debug = boolEnv(flygoDevDebug)
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoDevDebug, a.Config.Dev.Debug)
		})
	})
}

//set server host
func (a *App) setServerHost() {
	host := stringEnv(flygoServerHost)
	calls.NEmpty(host, func() {
		a.Config.Server.Host = host
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoServerHost, host)
		})
	})
}

//set server port
func (a *App) setServerPort() {
	port, err := intEnv(flygoServerPort)
	if err == nil {
		a.Config.Server.Port = port
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoServerPort, a.Config.Server.Port)
		})
	}
}

//set banner enable
func (a *App) setBannerEnable() {
	calls.NEmpty(stringEnv(flygoBannerEnable), func() {
		a.Config.Banner.Enable = boolEnv(flygoBannerEnable)
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoBannerEnable, a.Config.Banner.Enable)
		})
	})
}

//set banner type
func (a *App) setBannerType() {
	bannerType := stringEnv(flygoBannerType)
	if bannerType != "" {
		a.Config.Banner.Type = bannerType
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoBannerType, bannerType)
		})
	}
}

//set banner text
func (a *App) setBannerText() {
	bannerText := stringEnv(flygoBannerText)
	if bannerText != "" {
		a.Config.Banner.Text = bannerText
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoBannerText, bannerText)
		})
	}
}

//set banner file
func (a *App) setBannerFile() {
	bannerFile := stringEnv(flygoBannerFile)
	if bannerFile != "" {
		a.Config.Banner.File = bannerFile
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoBannerFile, bannerFile)
		})
	}
}

//set server tls enable
func (a *App) setServerTlsEnable() {
	calls.NEmpty(stringEnv(flygoServerTlsEnable), func() {
		a.Config.Server.TLS.Enable = boolEnv(flygoServerTlsEnable)
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoServerTlsEnable, a.Config.Server.TLS.Enable)
		})
	})
}

//set server tls cert file
func (a *App) setServerTlsCertFile() {
	serverTlsCertFile := stringEnv(flygoServerTlsCertFile)
	if serverTlsCertFile != "" {
		a.Config.Server.TLS.CertFile = serverTlsCertFile
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoServerTlsCertFile, serverTlsCertFile)
		})
	}
}

//set server tls key file
func (a *App) setServerTlsKeyFile() {
	serverTlsKeyFile := stringEnv(flygoServerTlsKeyFile)
	if serverTlsKeyFile != "" {
		a.Config.Server.TLS.KeyFile = serverTlsKeyFile
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoServerTlsKeyFile, serverTlsKeyFile)
		})
	}
}

//set template enable
func (a *App) setTemplateEnable() {
	calls.NEmpty(stringEnv(flygoTemplateEnable), func() {
		a.Config.Template.Enable = boolEnv(flygoTemplateEnable)
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoTemplateEnable, a.Config.Template.Enable)
		})
	})
}

//set template cache
func (a *App) setTemplateCache() {
	calls.NEmpty(stringEnv(flygoTemplateCache), func() {
		a.Config.Template.Cache = boolEnv(flygoTemplateCache)
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoTemplateCache, a.Config.Template.Cache)
		})
	})
}

//set template root
func (a *App) setTemplateRoot() {
	templateRoot := stringEnv(flygoTemplateRoot)
	calls.NEmpty(templateRoot, func() {
		a.Config.Template.Root = templateRoot
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoTemplateRoot, templateRoot)
		})
	})
}

//set template suffix
func (a *App) setTemplateSuffix() {
	templateSuffix := stringEnv(flygoTemplateSuffix)
	calls.NEmpty(templateSuffix, func() {
		a.Config.Template.Suffix = templateSuffix
		a.DebugTrace(func() {
			a.Logger.Info("[Env]Set [%v] = [%v]", flygoTemplateSuffix, templateSuffix)
		})
	})
}

//parseEnv
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

	//set server tls enable
	a.setServerTlsEnable()

	//set server tls cert file
	a.setServerTlsCertFile()

	//set server tls key file
	a.setServerTlsKeyFile()

	//set banner enable
	a.setBannerEnable()

	//set banner type
	a.setBannerType()

	//set banner type
	a.setBannerText()

	//set banner file
	a.setBannerFile()

	//set template enable
	a.setTemplateEnable()

	//set template cache
	a.setTemplateCache()

	//set template root
	a.setTemplateRoot()

	//set template suffix
	a.setTemplateSuffix()
}

//stringEnv
func stringEnv(key string) string {
	return os.Getenv(key)
}

//boolEnv
func boolEnv(key string) bool {
	be := os.Getenv(key)
	//	case "1", "t", "T", "true", "TRUE", "True":
	//	case "0", "f", "F", "false", "FALSE", "False":
	b, err := strconv.ParseBool(be)
	return err == nil && b
}

//intEnv
func intEnv(key string) (int, error) {
	be := os.Getenv(key)
	return strconv.Atoi(be)
}
