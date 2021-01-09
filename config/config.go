package config

import (
	"gopkg.in/yaml.v2"
	"html/template"
	"net/http"
	"time"
)

type Config struct {
	Server *YmlServerConfig `yaml:"server"`
	Flygo  *YmlConfig       `yaml:"flygo"`
}

func (yc *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, yc)
}

type YmlServerConfig struct {
	MaxHeaderSize int                     `yaml:"maxHeaderSize"`
	Timeout       *YmlServerConfigTimeout `yaml:"timeout"`
}

type YmlServerConfigTimeout struct {
	Read       time.Duration `yaml:"read"`
	ReadHeader time.Duration `yaml:"readHeader"`
	Write      time.Duration `yaml:"write"`
	Idle       time.Duration `yaml:"idle"`
}

type YmlConfig struct {
	Dev      *YmlConfigDev      `yaml:"dev"`
	Banner   *YmlConfigBanner   `yaml:"banner"`
	Server   *YmlConfigServer   `yaml:"server"`
	Template *YmlConfigTemplate `yaml:"template"`
}

type YmlConfigDev struct {
	Debug bool `yaml:"debug"`
}

type YmlConfigBanner struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
	File   string `yaml:"file"`
	Text   string `yaml:"text"`
}

type YmlConfigServer struct {
	Host string              `yaml:"host"`
	Port int                 `yaml:"port"`
	TLS  *YmlConfigServerTls `yaml:"tls"`
}

type YmlConfigTemplate struct {
	Enable  bool   `yaml:"enable"`
	Cache   bool   `yaml:"cache"`
	Root    string `yaml:"root"`
	Suffix  string `yaml:"suffix"`
	FuncMap template.FuncMap
}

type YmlConfigServerTls struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

func Default() *Config {
	return &Config{
		Server: &YmlServerConfig{
			MaxHeaderSize: http.DefaultMaxHeaderBytes,
			Timeout: &YmlServerConfigTimeout{
				Read:       time.Minute,
				ReadHeader: time.Minute,
				Write:      time.Minute,
				Idle:       time.Second,
			},
		},
		Flygo: &YmlConfig{
			Dev: &YmlConfigDev{
				Debug: false,
			}, Banner: &YmlConfigBanner{
				Enable: true,
				Type:   "default",
				File:   "banner.txt",
				Text:   "FLYGO",
			}, Server: &YmlConfigServer{
				Host: "0.0.0.0",
				Port: 80,
				TLS: &YmlConfigServerTls{
					Enable:   false,
					CertFile: "",
					KeyFile:  "",
				},
			},
			Template: &YmlConfigTemplate{
				Enable: true,
				Cache:  true,
				Root:   "./templates",
				Suffix: ".html",
			},
		}}
}
