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

func (a *App) parseConfig() *App {
	execdir, _ := os.Executable()
	execroot := filepath.Dir(execdir)
	rot := a.Config.Flygo.Template.Root

	if strings.HasPrefix(rot, "/") {

	} else {
		switch rot {
		case "", ".", "./":
			rot = execroot
		default:

			rot = strings.TrimPrefix(rot, "./")
			rot = filepath.Join(execroot, rot)
		}
	}
	a.Config.Flygo.Template.Root = rot
	return a
}
