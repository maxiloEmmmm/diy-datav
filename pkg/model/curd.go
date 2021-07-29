// Code generated by entc, DO NOT EDIT.

package model

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/assets"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/dataset"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/menu"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/user"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
	go_tool "github.com/maxiloEmmmm/go-tool"
	contact "github.com/maxiloEmmmm/go-web/contact"
)

func uuidId(id string) uuid.UUID {
	u, _ := uuid.Parse(id)
	return u
}

func newApi(client *Client, opt *ApiOption) *Api {
	obj := &Api{Client: client}

	if opt != nil {
		if len(opt.Fields) > 0 {
			af := ActionFields{}
			af.SetFields(opt.Fields["Create"])
			obj.Fields.Create = af

			af = ActionFields{}
			af.SetFields(opt.Fields["Update"])
			obj.Fields.Update = af
		}
	}
	return obj
}

type ApiOption struct {
	Fields map[string][]string
}

type Api struct {
	Fields struct {
		Create ActionFields
		Update ActionFields
	}
	Client *Client
}

type ActionFields struct {
	Has    bool
	Fields map[string]bool
}

func (af *ActionFields) SetFields(fields []string) {
	af.Has = true
	af.Fields = make(map[string]bool, len(fields))
	for _, field := range fields {
		af.Fields[field] = true
	}
}

type Endpoint interface {
	List(*contact.GinHelp)
	Update(*contact.GinHelp)
	Create(*contact.GinHelp)
	Get(*contact.GinHelp)
	Delete(*contact.GinHelp)
}

type CurdBuilder struct {
	Apis struct {
		Assets     *AssetsApi
		DataSet    *DataSetApi
		Menu       *MenuApi
		TypeConfig *TypeConfigApi
		User       *UserApi
		View       *ViewApi
		ViewBlock  *ViewBlockApi
	}
}

func NewCurdBuilder(client *Client) *CurdBuilder {
	cb := &CurdBuilder{}
	cb.Apis.Assets = NewAssetsApi(client, nil)
	cb.Apis.DataSet = NewDataSetApi(client, nil)
	cb.Apis.Menu = NewMenuApi(client, nil)
	cb.Apis.TypeConfig = NewTypeConfigApi(client, nil)
	cb.Apis.User = NewUserApi(client, nil)
	cb.Apis.View = NewViewApi(client, nil)
	cb.Apis.ViewBlock = NewViewBlockApi(client, nil)
	return cb
}

func (cb *CurdBuilder) Route(prefix string, r gin.IRouter, pick []string) *gin.RouterGroup {
	if !strings.HasPrefix("/", prefix) {
		prefix = go_tool.StringJoin("/", prefix)
	}
	g := r.Group(prefix)

	hasPick := pick != nil && len(pick) > 0
	if !hasPick || go_tool.InArray(pick, TypeAssets) {
		cb.Group(g, "assets", cb.Apis.Assets)
	}
	if !hasPick || go_tool.InArray(pick, TypeDataSet) {
		cb.Group(g, "dataset", cb.Apis.DataSet)
	}
	if !hasPick || go_tool.InArray(pick, TypeMenu) {
		cb.Group(g, "menu", cb.Apis.Menu)
	}
	if !hasPick || go_tool.InArray(pick, TypeTypeConfig) {
		cb.Group(g, "typeconfig", cb.Apis.TypeConfig)
	}
	if !hasPick || go_tool.InArray(pick, TypeUser) {
		cb.Group(g, "user", cb.Apis.User)
	}
	if !hasPick || go_tool.InArray(pick, TypeView) {
		cb.Group(g, "view", cb.Apis.View)
	}
	if !hasPick || go_tool.InArray(pick, TypeViewBlock) {
		cb.Group(g, "viewblock", cb.Apis.ViewBlock)
	}

	return g
}

func (cb *CurdBuilder) Group(group *gin.RouterGroup, path string, api Endpoint) {
	apiGroup := group.Group(path)
	apiGroup.GET("", contact.GinHelpHandle(api.List))
	apiGroup.GET("/:id", contact.GinHelpHandle(api.Get))
	apiGroup.POST("", contact.GinHelpHandle(api.Create))
	apiGroup.PATCH("/:id", contact.GinHelpHandle(api.Update))
	apiGroup.DELETE("/:id", contact.GinHelpHandle(api.Delete))
}

type AssetsApi struct {
	*Api
	Filter             AssetsApiFilter
	SkipCreateAutoEdge bool
	SkipUpdateAutoEdge bool
}

