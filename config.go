package flygo

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type YmlConfig struct {
	Flygo YmlConfigFlygo
}

func (yc *YmlConfig) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, yc)
}

func (a *App) checkConfig() {
	a.Config.Flygo.Server.WebRoot = strings.TrimRight(a.Config.Flygo.Server.WebRoot, string(filepath.Separator))
	a.Config.Flygo.Server.ContextPath = strings.TrimRight(a.Config.Flygo.Server.ContextPath, "/")

	a.Config.Flygo.Static.Pattern = strings.Trim(a.Config.Flygo.Static.Pattern, "/")
	a.Config.Flygo.Static.Prefix = strings.Trim(a.Config.Flygo.Static.Prefix, string(filepath.Separator))

	a.Config.Flygo.View.Prefix = strings.Trim(a.Config.Flygo.View.Prefix, string(filepath.Separator))
	a.Config.Flygo.View.Suffix = strings.TrimLeft(a.Config.Flygo.View.Suffix, ".")
}

type YmlConfigFlygo struct {
	Dev      YmlConfigDev
	Banner   YmlConfigBanner
	Server   YmlConfigServer
	Static   YmlConfigStatic
	View     YmlConfigView
	Template YmlConfigTemplate
	Validate YmlConfigValidate
	Log      YmlConfigFlygoLog
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
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	WebRoot     string `yaml:"webRoot"`
	ContextPath string `yaml:"contextPath"`
	Tls         YmlConfigServerTls
}

type YmlConfigServerTls struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

type YmlConfigGlobal struct {
}

type YmlConfigStatic struct {
	Enable  bool              `yaml:"enable"`
	Pattern string            `yaml:"pattern"`
	Prefix  string            `yaml:"prefix"`
	Cache   bool              `yaml:"cache"`
	Mimes   map[string]string `yaml:"mimes"`
	Favicon YmlConfigStaticFavicon
}

type YmlConfigStaticFavicon struct {
	Enable bool `yaml:"enable"`
}

type YmlConfigView struct {
	Enable bool   `yaml:"enable"`
	Prefix string `yaml:"prefix"`
	Suffix string `yaml:"suffix"`
	Cache  bool   `yaml:"cache"`
}

type YmlConfigTemplate struct {
	Enable bool `yaml:"enable"`
	Delims YmlConfigTemplateDelims
}

type YmlConfigTemplateDelims struct {
	Left  string `yaml:"left"`
	Right string `yaml:"right"`
}

type YmlConfigValidate struct {
	Code int `yaml:"code"`
}

type YmlConfigFlygoLog struct {
	Type string `yaml:"type"`
	File YmlConfigFlygoLogFile
}

type YmlConfigFlygoLogFile struct {
	Out string `yaml:"out"`
	Err string `yaml:"err"`
}

func (a *App) parseYml() {
	file := a.ConfigFile
	if file != "" {
		bytes, err := ioutil.ReadFile(file)
		a.DebugTrace(func() {
			if err != nil {
				a.Logger.Warn("[parseYml]%v", err)
			}
		})
		if err == nil {
			err = a.Config.Unmarshal(bytes)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	a.inita()
}

func (a *App) printConfig() {
	bytes, _ := yaml.Marshal(a.Config)
	a.Logger.Info("\n%v", string(bytes))
}

func defaultYmlConfig() *YmlConfig {
	return &YmlConfig{
		Flygo: YmlConfigFlygo{
			Dev: YmlConfigDev{
				Debug: false,
			},
			Banner: YmlConfigBanner{
				Enable: true,
				Type:   "default",
				File:   "banner.txt",
				Text:   "FLYGO",
			},
			Server: YmlConfigServer{
				Host:        "",
				Port:        80,
				WebRoot:     "",
				ContextPath: "",
				Tls: YmlConfigServerTls{
					Enable:   false,
					CertFile: "",
					KeyFile:  "",
				},
			},
			Static: YmlConfigStatic{
				Enable:  false,
				Pattern: "static",
				Prefix:  "static",
				Cache:   false,
				Mimes: map[string]string{
					"txt":  "text/plain;charset=utf-8",
					"html": "text/html;charset=utf-8",
					"js":   "text/javascript;charset=utf-8",
					"css":  "text/css; charset=utf-8",
					"json": "application/json;charset=utf-8",
					"xml":  "text/xml;charset=utf-8",
					"bmp":  "image/bmp",
					"jpg":  "image/jpg",
					"png":  "image/png",
					"gif":  "image/gif",
					"ico":  "image/x-icon",
				},
				Favicon: YmlConfigStaticFavicon{
					Enable: false,
				},
			},
			View: YmlConfigView{
				Enable: false,
				Prefix: "templates",
				Suffix: "html",
				Cache:  false,
			},
			Template: YmlConfigTemplate{
				Enable: false,
				Delims: YmlConfigTemplateDelims{
					Left:  "{{",
					Right: "}}",
				},
			},
			Validate: YmlConfigValidate{
				Code: 1,
			},
			Log: YmlConfigFlygoLog{
				Type: "stdout",
				File: YmlConfigFlygoLogFile{
					Out: "out.log",
					Err: "err.log",
				},
			},
		},
	}
}
