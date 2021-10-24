package types

import "github.com/maxiloEmmmm/diy-datav/pkg/model"

type View struct {
	Id         int          `json:"id"`
	Desc       string       `json:"desc"`
	Config     string       `json:"config"`
	BgAssetsID int          `json:"bgAssetsID"`
	Blocks     []*ViewBlock `json:"blocks"`
}

type ViewBlock struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Config string `json:"config"`
}

type ViewBlockConfig struct {
	Type   interface{}           `json:"type"`
	Common ViewBlockCommonConfig `json:"common"`
}

type ViewBlockCommonConfig struct {
	Input    []*DataSet      `json:"input"`
	Refresh  float64         `json:"refresh"`
	Position *CommonPosition `json:"position"`
	ZIndex   int             `json:"zIndex"`
	Desc     interface{}     `json:"desc"`
	Border   interface{}     `json:"border"`
	Bg       string          `json:"bg"`
	Grid     string          `json:"grid"`
}

type CommonPosition struct {
	Left   float64 `json:"left"`
	Top    float64 `json:"top"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type DataSet struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Config string `json:"config"`
	Title  string `json:"title"`
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

type Permission struct {
	Add    []string `json:"add"`
	Remove []string `json:"remove"`
}

type MenuPermission struct {
	Id    string             `json:"id"`
	Url   string             `json:"url"`
	Title string             `json:"title"`
	Eft   bool               `json:"eft"`
	Edges MenuPermissionEdge `json:"edges"`
}

type MenuPermissionEdge struct {
	Children []*MenuPermission `json:"children"`
}

type ViewPermission struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Eft   bool   `json:"eft"`
}
