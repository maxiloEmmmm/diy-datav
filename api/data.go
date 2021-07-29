package api

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
)

func init() {
	Apis = append(Apis, newAuthApi(http.MethodGet, "data/:id", Data))
	//data-tmp-echo
	Apis = append(Apis, newAuthApi(http.MethodPost, "data-tmp-echo", TmpEchoData))
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

func TmpEchoData(c *contact.GinHelp) {
	body := &types.TmpEcho{}
	c.InValidBind(body)

	data, err := service.NewDataSetService(c.AppContext).LoadTmpEcho(body)
	c.AssetsInValid("load", err)
	c.Resource(data)
}
