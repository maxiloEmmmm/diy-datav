package types

import "github.com/maxiloEmmmm/diy-datav/pkg/model"

type View struct {
	Id         int
	Desc       string
	Config     string
	BgAssetsID int
	Blocks     []*ViewBlock
}

type ViewBlock struct {
	Id     int
	Type   string
	Config string
}

type ViewBlockConfig struct {
	Type   interface{}
	Common ViewBlockCommonConfig
}

type ViewBlockCommonConfig struct {
	Input    []*DataSet
	Refresh  float64
	Position *CommonPosition
}

type CommonPosition struct {
	Left   string
	Top    string
	Right  string
	Bottom string
}

type DataSet struct {
	Id     int
	Type   string
	Config string
	Title  string
}

type UploadResource struct {
	Id   int
	Path string
}

type UploadRequest struct {
	Type string `form:"type"`
}

type BgRequest struct {
	Id int `uri:"id"`
}

type TypeConfig struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Config string `json:"config"`
}

func NewTypeConfig(config *model.TypeConfig) *TypeConfig {
	return &TypeConfig{
		Id:     config.ID,
		Config: config.Config,
		Title:  config.Title,
	}
}

type Enum struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
	V     interface{}
}

type TmpEcho struct {
	Type   string `json:"type"`
	Config string `json:"config"`
}
