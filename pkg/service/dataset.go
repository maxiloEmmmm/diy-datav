package service

import (
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	datasetUtil "github.com/maxiloEmmmm/diy-datav/pkg/dataset"
)

type DataSetServiceI7e interface {
	Load(sdId int) (interface{}, error)
}

type DataSetService struct {
	Context
}

func NewDataSetService(context Context) *DataSetService {
	return &DataSetService{Context: context}
}

func (d *DataSetService) Load(sdId int) (interface{}, error) {
	dataset, err := app.Db.DataSet.Get(d.Context, sdId)
	if err != nil {
		return nil, err
	}

	return datasetUtil.Load(dataset.Type, dataset.Config)
}



