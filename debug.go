package flygo

import (
	"gopkg.in/yaml.v2"
	"reflect"
)

func (a *App) debugTrace() {
	a.printConfigs()
	a.printRestControllers()
	a.printRouters()
	a.printMiddlewares()
}

func (a *App) printConfigs() {
	bytes, err := yaml.Marshal(a.Config)
	if err != nil {
		a.Logger.Debugf("Config: %v", err)
	} else {
		a.Logger.Debugf("Config: \n%v", string(bytes))
	}
}

func (a *App) printRestControllers() {
	for _, c := range a.controllers {
		a.Logger.Debugf("REST]controller registered %v: ", reflect.TypeOf(c))
	}
}

func (a *App) printRouters() {
	for k := range a.parsedRouters.Simples {
		a.Logger.Debugf("router: simple routed %v: ", k)
	}
	for k, v := range a.parsedRouters.Dynamics {
		for kk, vv := range v {
			a.Logger.Debugf("router: dynamic routed %v:%v Pos: %v", kk, k, vv.Pos)
		}
	}
}

func (a *App) printMiddlewares() {
	for _, v := range a.middlewares {
		if v != nil {
			a.Logger.Debugf("middleware: used %v", v.Name())
		}
	}
}
