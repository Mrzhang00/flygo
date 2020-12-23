package config

import (
	"gopkg.in/yaml.v2"
	"html/template"
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
	Dev      *YmlConfigDev      `yaml:"dev"`
	Banner   *YmlConfigBanner   `yaml:"banner"`
	Server   *YmlConfigServer   `yaml:"server"`
	Template *YmlConfigTemplate `yaml:"template"`
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

//Define YmlConfigTemplate struct
type YmlConfigTemplate struct {
	Enable  bool   `yaml:"enable"`
	Cache   bool   `yaml:"cache"`
	Root    string `yaml:"root"`
	Suffix  string `yaml:"suffix"`
	FuncMap template.FuncMap
}

//Define YmlConfigServerTls struct
type YmlConfigServerTls struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

func Default() *Config {
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
		&YmlConfigTemplate{
			Enable: true,
			Cache:  true,
			Root:   "./templates",
			Suffix: ".html",
		},
	}}
}
