package flygo

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

type YmlConfig struct {
	Flygo YmlConfigFlygo
}

func (yc *YmlConfig) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, yc)
}

func (a *App) checkConfig() {
	a.Config.Flygo.Server.WebRoot = strings.Trim(app.Config.Flygo.Server.WebRoot, string(filepath.Separator))
	a.Config.Flygo.Server.ContextPath = strings.Trim(app.Config.Flygo.Server.WebRoot, "/")

	a.Config.Flygo.Static.Pattern = strings.Trim(app.Config.Flygo.Static.Pattern, "/")
	a.Config.Flygo.Static.Prefix = strings.Trim(app.Config.Flygo.Static.Prefix, string(filepath.Separator))

	a.Config.Flygo.View.Prefix = strings.Trim(app.Config.Flygo.View.Prefix, string(filepath.Separator))
	a.Config.Flygo.View.Suffix = strings.TrimSpace(app.Config.Flygo.View.Suffix)
}

type YmlConfigFlygo struct {
	Dev      YmlConfigDev
	Banner   YmlConfigBanner
	Server   YmlConfigServer
	Static   YmlConfigStatic
	View     YmlConfigView
	Template YmlConfigTemplate
	Session  YmlConfigSession
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

type YmlConfigSession struct {
	Enable  bool          `yaml:"enable"`
	Timeout time.Duration `yaml:"timeout"`
}

type YmlConfigValidate struct {
	Err YmlConfigValidateErr
}

type YmlConfigValidateErr struct {
	Code int `yaml:"code"`
}

func (a *App) parseYml() {
	file := a.ConfigFile
	if file == "" {
		return
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = a.Config.Unmarshal(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	a.inita()
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
					"css":  "text/css;charset=utf-8",
					"json": "application/json;charset=utf-8",
					"jpg":  "image/jpg",
					"png":  "image/png",
					"gif":  "image/gif",
					"ico":  "image/x-icon",
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
			Session: YmlConfigSession{
				Enable:  false,
				Timeout: time.Hour,
			},
			Validate: YmlConfigValidate{
				Err: YmlConfigValidateErr{
					Code: 1,
				},
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

type YmlConfigFlygoLog struct {
	Type string `yaml:"type"`
	File YmlConfigFlygoLogFile
}

type YmlConfigFlygoLogFile struct {
	Out string `yaml:"out"`
	Err string `yaml:"err"`
}
