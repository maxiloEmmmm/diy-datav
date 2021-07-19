package service

import (
	"context"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/dataset"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
)

type TypeConfigServiceI7e interface {
	//Create()
	//Update()
	//List()
	//Get()
	//Delete()

	HttpKind() ([]*types.TypeConfig, error)
	StaticKind() ([]*types.TypeConfig, error)
	MysqlKind() ([]*types.TypeConfig, error)
}

type TypeConfigService struct {
	context.Context
}

func (t TypeConfigService) HttpKind() ([]*types.TypeConfig, error) {
	tc, err := app.Db.TypeConfig.Query().Where(
		typeconfig.Type(dataset.Http),
	).All(t.Context)
	if err != nil {
		return nil, err
	}

	data := make([]*types.TypeConfig, 0, len(tc))
	for _, t := range tc {
		data = append(data, types.NewTypeConfig(t))
	}

	return data, nil
}

func (t TypeConfigService) StaticKind() ([]*types.TypeConfig, error) {
	tc, err := app.Db.TypeConfig.Query().Where(
		typeconfig.Type(dataset.Static),
	).All(t.Context)
	if err != nil {
		return nil, err
	}

	data := make([]*types.TypeConfig, 0, len(tc))
	for _, t := range tc {
		data = append(data, types.NewTypeConfig(t))
	}

	return data, nil
}

func (t TypeConfigService) MysqlKind() ([]*types.TypeConfig, error) {
	tc, err := app.Db.TypeConfig.Query().Where(
		typeconfig.Type(dataset.Mysql),
	).All(t.Context)
	if err != nil {
		return nil, err
	}

	data := make([]*types.TypeConfig, 0, len(tc))
	for _, t := range tc {
		data = append(data, types.NewTypeConfig(t))
	}

	return data, nil
}

func NewTypeConfigService(context context.Context) TypeConfigServiceI7e {
	return &TypeConfigService{Context: context}
}
