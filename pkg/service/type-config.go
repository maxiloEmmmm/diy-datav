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

	Kind() ([]*types.Enum, error)
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
		typeconfig.Type(dataset.Sql),
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

func (t TypeConfigService) Kind() ([]*types.Enum, error) {
	return []*types.Enum{
		{Label: dataset.Http, Value: dataset.Http},
		{Label: dataset.Sql, Value: dataset.Sql},
		{Label: dataset.Static, Value: dataset.Static},
	}, nil
}

func NewTypeConfigService(context context.Context) TypeConfigServiceI7e {
	return &TypeConfigService{Context: context}
}
