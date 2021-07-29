package api

import (
	"context"
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/menu"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/user"
	"github.com/maxiloEmmmm/diy-datav/pkg/permission"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
	"time"
)

func init() {
	Apis = append(Apis, newApi(http.MethodPost, "auth/login", AuthLogin))

	Apis = append(Apis, newAuthApi(http.MethodGet, "auth/refresh_token", RefreshToken))
	Apis = append(Apis, newAuthApi(http.MethodGet, "auth/menu", AuthMenu))
}

func AuthLogin(c *contact.GinHelp) {
	jsonX := &struct {
		Username string `json:"username",binding:"required"`
		Password string `json:"password",binding:"required"`
	}{}
	c.InValidBind(&jsonX)

	userX := app.Db.User.Query().Where(user.UsernameEQ(jsonX.Username)).FirstX(c.AppContext)

	if userX.Password == go_tool.Md5(jsonX.Password) {
		c.Resource(contact.H{"token": contact.JwtNew().
			SetExpiresAtWeek().
			SetPrimaryKey(fmt.Sprintf("%d", userX.ID)).
			GenerateToken()})
	} else {
		c.InValid("auth", "password error")
	}
}

func RefreshToken(c *contact.GinHelp) {
	auth, _ := c.Get("auth")
	c.Resource(contact.H{"token": auth.(*contact.JwtLib).RefreshToken(7 * 24 * time.Hour)})
}

func AuthMenu(c *contact.GinHelp) {
	c.Resource(DepMenuViewPermission(
		c.AppContext, app.User(c.AppContext),
		app.Db.Menu.Query().WithChildren().Where(menu.Not(menu.HasParent())).AllX(c.AppContext),
	))
}

func DepMenuViewPermission(ctx context.Context, sub string, menus []*model.Menu) []*types.MenuPermission {
	mp := make([]*types.MenuPermission, 0, len(menus))
	has := false
	for _, menu := range menus {
		ok := permission.Pass(sub, &permission.PRMenu{Menu: menu}, permission.AccessMenuAction)
		obj := &types.MenuPermission{
			Id:    fmt.Sprintf("%d", menu.ID),
			Url:   menu.URL,
			Title: menu.Title,
			Edges: types.MenuPermissionEdge{},
		}
		for _, child := range menu.Edges.Children {
			child.Edges.Children = child.QueryChildren().WithChildren().AllX(ctx)
		}
		obj.Edges.Children = append(obj.Edges.Children, DepMenuViewPermission(ctx, sub, menu.Edges.Children)...)

		if len(obj.Edges.Children) != 0 {
			ok = true
		}

		if ok {
			mp = append(mp, obj)
		}

		has = has || ok
	}

	return mp
}