type AssetsApiFilter struct {
	CreatePipe   func(help *contact.GinHelp, createPipe *AssetsCreate, edges AssetsEdges)
	CreateAfter  func(help *contact.GinHelp, item *Assets, edges AssetsEdges)
	UpdatePipe   func(help *contact.GinHelp, old *Assets, updatePipe *AssetsUpdateOne, edges AssetsEdges)
	UpdateAfter  func(help *contact.GinHelp, old *Assets, item *Assets, edges AssetsEdges)
	ListPipe     func(help *contact.GinHelp, listPipe *AssetsQuery)
	ListData     func(help *contact.GinHelp, items []*Assets) interface{}
	DeleteBefore func(help *contact.GinHelp, item *Assets)
	GetPipe      func(help *contact.GinHelp, getPipe *AssetsQuery)
	GetData      func(help *contact.GinHelp, item *Assets) *Assets
}

func NewAssetsApi(client *Client, opt *ApiOption) *AssetsApi {
	return &AssetsApi{Api: newApi(client, opt)}
}

func (c *AssetsApi) List(help *contact.GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := c.Client.Assets.Query()
		if c.Filter.ListPipe != nil {
			c.Filter.ListPipe(help, pipe)
		}
		clonePipe := pipe.Clone()

		pipe = pipe.Offset(start).Limit(size)
		items := pipe.AllX(help.AppContext)

		var data interface{} = items
		if c.Filter.ListData != nil {
			data = c.Filter.ListData(help, items)
		}
		return data, clonePipe.CountX(help.AppContext)
	})
}

func (c *AssetsApi) Delete(help *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	item := c.Client.Assets.GetX(help.AppContext, uri.Id)
	if c.Filter.DeleteBefore != nil {
		c.Filter.DeleteBefore(help, item)
	}
	c.Client.Assets.DeleteOne(item).ExecX(help.AppContext)
	help.ResourceDelete()
}

