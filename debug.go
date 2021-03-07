package flygo

import (
	"gopkg.in/yaml.v2"
	"reflect"
)

// DebugTrace call
func (a *App) DebugTrace(call func()) *App {
	if a.Config.Flygo.Dev.Debug {
		call()
	}
	return a
}

func (a *App) debugTrace() {
	a.DebugTrace(func() {
		a.printConfigs()
		a.printRestControllers()
		a.printRouters()
		a.printMiddlewares()
	})
}

func (a *App) printConfigs() {
	bytes, err := yaml.Marshal(a.Config)
	if err != nil {
		a.Logger.Warn("[Config]%v", err)
	} else {
		a.Logger.Info("[Config]\n%v", string(bytes))
	}
}

func (a *App) printRestControllers() {
	for _, c := range a.controllers {
		a.Logger.Info("[REST]controller registered [%v]", reflect.TypeOf(c))
	}
}

func (a *App) printRouters() {
	for k := range a.parsedRouters.Simples {
		a.Logger.Info("[Router]simple routed [%v]", k)
	}
	for k, v := range a.parsedRouters.Dynamics {
		for kk, vv := range v {
			a.Logger.Info("[Router]dynamic routed [%v:%v] Pos[%v]", kk, k, vv.Pos)
		}
	}
}

func (a *App) printMiddlewares() {
	for _, v := range a.middlewares {
		if v != nil {
			a.Logger.Info("[Middleware]used [%v]", v.Name())
		}
	}
}
