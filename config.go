package flygo

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (a *App) parseYml() {
	file := a.ConfigFile
	if file != "" {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			a.Logger.Debugf("[parseYml]%v", err)
		} else {
			err = a.Config.Unmarshal(bytes)
			if err != nil {
				if err != nil {
					a.Logger.Debugf("[parseYml]%v", err)
				}
			}
		}
	}
}

func (a *App) parseConfig() *App {
	rootDir, _ := os.Getwd()
	rot := a.Config.Flygo.Template.Root
	if strings.HasPrefix(rot, "/") {
	} else {
		switch rot {
		case "", ".", "./":
			rot = rootDir
		default:
			rot = strings.TrimPrefix(rot, "./")
			rot = filepath.Join(rootDir, rot)
		}
	}
	a.Config.Flygo.Template.Root = rot
	return a
}
