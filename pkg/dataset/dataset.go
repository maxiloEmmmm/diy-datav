package dataset

import (
	"context"
	"fmt"
)

type DataSet interface {
	Load(ctx context.Context, config string) (interface{}, error)
}

type EmptyData []*DataDesc

type DataDesc map[string]interface{}

var Reg = make(map[string]DataSet, 0)

func NewDataSet(typ string) (DataSet, error) {
	if instance, exist := Reg[typ]; exist {
		return instance, nil
	}

	return nil, fmt.Errorf("not found dataset type for %s", typ)
}

