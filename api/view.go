package api

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
)

func init() {
	Apis = append(Apis, newApi(http.MethodPut, "view/:id", ViewStore))
}

func ViewStore(c *contact.GinHelp) {
	view := &types.View{}
	c.InValidBind(view)

	view, err := service.NewViewService(c.AppContext).Store(view)
	c.AssetsInValid("store", err)

	c.Resource(view)
}


