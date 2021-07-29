package api

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
)

func init() {
	Apis = append(Apis, newAuthApi(http.MethodGet, "assets-type-assets", AssetsType))
	Apis = append(Apis, newAuthApi(http.MethodGet, "assets-file/:id", AssetsFile))
}

func AssetsType(c *contact.GinHelp) {
	data, err := service.NewAssetsService(c.AppContext).TypeList()
	c.AssetsInValid("load", err)
	c.Resource(data)
}

func AssetsFile(c *contact.GinHelp) {
	view := &types.BgRequest{}
	c.InValidBindUri(view)
	_ = service.NewAssetsService(c.AppContext).File(view.Id, c.Writer, c.Request)
}
