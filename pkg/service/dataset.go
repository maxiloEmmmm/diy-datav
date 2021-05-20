package service

import (
	"context"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	datasetUtil "github.com/maxiloEmmmm/diy-datav/pkg/dataset"
)

type DataSetServiceI7e interface {
	Load(sdId int) (interface{}, error)
}

type DataSetService struct {
	context.Context
}

func NewDataSetService(context context.Context) DataSetServiceI7e {
	return &DataSetService{Context: context}
}

func (d *DataSetService) Load(sdId int) (interface{}, error) {
	dataset, err := app.Db.DataSet.Get(d.Context, sdId)
	if err != nil {
		return nil, err
	}

	return datasetUtil.Load(d.Context, dataset.Type, dataset.Config)
}



