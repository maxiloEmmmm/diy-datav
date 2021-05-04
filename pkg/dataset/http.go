package dataset

const Http = "dataset.http"

func init() {
	Reg[Http] = &http{}
}

type http struct {}

func (h *http) Load(config string) (interface{}, error) {
	panic("implement me")
}
