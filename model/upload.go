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

	me.Filename = fmt.Sprintf("%+v", handler.Filename)
	me.Size = fmt.Sprintf("%+v", handler.Size)
	me.Header = fmt.Sprintf("%+v", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("/tmp", "upload-*.png")
	defer tempFile.Close()
	if err != nil {
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	_, err = tempFile.Write(fileBytes)

	return
}
