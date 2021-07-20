package api

import (
	"fmt"
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
		newApi(http.MethodPost, "view/bg/upload", ViewUploadBg),
		newApi(http.MethodGet, "view/:id/bg", ViewBg),
	)
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
