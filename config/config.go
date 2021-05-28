package config

import (
	"gopkg.in/yaml.v2"
	"html/template"
	"net/http"
	"time"
)

// Config struct
type Config struct {
	Server *YmlServerConfig `yaml:"server"`
	Flygo  *YmlConfig       `yaml:"flygo"`
}

// Unmarshal yaml
func (yc *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, yc)
}

// YmlServerConfig struct
type YmlServerConfig struct {
	MaxHeaderSize int                     `yaml:"maxHeaderSize"`
	Timeout       *YmlServerConfigTimeout `yaml:"timeout"`
}

// YmlServerConfigTimeout struct
type YmlServerConfigTimeout struct {
	Read       time.Duration `yaml:"read"`
	ReadHeader time.Duration `yaml:"readHeader"`
	Write      time.Duration `yaml:"write"`
	Idle       time.Duration `yaml:"idle"`
}

// YmlConfig struct
type YmlConfig struct {
	Banner   *YmlConfigBanner   `yaml:"banner"`
	Server   *YmlConfigServer   `yaml:"server"`
	Template *YmlConfigTemplate `yaml:"template"`
}

// YmlConfigBanner struct
type YmlConfigBanner struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
	File   string `yaml:"file"`
	Text   string `yaml:"text"`
}

// YmlConfigServer struct
type YmlConfigServer struct {
	Host string              `yaml:"host"`
	Port int                 `yaml:"port"`
	TLS  *YmlConfigServerTls `yaml:"tls"`
}

// YmlConfigTemplate struct
type YmlConfigTemplate struct {
	Enable  bool   `yaml:"enable"`
	Cache   bool   `yaml:"cache"`
	Root    string `yaml:"root"`
	Suffix  string `yaml:"suffix"`
	FuncMap template.FuncMap
}

// YmlConfigServerTls struct
type YmlConfigServerTls struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

// Default Config
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
			Banner: &YmlConfigBanner{
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
