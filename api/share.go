package api

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
	"time"
)

func init() {
	Apis = append(Apis,
		newApi(http.MethodGet, "share.view/:id", ShareGet),
		newApi(http.MethodGet, "share.data/:sid/:id", ShareData),
	)
}

func ShareGet(c *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	c.InValidBindUri(uri)

	share, err := app.Db.Share.Get(c.AppContext, uri.Id)
	c.AssetsInValid("find.share", err)

	if share.EndAt.Before(time.Now()) {
		c.InValid("share.invalid", "")
	}

	v, err := share.QueryView().First(c.AppContext)
	c.AssetsInValid(`view.get`, err)

	vi, err := app.Db.View.Query().WithBg().WithBlocks().Where(view.ID(v.ID)).First(c.AppContext)
	c.AssetsInValid("get", err)

	vd, err := NewView(vi, false)
	c.AssetsInValid("new", err)

	c.Resource(vd)
}

func ShareData(c *contact.GinHelp) {
	uri := &struct {
		Id  int `uri:"id"`
		Sid int `uri:"sid"`
	}{}
	c.InValidBindUri(uri)

	data, err := service.NewDataSetService(c.AppContext).Load(uri.Id, func(view *model.View) bool {
		share, err := app.Db.Share.Get(c.AppContext, uri.Sid)
		c.AssetsInValid("find.share", err)

		if share.EndAt.Before(time.Now()) {
			c.InValid("share.invalid", "")
		}

		v, err := share.QueryView().First(c.AppContext)
		c.AssetsInValid(`view.get`, err)

		return view.ID == v.ID
	})
	c.AssetsInValid("load", err)
	c.Resource(data)
}
