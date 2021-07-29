package permission

import (
	"github.com/maxiloEmmmm/go-web/contact"
)

type Action struct {
	Title string
	Key   string
}

func NewAction(title, key string) *Action {
	return &Action{Title: title, Key: key}
}

type Resource interface {
	Key() string
}

func Pass(sub string, r Resource, action *Action) bool {
	ok, _ := contact.Permission.Enforce(sub, r.Key(), action.Key)
	return ok
}

func Has(sub string, r Resource, action *Action) bool {
	return contact.Permission.HasPolicy(sub, r.Key(), action.Key, "allow")
}
