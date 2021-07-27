package api

import (
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/view"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
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
		newApi(http.MethodPut, "view", ViewStore),
		newApi(http.MethodGet, "view/:id", ViewGet),
		newApi(http.MethodGet, "view", ViewList),
		newApi(http.MethodDelete, "view/:id", ViewDelete),
		newApi(http.MethodPost, "view/bg/upload", ViewUploadBg),
		newApi(http.MethodGet, "view/:id/bg", ViewBg),
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
		Id int `json:"id"`
	}{}
	c.InValidBindUri(uri)

	c.Resource(app.Db.View.Query().Where(view.ID(uri.Id)).FirstX(c.AppContext))
}

func ViewDelete(c *contact.GinHelp) {
	uri := &struct {
		Id int `json:"id"`
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

	view, err := service.NewViewService(c.AppContext).Store(view)
	c.AssetsInValid("store", err)

	c.Resource(view)
}

func ViewUploadBg(c *contact.GinHelp) {
	fileHeader, err := c.FormFile("bg")
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
	c.InValidBind(view)

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
