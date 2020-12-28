package flygo

import (
	"github.com/billcoding/calls"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

//parseConfig
func (a *App) parseConfig() *App {
	rot := a.Config.Template.Root
	// "" "." "./" "./templates"
	calls.True(rot == "" || rot == "." || rot == "./", func() {
		execdir, _ := os.Executable()
		rot = filepath.Dir(execdir)
	})
	calls.NEmpty(strings.TrimPrefix(rot, "./"), func() {
		execdir, _ := os.Executable()
		rot = filepath.Join(filepath.Dir(execdir), strings.TrimPrefix(rot, "./"))
	})
	a.Config.Template.Root = rot
	return a
}
