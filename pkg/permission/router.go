package permission

import (
	"fmt"
)

const (
	RouterPR = "pr-router"
)

type PRRouter struct {
	Path string
}

func (pr *PRRouter) Key() string {
	return fmt.Sprintf("%s-%s", RouterPR, pr.Path)
}
