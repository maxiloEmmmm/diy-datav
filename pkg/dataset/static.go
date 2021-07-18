package dataset

import (
	"context"
	"encoding/json"
)

const Static = "dataset.static"

func init() {
	Reg[Static] = &static{}
}

type static struct{}

type StaticConfig struct {
	Id int
}

func (h *static) Load(ctx context.Context, config string) (interface{}, error) {
	hc := &StaticConfig{}
	err := json.Unmarshal([]byte(config), hc)
	if err != nil {
		return EmptyData{}, err
	}

	// TODO: read static data from db by id
	return EmptyData{}, nil
}
