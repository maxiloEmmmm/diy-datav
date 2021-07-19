package dataset

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
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
	sc := &StaticConfig{}
	err := json.Unmarshal([]byte(config), sc)
	if err != nil {
		return EmptyData{}, err
	}

	tc, err := app.Db.TypeConfig.Get(ctx, sc.Id)
	if err != nil {
		return nil, err
	}

	if tc.Type != Static {
		return nil, errors.New("engine type not eq")
	}

	td := EmptyData{}
	err = json.Unmarshal([]byte(tc.Config), &td)
	if err != nil {
		return EmptyData{}, err
	}

	return td, nil
}
