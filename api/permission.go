package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/menu"
	"github.com/maxiloEmmmm/diy-datav/pkg/permission"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
	"strings"
)

func init() {
	Apis = append(Apis, newAuthApi(http.MethodPost, "permission/change", PermissionChange))
	Apis = append(Apis, newAuthApi(http.MethodGet, "permission/api", PermissionApi))
	Apis = append(Apis, newAuthApi(http.MethodGet, "permission/menu", PermissionMenu))
	Apis = append(Apis, newAuthApi(http.MethodGet, "permission/view", PermissionView))
	Apis = append(Apis, newAuthFuncApi(func(route gin.IRouter) {
		contact.PermissionRoute(route, "")
	}))
}

func PermissionChange(c *contact.GinHelp) {
	body := &types.Permission{}
	c.InValidBind(body)

	for _, add := range body.Add {
		rules := strings.Split(add, ",")
		contact.Permission.AddPolicy(rules)
		rules[0] = app.Config.Permission.Role
		contact.Permission.AddPolicy(rules)
	}
	for _, remove := range body.Remove {
		contact.Permission.RemovePolicy(strings.Split(remove, ","))
	}
	c.ResourceCreate(nil)
}

func PermissionView(c *contact.GinHelp) {
	sub := c.DefaultQuery("sub", "")
	if len(sub) == 0 {
		c.InValid("resource", "not found")
	}

	c.Resource(DepViewPermission(
		c.AppContext, sub,
		app.Db.View.Query().AllX(c.AppContext),
	))
}

func DepViewPermission(ctx context.Context, sub string, views []*model.View) []*types.ViewPermission {
	mp := make([]*types.ViewPermission, 0, len(views))
	for _, view := range views {
		mp = append(mp, &types.ViewPermission{
			Id:    fmt.Sprintf("%d", view.ID),
			Title: view.Desc,
			Eft:   permission.Has(sub, &permission.PRView{View: view}, permission.GetInfoAction),
		})
	}
	return mp
}

func PermissionApi(c *contact.GinHelp) {
	sub := c.DefaultQuery("sub", "")
	if len(sub) == 0 {
		c.InValid("resource", "not found")
	}

	ma := go_tool.ArrayKeyBy(go_tool.ArrayMap(app.Engine.Routes(), func(d interface{}) interface{} {
		tmp := d.(gin.RouteInfo)
		return contact.H{
			"method": tmp.Method,
			"path":   tmp.Path,
			// index 1, use after `api` prefix
			"group": strings.Split(strings.Trim(tmp.Path, "/"), "/")[1],
			"eft":   permission.Has(sub, &permission.PRRouter{Path: tmp.Path}, permission.NewAction(tmp.Method, tmp.Method)),
		}
	}), "group")
	rs := make(map[string]interface{}, len(ma))
	for k, v := range ma {
		rs[k.(string)] = v
	}
	c.Resource(rs)
}

func PermissionMenu(c *contact.GinHelp) {
	sub := c.DefaultQuery("sub", "")
	if len(sub) == 0 {
		c.InValid("resource", "not found")
	}

	c.Resource(DepMenuPermission(
		c.AppContext, sub,
		app.Db.Menu.Query().WithChildren().Where(menu.Not(menu.HasParent())).AllX(c.AppContext),
	))
}

func DepMenuPermission(ctx context.Context, sub string, menus []*model.Menu) []*types.MenuPermission {
	mp := make([]*types.MenuPermission, 0, len(menus))
	for _, menu := range menus {
		obj := &types.MenuPermission{
			Id:    fmt.Sprintf("%d", menu.ID),
			Url:   menu.URL,
			Title: menu.Title,
			Eft:   permission.Has(sub, &permission.PRMenu{Menu: menu}, permission.AccessMenuAction),
			Edges: types.MenuPermissionEdge{},
		}
		for _, child := range menu.Edges.Children {
			child.Edges.Children = child.QueryChildren().WithChildren().AllX(ctx)
		}
		obj.Edges.Children = append(obj.Edges.Children, DepMenuPermission(ctx, sub, menu.Edges.Children)...)

		mp = append(mp, obj)
	}
	return mp
}
