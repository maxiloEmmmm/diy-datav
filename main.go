package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
)

func main() {
	engine := gin.Default()
	curd := model.NewCurdBuilder(app.Db)
	curd.Route("/api", engine, []string{model.TypeTypeConfig})
	engine.Run(":8000")
}
