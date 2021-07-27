package main

import (
	"github.com/gin-gonic/gin"
	appApi "github.com/maxiloEmmmm/diy-datav/api"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/dataset"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"github.com/maxiloEmmmm/go-web/contact"
)

func main() {
	contact.InitLog()

	defer app.Db.Close()

	engine := gin.Default()
	engine.Use(contact.GinCors())

	apiGroup := engine.Group("/api")

	for _, api := range appApi.Apis {
		api.Handle(apiGroup)
	}

	curd := model.NewCurdBuilder(app.Db)
	curd.Apis.TypeConfig.Filter.ListPipe = func(help *contact.GinHelp, listPipe *model.TypeConfigQuery) {
		if val, exist := help.GetQuery("type"); exist && val != "" {
			listPipe.Where(typeconfig.Type(val))
		}

		if val, exist := help.GetQuery("title"); exist && val != "" {
			listPipe.Where(typeconfig.TitleContains(val))
		}
	}
	curd.Apis.TypeConfig.Filter.UpdateAfter = func(help *contact.GinHelp, item *model.TypeConfig) {
		if item.Type == dataset.Sql {
			_ = dataset.RemoveConnect(item.ID)
		}
	}
	curd.Route("/", apiGroup, []string{model.TypeTypeConfig})
	engine.Run(":8000")
}
