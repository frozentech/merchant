package controller

import (
	"net/http"

	merchantError "github.com/frozentech/merchant/error"
	"github.com/frozentech/merchant/model"
)

// Upload ...
func Upload(w http.ResponseWriter, r *http.Request) {
	var (
		upload = model.NewUpload()
		resp   = model.Response{}
	)

	err := upload.Upload("myfile", r)
	if err != nil {
		resp = merchantError.StatusRecord(merchantError.UploadFailed)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)

	resp.Result = upload
	Output(w, http.StatusCreated, resp)
	return
}