func (c *AssetsApi) Create(help *contact.GinHelp) {
	body := &Assets{}
	help.InValidBind(body)

	pipe := c.Client.Assets.Create()
	if !c.Fields.Create.Has || c.Fields.Create.Fields[assets.FieldPath] {
		pipe.SetPath(body.Path)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[assets.FieldExt] {
		pipe.SetExt(body.Ext)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[assets.FieldType] {
		pipe.SetType(body.Type)
	}
	if !c.SkipCreateAutoEdge {
		if body.Edges.View != nil {
			pipe.AddView(body.Edges.View...)
		}
	}

	if c.Filter.CreatePipe != nil {
		c.Filter.CreatePipe(help, pipe, body.Edges)
	}

	item := pipe.SaveX(help.AppContext)

	if c.Filter.CreateAfter != nil {
		c.Filter.CreateAfter(help, item, body.Edges)
	}

	help.Resource(item)
}

func (c *AssetsApi) Update(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &Assets{}
	help.InValidBind(body)

	item := c.Client.Assets.GetX(help.AppContext, uri.Id)
	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := item.Update()
		if !c.Fields.Update.Has || c.Fields.Update.Fields[assets.FieldPath] {
			pipe.SetPath(body.Path)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[assets.FieldExt] {
			pipe.SetExt(body.Ext)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[assets.FieldType] {
			pipe.SetType(body.Type)
		}
		if !c.SkipUpdateAutoEdge {
			if body.Edges.View != nil {
				pipe.AddView(body.Edges.View...)
			}
		}

		if c.Filter.UpdatePipe != nil {
			c.Filter.UpdatePipe(help, item, pipe, body.Edges)
		}
		currentItem := pipe.SaveX(help.AppContext)
		if c.Filter.UpdateAfter != nil {
			c.Filter.UpdateAfter(help, item, currentItem, body.Edges)
		}

		item = currentItem
	}
	help.Resource(item)
}

func (c *AssetsApi) Get(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	pipe := c.Client.Assets.Query().Where(assets.ID(uri.Id))
	if c.Filter.GetPipe != nil {
		c.Filter.GetPipe(help, pipe)
	}

	item := pipe.FirstX(help.AppContext)
	if c.Filter.GetData != nil {
		item = c.Filter.GetData(help, item)
	}

	help.Resource(item)
}

type DataSetApi struct {
	*Api
	Filter             DataSetApiFilter
	SkipCreateAutoEdge bool
	SkipUpdateAutoEdge bool
}

type DataSetApiFilter struct {
	CreatePipe   func(help *contact.GinHelp, createPipe *DataSetCreate, edges DataSetEdges)
	CreateAfter  func(help *contact.GinHelp, item *DataSet, edges DataSetEdges)
	UpdatePipe   func(help *contact.GinHelp, old *DataSet, updatePipe *DataSetUpdateOne, edges DataSetEdges)
	UpdateAfter  func(help *contact.GinHelp, old *DataSet, item *DataSet, edges DataSetEdges)
	ListPipe     func(help *contact.GinHelp, listPipe *DataSetQuery)
	ListData     func(help *contact.GinHelp, items []*DataSet) interface{}
	DeleteBefore func(help *contact.GinHelp, item *DataSet)
	GetPipe      func(help *contact.GinHelp, getPipe *DataSetQuery)
	GetData      func(help *contact.GinHelp, item *DataSet) *DataSet
}

func NewDataSetApi(client *Client, opt *ApiOption) *DataSetApi {
	return &DataSetApi{Api: newApi(client, opt)}
}

func (c *DataSetApi) List(help *contact.GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := c.Client.DataSet.Query()
		if c.Filter.ListPipe != nil {
			c.Filter.ListPipe(help, pipe)
		}
		clonePipe := pipe.Clone()

		pipe = pipe.Offset(start).Limit(size)
		items := pipe.AllX(help.AppContext)

		var data interface{} = items
		if c.Filter.ListData != nil {
			data = c.Filter.ListData(help, items)
		}
		return data, clonePipe.CountX(help.AppContext)
	})
}

func (c *DataSetApi) Delete(help *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	item := c.Client.DataSet.GetX(help.AppContext, uri.Id)
	if c.Filter.DeleteBefore != nil {
		c.Filter.DeleteBefore(help, item)
	}
	c.Client.DataSet.DeleteOne(item).ExecX(help.AppContext)
	help.ResourceDelete()
}

func (c *DataSetApi) Create(help *contact.GinHelp) {
	body := &DataSet{}
	help.InValidBind(body)

	pipe := c.Client.DataSet.Create()
	if !c.Fields.Create.Has || c.Fields.Create.Fields[dataset.FieldType] {
		pipe.SetType(body.Type)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[dataset.FieldTitle] {
		pipe.SetTitle(body.Title)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[dataset.FieldConfig] {
		pipe.SetConfig(body.Config)
	}
	if !c.SkipCreateAutoEdge {
		if body.Edges.Block != nil {
			pipe.SetBlock(body.Edges.Block)
		}
	}

	if c.Filter.CreatePipe != nil {
		c.Filter.CreatePipe(help, pipe, body.Edges)
	}

	item := pipe.SaveX(help.AppContext)

	if c.Filter.CreateAfter != nil {
		c.Filter.CreateAfter(help, item, body.Edges)
	}

	help.Resource(item)
}

func (c *DataSetApi) Update(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &DataSet{}
	help.InValidBind(body)

	item := c.Client.DataSet.GetX(help.AppContext, uri.Id)
	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := item.Update()
		if !c.Fields.Update.Has || c.Fields.Update.Fields[dataset.FieldType] {
			pipe.SetType(body.Type)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[dataset.FieldTitle] {
			pipe.SetTitle(body.Title)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[dataset.FieldConfig] {
			pipe.SetConfig(body.Config)
		}
		if !c.SkipUpdateAutoEdge {
			if body.Edges.Block != nil {
				pipe.SetBlock(body.Edges.Block)
			}
		}

		if c.Filter.UpdatePipe != nil {
			c.Filter.UpdatePipe(help, item, pipe, body.Edges)
		}
		currentItem := pipe.SaveX(help.AppContext)
		if c.Filter.UpdateAfter != nil {
			c.Filter.UpdateAfter(help, item, currentItem, body.Edges)
		}

		item = currentItem
	}
	help.Resource(item)
}

func (c *DataSetApi) Get(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	pipe := c.Client.DataSet.Query().Where(dataset.ID(uri.Id))
	if c.Filter.GetPipe != nil {
		c.Filter.GetPipe(help, pipe)
	}

	item := pipe.FirstX(help.AppContext)
	if c.Filter.GetData != nil {
		item = c.Filter.GetData(help, item)
	}

	help.Resource(item)
}

type MenuApi struct {
	*Api
	Filter             MenuApiFilter
	SkipCreateAutoEdge bool
	SkipUpdateAutoEdge bool
}

type MenuApiFilter struct {
	CreatePipe   func(help *contact.GinHelp, createPipe *MenuCreate, edges MenuEdges)
	CreateAfter  func(help *contact.GinHelp, item *Menu, edges MenuEdges)
	UpdatePipe   func(help *contact.GinHelp, old *Menu, updatePipe *MenuUpdateOne, edges MenuEdges)
	UpdateAfter  func(help *contact.GinHelp, old *Menu, item *Menu, edges MenuEdges)
	ListPipe     func(help *contact.GinHelp, listPipe *MenuQuery)
	ListData     func(help *contact.GinHelp, items []*Menu) interface{}
	DeleteBefore func(help *contact.GinHelp, item *Menu)
	GetPipe      func(help *contact.GinHelp, getPipe *MenuQuery)
	GetData      func(help *contact.GinHelp, item *Menu) *Menu
}

func NewMenuApi(client *Client, opt *ApiOption) *MenuApi {
	return &MenuApi{Api: newApi(client, opt)}
}

func (c *MenuApi) List(help *contact.GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := c.Client.Menu.Query()
		if c.Filter.ListPipe != nil {
			c.Filter.ListPipe(help, pipe)
		}
		clonePipe := pipe.Clone()

		pipe = pipe.Offset(start).Limit(size)
		items := pipe.AllX(help.AppContext)

		var data interface{} = items
		if c.Filter.ListData != nil {
			data = c.Filter.ListData(help, items)
		}
		return data, clonePipe.CountX(help.AppContext)
	})
}

func (c *MenuApi) Delete(help *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	item := c.Client.Menu.GetX(help.AppContext, uri.Id)
	if c.Filter.DeleteBefore != nil {
		c.Filter.DeleteBefore(help, item)
	}
	c.Client.Menu.DeleteOne(item).ExecX(help.AppContext)
	help.ResourceDelete()
}

func (c *MenuApi) Create(help *contact.GinHelp) {
	body := &Menu{}
	help.InValidBind(body)

	pipe := c.Client.Menu.Create()
	if !c.Fields.Create.Has || c.Fields.Create.Fields[menu.FieldTitle] {
		pipe.SetTitle(body.Title)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[menu.FieldURL] {
		pipe.SetURL(body.URL)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[menu.FieldCreatedAt] {
		pipe.SetCreatedAt(body.CreatedAt)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[menu.FieldUpdatedAt] {
		pipe.SetUpdatedAt(body.UpdatedAt)
	}
	if !c.SkipCreateAutoEdge {
		if body.Edges.Parent != nil {
			pipe.SetParent(body.Edges.Parent)
		}
		if body.Edges.Children != nil {
			pipe.AddChildren(body.Edges.Children...)
		}
	}

	if c.Filter.CreatePipe != nil {
		c.Filter.CreatePipe(help, pipe, body.Edges)
	}

	item := pipe.SaveX(help.AppContext)

	if c.Filter.CreateAfter != nil {
		c.Filter.CreateAfter(help, item, body.Edges)
	}

	help.Resource(item)
}

func (c *MenuApi) Update(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &Menu{}
	help.InValidBind(body)

	item := c.Client.Menu.GetX(help.AppContext, uri.Id)
	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := item.Update()
		if !c.Fields.Update.Has || c.Fields.Update.Fields[menu.FieldTitle] {
			pipe.SetTitle(body.Title)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[menu.FieldURL] {
			pipe.SetURL(body.URL)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[menu.FieldCreatedAt] {
			pipe.SetCreatedAt(body.CreatedAt)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[menu.FieldUpdatedAt] {
			pipe.SetUpdatedAt(body.UpdatedAt)
		}
		if !c.SkipUpdateAutoEdge {
			if body.Edges.Parent != nil {
				pipe.SetParent(body.Edges.Parent)
			}
			if body.Edges.Children != nil {
				pipe.AddChildren(body.Edges.Children...)
			}
		}

		if c.Filter.UpdatePipe != nil {
			c.Filter.UpdatePipe(help, item, pipe, body.Edges)
		}
		currentItem := pipe.SaveX(help.AppContext)
		if c.Filter.UpdateAfter != nil {
			c.Filter.UpdateAfter(help, item, currentItem, body.Edges)
		}

		item = currentItem
	}
	help.Resource(item)
}

func (c *MenuApi) Get(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	pipe := c.Client.Menu.Query().Where(menu.ID(uri.Id))
	if c.Filter.GetPipe != nil {
		c.Filter.GetPipe(help, pipe)
	}

	item := pipe.FirstX(help.AppContext)
	if c.Filter.GetData != nil {
		item = c.Filter.GetData(help, item)
	}

	help.Resource(item)
}

type TypeConfigApi struct {
	*Api
	Filter             TypeConfigApiFilter
	SkipCreateAutoEdge bool
	SkipUpdateAutoEdge bool
}

type TypeConfigApiFilter struct {
	CreatePipe   func(help *contact.GinHelp, createPipe *TypeConfigCreate)
	CreateAfter  func(help *contact.GinHelp, item *TypeConfig)
	UpdatePipe   func(help *contact.GinHelp, old *TypeConfig, updatePipe *TypeConfigUpdateOne)
	UpdateAfter  func(help *contact.GinHelp, old *TypeConfig, item *TypeConfig)
	ListPipe     func(help *contact.GinHelp, listPipe *TypeConfigQuery)
	ListData     func(help *contact.GinHelp, items []*TypeConfig) interface{}
	DeleteBefore func(help *contact.GinHelp, item *TypeConfig)
	GetPipe      func(help *contact.GinHelp, getPipe *TypeConfigQuery)
	GetData      func(help *contact.GinHelp, item *TypeConfig) *TypeConfig
}

func NewTypeConfigApi(client *Client, opt *ApiOption) *TypeConfigApi {
	return &TypeConfigApi{Api: newApi(client, opt)}
}

func (c *TypeConfigApi) List(help *contact.GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := c.Client.TypeConfig.Query()
		if c.Filter.ListPipe != nil {
			c.Filter.ListPipe(help, pipe)
		}
		clonePipe := pipe.Clone()

		pipe = pipe.Offset(start).Limit(size)
		items := pipe.AllX(help.AppContext)

		var data interface{} = items
		if c.Filter.ListData != nil {
			data = c.Filter.ListData(help, items)
		}
		return data, clonePipe.CountX(help.AppContext)
	})
}

func (c *TypeConfigApi) Delete(help *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	item := c.Client.TypeConfig.GetX(help.AppContext, uri.Id)
	if c.Filter.DeleteBefore != nil {
		c.Filter.DeleteBefore(help, item)
	}
	c.Client.TypeConfig.DeleteOne(item).ExecX(help.AppContext)
	help.ResourceDelete()
}

func (c *TypeConfigApi) Create(help *contact.GinHelp) {
	body := &TypeConfig{}
	help.InValidBind(body)

	pipe := c.Client.TypeConfig.Create()
	if !c.Fields.Create.Has || c.Fields.Create.Fields[typeconfig.FieldType] {
		pipe.SetType(body.Type)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[typeconfig.FieldTitle] {
		pipe.SetTitle(body.Title)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[typeconfig.FieldConfig] {
		pipe.SetConfig(body.Config)
	}

	if c.Filter.CreatePipe != nil {
		c.Filter.CreatePipe(help, pipe)
	}

	item := pipe.SaveX(help.AppContext)

	if c.Filter.CreateAfter != nil {
		c.Filter.CreateAfter(help, item)
	}

	help.Resource(item)
}

func (c *TypeConfigApi) Update(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &TypeConfig{}
	help.InValidBind(body)

	item := c.Client.TypeConfig.GetX(help.AppContext, uri.Id)
	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := item.Update()
		if !c.Fields.Update.Has || c.Fields.Update.Fields[typeconfig.FieldType] {
			pipe.SetType(body.Type)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[typeconfig.FieldTitle] {
			pipe.SetTitle(body.Title)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[typeconfig.FieldConfig] {
			pipe.SetConfig(body.Config)
		}

		if c.Filter.UpdatePipe != nil {
			c.Filter.UpdatePipe(help, item, pipe)
		}
		currentItem := pipe.SaveX(help.AppContext)
		if c.Filter.UpdateAfter != nil {
			c.Filter.UpdateAfter(help, item, currentItem)
		}

		item = currentItem
	}
	help.Resource(item)
}

func (c *TypeConfigApi) Get(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	pipe := c.Client.TypeConfig.Query().Where(typeconfig.ID(uri.Id))
	if c.Filter.GetPipe != nil {
		c.Filter.GetPipe(help, pipe)
	}

	item := pipe.FirstX(help.AppContext)
	if c.Filter.GetData != nil {
		item = c.Filter.GetData(help, item)
	}

	help.Resource(item)
}

type UserApi struct {
	*Api
	Filter             UserApiFilter
	SkipCreateAutoEdge bool
	SkipUpdateAutoEdge bool
}

type UserApiFilter struct {
	CreatePipe   func(help *contact.GinHelp, createPipe *UserCreate)
	CreateAfter  func(help *contact.GinHelp, item *User)
	UpdatePipe   func(help *contact.GinHelp, old *User, updatePipe *UserUpdateOne)
	UpdateAfter  func(help *contact.GinHelp, old *User, item *User)
	ListPipe     func(help *contact.GinHelp, listPipe *UserQuery)
	ListData     func(help *contact.GinHelp, items []*User) interface{}
	DeleteBefore func(help *contact.GinHelp, item *User)
	GetPipe      func(help *contact.GinHelp, getPipe *UserQuery)
	GetData      func(help *contact.GinHelp, item *User) *User
}

func NewUserApi(client *Client, opt *ApiOption) *UserApi {
	return &UserApi{Api: newApi(client, opt)}
}

func (c *UserApi) List(help *contact.GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := c.Client.User.Query()
		if c.Filter.ListPipe != nil {
			c.Filter.ListPipe(help, pipe)
		}
		clonePipe := pipe.Clone()

		pipe = pipe.Offset(start).Limit(size)
		items := pipe.AllX(help.AppContext)

		var data interface{} = items
		if c.Filter.ListData != nil {
			data = c.Filter.ListData(help, items)
		}
		return data, clonePipe.CountX(help.AppContext)
	})
}

func (c *UserApi) Delete(help *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	item := c.Client.User.GetX(help.AppContext, uri.Id)
	if c.Filter.DeleteBefore != nil {
		c.Filter.DeleteBefore(help, item)
	}
	c.Client.User.DeleteOne(item).ExecX(help.AppContext)
	help.ResourceDelete()
}

func (c *UserApi) Create(help *contact.GinHelp) {
	body := &User{}
	help.InValidBind(body)

	pipe := c.Client.User.Create()
	if !c.Fields.Create.Has || c.Fields.Create.Fields[user.FieldUsername] {
		pipe.SetUsername(body.Username)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[user.FieldPassword] {
		pipe.SetPassword(body.Password)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[user.FieldEnable] {
		pipe.SetEnable(body.Enable)
	}

	if c.Filter.CreatePipe != nil {
		c.Filter.CreatePipe(help, pipe)
	}

	item := pipe.SaveX(help.AppContext)

	if c.Filter.CreateAfter != nil {
		c.Filter.CreateAfter(help, item)
	}

	help.Resource(item)
}

func (c *UserApi) Update(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &User{}
	help.InValidBind(body)

	item := c.Client.User.GetX(help.AppContext, uri.Id)
	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := item.Update()
		if !c.Fields.Update.Has || c.Fields.Update.Fields[user.FieldUsername] {
			pipe.SetUsername(body.Username)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[user.FieldPassword] {
			pipe.SetPassword(body.Password)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[user.FieldEnable] {
			pipe.SetEnable(body.Enable)
		}

		if c.Filter.UpdatePipe != nil {
			c.Filter.UpdatePipe(help, item, pipe)
		}
		currentItem := pipe.SaveX(help.AppContext)
		if c.Filter.UpdateAfter != nil {
			c.Filter.UpdateAfter(help, item, currentItem)
		}

		item = currentItem
	}
	help.Resource(item)
}

func (c *UserApi) Get(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	pipe := c.Client.User.Query().Where(user.ID(uri.Id))
	if c.Filter.GetPipe != nil {
		c.Filter.GetPipe(help, pipe)
	}

	item := pipe.FirstX(help.AppContext)
	if c.Filter.GetData != nil {
		item = c.Filter.GetData(help, item)
	}

	help.Resource(item)
}

type ViewApi struct {
	*Api
	Filter             ViewApiFilter
	SkipCreateAutoEdge bool
	SkipUpdateAutoEdge bool
}

type ViewApiFilter struct {
	CreatePipe   func(help *contact.GinHelp, createPipe *ViewCreate, edges ViewEdges)
	CreateAfter  func(help *contact.GinHelp, item *View, edges ViewEdges)
	UpdatePipe   func(help *contact.GinHelp, old *View, updatePipe *ViewUpdateOne, edges ViewEdges)
	UpdateAfter  func(help *contact.GinHelp, old *View, item *View, edges ViewEdges)
	ListPipe     func(help *contact.GinHelp, listPipe *ViewQuery)
	ListData     func(help *contact.GinHelp, items []*View) interface{}
	DeleteBefore func(help *contact.GinHelp, item *View)
	GetPipe      func(help *contact.GinHelp, getPipe *ViewQuery)
	GetData      func(help *contact.GinHelp, item *View) *View
}

func NewViewApi(client *Client, opt *ApiOption) *ViewApi {
	return &ViewApi{Api: newApi(client, opt)}
}

func (c *ViewApi) List(help *contact.GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := c.Client.View.Query()
		if c.Filter.ListPipe != nil {
			c.Filter.ListPipe(help, pipe)
		}
		clonePipe := pipe.Clone()

		pipe = pipe.Offset(start).Limit(size)
		items := pipe.AllX(help.AppContext)

		var data interface{} = items
		if c.Filter.ListData != nil {
			data = c.Filter.ListData(help, items)
		}
		return data, clonePipe.CountX(help.AppContext)
	})
}

func (c *ViewApi) Delete(help *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	item := c.Client.View.GetX(help.AppContext, uri.Id)
	if c.Filter.DeleteBefore != nil {
		c.Filter.DeleteBefore(help, item)
	}
	c.Client.View.DeleteOne(item).ExecX(help.AppContext)
	help.ResourceDelete()
}

func (c *ViewApi) Create(help *contact.GinHelp) {
	body := &View{}
	help.InValidBind(body)

	pipe := c.Client.View.Create()
	if !c.Fields.Create.Has || c.Fields.Create.Fields[view.FieldDesc] {
		pipe.SetDesc(body.Desc)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[view.FieldConfig] {
		pipe.SetConfig(body.Config)
	}
	if !c.SkipCreateAutoEdge {
		if body.Edges.Bg != nil {
			pipe.SetBg(body.Edges.Bg)
		}
		if body.Edges.Blocks != nil {
			pipe.AddBlocks(body.Edges.Blocks...)
		}
	}

	if c.Filter.CreatePipe != nil {
		c.Filter.CreatePipe(help, pipe, body.Edges)
	}

	item := pipe.SaveX(help.AppContext)

	if c.Filter.CreateAfter != nil {
		c.Filter.CreateAfter(help, item, body.Edges)
	}

	help.Resource(item)
}

func (c *ViewApi) Update(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &View{}
	help.InValidBind(body)

	item := c.Client.View.GetX(help.AppContext, uri.Id)
	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := item.Update()
		if !c.Fields.Update.Has || c.Fields.Update.Fields[view.FieldDesc] {
			pipe.SetDesc(body.Desc)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[view.FieldConfig] {
			pipe.SetConfig(body.Config)
		}
		if !c.SkipUpdateAutoEdge {
			if body.Edges.Bg != nil {
				pipe.SetBg(body.Edges.Bg)
			}
			if body.Edges.Blocks != nil {
				pipe.AddBlocks(body.Edges.Blocks...)
			}
		}

		if c.Filter.UpdatePipe != nil {
			c.Filter.UpdatePipe(help, item, pipe, body.Edges)
		}
		currentItem := pipe.SaveX(help.AppContext)
		if c.Filter.UpdateAfter != nil {
			c.Filter.UpdateAfter(help, item, currentItem, body.Edges)
		}

		item = currentItem
	}
	help.Resource(item)
}

func (c *ViewApi) Get(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	pipe := c.Client.View.Query().Where(view.ID(uri.Id))
	if c.Filter.GetPipe != nil {
		c.Filter.GetPipe(help, pipe)
	}

	item := pipe.FirstX(help.AppContext)
	if c.Filter.GetData != nil {
		item = c.Filter.GetData(help, item)
	}

	help.Resource(item)
}

type ViewBlockApi struct {
	*Api
	Filter             ViewBlockApiFilter
	SkipCreateAutoEdge bool
	SkipUpdateAutoEdge bool
}

type ViewBlockApiFilter struct {
	CreatePipe   func(help *contact.GinHelp, createPipe *ViewBlockCreate, edges ViewBlockEdges)
	CreateAfter  func(help *contact.GinHelp, item *ViewBlock, edges ViewBlockEdges)
	UpdatePipe   func(help *contact.GinHelp, old *ViewBlock, updatePipe *ViewBlockUpdateOne, edges ViewBlockEdges)
	UpdateAfter  func(help *contact.GinHelp, old *ViewBlock, item *ViewBlock, edges ViewBlockEdges)
	ListPipe     func(help *contact.GinHelp, listPipe *ViewBlockQuery)
	ListData     func(help *contact.GinHelp, items []*ViewBlock) interface{}
	DeleteBefore func(help *contact.GinHelp, item *ViewBlock)
	GetPipe      func(help *contact.GinHelp, getPipe *ViewBlockQuery)
	GetData      func(help *contact.GinHelp, item *ViewBlock) *ViewBlock
}

func NewViewBlockApi(client *Client, opt *ApiOption) *ViewBlockApi {
	return &ViewBlockApi{Api: newApi(client, opt)}
}

func (c *ViewBlockApi) List(help *contact.GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := c.Client.ViewBlock.Query()
		if c.Filter.ListPipe != nil {
			c.Filter.ListPipe(help, pipe)
		}
		clonePipe := pipe.Clone()

		pipe = pipe.Offset(start).Limit(size)
		items := pipe.AllX(help.AppContext)

		var data interface{} = items
		if c.Filter.ListData != nil {
			data = c.Filter.ListData(help, items)
		}
		return data, clonePipe.CountX(help.AppContext)
	})
}

func (c *ViewBlockApi) Delete(help *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	item := c.Client.ViewBlock.GetX(help.AppContext, uri.Id)
	if c.Filter.DeleteBefore != nil {
		c.Filter.DeleteBefore(help, item)
	}
	c.Client.ViewBlock.DeleteOne(item).ExecX(help.AppContext)
	help.ResourceDelete()
}

func (c *ViewBlockApi) Create(help *contact.GinHelp) {
	body := &ViewBlock{}
	help.InValidBind(body)

	pipe := c.Client.ViewBlock.Create()
	if !c.Fields.Create.Has || c.Fields.Create.Fields[viewblock.FieldType] {
		pipe.SetType(body.Type)
	}
	if !c.Fields.Create.Has || c.Fields.Create.Fields[viewblock.FieldConfig] {
		pipe.SetConfig(body.Config)
	}
	if !c.SkipCreateAutoEdge {
		if body.Edges.View != nil {
			pipe.SetView(body.Edges.View)
		}
		if body.Edges.Dataset != nil {
			pipe.AddDataset(body.Edges.Dataset...)
		}
	}

	if c.Filter.CreatePipe != nil {
		c.Filter.CreatePipe(help, pipe, body.Edges)
	}

	item := pipe.SaveX(help.AppContext)

	if c.Filter.CreateAfter != nil {
		c.Filter.CreateAfter(help, item, body.Edges)
	}

	help.Resource(item)
}

func (c *ViewBlockApi) Update(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &ViewBlock{}
	help.InValidBind(body)

	item := c.Client.ViewBlock.GetX(help.AppContext, uri.Id)
	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := item.Update()
		if !c.Fields.Update.Has || c.Fields.Update.Fields[viewblock.FieldType] {
			pipe.SetType(body.Type)
		}
		if !c.Fields.Update.Has || c.Fields.Update.Fields[viewblock.FieldConfig] {
			pipe.SetConfig(body.Config)
		}
		if !c.SkipUpdateAutoEdge {
			if body.Edges.View != nil {
				pipe.SetView(body.Edges.View)
			}
			if body.Edges.Dataset != nil {
				pipe.AddDataset(body.Edges.Dataset...)
			}
		}

		if c.Filter.UpdatePipe != nil {
			c.Filter.UpdatePipe(help, item, pipe, body.Edges)
		}
		currentItem := pipe.SaveX(help.AppContext)
		if c.Filter.UpdateAfter != nil {
			c.Filter.UpdateAfter(help, item, currentItem, body.Edges)
		}

		item = currentItem
	}
	help.Resource(item)
}

func (c *ViewBlockApi) Get(help *contact.GinHelp) {
	uri := struct {
		Id int `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	pipe := c.Client.ViewBlock.Query().Where(viewblock.ID(uri.Id))
	if c.Filter.GetPipe != nil {
		c.Filter.GetPipe(help, pipe)
	}

	item := pipe.FirstX(help.AppContext)
	if c.Filter.GetData != nil {
		item = c.Filter.GetData(help, item)
	}

	help.Resource(item)
}
