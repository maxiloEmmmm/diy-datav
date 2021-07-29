package api

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/go-web/contact"
	"net/http"
)

func init() {
	Apis = append(Apis, newApi(http.MethodGet, "tc/example/http", ExampleHttp))
	Apis = append(Apis, newAuthApi(http.MethodGet, "tc/kind/static", StaticKind))
	Apis = append(Apis, newAuthApi(http.MethodGet, "tc/kind/sql", MysqlKind))
	Apis = append(Apis, newAuthApi(http.MethodGet, "tc/kind/http", HttpKind))
	Apis = append(Apis, newAuthApi(http.MethodGet, "tc/kind", Kind))
}

type exampleHttpItem struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

func ExampleHttp(c *contact.GinHelp) {
	c.Resource([]*exampleHttpItem{
		{Year: 2000, Value: 5},
		{Year: 2001, Value: 30},
		{Year: 2002, Value: 10},
		{Year: 2003, Value: 15},
		{Year: 2004, Value: 9},
		{Year: 2005, Value: 31},
	})
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
