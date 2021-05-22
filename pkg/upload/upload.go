package upload

import (
	"fmt"
	"io"
	"net/http"
)

type Upload interface {
	Upload(dirPath string, reader io.Reader) (string, error)
	Remove(path string) error
	WebDownload(path, ext string, rw http.ResponseWriter, req *http.Request) error
}

var Reg = make(map[string]Upload, 0)

func NewUpload(typ string) (Upload, error) {
	engine, ok := Reg[typ]
	if ok {
		return engine, nil
	}

	return nil, fmt.Errorf("not found upload type for %s", typ)
}