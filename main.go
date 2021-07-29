package main

import (
	"github.com/gin-gonic/gin"
	appApi "github.com/maxiloEmmmm/diy-datav/api"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/dataset"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/user"
	go_tool "github.com/maxiloEmmmm/go-tool"
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
	curd.Apis.User.Filter.CreatePipe = func(help *contact.GinHelp, createPipe *model.UserCreate) {
		username, _ := createPipe.Mutation().Username()
		if app.Db.User.Query().Where(user.Username(username)).ExistX(help.AppContext) {
			help.InValid("username", "dup")
		}

		password, _ := createPipe.Mutation().Password()
		createPipe.SetPassword(go_tool.Md5(password))
	}
	curd.Apis.User.Filter.UpdatePipe = func(help *contact.GinHelp, old *model.User, updatePipe *model.UserUpdateOne) {
		password, _ := updatePipe.Mutation().Password()
		if password == "" {
			updatePipe.SetPassword(old.Password)
		}
	}
	curd.Apis.User.Filter.ListData = func(help *contact.GinHelp, items []*model.User) interface{} {
		return go_tool.ArrayMap(items, func(i interface{}) interface{} {
			u := i.(*model.User)
			u.Password = ""
			return u
		})
	}
	curd.Apis.User.Filter.GetData = func(help *contact.GinHelp, item *model.User) *model.User {
		item.Password = ""
		return item
	}

	curd.Apis.TypeConfig.Filter.ListPipe = func(help *contact.GinHelp, listPipe *model.TypeConfigQuery) {
		if val, exist := help.GetQuery("type"); exist && val != "" {
			listPipe.Where(typeconfig.Type(val))
		}

		if val, exist := help.GetQuery("title"); exist && val != "" {
			listPipe.Where(typeconfig.TitleContains(val))
		}
	}
	curd.Apis.TypeConfig.Filter.UpdateAfter = func(help *contact.GinHelp, old *model.TypeConfig, item *model.TypeConfig) {
		if item.Type == dataset.Sql {
			_ = dataset.RemoveConnect(item.ID)
		}
	}
	curd.Route("/", apiGroup, []string{model.TypeTypeConfig, model.TypeAssets, model.TypeUser})
	engine.Run(":8000")
}
