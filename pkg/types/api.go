package types

type View struct {
	Id int
	Desc string
	Config string
	BgAssetsID int
	Blocks []*ViewBlock
}

type ViewBlock struct {
	Id int
	Type string
	Config string
}

type ViewBlockConfigInput struct {
	Input []*DataSet
}

type DataSet struct {
	Id int
	Type string
	Config string
}
