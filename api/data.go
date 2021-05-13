package api

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
)

func init() {
	Apis = append(Apis, newApi(http.MethodGet, "data/:id", Data))
}

func Data(c *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	c.InValidBindUri(uri)

	data, err := service.NewDataSetService(c.AppContext).Load(uri.Id)
	c.AssetsInValid("load", err)
	c.Resource(data)
}
