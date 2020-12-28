package flygo

import (
	"os"
	"strconv"
	"time"
)

const (
	flggoConfig = "FLYGO_CONFIG" //env config file

	flygoDevDebug = "FLYGO_DEV_DEBUG" //env dev debug

	flygoServerHost = "FLYGO_SERVER_HOST" //env server host
	flygoServerPort = "FLYGO_SERVER_PORT" //env server port

	flygoBannerEnable = "FLYGO_BANNER_ENABLE" //env banner enable
	flygoBannerType   = "FLYGO_BANNER_TYPE"   //env banner type
	flygoBannerText   = "FLYGO_BANNER_TEXT"   //env banner text
	flygoBannerFile   = "FLYGO_BANNER_ENABLE" //env banner file

	flygoServerTlsEnable   = "FLYGO_SERVER_TLS_ENABLE"    //env server tls enable
	flygoServerTlsCertFile = "FLYGO_SERVER_TLS_CERT_FILE" //env server tls cert file
	flygoServerTlsKeyFile  = "FLYGO_SERVER_TLS_KEY_FILE"  //env server tls key file
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
		a.Config.Dev.Debug = boolEnv(debug)
	}
}

//set server host
func (a *App) setServerHost() {
	host := stringEnv(flygoServerHost)
	if host != "" {
		a.Config.Server.Host = host
	}
}

//set server port
func (a *App) setServerPort() {
	port, err := intEnv(flygoServerPort)
	if err == nil {
		a.Config.Server.Port = port
	}
}

//set banner enable
func (a *App) setBannerEnable() {
	bannerEnable := stringEnv(flygoBannerEnable)
	if bannerEnable != "" {
		a.Config.Banner.Enable = boolEnv(flygoBannerEnable)
	}
}

//set banner type
func (a *App) setBannerType() {
	bannerType := stringEnv(flygoBannerType)
	if bannerType != "" {
		a.Config.Banner.Type = bannerType
	}
}

//set banner text
func (a *App) setBannerText() {
	if stringEnv(flygoBannerText) != "" {
		a.Config.Banner.Text = stringEnv(flygoBannerText)
	}
}

//set banner file
func (a *App) setBannerFile() {
	bannerFile := stringEnv(flygoBannerFile)
	if bannerFile != "" {
		a.Config.Banner.File = bannerFile
	}
}

//set server tls enable
func (a *App) setServerTlsEnable() {
	serverTlsEnable := stringEnv(flygoServerTlsEnable)
	if serverTlsEnable != "" {
		a.Config.Server.TLS.Enable = boolEnv(flygoServerTlsEnable)
	}
}

//set server tls cert file
func (a *App) setServerTlsCertFile() {
	serverTlsCertFile := stringEnv(flygoServerTlsCertFile)
	if serverTlsCertFile != "" {
		a.Config.Server.TLS.CertFile = serverTlsCertFile
	}
}

//set server tls key file
func (a *App) setServerTlsKeyFile() {
	serverTlsKeyFile := stringEnv(flygoServerTlsKeyFile)
	if serverTlsKeyFile != "" {
		a.Config.Server.TLS.KeyFile = serverTlsKeyFile
	}
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
}

//stringEnv
func stringEnv(key string) string {
	return os.Getenv(key)
}

//boolEnv
func boolEnv(key string) bool {
	be := os.Getenv(key)
	if be == "" {
		return false
	}
	return be == "ON" || be == "on" || be == "1" || be == "true" || be == "TRUE"
}

//intEnv
func intEnv(key string) (int, error) {
	be := os.Getenv(key)
	return strconv.Atoi(be)
}

//durationEnv
func durationEnv(key string) (time.Duration, error) {
	be := os.Getenv(key)
	return time.ParseDuration(be)
}
