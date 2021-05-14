package service

import (
	"context"
	"encoding/json"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/dataset"
	modelView "github.com/maxiloEmmmm/diy-datav/pkg/model/view"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/viewblock"
	"github.com/maxiloEmmmm/diy-datav/pkg/types"
)

type ViewServiceI7e interface {
	Load(sdId int) (interface{}, error)
}

type ViewService struct {
	context.Context
}

func NewViewService(context context.Context) *ViewService {
	return &ViewService{Context: context}
}

func(v *ViewService)Store(view *types.View) (*types.View, error) {
	return view, app.WithTx(v.Context, func(tx *model.Tx) error {
		var (
			err error
			viewModel *model.View
		)

		if view.Id != 0 {
			viewModel, err = tx.View.UpdateOneID(view.Id).
				SetConfig(view.Config).SetDesc(view.Desc).SetBgID(view.BgAssetsID).Save(v.Context)
		}else {
			viewModel, err = tx.View.Create().
				SetConfig(view.Config).SetDesc(view.Desc).SetBgID(view.BgAssetsID).Save(v.Context)
		}

		if err != nil {
			return err
		}

		view.Id = viewModel.ID

		oldBlocks, err := viewModel.QueryBlocks().IDs(v.Context)
		if err != nil {
			return err
		}

		// 删除旧block
		_, err = tx.ViewBlock.Delete().Where(viewblock.HasViewWith(modelView.ID(view.Id))).Exec(v.Context)
		if err != nil {
			return err
		}

		// 删除旧dataset
		_, err = tx.DataSet.Delete().Where(dataset.HasBlockWith(viewblock.IDIn(oldBlocks...))).Exec(v.Context)
		if err != nil {
			return err
		}

		for _, block := range view.Blocks {
			blockInstance, err := tx.ViewBlock.Create().SetConfig(block.Config).SetType(block.Type).SetView(viewModel).Save(v.Context)
			if err != nil {
				return err
			}

			block.Id = blockInstance.ID

			vbci := &types.ViewBlockConfigInput{}
			err = json.Unmarshal([]byte(block.Config), vbci)
			if err != nil {
				return err
			}

			for _, dsc := range vbci.Input {
				dsInstance, err := tx.DataSet.Create().SetType(dsc.Type).SetConfig(dsc.Config).SetBlock(blockInstance).Save(v.Context)
				if err != nil {
					return err
				}
				dsc.Id = dsInstance.ID
			}

			lastConfig, err := json.Marshal(vbci)
			if err != nil {
				return err
			}

			block.Config = string(lastConfig)
			if err = tx.ViewBlock.UpdateOne(blockInstance).SetConfig(block.Config).Exec(v.Context); err != nil {
				return err
			}
		}

		return nil
	})
}
