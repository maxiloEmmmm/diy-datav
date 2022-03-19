package api

import (
	"encoding/json"
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
	"github.com/maxiloEmmmm/diy-datav/pkg/permission"
	"github.com/maxiloEmmmm/diy-datav/pkg/service"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
	"github.com/maxiloEmmmm/go-web/contact"
	"image"
	"image/color"
	"image/png"
	"net/http"
)

func init() {
	Apis = append(Apis,
		newAuthApi(http.MethodPut, "view", ViewStore),
		newAuthApi(http.MethodGet, "view/:id", ViewGet),
		newAuthApi(http.MethodGet, "view", ViewList),
		newAuthApi(http.MethodDelete, "view/:id", ViewDelete),
		newAuthApi(http.MethodPost, "view/bg/upload", ViewUploadBg),
		newAuthApi(http.MethodGet, "view-bg-assets", ViewBgAssets),
		newAuthApi(http.MethodGet, "view/:id/bg", ViewBg),
	)
}

func ViewList(c *contact.GinHelp) {
	c.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := app.Db.View.Query()
		if val, exist := c.GetQuery("desc"); exist && val != "" {
			pipe.Where(view.DescContains(val))
		}
		return pipe.Limit(size).Offset((start - 1) * size).AllX(c.AppContext), pipe.CountX(c.AppContext)
	})
}

func ViewGet(c *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	c.InValidBindUri(uri)

	query := &struct {
		Type string `form:"type"`
	}{}
	c.InValidBindQuery(query)

	isFull := query.Type == "full"

	v, err := app.Db.View.Query().WithBg().WithBlocks().Where(view.ID(uri.Id)).First(c.AppContext)
	c.AssetsInValid("get", err)

	if !permission.Pass(app.User(c), &permission.PRView{v}, permission.GetInfoAction) {
		c.InValid("403", "")
	}

	vd, err := NewView(v, isFull)
	c.AssetsInValid("new", err)

	c.Resource(vd)
}

func NewView(v *model.View, isFull bool) (*types.View, error) {
	vd := &types.View{}
	vd.Id = v.ID
	vd.BgAssetsID = v.Edges.Bg.ID
	vd.Desc = v.Desc
	vd.Blocks = make([]*types.ViewBlock, 0, len(v.Edges.Blocks))
	for _, block := range v.Edges.Blocks {
		c := new(types.ViewBlockConfig)
		err := json.Unmarshal([]byte(block.Config), c)
		if err != nil {
			return nil, err
		}
		b := &types.ViewBlock{
			Id:     block.ID,
			Type:   block.Type,
			Config: c,
		}

		if !isFull {
			for _, input := range c.Common.Input {
				*input = types.DataSet{
					Id: input.Id,
				}
			}
		}

		vd.Blocks = append(vd.Blocks, b)
	}

	return vd, nil
}

func ViewDelete(c *contact.GinHelp) {
	uri := &struct {
		Id int `uri:"id"`
	}{}
	c.InValidBindUri(uri)

	viewObj, err := app.Db.View.Query().Where(view.ID(uri.Id)).First(c.AppContext)
	c.AssetsInValid("find.view", err)

	_, err = app.Db.ViewBlock.Delete().Where(viewblock.HasViewWith(view.ID(viewObj.ID))).Exec(c.AppContext)
	c.AssetsInValid("remove.block", err)

	c.Resource(app.Db.View.DeleteOne(viewObj).Exec(c.AppContext))
}

func ViewStore(c *contact.GinHelp) {
	view := &types.View{}
	c.InValidBind(view)

	v, err := service.NewViewService(c.AppContext).Store(view)
	c.AssetsInValid("store", err)

	c.Resource(v)
}

func ViewBgAssets(c *contact.GinHelp) {
	c.Resource(app.Db.Assets.Query().AllX(c.AppContext))
}

func ViewUploadBg(c *contact.GinHelp) {
	fileHeader, err := c.FormFile("file")
	c.AssetsInValid("form.file", err)

	bgUpload := &types.UploadRequest{}
	c.InValidBindQuery(bgUpload)

	file, err := fileHeader.Open()
	c.AssetsInValid("form.file.open", err)
	defer file.Close()

	resource, err := service.NewViewService(c.AppContext).Upload(bgUpload.Type, file, fileHeader.Header.Get("Content-Type"))
	c.AssetsInValid("upload", err)

	c.Resource(resource)
}

func ViewBg(c *contact.GinHelp) {
	view := &types.BgRequest{}
	c.InValidBindUri(view)

	err := service.NewViewService(c.AppContext).WebDownloadBG(view.Id, c.Writer, c.Request)
	if err != nil {
		contact.Warning.Log("web.download", fmt.Sprintf("view id: %d, err: %v", view.Id, err))
		//	fallback to write black
		c.Writer.Header().Set("Content-Type", "image/png")

		alpha := image.NewAlpha(image.Rect(0, 0, 800, 764))
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				alpha.Set(x, y, color.Alpha{A: uint8(x % 256)})
			}
		}
		_ = png.Encode(c.Writer, alpha)
	}
}
