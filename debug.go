package flygo

import (
	"github.com/billcoding/calls"
	"gopkg.in/yaml.v2"
	"reflect"
)

//DebugTrace
func (a *App) DebugTrace(call func()) *App {
	if a.Config.Dev.Debug {
		call()
	}
	return a
}

//debugTrace
func (a *App) debugTrace() {
	a.DebugTrace(func() {
		a.printConfigs()
		a.printRestControllers()
		a.printRouters()
		a.printMiddlewares()
	})
}

//printConfigs
func (a *App) printConfigs() {
	bytes, err := yaml.Marshal(a.Config)
	calls.NNil(err, func() {
		a.Logger.Warn("[Config]%v", err)
	})
	calls.Nil(err, func() {
		a.Logger.Info("[Config]\n%v", string(bytes))
	})
}

//printRestControllers
func (a *App) printRestControllers() {
	for _, c := range a.controllers {
		a.Logger.Info("[REST]controller registered [%v]", reflect.TypeOf(c))
	}
}

//printRouters
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

//printMiddlewares
func (a *App) printMiddlewares() {
	for _, v := range a.middlewares {
		calls.NNil(v, func() {
			a.Logger.Info("[Middleware]used [%v]", v.Name())
		})
	}
}
