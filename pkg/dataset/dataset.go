package dataset

import (
	"context"
	"errors"
)

type I7e interface {
	Load(ctx context.Context, config string) (interface{}, error)
}

type EmptyData []struct{}

type DataDesc map[string]interface{}

var Reg = make(map[string]I7e, 0)

func Load(ctx context.Context, typ string, config string) (interface{}, error) {
	if instance, exist := Reg[typ]; exist {
		return instance.Load(ctx, config)
	}
	return nil, errors.New("dataset type not found")
}



