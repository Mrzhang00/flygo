package flygo

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Define Config struct
type Config struct {
	*YmlConfig `yaml:"flygo"`
}

//Unmarshal
func (yc *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, yc)
}

//Define YmlConfig struct
type YmlConfig struct {
	*YmlConfigDev    `yaml:"dev"`
	*YmlConfigBanner `yaml:"banner"`
	*YmlConfigServer `yaml:"server"`
}

//Define YmlConfigDev struct
type YmlConfigDev struct {
	Debug bool `yaml:"debug"`
}

//Define YmlConfigBanner struct
type YmlConfigBanner struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
	File   string `yaml:"file"`
	Text   string `yaml:"text"`
}

//Define YmlConfigServer struct
type YmlConfigServer struct {
	Host string              `yaml:"host"`
	Port int                 `yaml:"port"`
	TLS  *YmlConfigServerTls `yaml:"tls"`
}

//Define YmlConfigServerTls struct
type YmlConfigServerTls struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

//parseYml
func (a *App) parseYml() {
	file := a.ConfigFile
	if file != "" {
		bytes, err := ioutil.ReadFile(file)

		if err != nil {
			a.DebugTrace(func() {
				a.Logger.Warn("[parseYml]%v", err)
			})
		}

		if err == nil {
			err = a.Config.Unmarshal(bytes)
			if err != nil {

				a.DebugTrace(func() {
					if err != nil {
						a.Logger.Warn("[parseYml]%v", err)
					}
				})

			}
		}
	}
}

func (a *App) printConfig() {
	bytes, err := yaml.Marshal(a.Config)

	a.DebugTrace(func() {
		if err != nil {
			a.Logger.Warn("[printConfig]%v", err)
		}
	})

	a.DebugTrace(func() {
		if err != nil {
			a.Logger.Info("[printConfig]%v", string(bytes))
		}
	})
}

func defaultConfig() *Config {
	return &Config{&YmlConfig{
		&YmlConfigDev{
			Debug: false,
		}, &YmlConfigBanner{
			Enable: true,
			Type:   "default",
			File:   "banner.txt",
			Text:   "FLYGO",
		}, &YmlConfigServer{
			Host: "",
			Port: 80,
			TLS: &YmlConfigServerTls{
				Enable:   false,
				CertFile: "",
				KeyFile:  "",
			},
		},
	}}
}
