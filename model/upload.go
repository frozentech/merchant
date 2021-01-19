package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// LimitInBytes ...
	LimitInBytes = 50000
)

// Upload ...
type Upload struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
	Header   string `json:"header"`
}

// NewUpload ...
func NewUpload() *Upload {
	return &Upload{}
}

// Upload ...
func (me *Upload) Upload(name string, r *http.Request) (err error) {
	r.ParseMultipartForm(LimitInBytes)
	file, handler, err := r.FormFile(name)
	defer file.Close()
	if err != nil {
		return
	}

	filename := fmt.Sprintf("upload-%s.png", GenerateUUID())

	me.Filename = fmt.Sprintf("/assets/%s", filename)
	me.Size = fmt.Sprintf("%+v", handler.Size)
	me.Header = fmt.Sprintf("%+v", handler.Header)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("/tmp/"+filename, fileBytes, 0644)

	return
}
