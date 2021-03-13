package flygo

import (
	"fmt"
	"github.com/billcoding/flygo/rest"
	"github.com/billcoding/flygo/util"
	"net/http"
)

// REST Route REST ful Controller
func (a *App) REST(c ...rest.Controller) *App {
	a.controllers = append(a.controllers, c...)
	return a
}

func (a *App) routeRestControllers() *App {
	for _, c := range a.controllers {
		prefix := util.TrimSpecialChars(c.Prefix())

		if c.GET() != nil {
			getPattern := fmt.Sprintf("%s/{RESTFUL_ID}", prefix)
			a.Route(http.MethodGet, getPattern, c.GET())
		}

		if c.GETS() != nil {
			a.Route(http.MethodGet, prefix, c.GETS())
		}

		if c.POST() != nil {
			a.Route(http.MethodPost, prefix, c.POST())
		}

		if c.PUT() != nil {
			a.Route(http.MethodPut, prefix, c.PUT())
		}

		if c.DELETE() != nil {
			deletePattern := fmt.Sprintf("%s/{RESTFUL_ID}", prefix)
			a.Route(http.MethodDelete, deletePattern, c.DELETE())
		}

	}
	return a
}
