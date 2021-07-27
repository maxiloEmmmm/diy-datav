package dataset

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
	"io/ioutil"
	httpUtil "net/http"
)

const Http = "dataset.http"

func init() {
	Reg[Http] = &http{}
}

type http struct{}

type HttpConfig struct {
	Url string `json:"url"`
	Ref int    `json:"ref"`
}

func (h *http) Load(ctx context.Context, config string) (interface{}, error) {
	hc := &HttpConfig{}
	err := json.Unmarshal([]byte(config), hc)
	if err != nil {
		return EmptyData{}, err
	}

	if hc.Ref != 0 {
		ref, err := app.Db.TypeConfig.Query().Where(typeconfig.ID(hc.Ref)).First(ctx)
		if err != nil {
			return EmptyData{}, fmt.Errorf("http.Load: get ref(%d) err %w", hc.Ref, err)
		}

		hc = &HttpConfig{}
		err = json.Unmarshal([]byte(ref.Config), hc)
		if err != nil {
			return EmptyData{}, fmt.Errorf("http.Load: parse ref(%d)'s config err %w", hc.Ref, err)
		}
	}

	client := &httpUtil.Client{}
	resp, err := client.Get(hc.Url)
	if err != nil {
		return EmptyData{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EmptyData{}, err
	}

	td := EmptyData{}
	err = json.Unmarshal(data, &td)
	if err != nil {
		return EmptyData{}, err
	}

	return td, nil
}
