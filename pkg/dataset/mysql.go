package dataset

const Mysql = "dataset.mysql"

func init() {
	Reg[Mysql] = &http{}
}

type mysql struct {}

func (h *mysql) Load(config string) (interface{}, error) {
	panic("implement me")
}
