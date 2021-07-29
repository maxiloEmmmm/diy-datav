package permission

import (
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
)

//*model.Menu
const (
	MenuPR              = "pr-menu"
	AccessMenuActionKey = "AccessMenuActionKey"
)

var (
	AccessMenuAction = NewAction("访问", AccessMenuActionKey)
)

var MenuActions = []*Action{
	GetInfoAction,
}

type PRMenu struct {
	*model.Menu
}

func (pr *PRMenu) Key() string {
	return fmt.Sprintf("%s-%s", MenuPR, pr.URL)
}
