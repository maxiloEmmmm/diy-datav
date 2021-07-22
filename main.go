package main

import (
	"github.com/gin-gonic/gin"
	appApi "github.com/maxiloEmmmm/diy-datav/api"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"github.com/maxiloEmmmm/go-web/contact"
)

func main() {
	contact.InitLog()

	defer app.Db.Close()

	engine := gin.Default()

	apiGroup := engine.Group("/api")

	for _, api := range appApi.Apis {
		api.Handle(apiGroup)
	}

	curd := model.NewCurdBuilder(app.Db)
	curd.Apis.TypeConfig.Filter.GetPipe = func(help *contact.GinHelp, getPipe *model.TypeConfigQuery) {
		if val, exist := help.GetQuery("type"); exist && val != "" {
			getPipe.Where(typeconfig.Type(val))
		}

		if val, exist := help.GetQuery("title"); exist && val != "" {
			getPipe.Where(typeconfig.TitleContains(val))
		}
	}
	curd.Route("/", apiGroup, []string{model.TypeTypeConfig})
	engine.Run(":8000")
}
