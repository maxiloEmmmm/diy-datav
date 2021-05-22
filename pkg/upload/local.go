package upload

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const Local = "upload.local"

func init() {
	Reg[Local] = &local{}
}

type local struct {

}

func(l *local)Upload(dirPath string, reader io.Reader) (string, error) {
	path := filepath.Join(dirPath, uuid.NewString())

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	n, err := file.Write(data)
	if err != nil || n != len(data) {
		return "", err
	}

	return path, nil
}

func(l *local)Remove(path string) error {
	return os.Remove(path)
}

func(l *local)WebDownload(path, ext string, rw http.ResponseWriter, req *http.Request) error {
	rw.Header().Set("Content-Type", ext)
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	n, err := rw.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return fmt.Errorf("web download failed with assets: %s", path)
	}

	return nil
}
