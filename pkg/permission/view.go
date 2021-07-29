package permission

import (
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
)

const (
	ViewPR           = "pr-view"
	GetInfoActionKey = "GetInfoActionKey"
)

var (
	GetInfoAction = NewAction("获取信息", GetInfoActionKey)
)

var ViewActions = []*Action{
	GetInfoAction,
}

type PRView struct {
	*model.View
}

func (pr *PRView) Key() string {
	return fmt.Sprintf("%s-%d", ViewPR, pr.ID)
}
