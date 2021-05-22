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

type ViewBlockConfig struct {
	Type string
	Common ViewBlockCommonConfig
}

type ViewBlockCommonConfig struct {
	Input []*DataSet
	Refresh float64
	Position *CommonPosition
}

type CommonPosition struct {
	Left string
	Top string
	Right string
	Bottom string
}

type DataSet struct {
	Id int
	Type string
	Config string
}

type UploadResource struct {
	Id int
	Path string
}

type UploadRequest struct {
	Type string `form:"type"`
}

type BgRequest struct {
	Id int `uri:"id"`
}
