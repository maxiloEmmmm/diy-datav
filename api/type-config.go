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
	Year  int  `json:"year"`
	Value int  `json:"value"`
	Num   int  `json:"num"`
	Sex   bool `json:"sex"`
	Age   int  `json:"age"`
}

func ExampleHttp(c *contact.GinHelp) {
	c.Resource([]*exampleHttpItem{
		{Year: 2000, Value: 4, Sex: true, Age: 18},
		{Year: 2001, Value: 15, Sex: true, Age: 21},
		{Year: 2002, Value: 8, Sex: false, Age: 12},
		{Year: 2003, Value: 15, Sex: false, Age: 25},
		{Year: 2004, Value: 9, Sex: false, Age: 1},
		{Year: 2005, Value: 15, Sex: true, Age: 12},
		{Year: 2006, Value: 3, Sex: false, Age: 16},
		{Year: 2007, Value: 19, Sex: true, Age: 44},
		{Year: 2008, Value: 1, Sex: true, Age: 55},
		{Year: 2009, Value: 2, Sex: true, Age: 60},
		{Year: 2010, Value: 3, Sex: false, Age: 78},
		{Year: 2011, Value: 16, Sex: false, Age: 100},
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
