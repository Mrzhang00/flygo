package flygo

import (
	"fmt"
	rf "github.com/billcoding/flygo/rest"
	"net/http"
)

//REST
func (a *App) REST(c ...rf.Controller) *App {
	a.controllers = append(a.controllers, c...)
	return a
}

//routeRestControllers
func (a *App) routeRestControllers() *App {
	for _, c := range a.controllers {
		prefix := c.Prefix()

		if c.GET() != nil {
			getPattern := fmt.Sprintf("%s/{RESTFUL_ID}", prefix)
			a.Route(http.MethodGet, getPattern, c.GET())
		}

		if c.GETS() != nil {
			getsPattern := fmt.Sprintf("%s", prefix)
			a.Route(http.MethodGet, getsPattern, c.GETS())
		}

		if c.POST() != nil {
			postPattern := fmt.Sprintf("%s", prefix)
			a.Route(http.MethodPost, postPattern, c.POST())
		}

		if c.PUT() != nil {
			putPattern := fmt.Sprintf("%s", prefix)
			a.Route(http.MethodPut, putPattern, c.PUT())
		}

		if c.DELETE() != nil {
			deletePattern := fmt.Sprintf("%s/{id}", prefix)
			a.Route(http.MethodDelete, deletePattern, c.DELETE())
		}

	}
	return a
}
