package dataset

import (
	"encoding/json"
	"io/ioutil"
	httpUtil "net/http"
)

const Http = "dataset.http"

func init() {
	Reg[Http] = &http{}
}

type http struct {}

type HttpConfig struct {
	Url string
}

func (h *http) Load(config string) (interface{}, error) {
	hc := &HttpConfig{}
	err := json.Unmarshal([]byte(config), hc)
	if err != nil {
		return EmptyData{}, err
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

	td := &DataDesc{}
	err = json.Unmarshal(data, td)
	if err != nil {
		return EmptyData{}, err
	}

	return td, nil
}
