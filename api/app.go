package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maxiloEmmmm/go-web/contact"
)

type api struct {
	method    string
	path      string
	do        contact.GinHelpHandlerFunc
	auth      bool
	customize func(route gin.IRouter)
}

func newApi(method, path string, do contact.GinHelpHandlerFunc) *api {
	return &api{
		method: method,
		path:   path,
		do:     do,
	}
}

func newAuthApi(method, path string, do contact.GinHelpHandlerFunc) *api {
	return &api{
		method: method,
		path:   path,
		do:     do,
		auth:   true,
	}
}

func newFuncApi(cb func(route gin.IRouter)) *api {
	return &api{
		customize: cb,
	}
}

func newAuthFuncApi(cb func(route gin.IRouter)) *api {
	return &api{
		customize: cb,
		auth:      true,
	}
}

func (a *api) Handle(route gin.IRouter) {
	if a.customize != nil {
		a.customize(route)
	} else {
		route.Handle(a.method, a.path, contact.GinHelpHandle(a.do))
	}
}

func (a *api) ShouldAuth() bool {
	return a.auth
}

var Apis []*api
