package flygo

import (
	"fmt"
	"time"
)

//Define appInfo struct
type AppInfo struct {
	App   *App
	Id    string
	Name  string
	Ctime time.Time
}

//Return new appInfo
func newAppInfo(id string, app *App) *AppInfo {
	return &AppInfo{
		App:   app,
		Id:    id,
		Name:  fmt.Sprintf("%s%s", AppGroup.prefix, id),
		Ctime: time.Now(),
	}
}
