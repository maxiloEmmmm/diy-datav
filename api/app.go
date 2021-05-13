package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maxiloEmmmm/go-web/contact"
)

type api struct {
	method string
	path string
	do contact.GinHelpHandlerFunc
}

func newApi(method, path string, do contact.GinHelpHandlerFunc) *api {
	return &api{
		method: method,
		path: path,
		do: do,
	}
}

func(a *api)Handle(route gin.IRoutes) {
	route.Handle(a.method, a.path, contact.GinHelpHandle(a.do))
}

var Apis []*api
