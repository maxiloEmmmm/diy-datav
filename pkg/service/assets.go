package service

import (
	"context"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/assets"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
	"github.com/maxiloEmmmm/diy-datav/pkg/upload"
	"net/http"
)

type AssetsServiceI7e interface {
	TypeList() ([]*types.Enum, error)
	File(id int, rw http.ResponseWriter, req *http.Request) error
}

type AssetsService struct {
	context.Context
}

func (t AssetsService) TypeList() ([]*types.Enum, error) {
	return []*types.Enum{
		{Label: upload.Local, Value: upload.Local},
	}, nil
}

func (t AssetsService) File(id int, rw http.ResponseWriter, req *http.Request) error {
	bg, err := app.Db.Assets.Query().Where(assets.ID(id)).First(t.Context)
	if err != nil {
		return err
	}

	uploadUtil, err := upload.NewUpload(bg.Type)
	if err != nil {
		return err
	}

	return uploadUtil.WebDownload(bg.Path, bg.Ext, rw, req)
}

func NewAssetsService(context context.Context) AssetsServiceI7e {
	return &AssetsService{Context: context}
}
