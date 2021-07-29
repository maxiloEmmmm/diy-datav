package main

import (
	"context"
	"github.com/gin-gonic/gin"
	appApi "github.com/maxiloEmmmm/diy-datav/api"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/permission"
	"github.com/maxiloEmmmm/go-web/contact"
)

func main() {
	ctx := context.Background()

	app.Init(ctx)
	defer app.Close()

	engine := gin.Default()
	engine.Use(contact.GinCors())

	apiGroup := engine.Group("/api")

	auth := apiGroup.Group("")
	auth.Use(contact.GinRouteAuth())
	auth.Use()
	auth.Use(contact.GinCasbin(contact.GinCasbinOption{
		Sub: func(c *gin.Context) string {
			return app.User(c)
		},
		Obj: func(context *gin.Context) string {
			return (&permission.PRRouter{Path: context.Request.URL.Path}).Key()
		},
		Act: func(context *gin.Context) string {
			return context.Request.Method
		},
	}))

	for _, api := range appApi.Apis {
		if api.ShouldAuth() {
			api.Handle(auth)
		} else {
			api.Handle(apiGroup)
		}
	}

	app.Engine = engine
	app.InitPermission(ctx, engine)

	engine.Run(":8000")
}
