package config

import (
	"gopkg.in/yaml.v3"
	t "html/template"
	"net/http"
	"time"
)

// Config struct
type Config struct {
	Server   *Server   `yaml:"server"`
	Banner   *Banner   `yaml:"banner"`
	Template *Template `yaml:"template"`
}

// Unmarshal yaml
func (yc *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, yc)
}

type Server struct {
	Host              string        `yaml:"host"`
	Port              int           `yaml:"port"`
	TLS               *TLS          `yaml:"tls"`
	MaxHeaderSize     int           `yaml:"max_header_size"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
}

type Banner struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
	File   string `yaml:"file"`
	Text   string `yaml:"text"`
}

type Template struct {
	Enable  bool   `yaml:"enable"`
	Cache   bool   `yaml:"cache"`
	Root    string `yaml:"root"`
	Suffix  string `yaml:"suffix"`
	FuncMap t.FuncMap
}

type TLS struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}

// Default Config
func Default() *Config {
	return &Config{
		Server: &Server{
			Host: "0.0.0.0",
			Port: 80,
			TLS: &TLS{
				Enable:   false,
				CertFile: "",
				KeyFile:  "",
			},
			MaxHeaderSize:     http.DefaultMaxHeaderBytes,
			ReadTimeout:       time.Minute,
			ReadHeaderTimeout: time.Minute,
			WriteTimeout:      time.Minute,
			IdleTimeout:       time.Second,
		},
		Banner: &Banner{
			Enable: true,
			Type:   "default",
			File:   "banner.txt",
			Text:   "FLY GO GO",
		},
		Template: &Template{
			Enable: true,
			Cache:  true,
			Root:   "templates",
			Suffix: ".html",
		},
	}
}
