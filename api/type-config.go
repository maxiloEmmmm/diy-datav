package api

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
)

func init() {
	Apis = append(Apis, newApi(http.MethodGet, "tc/kind/static", StaticKind))
	Apis = append(Apis, newApi(http.MethodGet, "tc/kind/sql", MysqlKind))
	Apis = append(Apis, newApi(http.MethodGet, "tc/kind/http", HttpKind))
	Apis = append(Apis, newApi(http.MethodGet, "tc/kind", Kind))
	//Apis = append(Apis, newApi(http.MethodGet, "tc", Static))
	//Apis = append(Apis, newApi(http.MethodPost, "tc", Static))
	//Apis = append(Apis, newApi(http.MethodGet, "tc/:id", Static))
	//Apis = append(Apis, newApi(http.MethodPatch, "tc/:id", Static))
	//Apis = append(Apis, newApi(http.MethodDelete, "tc/:id", Static))
}

func StaticKind(c *contact.GinHelp) {
	data, err := service.NewTypeConfigService(c.AppContext).StaticKind()
	c.AssetsInValid("load", err)
	c.Resource(data)
}

func MysqlKind(c *contact.GinHelp) {
	data, err := service.NewTypeConfigService(c.AppContext).MysqlKind()
	c.AssetsInValid("load", err)
	c.Resource(data)
}

func HttpKind(c *contact.GinHelp) {
	data, err := service.NewTypeConfigService(c.AppContext).HttpKind()
	c.AssetsInValid("load", err)
	c.Resource(data)
}

func Kind(c *contact.GinHelp) {
	data, err := service.NewTypeConfigService(c.AppContext).Kind()
	c.AssetsInValid("load", err)
	c.Resource(data)
}
