package permission

import "fmt"

type Resource interface {
	Key() string
}

func Pass(u string, action string, r Resource) string {
	return fmt.Sprintf("%s.%s.%s", u, action, r.Key())
}
