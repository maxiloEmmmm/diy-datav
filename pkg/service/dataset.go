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

	datasetEngine, err := datasetUtil.NewDataSet(dataset.Type)
	if err != nil {
		return nil, err
	}

	return datasetEngine.Load(d.Context, dataset.Config)
}



