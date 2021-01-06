package flygo

import (
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

//parseConfig
func (a *App) parseConfig() *App {
	execdir, _ := os.Executable()
	execroot := filepath.Dir(execdir)
	rot := a.Config.Template.Root
	// "" "." "./" "./templates" "templates/" "/templates" "/templates/"
	if strings.HasPrefix(rot, "/") {
		//Absolute path
	} else {
		switch rot {
		case "", ".", "./":
			rot = execroot
		default:
			//"./templates" "templates/"
			rot = strings.TrimPrefix(rot, "./")
			rot = filepath.Join(execroot, rot)
		}
	}
	a.Config.Template.Root = rot
	return a
}
