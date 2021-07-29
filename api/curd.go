package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/dataset"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/menu"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/user"
	"github.com/maxiloEmmmm/diy-datav/pkg/permission"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"github.com/maxiloEmmmm/go-web/contact"
)

func init() {
	Apis = append(Apis, newAuthFuncApi(func(route gin.IRouter) {
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
		curd.Apis.Menu.Api.Fields.Create.SetFields([]string{menu.FieldTitle, menu.FieldURL})
		curd.Apis.Menu.Filter.CreatePipe = func(help *contact.GinHelp, createPipe *model.MenuCreate, edges model.MenuEdges) {
			if edges.Parent != nil {
				item, err := app.Db.Menu.Get(help.AppContext, edges.Parent.ID)
				help.AssetsInValid("menu.parent.found", err)
				createPipe.SetParent(item)
			}
		}
		curd.Apis.Menu.Filter.CreateAfter = func(help *contact.GinHelp, item *model.Menu, edges model.MenuEdges) {
			app.AddAdminPolicy(&permission.PRMenu{Menu: item}, permission.AccessMenuAction)
		}
		curd.Apis.Menu.Filter.ListPipe = func(help *contact.GinHelp, listPipe *model.MenuQuery) {
			listPipe.WithChildren().Where(menu.Not(menu.HasParent()))
		}
		curd.Apis.Menu.Filter.ListData = func(help *contact.GinHelp, items []*model.Menu) interface{} {
			DepMenu(help.AppContext, items)
			return items
		}
		curd.Apis.Menu.Filter.DeleteBefore = func(help *contact.GinHelp, item *model.Menu) {
			app.Db.Menu.Delete().Where(menu.HasParentWith(menu.ID(item.ID))).ExecX(help.AppContext)
			app.RemoveAdminPolicy(&permission.PRMenu{Menu: item}, permission.AccessMenuAction)
		}
		curd.Route("/", route, []string{model.TypeTypeConfig, model.TypeAssets, model.TypeUser, model.TypeMenu})
	}))
}

func DepMenu(ctx context.Context, menus []*model.Menu) {
	for _, menu := range menus {
		for _, child := range menu.Edges.Children {
			child.Edges.Children = child.QueryChildren().WithChildren().AllX(ctx)
			DepMenu(ctx, child.Edges.Children)
		}
	}
}
